package player

import "engine/services/ecs"

type OwnerComponent struct {
	Owner ecs.EntityID
}

func NewOwner(owner ecs.EntityID) OwnerComponent {
	return OwnerComponent{owner}
}

//

type Service interface {
	Owner() ecs.ComponentsArray[OwnerComponent]
}
