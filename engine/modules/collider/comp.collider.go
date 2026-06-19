package collider

import (
	"engine/modules/ecs"
)

type Component struct{ ID ecs.EntityID }

func NewCollider(id ecs.EntityID) Component {
	return Component{ID: id}
}
