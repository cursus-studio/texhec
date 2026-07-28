// allows us to copy entity with copyable components
//
// its to create copies of entities. its equivalent of unity prefabs (unity semantics)
package prototype

import "engine/modules/ecs"

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
