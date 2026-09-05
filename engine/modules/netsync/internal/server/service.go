package server

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/loop"
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
}

func NewService(c ioc.Dic, config config.Config) *Service {
	s := ioc.GetServices[*Service](c)
	s.Config = config
	s.recordedEventUUID = nil
	s.recordingID = 0

	// listen to server messages
	listeners := map[reflect.Type]func(ecs.EntityID, any){
		reflect.TypeFor[clienttypes.FetchStateDTO](): func(entity ecs.EntityID, a any) {
			s.ListenFetchState(entity, a.(clienttypes.FetchStateDTO))
		},
		reflect.TypeFor[clienttypes.EmitEventDTO](): func(entity ecs.EntityID, a any) {
			s.ListenEmitEvent(entity, a.(clienttypes.EmitEventDTO))
		},
		reflect.TypeFor[clienttypes.TransparentEventDTO](): func(entity ecs.EntityID, a any) {
			s.ListenTransparentEvent(entity, a.(clienttypes.TransparentEventDTO))
		},
	}
	events.Listen(s.EventsBuilder(), func(loop.FrameEvent) {
		for _, clients := range s.NetSync().Client().GetEntities() {
			for _, client := range s.Hierarchy().Children(clients).GetIndices() {
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
		}
	})

	return s
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

// func (t *Service) loadConnections() {
// 	for _, entity := range t.dirtySet.Get() {
// 		if ok := t.listeners.Get(entity); ok {
// 			continue
// 		}
// 		t.listeners.Add(entity)
// 		if _, ok := t.NetSync().Client().Get(entity); !ok {
// 			continue
// 		}
//
// 		comp, ok := t.Connection().Component().Get(entity)
// 		if !ok {
// 			continue
// 		}
// 		messages := comp.Conn().Messages()
// 		go func(entity ecs.EntityID) {
// 			for {
// 				message, ok := <-messages
// 				if !ok {
// 					break
// 				}
// 				t.mutex.Lock()
// 				t.messagesSentFromClient = append(t.messagesSentFromClient, clientMessage{
// 					Client:  entity,
// 					Message: message,
// 				})
// 				t.mutex.Unlock()
// 			}
// 			t.mutex.Lock()
// 			t.toRemove = append(t.toRemove, entity)
// 			t.mutex.Unlock()
// 		}(entity)
// 	}
// }

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
