package ecs

type ApplyEntityEvent interface {
	ApplyEntity(entityEmitting EntityID) (event any)
}

//

type RemoveEntityEvent struct{ Entity EntityID }

func NewRemoveEntityEvent(entity EntityID) RemoveEntityEvent {
	return RemoveEntityEvent{entity}
}
