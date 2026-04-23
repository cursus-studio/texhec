package test

import (
	"engine/modules/hierarchy"
	enginepkg "engine/pkg"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type Setup struct {
	World   ecs.World
	Service hierarchy.Service
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		enginepkg.Pkg,
	)
	return Setup{
		ioc.Get[ecs.World](c),
		ioc.Get[hierarchy.Service](c),
	}
}
