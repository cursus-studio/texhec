package hierarchyservice

import (
	"engine"
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"engine/modules/hierarchy"

	"github.com/ogiusek/ioc/v2"
)

type parentComponent struct{}

type service struct {
	engine.EngineWorld `inject:""`
	hierarchyArray     ecs.ComponentArray[hierarchy.Component]
	parentArray        ecs.ComponentArray[parentComponent]

	parents      datastructures.SparseArray[ecs.EntityID, ecs.EntityID]
	children     datastructures.SparseArray[ecs.EntityID, datastructures.SparseSet[ecs.EntityID]]
	flatChildren datastructures.SparseArray[ecs.EntityID, datastructures.SparseSet[ecs.EntityID]]
}

func NewService(c ioc.Dic) hierarchy.Service {
	s := ioc.GetServices[*service](c)

	s.hierarchyArray = ecs.GetComponentArray[hierarchy.Component](s.World())
	s.parentArray = ecs.GetComponentArray[parentComponent](s.World())
	s.parents = datastructures.NewSparseArray[ecs.EntityID, ecs.EntityID]()
	s.children = datastructures.NewSparseArray[ecs.EntityID, datastructures.SparseSet[ecs.EntityID]]()
	s.flatChildren = datastructures.NewSparseArray[ecs.EntityID, datastructures.SparseSet[ecs.EntityID]]()

	s.hierarchyArray.OnUpsert(s.handleHierarchyChange)
	s.hierarchyArray.OnRemove(s.handleHierarchyChange)
	s.parentArray.OnRemove(s.handleParentRemoval)

	return s
}

func (s *service) Component() ecs.ComponentArray[hierarchy.Component] {
	return s.hierarchyArray
}

func (s *service) IsChildOf(child ecs.EntityID, wantedParent ecs.EntityID) bool {
	for {
		parent, ok := s.Parent(child)
		if !ok {
			return false
		}
		child = parent
		if child == wantedParent {
			return true
		}
	}
}

func (s *service) SetParent(child ecs.EntityID, parent ecs.EntityID) {
	s.hierarchyArray.Set(child, hierarchy.NewParent(parent))
}

func (s *service) Parent(child ecs.EntityID) (ecs.EntityID, bool) {
	comp, ok := s.hierarchyArray.Get(child)
	return comp.Parent, ok
}

//

func (s *service) GetParents(child ecs.EntityID) datastructures.SparseSetReader[ecs.EntityID] {
	orderedParents := s.GetOrderedParents(child)

	parents := datastructures.NewSparseSetWithPaging[ecs.EntityID]()
	for _, parent := range orderedParents {
		parents.Add(parent)
	}
	return parents
}

func (s *service) GetOrderedParents(child ecs.EntityID) []ecs.EntityID {
	parents := []ecs.EntityID{child}
	for {
		parent, ok := s.hierarchyArray.Get(child)
		if !ok {
			return parents[1:]
		}
		parents = append(parents, parent.Parent)
		if len(parents) != 1 && parents[0] == parent.Parent {
			return nil
		}
		child = parent.Parent
	}
}

func (s *service) GetOrderedPreviousParents(child ecs.EntityID) []ecs.EntityID {
	parents := []ecs.EntityID{child}
	for {
		parent, ok := s.parents.Get(child)
		if !ok {
			return parents[1:]
		}
		parents = append(parents, parent)
		if len(parents) != 1 && parents[0] == parent {
			return nil
		}
		child = parent
	}
}

//

func (s *service) SetChildren(parent ecs.EntityID, children ...ecs.EntityID) {
	previousChildren := s.Children(parent).GetIndices()
	for _, child := range previousChildren {
		s.hierarchyArray.Remove(child)
	}

	for i := range len(children) {
		s.SetParent(children[i], parent)
	}
}

//

func (s *service) Children(parent ecs.EntityID) datastructures.SparseSetReader[ecs.EntityID] {
	children, ok := s.children.Get(parent)
	if !ok {
		return datastructures.NewSparseSetWithPaging[ecs.EntityID]()
	}
	return children
}

func (s *service) GetFlatChildren(parent ecs.EntityID) datastructures.SparseSetReader[ecs.EntityID] {
	if flatChildren, ok := s.flatChildren.Get(parent); ok {
		return flatChildren
	}
	flatChildren := datastructures.NewSparseSetWithPaging[ecs.EntityID]()

	children, ok := s.children.Get(parent)
	if !ok {
		return flatChildren
	}

	childrens := []datastructures.SparseSet[ecs.EntityID]{children}

	for len(childrens) != 0 {
		children := childrens[0]
		childrens = childrens[1:]

		for _, child := range children.GetIndices() {
			if added := flatChildren.Add(child); !added {
				continue
			}
			children, ok := s.children.Get(child)
			if !ok {
				continue
			}
			childrens = append(childrens, children)
		}
	}

	s.flatChildren.Set(parent, flatChildren)
	return flatChildren
}

func (s *service) FlatChildren(parent ecs.EntityID) datastructures.SparseSetReader[ecs.EntityID] {
	return s.GetFlatChildren(parent)
}

//

func (s *service) handleHierarchyChange(child ecs.EntityID) {
	previousParent, previousParentOk := s.parents.Get(child)
	hierarchy, nextParentOk := s.hierarchyArray.Get(child)
	if previousParentOk == nextParentOk && hierarchy.Parent == previousParent {
		return
	}

	if previousParentOk { // remove in parents
		s.parents.Remove(child)

		for _, parent := range s.GetOrderedPreviousParents(child) {
			s.flatChildren.Remove(parent)
		}

		// remove as a child
		children, ok := s.children.Get(previousParent)
		if !ok { // this shouldn't occur and means invalid internal state
			goto addCurrentParent
		}
		children.Remove(child)
		if len(children.GetIndices()) == 0 {
			s.children.Remove(previousParent)
		}
	}

addCurrentParent:
	nextParent := hierarchy.Parent
	if nextParentOk { // add in parents
		// add parent
		s.parents.Set(child, nextParent)

		// add as parent
		parentChildren, ok := s.children.Get(nextParent)
		if !ok {
			// mark as parent
			s.parentArray.Set(nextParent, parentComponent{})

			// add children
			parentChildren = datastructures.NewSparseSetWithPaging[ecs.EntityID]()
			s.children.Set(nextParent, parentChildren)
		}
		parentChildren.Add(child)
	}
}

func (s *service) handleParentRemoval(parent ecs.EntityID) {
	if _, isParent := s.parentArray.Get(parent); isParent {
		return
	}

	children := s.GetFlatChildren(parent)

	for _, parent := range s.GetOrderedParents(parent) {
		s.flatChildren.Remove(parent)
	}

	s.children.Remove(parent)
	s.flatChildren.Remove(parent)
	for _, child := range children.GetIndices() {
		s.flatChildren.Remove(child)
		s.children.Remove(child)
		s.parents.Remove(child)
		s.World().RemoveEntity(child)
	}
}
