package prototype

import "engine/services/ecs"

type CloneEvent struct {
	Cloned,
	Clone ecs.EntityID
}

func NewCloneEvent(cloned, clone ecs.EntityID) CloneEvent {
	return CloneEvent{cloned, clone}
}

//

type Service interface {
	Clone(cloned ecs.EntityID) ecs.EntityID
	CloneTo(cloned, clone ecs.EntityID)
}
