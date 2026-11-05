package config

import (
	"engine/modules/ecs"
	"engine/modules/netsync"
	"reflect"
)

// can client emit event
func (config *config) Auth(client ecs.EntityID, event any) (any, error) {
	eventValue := reflect.ValueOf(event)
	eventType := eventValue.Type()
	if _, ok := config.AllowedClientEvents[eventType]; !ok {
		return event, nil
	}
	eventPointerValue := reflect.New(eventType)
	eventPointerValue.Elem().Set(eventValue)

	eventPointer := eventPointerValue.Interface()
	if validatedEvent, ok := eventPointer.(netsync.ValidatedEvent); ok {
		validatedEvent.SetConnection(client)
	}

	for _, validation := range config.GenericValidateEvent {
		if err := validation(client, event); err != nil {
			return nil, err
		}
	}

	event = eventPointerValue.Elem().Interface()
	if handler, ok := config.ValidateEvent[eventType]; ok {
		if err := handler(event); err != nil {
			return nil, err
		}
	}
	return event, nil
}

func (c *config) IsTickEvent(event any) bool {
	config, ok := c.Events[reflect.TypeOf(event)]
	return ok && config.IsTickEvent
}
