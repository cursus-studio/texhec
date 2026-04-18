package test

import (
	"engine/modules/groups"
	groupspkg "engine/modules/groups/pkg"
	"engine/modules/hierarchy"
	hierarchypkg "engine/modules/hierarchy/pkg"
	"engine/services/clock"
	"engine/services/ecs"
	"engine/services/logger"
	"testing"
	"time"

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
		logger.Pkg(logger.NewConfig(true, func(c ioc.Dic, message string) { print(message) })),
		clock.Pkg(time.RFC3339Nano),
		ecs.Pkg,
		hierarchypkg.Pkg,
		groupspkg.Pkg,
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
