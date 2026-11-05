package record

import (
	"engine/modules/ecs"
	"reflect"
)

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
	zero := func() Component {
		var zero Component
		return zero
	}
	componentType := reflect.TypeFor[Component]()
	i, ok := config.ComponentsIndices[componentType]
	if ok {
		goto cleanup
	}
	i = len(*config.ComponentsOrder)
	*config.ComponentsOrder = append(*config.ComponentsOrder, componentType)
	config.ComponentsIndices[componentType] = i
	config.RecordedComponents[componentType] = func(w ecs.World) ecs.AnyComponentArray {
		return ecs.GetComponentArray[Component](w)
	}
	config.InheritZero[componentType] = func(inherit ecs.World) {
		zero = ecs.GetComponentArray[Component](inherit).GetEmpty
	}

cleanup:
	return func(components []any) (Component, bool) {
		if len(components) == 0 {
			return zero(), false
		}
		if components[i] == nil {
			return zero(), false
		}
		return components[i].(Component), true
	}
}

//
