package internal

import (
	"engine"
	"engine/modules/groups"
	"engine/modules/hierarchy"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	HierarchyT         hierarchy.ServiceT[groups.GroupsComponent] `inject:""`

	groupsArray ecs.ComponentsArray[groups.GroupsComponent]
}

func NewService(c ioc.Dic) groups.Service {
	t := ioc.GetServices[*service](c)

	t.groupsArray = ecs.GetComponentsArray[groups.GroupsComponent](t.World())
	t.Init()
	return t
}

func (s *service) Component() ecs.ComponentsArray[groups.GroupsComponent] {
	return s.groupsArray
}
func (s *service) Inherit() ecs.ComponentsArray[hierarchy.InheritComponent[groups.GroupsComponent]] {
	return s.HierarchyT.Inherit()
}
func (s *service) InheritGroups(entity ecs.EntityID) {
	s.HierarchyT.Inherit().Set(entity, hierarchy.InheritComponent[groups.GroupsComponent]{})
}
