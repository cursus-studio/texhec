package server

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/loop"
	"engine/modules/netsync"
	"engine/modules/netsync/internal/clienttypes"
	"engine/modules/netsync/internal/config"
	"engine/modules/netsync/internal/servertypes"
	"engine/modules/record"
	"engine/modules/uuid"
	"fmt"
	"reflect"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type Service struct {
	engine.EngineWorld `inject:""`
	config.Config

	recordedEventUUID *uuid.UUID
	recordingID       record.UUIDRecordingID
	onTick            []func()
}

func NewService(c ioc.Dic, config config.Config) *Service {
	s := ioc.GetServices[*Service](c)
	s.Config = config
	s.recordedEventUUID = nil
	s.recordingID = 0
	s.Connection().Component().OnUpsert(s.OnConnectionUpsert)

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
	listeners := map[reflect.Type]func(ecs.EntityID, any){
		reflect.TypeFor[clienttypes.FetchStateDTO](): func(entity ecs.EntityID, a any) {
			s.onTick = append(s.onTick, func() {
				s.ListenFetchState(entity, a.(clienttypes.FetchStateDTO))
			})
		},
		reflect.TypeFor[clienttypes.EmitEventDTO](): func(entity ecs.EntityID, a any) {
			s.ListenEmitEvent(entity, a.(clienttypes.EmitEventDTO))
		},
		reflect.TypeFor[clienttypes.TransparentEventDTO](): func(entity ecs.EntityID, a any) {
			s.ListenTransparentEvent(entity, a.(clienttypes.TransparentEventDTO))
		},
	}

	events.Listen(s.EventsBuilder(), func(loop.FrameEvent) {
		for _, client := range s.NetSync().Client().GetEntities() {
			conn, ok := s.Connection().Component().Get(client)
			if !ok {
				s.Logger().Warn(fmt.Errorf("not connected to server"))
				continue
			}
			messages := conn.Conn().Messages()
			for _, msg := range messages {
				messageType := reflect.TypeOf(msg)
				listener, ok := listeners[messageType]
				if !ok {
					s.Logger().Log(fmt.Errorf("invalid listener called there is no %v type", messageType.String()))
					continue
				}
				listener(client, msg)
			}
		}
	})
	events.Listen(s.EventsBuilder(), func(loop.TickEvent) {
		for _, listener := range s.onTick {
			listener()
		}
		s.onTick = s.onTick[:0]
	})
}

// public methods

func (s *Service) BeforeEvent(event any) {
	if len(s.NetSync().Client().GetEntities()) == 0 {
		return
	}

	if s.recordedEventUUID == nil {
		uuid := s.UUID().NewUUID()
		s.recordedEventUUID = &uuid
	}
	s.recordingID = s.Record().UUID().StartRecording(s.RecordConfig)
}

func (s *Service) AfterEvent(event any) {
	if len(s.NetSync().Client().GetEntities()) == 0 {
		return
	}

	if recording, ok := s.Record().UUID().Stop(s.recordingID); ok && s.recordedEventUUID != nil {
		s.emitChanges(*s.recordedEventUUID, recording)
	}
	s.recordingID = 0
}

func (s *Service) OnTransparentEvent(event any) {
	if len(s.NetSync().Client().GetEntities()) == 0 {
		return
	}

	for _, client := range s.NetSync().Client().GetEntities() {
		connComp, ok := s.Connection().Component().Get(client)
		if !ok {
			return
		}
		s.Logger().Log(connComp.Conn().Send(servertypes.TransparentEventDTO{Event: event}))
	}
}

func (s *Service) ListenFetchState(entity ecs.EntityID, dto clienttypes.FetchStateDTO) {
	state := s.Record().UUID().GetState(s.RecordConfig)
	s.sendVisible(entity, nil, state)
}

func (s *Service) ListenEmitEvent(entity ecs.EntityID, dto clienttypes.EmitEventDTO) {
	conn, ok := s.Connection().Component().Get(entity)
	if !ok {
		return
	}
	event, err := s.Auth(entity, dto.Event)
	if err != nil {
		err := conn.Conn().Send(servertypes.SendChangeDTO{Error: err})
		s.Logger().Log(err)
		return
	}
	s.recordedEventUUID = &dto.ID
	events.EmitAny(s.Events(), event)
}

func (s *Service) ListenTransparentEvent(entity ecs.EntityID, dto clienttypes.TransparentEventDTO) {
	conn, ok := s.Connection().Component().Get(entity)
	if !ok {
		return
	}
	event, err := s.Auth(entity, dto.Event)
	if err != nil {
		err := conn.Conn().Send(servertypes.TransparentEventDTO{Error: err})
		s.Logger().Log(err)
		return
	}
	events.EmitAny(s.Events(), event)
}

// private methods

func (s *Service) sendVisible(client ecs.EntityID, eventUUID *uuid.UUID, changes record.UUIDRecording) {
	connComp, ok := s.Connection().Component().Get(client)
	if !ok {
		return
	}

	// TODO manage visibility
	sentChanges := changes
	// for uuid, _ := range changes.Entities {
	// 	// if cannot use remove it
	// 	delete(changes.Entities, uuid)
	// }

	if len(sentChanges.Entities) == 0 {
		return
	}

	go func() {
		if eventUUID != nil {
			err := connComp.Conn().Send(servertypes.SendChangeDTO{
				EventID: *eventUUID,
				Changes: sentChanges,
			})
			s.Logger().Warn(err)
		} else {
			err := connComp.Conn().Send(servertypes.SendStateDTO{
				State: sentChanges,
			})
			s.Logger().Warn(err)
		}
	}()
}

func (s *Service) emitChanges(eventUUID uuid.UUID, changes record.UUIDRecording) {
	for _, client := range s.NetSync().Client().GetEntities() {
		s.sendVisible(client, &eventUUID, changes)
	}
}
