package test

import (
	"engine/mock"
	"engine/modules/groups"
	"engine/modules/hierarchy"
	"engine/services/ecs"
	"testing"

	"github.com/ogiusek/ioc/v2"
)

type Setup struct {
	world     ecs.World
	hierarchy hierarchy.Service
	groups    groups.Service
	T         *testing.T
}

func NewSetup(t *testing.T) Setup {
	c := ioc.NewContainer(
		mock.Pkg,
	)

	return Setup{
		ioc.Get[ecs.World](c),
		ioc.Get[hierarchy.Service](c),
		ioc.Get[groups.Service](c),
		t,
	}
}

func (setup *Setup) expectGroups(entity ecs.EntityID, expectedGroups groups.GroupsComponent) {
	groups, _ := setup.groups.Component().Get(entity)
	if groups != expectedGroups {
		setup.T.Errorf("expected pos %v but has %v", expectedGroups, groups)
	}
}
