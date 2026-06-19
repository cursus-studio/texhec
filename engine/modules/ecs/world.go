package ecs

import (
	"engine/modules/ecs/internal"
	"engine/modules/ecs/internal/ecstypes"
)

// dirty set
type DirtySet = ecstypes.DirtySet

func NewDirtySet() DirtySet { return internal.NewDirtySet() }

// world
type EntityID = ecstypes.EntityID
type Component = ecstypes.Component
type World = ecstypes.World

func NewWorld() World {
	return internal.NewWorld()
}

// component array getter
func GetComponentArray[Component any](world World) ComponentArray[Component] {
	return internal.GetComponentArray[Component](world)
}

// component array
type AnyComponentArray = ecstypes.AnyComponentArray
type ComponentArray[Component any] = ecstypes.ComponentArray[Component]
