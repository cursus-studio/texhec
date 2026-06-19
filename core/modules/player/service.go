// allowes objects to be owned
package player

import "engine/modules/ecs"

type OwnerComponent struct {
	Owner ecs.EntityID
}

func NewOwner(owner ecs.EntityID) OwnerComponent {
	return OwnerComponent{owner}
}

//

type Service interface {
	Owner() ecs.ComponentArray[OwnerComponent]
}
