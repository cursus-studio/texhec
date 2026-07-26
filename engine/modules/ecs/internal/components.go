package internal

import (
	"engine/modules/ecs/internal/ecstypes"
	"reflect"
)

// interface

type componentType struct {
	componentType reflect.Type
}

func (t *componentType) String() string { return t.componentType.String() }

func newComponentType(t reflect.Type) componentType {
	return componentType{componentType: t}
}

//

type Component any

func getComponentType(component Component) componentType {
	typeOfComponent := reflect.TypeOf(component)
	kind := typeOfComponent.Kind()
	if kind != reflect.Struct && kind != reflect.Array {
		panic("component has to be a struct or array (cannot use pointers under the hood)")
	}
	return newComponentType(typeOfComponent)
}

//
//
//

// impl

type componentsImpl struct {
	storage *componentsStorage
}

func (components *componentsImpl) Components() ComponentsStorage { return components.storage }

func (components *componentsImpl) RemoveEntity(entity ecstypes.EntityID) {
	for _, arr := range components.storage.arraySlice {
		arr.Remove(entity)
	}
}

func newComponents() *componentsImpl {
	return &componentsImpl{
		storage: newComponentsStorage(),
	}
}

//

type arraysSharedInterface interface {
	ecstypes.AnyComponentArray
	// this adds listeners for change and remove
}

type componentsStorage struct {
	arrays              map[componentType]arraysSharedInterface // any is *componentsArray[ComponentType]
	arraySlice          []arraysSharedInterface
	onArrayAddListeners map[componentType][]func(arraysSharedInterface)
}

type ComponentsStorage *componentsStorage

func newComponentsStorage() ComponentsStorage {
	return &componentsStorage{
		arrays:              make(map[componentType]arraysSharedInterface),
		arraySlice:          make([]arraysSharedInterface, 0),
		onArrayAddListeners: make(map[componentType][]func(arraysSharedInterface)),
	}
}

func GetComponentArray[Component any](rawWorld ecstypes.World) ecstypes.ComponentArray[Component] {
	world := rawWorld.(*world)
	components := world.Components()
	var zero Component
	componentType := getComponentType(zero)

	if array, ok := components.arrays[componentType]; ok {
		return array.(ecstypes.ComponentArray[Component])
	}
	array := newComponentArray[Component](*world.entitiesImpl)
	components.arrays[componentType] = array
	components.arraySlice = append(components.arraySlice, array)
	//
	listeners := components.onArrayAddListeners[componentType]
	for _, listener := range listeners {
		listener(array)
	}
	delete(components.onArrayAddListeners, componentType)
	return array
}

func SaveComponent[Component any](
	w ecstypes.World,
	entity ecstypes.EntityID,
	component Component,
) {
	GetComponentArray[Component](w).Set(entity, component)
}

func GetComponent[Component any](
	w ecstypes.World,
	entity ecstypes.EntityID,
) (Component, bool) {
	return GetComponentArray[Component](w).
		Get(entity)
}

func RemoveComponent[Component any](
	w ecstypes.World,
	entity ecstypes.EntityID,
) {
	GetComponentArray[Component](w).
		Remove(entity)
}

func GetEntitiesWithComponents(
	components ComponentsStorage,
	componentTypes ...componentType,
) []ecstypes.EntityID {
	if len(componentTypes) == 0 {
		return nil
	}

	var arrays []arraysSharedInterface
	for _, componentType := range componentTypes {
		array, ok := components.arrays[componentType]
		if !ok {
			return nil
		}
		arrays = append(arrays, array)
	}

	arrayEntities := arrays[0].GetEntities()
	arrays = arrays[1:]
	finalEntities := []ecstypes.EntityID{}
arrayEntities:
	for _, entity := range arrayEntities {
		for _, array := range arrays {
			if _, ok := array.GetAny(entity); !ok {
				continue arrayEntities
			}
		}
		finalEntities = append(finalEntities, entity)
	}
	return finalEntities
}
