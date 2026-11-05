package collider

import (
	"engine/modules/ecs"
)

type ColliderComponent struct{ ID ecs.EntityID }

func NewCollider(id ecs.EntityID) ColliderComponent {
	return ColliderComponent{ID: id}
}
