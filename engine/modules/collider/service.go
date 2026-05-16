// this module allows us to check collision between objects and rays
// current algorithm used under the hood is spatial algorithm.
package collider

import (
	"engine/services/ecs"
)

type Service interface {
	Component() ecs.ComponentsArray[Component]

	// todo add collision groups
	// narrow
	CollidesWithRay(ecs.EntityID, Ray) *ObjectRayCollision
	CollidesWithObject(ecs.EntityID, ecs.EntityID) *ObjectObjectCollision

	// broad
	Raycast(Ray) *ObjectRayCollision
	RaycastAll(Ray) []ObjectRayCollision
	NarrowCollisions(ecs.EntityID) []ecs.EntityID

	AddRayFallThroughPolicy(FallTroughPolicy)
}

type FallTroughPolicy interface {
	// position is normalized to be between -1 and 1
	FallThrough(target ObjectRayCollision) bool
}
