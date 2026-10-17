package config

import (
	"engine/modules/ecs"
	"engine/modules/record"
	"fmt"
	"reflect"

	"github.com/ogiusek/events"
)

type EventConfig = *eventConfig
type eventConfig struct {
	IsTickEvent bool
	Builder     func(events.Builder, func(any))
}

func NewEventConfig(isTickEvent bool, builder func(events.Builder, func(any))) EventConfig {
	return &eventConfig{
		IsTickEvent: isTickEvent,
		Builder:     builder,
	}
}

//

type Config = *config
type config struct {
	EventTypes   []reflect.Type
	EventConfigs []EventConfig
	Events       map[reflect.Type]EventConfig

	TickEvent []int

	RecordConfig record.Config

	// client
	MaxPredictions uint8

	// auth
	AllowedClientEvents  map[reflect.Type]struct{}
	GenericValidateEvent []func(ecs.EntityID, any) error
	ValidateEvent        map[reflect.Type]func(any) error
}

func NewConfig() Config {
	return &config{
		EventTypes:   make([]reflect.Type, 0),
		EventConfigs: make([]EventConfig, 0),
		Events:       make(map[reflect.Type]EventConfig),

		TickEvent: make([]int, 0),

		RecordConfig: record.NewConfig(),

		// client
		MaxPredictions: 255,

		// auth
		AllowedClientEvents:  make(map[reflect.Type]struct{}),
		GenericValidateEvent: make([]func(ecs.EntityID, any) error, 0),
		ValidateEvent:        make(map[reflect.Type]func(any) error),
	}
}

func AddEvent[EventType any](config Config) {
	eventType := reflect.TypeFor[EventType]()
	if _, ok := config.Events[eventType]; ok {
		panic(fmt.Sprintf("tried to register \"%v\" second time", eventType.String()))
	}

	eventConfig := NewEventConfig(false, func(b events.Builder, f func(any)) {
		events.Listen(b, func(e EventType) { f(e) })
	})
	config.EventTypes = append(config.EventTypes, eventType)
	config.EventConfigs = append(config.EventConfigs, eventConfig)
	config.Events[eventType] = eventConfig
	config.AllowedClientEvents[eventType] = struct{}{}
}

func AddTickEvent[EventType any](config Config) {
	AddEvent[EventType](config)
	config.Events[reflect.TypeFor[EventType]()].IsTickEvent = true
}

func AddGenericEventValidation(config Config, handler func(client ecs.EntityID, event any) error) {
	config.GenericValidateEvent = append(config.GenericValidateEvent, handler)
}

func AddEventValidation[EventType any](config Config, handler func(event EventType) error) {
	eventType := reflect.TypeFor[EventType]()
	config.ValidateEvent[eventType] = func(a any) error {
		return handler(a.(EventType))
	}
}
