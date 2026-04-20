package internal

import (
	"engine"
	"engine/modules/groups"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`

	inheritArray ecs.ComponentsArray[groups.InheritGroupsComponent]
	groupsArray  ecs.ComponentsArray[groups.GroupsComponent]
}

func NewService(c ioc.Dic) groups.Service {
	t := ioc.GetServices[*service](c)

	t.inheritArray = ecs.GetComponentsArray[groups.InheritGroupsComponent](t.World())
	t.groupsArray = ecs.GetComponentsArray[groups.GroupsComponent](t.World())
	t.Init()
	return t
}

func (t *service) Component() ecs.ComponentsArray[groups.GroupsComponent] {
	return t.groupsArray
}
func (t *service) Inherit() ecs.ComponentsArray[groups.InheritGroupsComponent] {
	return t.inheritArray
}
