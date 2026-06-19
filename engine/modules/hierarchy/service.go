// defines child-parent relationship.
// this is one of core modules on which relies most of the engine.
//
// service stores separate relation cache and updates it on changes to the hierarchy.
// this allows us to have O(1) access time to children and parents
package hierarchy

import (
	"engine/modules/ecs"
	"engine/services/datastructures"
	"errors"
)

var (
	ErrParentCycle error = errors.New("parent cycle is not allowed")
)

type Component struct {
	Parent ecs.EntityID
}

func NewParent(parent ecs.EntityID) Component { return Component{parent} }

//

type Service interface {
	Component() ecs.ComponentArray[Component]

	// returns true if is child of any parent doesn't matter the depth
	IsChildOf(child ecs.EntityID, parent ecs.EntityID) bool
	SetParent(child ecs.EntityID, parent ecs.EntityID)
	Parent(child ecs.EntityID) (ecs.EntityID, bool)

	// from closest to furthest
	GetParents(child ecs.EntityID) datastructures.SparseSetReader[ecs.EntityID]
	GetOrderedParents(child ecs.EntityID) []ecs.EntityID

	// maintains order of children and adds component to children
	// even if children doesn't exist
	SetChildren(parent ecs.EntityID, children ...ecs.EntityID)

	Children(parent ecs.EntityID) datastructures.SparseSetReader[ecs.EntityID]
	// includes children of children
	FlatChildren(parent ecs.EntityID) datastructures.SparseSetReader[ecs.EntityID]
}
