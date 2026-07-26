package internal

import (
	"engine/modules/datastructures"
	"engine/modules/ecs/internal/ecstypes"
)

// impl

type entitiesImpl struct {
	counter  ecstypes.EntityID
	holes    datastructures.SparseSet[ecstypes.EntityID]
	entities datastructures.SparseSet[ecstypes.EntityID]
}

func newEntities() *entitiesImpl {
	return &entitiesImpl{
		counter:  0,
		holes:    datastructures.NewSparseSet[ecstypes.EntityID](),
		entities: datastructures.NewSparseSet[ecstypes.EntityID](),
	}
}

func (entitiesStorage *entitiesImpl) GetEntities() []ecstypes.EntityID {
	return entitiesStorage.entities.GetIndices()
}

func (entitiesStorage *entitiesImpl) EntityExists(entity ecstypes.EntityID) bool {
	return entitiesStorage.entities.Get(entity)
}

func (entitiesStorage *entitiesImpl) NewEntity() ecstypes.EntityID {
	if holes := entitiesStorage.holes.GetIndices(); len(holes) != 0 {
		id := holes[0]
		_ = entitiesStorage.holes.Remove(id)
		entitiesStorage.entities.Add(id)
		return id
	}
	entitiesStorage.counter += 1
	id := entitiesStorage.counter
	entitiesStorage.entities.Add(id)
	return id
}

func (entitiesStorage *entitiesImpl) EnsureExists(entity ecstypes.EntityID) {
	if ok := entitiesStorage.entities.Get(entity); ok {
		return
	}

	for entitiesStorage.counter < entity {
		entitiesStorage.counter++
		if id := entitiesStorage.counter; !entitiesStorage.entities.Get(id) {
			entitiesStorage.holes.Add(id)
		}
	}
	entitiesStorage.holes.Remove(entity)
	entitiesStorage.entities.Add(entity)
}

func (entitiesStorage *entitiesImpl) RemoveEntity(entity ecstypes.EntityID) {
	entitiesStorage.holes.Add(entity)
	entitiesStorage.entities.Remove(entity)
}
