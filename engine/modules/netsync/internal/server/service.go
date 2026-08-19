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
	t := ioc.GetServices[*Service](c)
	t.Config = config
	t.recordedEventUUID = nil
	t.recordingID = 0

	// listen to server messages
	listeners := map[reflect.Type]func(ecs.EntityID, any){
		reflect.TypeFor[clienttypes.FetchStateDTO](): func(entity ecs.EntityID, a any) {
			t.ListenFetchState(entity, a.(clienttypes.FetchStateDTO))
		},
		reflect.TypeFor[clienttypes.EmitEventDTO](): func(entity ecs.EntityID, a any) {
			t.ListenEmitEvent(entity, a.(clienttypes.EmitEventDTO))
		},
		reflect.TypeFor[clienttypes.TransparentEventDTO](): func(entity ecs.EntityID, a any) {
			t.ListenTransparentEvent(entity, a.(clienttypes.TransparentEventDTO))
		},
	}
	events.Listen(t.EventsBuilder(), func(loop.FrameEvent) {
		for _, clients := range t.NetSync().Client().GetEntities() {
			for _, client := range t.Hierarchy().Children(clients).GetIndices() {
				conn, ok := t.Connection().Component().Get(client)
				if !ok {
					t.Logger().Warn(fmt.Errorf("not connected to server"))
					continue
				}
				messages := conn.Conn().Messages()
				for _, msg := range messages {
					messageType := reflect.TypeOf(msg)
					listener, ok := listeners[messageType]
					if !ok {
						t.Logger().Log(fmt.Errorf("invalid listener called there is no %v type", messageType.String()))
						continue
					}
					listener(client, msg)
				}
			}
		}
	})

	return t
}

// public methods

func (t *Service) BeforeEvent(event any) {
	if len(t.NetSync().Client().GetEntities()) == 0 {
		return
	}

	if t.recordedEventUUID == nil {
		uuid := t.UUID().NewUUID()
		t.recordedEventUUID = &uuid
	}
	t.recordingID = t.Record().UUID().StartRecording(t.RecordConfig)
}

func (t *Service) AfterEvent(event any) {
	if len(t.NetSync().Client().GetEntities()) == 0 {
		return
	}

	if recording, ok := t.Record().UUID().Stop(t.recordingID); ok && t.recordedEventUUID != nil {
		t.emitChanges(*t.recordedEventUUID, recording)
	}
	t.recordingID = 0
}

func (t *Service) OnTransparentEvent(event any) {
	if len(t.NetSync().Client().GetEntities()) == 0 {
		return
	}

	for _, client := range t.NetSync().Client().GetEntities() {
		connComp, ok := t.Connection().Component().Get(client)
		if !ok {
			return
		}
		t.Logger().Log(connComp.Conn().Send(servertypes.TransparentEventDTO{Event: event}))
	}
}

func (t *Service) ListenFetchState(entity ecs.EntityID, dto clienttypes.FetchStateDTO) {
	state := t.Record().UUID().GetState(t.RecordConfig)
	t.sendVisible(entity, nil, state)
}

func (t *Service) ListenEmitEvent(entity ecs.EntityID, dto clienttypes.EmitEventDTO) {
	conn, ok := t.Connection().Component().Get(entity)
	if !ok {
		return
	}
	event, err := t.Auth(entity, dto.Event)
	if err != nil {
		err := conn.Conn().Send(servertypes.SendChangeDTO{Error: err})
		t.Logger().Log(err)
		return
	}
	t.recordedEventUUID = &dto.ID
	events.EmitAny(t.Events(), event)
}

func (t *Service) ListenTransparentEvent(entity ecs.EntityID, dto clienttypes.TransparentEventDTO) {
	conn, ok := t.Connection().Component().Get(entity)
	if !ok {
		return
	}
	event, err := t.Auth(entity, dto.Event)
	if err != nil {
		err := conn.Conn().Send(servertypes.TransparentEventDTO{Error: err})
		t.Logger().Log(err)
		return
	}
	events.EmitAny(t.Events(), event)
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

func (t *Service) sendVisible(client ecs.EntityID, eventUUID *uuid.UUID, changes record.UUIDRecording) {
	connComp, ok := t.Connection().Component().Get(client)
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
			t.Logger().Warn(err)
		} else {
			err := connComp.Conn().Send(servertypes.SendStateDTO{
				State: sentChanges,
			})
			t.Logger().Warn(err)
		}
	}()
}

func (t *Service) emitChanges(eventUUID uuid.UUID, changes record.UUIDRecording) {
	for _, client := range t.NetSync().Client().GetEntities() {
		t.sendVisible(client, &eventUUID, changes)
	}
}
