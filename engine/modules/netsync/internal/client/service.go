package client

import (
	"engine"
	"engine/modules/connection"
	"engine/modules/ecs"
	"engine/modules/loop"
	"engine/modules/netsync/internal/clienttypes"
	"engine/modules/netsync/internal/config"
	"engine/modules/netsync/internal/servertypes"
	"engine/modules/record"
	"fmt"
	"reflect"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type savedPrediction struct {
	PredictedEvent clienttypes.PredictedEvent
	Snapshot       record.UUIDRecording
}

type recordedPrediction struct {
	PredictedEvent clienttypes.PredictedEvent
}

// can:
// - apply server event
// - apply predicted event (starts and ends prediction)
type Service struct {
	engine.EngineWorld `inject:""`
	config.Config

	recordNextEvent    bool
	predictions        []savedPrediction
	recordedPrediction *recordedPrediction
	sentTransparentEvent,
	receivedTransparentEvent bool
	recordingID record.UUIDRecordingID

	dirtySet ecs.DirtySet
}

func NewService(c ioc.Dic, config config.Config) *Service {
	s := ioc.GetServices[*Service](c)
	s.Config = config
	s.recordNextEvent = true
	s.predictions = make([]savedPrediction, 0)
	s.recordedPrediction = nil
	s.sentTransparentEvent = false
	s.receivedTransparentEvent = false
	s.recordingID = 0

	s.dirtySet = ecs.NewDirtySet()

	s.NetSync().Server().AddDirtySet(s.dirtySet)
	s.Connection().Component().AddDirtySet(s.dirtySet)

	listeners := map[reflect.Type]func(any){
		reflect.TypeFor[servertypes.SendStateDTO](): func(a any) {
			s.ListenSendState(a.(servertypes.SendStateDTO))
		},
		reflect.TypeFor[servertypes.SendChangeDTO](): func(a any) {
			s.ListenSendChange(a.(servertypes.SendChangeDTO))
		},
		reflect.TypeFor[servertypes.TransparentEventDTO](): func(a any) {
			s.ListenTransparentEvent(a.(servertypes.TransparentEventDTO))
		},
	}
	events.Listen(s.EventsBuilder(), func(loop.FrameEvent) {
		for _, entity := range s.dirtySet.Get() {
			if _, ok := s.NetSync().Server().Get(entity); !ok {
				continue
			}
			conn, ok := s.Connection().Component().Get(entity)
			if !ok {
				continue
			}
			err := conn.Conn().Send(clienttypes.FetchStateDTO{})
			s.Logger().Warn(err)
		}
		conn := s.getConnection()
		if conn == nil {
			return
		}

		for _, server := range s.NetSync().Server().GetEntities() {
			conn, ok := s.Connection().Component().Get(server)
			if !ok {
				s.Logger().Warn(fmt.Errorf("not connected to server"))
				continue
			}
			for _, msg := range conn.Conn().Messages() {
				messageType := reflect.TypeOf(msg)
				listener, ok := listeners[messageType]
				if !ok {
					s.Logger().Log(fmt.Errorf("invalid listener of type '%v' called", messageType.String()))
					_ = conn.Conn().Close()
					return
				}
				listener(msg)
			}
		}
	})

	return s
}

// public methods

// doesn't send event to server
func (t *Service) BeforeEventRecord(event any) {
	clientConn := t.getConnection()
	if clientConn == nil {
		return
	}

	if !t.recordNextEvent {
		t.recordNextEvent = true
		return
	}

	if len(t.predictions) > t.MaxPredictions {
		t.Logger().Log(ErrExceededPredictions)
		t.undoPredictions()
		// reconciliate
		if err := clientConn.Send(clienttypes.FetchStateDTO{}); err != nil {
			t.Logger().Log(err)
		}
		return
	}

	t.recordingID = t.Record().UUID().StartBackwardsRecording(t.RecordConfig)
	t.recordedPrediction = &recordedPrediction{
		PredictedEvent: clienttypes.PredictedEvent{
			ID:    t.UUID().NewUUID(),
			Event: event,
		},
	}
}

func (t *Service) BeforeEvent(event any) {
	clientConn := t.getConnection()
	if clientConn == nil {
		return
	}
	t.BeforeEventRecord(event)
	if t.recordedPrediction == nil {
		return
	}

	dto := clienttypes.EmitEventDTO(t.recordedPrediction.PredictedEvent)
	if err := clientConn.Send(dto); err != nil {
		t.Logger().Log(err)
	}
}

func (t *Service) AfterEvent(event any) {
	conn := t.getConnection()
	if conn == nil {
		return
	}

	if t.recordedPrediction == nil {
		return
	}

	recording, ok := t.Record().UUID().Stop(t.recordingID)
	t.recordingID = 0
	if !ok {
		return
	}
	newPrediction := savedPrediction{
		PredictedEvent: t.recordedPrediction.PredictedEvent,
		Snapshot:       recording,
	}
	t.recordedPrediction = nil

	t.predictions = append(t.predictions, newPrediction)
}

func (t *Service) OnTransparentEvent(event any) {
	if t.receivedTransparentEvent {
		t.receivedTransparentEvent = false
		return
	}
	conn := t.getConnection()
	if conn == nil {
		return
	}

	t.sentTransparentEvent = true
	err := conn.Send(clienttypes.TransparentEventDTO{Event: event})
	t.Logger().Log(err)
}

func (t *Service) ListenSendChange(dto servertypes.SendChangeDTO) {
	conn := t.getConnection()
	if conn == nil {
		return
	}
	if dto.Error != nil {
		predictedEvents := t.undoPredictions()
		// reApplied events are events without applied event
		reEmitedEvents := make([]clienttypes.PredictedEvent, 0, len(predictedEvents))
		for _, predictedEvent := range predictedEvents {
			if predictedEvent.ID != dto.EventID {
				reEmitedEvents = append(reEmitedEvents, predictedEvent)
			}
		}
		t.applyPredictedEvents(reEmitedEvents)
		t.Logger().Log(dto.Error)
		return
	}
	// check is event predicted. if is then remove first event from queue
	// if isn't then undo predictions, emit server event(as not recordable), emit all predicted events again
	if len(t.predictions) == 0 {
		t.Record().UUID().Apply(t.RecordConfig, dto.Changes)
		return
	}
	if t.predictions[0].PredictedEvent.ID == dto.EventID {
		t.predictions = t.predictions[1:]
		return
		// TODO later. add test is prediction correct
		// replace this with state comparer for first prediction
		// stateEqual := true
		// if !stateEqual {
		// 	t.logger.Warn(ErrInvalidPrediction)
		// 	predictedEvents := t.UndoPredictions()
		// 	t.ApplyState(dto.Changes)
		// 	t.ApplyPredictedEvents(predictedEvents[1:])
		// } else {
		// 	t.predictions = t.predictions[1:]
		// }
	}
	predictedEvents := t.undoPredictions()
	t.Record().UUID().Apply(t.RecordConfig, dto.Changes)
	// reApplied events are events without applied event
	reEmitedEvents := make([]clienttypes.PredictedEvent, 0, len(predictedEvents))
	for _, predictedEvent := range predictedEvents {
		if predictedEvent.ID != dto.EventID {
			reEmitedEvents = append(reEmitedEvents, predictedEvent)
		}
	}
	t.applyPredictedEvents(reEmitedEvents)
}

// reconciliate
func (s *Service) ListenSendState(dto servertypes.SendStateDTO) {
	conn := s.getConnection()
	if conn == nil {
		return
	}
	if dto.Error != nil {
		s.predictions = nil
		s.Logger().Log(dto.Error)
		_ = conn.Close()
		return
	}
	s.predictions = nil
	s.Record().UUID().Apply(s.RecordConfig, dto.State)
}

func (t *Service) ListenTransparentEvent(dto servertypes.TransparentEventDTO) {
	if t.sentTransparentEvent {
		t.sentTransparentEvent = false
		return
	}
	if dto.Error != nil {
		t.Logger().Log(dto.Error)
		return
	}
	t.receivedTransparentEvent = true
	events.EmitAny(t.Events(), dto.Event)
}

// private methods

func (t *Service) undoPredictions() []clienttypes.PredictedEvent {
	// add events to the list
	var unDoneEvents []clienttypes.PredictedEvent
	snapshots := make([]record.UUIDRecording, len(t.predictions))
	for _, prediction := range t.predictions {
		unDoneEvents = append(unDoneEvents, prediction.PredictedEvent)
		// snapshots = append([]record.UUIDRecording{prediction.Snapshot}, snapshots...)
		snapshots = append(snapshots, prediction.Snapshot)
	}
	t.Record().UUID().Apply(t.RecordConfig, snapshots...)
	t.predictions = nil
	return unDoneEvents
}

func (t *Service) applyPredictedEvents(predictedEvents []clienttypes.PredictedEvent) {
	for _, predictedEvent := range predictedEvents[1:] {
		t.recordNextEvent = false
		events.EmitAny(t.Events(), predictedEvent.Event)
	}
}

func (t *Service) getConnection() connection.Conn {
	var conn connection.Conn
	if entities := t.NetSync().Server().GetEntities(); len(entities) == 1 {
		server := entities[0]
		comp, ok := t.Connection().Component().Get(server)
		if ok {
			conn = comp.Conn()
		}
	}
	if conn == nil { // isn't client clear all client data
		t.recordedPrediction = nil
		t.predictions = nil
	}
	return conn
}
