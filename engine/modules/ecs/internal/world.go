package internal

import "engine/modules/ecs/internal/ecstypes"

type world struct {
	*entitiesImpl
	*componentsImpl
}

func NewWorld() ecstypes.World {
	entitiesImpl := newEntities()
	componentsImpl := newComponents(entitiesImpl.entities)

	return &world{
		entitiesImpl:   entitiesImpl,
		componentsImpl: componentsImpl,
	}
}

func (world *world) NewEntity() ecstypes.EntityID {
	entity := world.entitiesImpl.NewEntity()
	return entity
}

func (world *world) RemoveEntity(entity ecstypes.EntityID) {
	world.entitiesImpl.RemoveEntity(entity)

	for _, arr := range world.storage.arraySlice {
		arr.Remove(entity)
	}
}

func (world *world) GetEntities() []ecstypes.EntityID {
	return world.entitiesImpl.GetEntities()
}

func (world *world) EntityExists(entity ecstypes.EntityID) bool {
	return world.entitiesImpl.EntityExists(entity)
}

func (world *world) WarmUp() {
	for _, arr := range world.storage.arraySlice {
		arr.GetEntities()
	}
}
