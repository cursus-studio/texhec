package test

import (
	"engine/modules/hierarchy"
	"engine/modules/hierarchy/pkg"
	"engine/services/clock"
	"engine/services/ecs"
	"engine/services/logger"
	"time"

	"github.com/ogiusek/ioc/v2"
)

type Setup struct {
	World   ecs.World
	Service hierarchy.Service
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		logger.Pkg(logger.NewConfig(true, func(c ioc.Dic, message string) { print(message) })),
		clock.Pkg(time.RFC3339Nano),
		ecs.Pkg,
		hierarchypkg.Pkg,
	)
	return Setup{
		ioc.Get[ecs.World](c),
		ioc.Get[hierarchy.Service](c),
	}
}
