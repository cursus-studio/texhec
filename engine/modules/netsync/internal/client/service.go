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
func (s *Service) BeforeEventRecord(event any) {
	clientConn := s.getConnection()
	if clientConn == nil {
		return
	}

	if !s.recordNextEvent {
		s.recordNextEvent = true
		return
	}

	if len(s.predictions) > s.MaxPredictions {
		s.Logger().Log(ErrExceededPredictions)
		s.undoPredictions()
		// reconciliate
		if err := clientConn.Send(clienttypes.FetchStateDTO{}); err != nil {
			s.Logger().Log(err)
		}
		return
	}

	s.recordingID = s.Record().UUID().StartBackwardsRecording(s.RecordConfig)
	s.recordedPrediction = &recordedPrediction{
		PredictedEvent: clienttypes.PredictedEvent{
			ID:    s.UUID().NewUUID(),
			Event: event,
		},
	}
}

func (s *Service) BeforeEvent(event any) {
	clientConn := s.getConnection()
	if clientConn == nil {
		return
	}
	s.BeforeEventRecord(event)
	if s.recordedPrediction == nil {
		return
	}

	dto := clienttypes.EmitEventDTO(s.recordedPrediction.PredictedEvent)
	if err := clientConn.Send(dto); err != nil {
		s.Logger().Log(err)
	}
}

func (s *Service) AfterEvent(event any) {
	conn := s.getConnection()
	if conn == nil {
		return
	}

	if s.recordedPrediction == nil {
		return
	}

	recording, ok := s.Record().UUID().Stop(s.recordingID)
	s.recordingID = 0
	if !ok {
		return
	}
	newPrediction := savedPrediction{
		PredictedEvent: s.recordedPrediction.PredictedEvent,
		Snapshot:       recording,
	}
	s.recordedPrediction = nil

	s.predictions = append(s.predictions, newPrediction)
}

func (s *Service) OnTransparentEvent(event any) {
	if s.receivedTransparentEvent {
		s.receivedTransparentEvent = false
		return
	}
	conn := s.getConnection()
	if conn == nil {
		return
	}

	s.sentTransparentEvent = true
	err := conn.Send(clienttypes.TransparentEventDTO{Event: event})
	s.Logger().Log(err)
}

func (s *Service) ListenSendChange(dto servertypes.SendChangeDTO) {
	conn := s.getConnection()
	if conn == nil {
		return
	}
	if dto.Error != nil {
		predictedEvents := s.undoPredictions()
		// reApplied events are events without applied event
		reEmitedEvents := make([]clienttypes.PredictedEvent, 0, len(predictedEvents))
		for _, predictedEvent := range predictedEvents {
			if predictedEvent.ID != dto.EventID {
				reEmitedEvents = append(reEmitedEvents, predictedEvent)
			}
		}
		s.applyPredictedEvents(reEmitedEvents)
		s.Logger().Log(dto.Error)
		return
	}
	// check is event predicted. if is then remove first event from queue
	// if isn't then undo predictions, emit server event(as not recordable), emit all predicted events again
	if len(s.predictions) == 0 {
		s.Record().UUID().Apply(s.RecordConfig, dto.Changes)
		return
	}
	if s.predictions[0].PredictedEvent.ID == dto.EventID {
		s.predictions = s.predictions[1:]
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
	predictedEvents := s.undoPredictions()
	s.Record().UUID().Apply(s.RecordConfig, dto.Changes)
	// reApplied events are events without applied event
	reEmitedEvents := make([]clienttypes.PredictedEvent, 0, len(predictedEvents))
	for _, predictedEvent := range predictedEvents {
		if predictedEvent.ID != dto.EventID {
			reEmitedEvents = append(reEmitedEvents, predictedEvent)
		}
	}
	s.applyPredictedEvents(reEmitedEvents)
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

func (s *Service) ListenTransparentEvent(dto servertypes.TransparentEventDTO) {
	if s.sentTransparentEvent {
		s.sentTransparentEvent = false
		return
	}
	if dto.Error != nil {
		s.Logger().Log(dto.Error)
		return
	}
	s.receivedTransparentEvent = true
	events.EmitAny(s.Events(), dto.Event)
}

// private methods

func (s *Service) undoPredictions() []clienttypes.PredictedEvent {
	// add events to the list
	var unDoneEvents []clienttypes.PredictedEvent
	snapshots := make([]record.UUIDRecording, len(s.predictions))
	for _, prediction := range s.predictions {
		unDoneEvents = append(unDoneEvents, prediction.PredictedEvent)
		// snapshots = append([]record.UUIDRecording{prediction.Snapshot}, snapshots...)
		snapshots = append(snapshots, prediction.Snapshot)
	}
	s.Record().UUID().Apply(s.RecordConfig, snapshots...)
	s.predictions = nil
	return unDoneEvents
}

func (s *Service) applyPredictedEvents(predictedEvents []clienttypes.PredictedEvent) {
	for _, predictedEvent := range predictedEvents[1:] {
		s.recordNextEvent = false
		events.EmitAny(s.Events(), predictedEvent.Event)
	}
}

func (s *Service) getConnection() connection.Conn {
	var conn connection.Conn
	if entities := s.NetSync().Server().GetEntities(); len(entities) == 1 {
		server := entities[0]
		comp, ok := s.Connection().Component().Get(server)
		if ok {
			conn = comp.Conn()
		}
	}
	if conn == nil { // isn't client clear all client data
		s.recordedPrediction = nil
		s.predictions = nil
	}
	return conn
}
