package server

import (
	"engine"
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"engine/modules/loop"
	"engine/modules/netsync"
	"engine/modules/netsync/internal/clienttypes"
	"engine/modules/netsync/internal/config"
	"engine/modules/netsync/internal/servertypes"
	"engine/modules/record"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type Service struct {
	engine.EngineWorld `inject:""`
	Config             ioc.Lazy[config.Config] `inject:""`

	recordingID record.RecordingID
	state       record.Recording

	ignoreNext bool

	onTick []func()
}

func NewService(c ioc.Dic) *Service {
	s := ioc.GetServices[*Service](c)
	s.Connection().Component().OnUpsert(s.OnConnectionUpsert)

	s.recordingID = s.Record().StartRecording(s.Config().RecordConfig)

	s.state = record.Recording{
		Entities: datastructures.NewSparseArray[ecs.EntityID, []any](),
	}

	return s
}

func (s *Service) OnConnectionUpsert(entity ecs.EntityID) {
	parent, ok := s.Hierarchy().Parent(entity)
	if !ok {
		return
	}
	if _, ok := s.NetSync().Clients().Get(parent); !ok {
		return
	}
	s.NetSync().Client().Set(entity, netsync.ClientComponent{})
}

func (s *Service) AddBeforeListeners() {
	// listen to server messages
	events.Listen(s.EventsBuilder(), s.ListenReconcile)
	events.Listen(s.EventsBuilder(), s.ListenPredictedEvent)

	events.Listen(s.EventsBuilder(), func(loop.TickEvent) {
		for _, listener := range s.onTick {
			listener()
		}
		s.onTick = s.onTick[:0]
	})
}
func (s *Service) AddAfterListeners() {
	events.Listen(s.EventsBuilder(), s.AfterTick)
}
func (s *Service) BeforeEvent(event any) {
	if s.ignoreNext {
		s.ignoreNext = false
		return
	}
	for _, client := range s.NetSync().Client().GetEntities() {
		connComp, ok := s.Connection().Component().Get(client)
		if !ok {
			continue
		}
		response := servertypes.EventDTO{Event: event}
		s.Logger().Log(connComp.Conn().Send(response))
	}
}
func (s *Service) AfterTick(loop.TickEvent) {
	recording, ok := s.Record().Peek(s.recordingID)
	if !ok {
		return
	}
	for _, entity := range recording.Entities.GetIndices() {
		components, _ := recording.Entities.Get(entity)
		if len(components) == 0 {
			s.state.Entities.Remove(entity)
			continue
		}
		s.state.Entities.Set(entity, components)
	}
}

// public methods

func (s *Service) ListenReconcile(dto clienttypes.ReconcileDTO) {
	connComp, ok := s.Connection().Component().Get(dto.ConnEntity)
	if !ok {
		return
	}
	err := connComp.Conn().Send(servertypes.ReconcileDTO{
		TickUnixNano: s.Loop().LastTick().UnixNano,
		State:        s.Record().UUIDState(s.state),
	})
	s.Logger().Warn(err)
}

func (s *Service) ListenPredictedEvent(dto clienttypes.PredictedEventDTO) {
	conn, ok := s.Connection().Component().Get(dto.ConnEntity)
	if !ok {
		return
	}
	event, err := s.Config().Auth(dto.ConnEntity, dto.Event)
	if err != nil {
		response := servertypes.PredictionResponseDTO{ID: dto.ID, Accepted: false}
		s.Logger().Log(conn.Conn().Send(response))
		return
	}
	response := servertypes.PredictionResponseDTO{ID: dto.ID, Accepted: true}
	s.Logger().Log(conn.Conn().Send(response))

	for _, client := range s.NetSync().Client().GetEntities() {
		if client == dto.ConnEntity {
			continue
		}
		connComp, ok := s.Connection().Component().Get(client)
		if !ok {
			continue
		}
		response := servertypes.EventDTO{Event: dto.Event}
		s.Logger().Log(connComp.Conn().Send(response))
	}
	s.ignoreNext = true
	events.EmitAny(s.Events(), event)
}
