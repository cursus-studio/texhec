package record

import (
	"engine/modules/uuid"
	"engine/services/datastructures"
	"engine/services/ecs"
	"reflect"
)

type Service interface {
	Entity() EntityKeyedRecorder
	UUID() UUIDKeyedRecorder
}

//

type Config struct {
	ComponentsOrder    *[]reflect.Type
	ComponentsIndices  map[reflect.Type]int
	RecordedComponents map[reflect.Type]func(ecs.World) ecs.AnyComponentArray
	InheritZero        map[reflect.Type]func(ecs.World)
}

func NewConfig() Config {
	componentsOrder := make([]reflect.Type, 0)
	return Config{
		ComponentsOrder:    &componentsOrder,
		ComponentsIndices:  make(map[reflect.Type]int),
		RecordedComponents: make(map[reflect.Type]func(ecs.World) ecs.AnyComponentArray),
		InheritZero:        make(map[reflect.Type]func(ecs.World)),
	}
}

type ComponentGetter[Component any] func(components []any) (Component, bool)

func AddToConfig[Component any](config Config) ComponentGetter[Component] {
	var zero Component
	componentType := reflect.TypeFor[Component]()
	i, ok := config.ComponentsIndices[componentType]
	if ok {
		goto cleanup
	}
	i = len(*config.ComponentsOrder)
	*config.ComponentsOrder = append(*config.ComponentsOrder, componentType)
	config.ComponentsIndices[componentType] = i
	config.RecordedComponents[componentType] = func(w ecs.World) ecs.AnyComponentArray {
		return ecs.GetComponentsArray[Component](w)
	}
	config.InheritZero[componentType] = func(inherit ecs.World) {
		arr := ecs.GetComponentsArray[Component](inherit)
		arr.OnEmptyChange(func(c Component) {
			zero = c
		})
	}

cleanup:
	return func(components []any) (Component, bool) {
		if len(components) == 0 {
			return zero, false
		}
		if components[i] == nil {
			return zero, false
		}
		return components[i].(Component), true
	}
}

//

type EntityKeyedRecorder interface {
	// gets state as finished recording
	GetState(Config) Recording

	// starts opened recording (opened recording is recorded until stopped)
	// applying it on previous state will create current state
	StartRecording(Config) RecordingID
	// starts opened recording (opened recording is recorded until stopped)
	// applying it rewinds state.
	StartBackwardsRecording(Config) RecordingID
	// finishes recording if open (false is returned if recording isn't started)
	Stop(RecordingID) (r Recording, ok bool)

	Apply(Config, ...Recording)
}

type RecordingID uint16
type Recording struct {
	// [componentArrayLayoutID]any component
	// nil for removed entity
	Entities datastructures.SparseArray[ecs.EntityID, []any]
}

type UUIDKeyedRecorder interface {
	// gets state as finished recording
	GetState(Config) UUIDRecording

	// starts opened recording (opened recording is recorded until stopped)
	// applying it on previous state will create current state
	StartRecording(Config) UUIDRecordingID
	// starts opened recording (opened recording is recorded until stopped)
	// applying it rewinds state.
	StartBackwardsRecording(Config) UUIDRecordingID
	// finishes recording if open (false is returned if recording isn't started)
	Stop(UUIDRecordingID) (r UUIDRecording, ok bool)

	Apply(Config, ...UUIDRecording)
}

type UUIDRecordingID uint16
type UUIDRecording struct {
	// map[componentUUID][componentArrayLayoutID]any component
	// map[componentUUID]nil is when entity is removed
	Entities map[uuid.UUID][]any
}
