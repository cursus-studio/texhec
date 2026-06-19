package inheritservice

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/hierarchy"

	"github.com/ogiusek/ioc/v2"
)

type service[Component any] struct {
	engine.EngineWorld `inject:""`
	component          ecs.ComponentArray[Component]
	inherit            ecs.ComponentArray[hierarchy.InheritComponent[Component]]

	equal func(Component, Component) bool
}

func NewService[Component any](c ioc.Dic) hierarchy.ServiceT[Component] {
	s := ioc.GetServices[*service[Component]](c)
	s.component = ecs.GetComponentArray[Component](s.World())
	s.inherit = ecs.GetComponentArray[hierarchy.InheritComponent[Component]](s.World())

	s.equal = ecs.ComponentComparator[Component]()
	s.Init()
	return s
}

func (s *service[Component]) Inherit() ecs.ComponentArray[hierarchy.InheritComponent[Component]] {
	return s.inherit
}

//

func (s *service[Component]) calculateGroup(entity ecs.EntityID) (Component, bool) {
	def := s.component.GetEmpty()
	parent, ok := s.Hierarchy().Parent(entity)
	if !ok {
		return def, false
	}
	groups, ok := s.component.Get(parent)
	if !ok {
		return def, false
	}
	return groups, ok
}

type save[Component any] struct {
	entity ecs.EntityID
	comp   Component
}

func (s *service[Component]) Init() {
	dirtySet := ecs.NewDirtySet()
	s.component.AddDependency(s.inherit)
	s.component.AddDependency(s.Hierarchy().Component())

	s.component.AddDirtySet(dirtySet)

	s.component.BeforeGet(func() {
		entities := dirtySet.Get()
		if len(entities) == 0 {
			return
		}
		children := []ecs.EntityID{}

		saves := []save[Component]{}

		for len(entities) != 0 || len(children) != 0 {
			if len(entities) == 0 {
				entities = children
				for _, save := range saves {
					s.component.Set(save.entity, save.comp)
				}

				dirtySet.Clear()
				children = nil
				saves = nil
			}
			entity := entities[0]
			entities = entities[1:]

			comp, ok := s.calculateGroup(entity)
			if !ok {
				continue
			}
			if originalComp, ok := s.component.Get(entity); ok && s.equal(comp, originalComp) {
				continue
			}
			saves = append(saves, save[Component]{
				entity: entity,
				comp:   comp,
			})

			for _, child := range s.Hierarchy().Children(entity).GetIndices() {
				if _, ok := s.inherit.Get(child); ok {
					children = append(children, child)
				}
			}
		}

		for _, save := range saves {
			s.component.Set(save.entity, save.comp)
		}
		dirtySet.Clear()
	})
}
