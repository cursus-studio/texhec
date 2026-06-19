package internal

import (
	"engine/modules/ecs/internal/ecstypes"
	"engine/services/datastructures"
)

// impl

type entitiesImpl struct {
	counter ecstypes.EntityID
	holes   datastructures.SparseSet[ecstypes.EntityID]

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
	for entitiesStorage.counter+1 < entity {
		entitiesStorage.counter += 1
		holeEntity := entitiesStorage.counter
		entitiesStorage.holes.Add(holeEntity)
	}
	entitiesStorage.counter = max(entitiesStorage.counter, entity)
	entitiesStorage.entities.Add(entity)
	entitiesStorage.holes.Remove(entity)
}

func (entitiesStorage *entitiesImpl) RemoveEntity(entity ecstypes.EntityID) {
	entitiesStorage.holes.Add(entity)
	entitiesStorage.entities.Remove(entity)
}
