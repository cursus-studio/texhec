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
	Clone(ecs.EntityID) ecs.EntityID
}
