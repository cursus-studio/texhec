package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/groups"
	"engine/modules/hierarchy"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	HierarchyT         hierarchy.ServiceT[groups.GroupsComponent] `inject:""`

	groupsArray ecs.ComponentArray[groups.GroupsComponent]
}

func NewService(c ioc.Dic) groups.Service {
	s := ioc.GetServices[*service](c)

	s.groupsArray = ecs.GetComponentArray[groups.GroupsComponent](s.World())
	s.Init()
	return s
}

func (s *service) Component() ecs.ComponentArray[groups.GroupsComponent] {
	return s.groupsArray
}
func (s *service) Inherit() ecs.ComponentArray[hierarchy.InheritComponent[groups.GroupsComponent]] {
	return s.HierarchyT.Inherit()
}
func (s *service) InheritGroups(entity ecs.EntityID) {
	s.HierarchyT.Inherit().Set(entity, hierarchy.InheritComponent[groups.GroupsComponent]{})
}
