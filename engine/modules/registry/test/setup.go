package test

import (
	"engine/modules/registry"
	registrypkg "engine/modules/registry/pkg"
	uuidpkg "engine/modules/uuid/pkg"
	"engine/services/clock"
	"engine/services/ecs"
	"engine/services/logger"
	"time"

	"github.com/ogiusek/ioc/v2"
)

type Setup struct {
	World   ecs.World        `inject:"1"`
	Service registry.Service `inject:"1"`
}

type TagValueComponent struct {
	Value string
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		clock.Pkg(time.RFC3339Nano),
		logger.Pkg(logger.NewConfig(true, func(c ioc.Dic, message string) { print(message) })),
		ecs.Pkg,
		uuidpkg.Pkg,
		registrypkg.Pkg,
		func(b ioc.Builder) {
			ioc.Wrap(b, func(c ioc.Dic, registry registry.Service) {
				world := ioc.Get[ecs.World](c)
				registry.Register("tag", func(entity ecs.EntityID, structTagValue string) {
					ecs.SaveComponent(world, entity, TagValueComponent{structTagValue})
				})
			})
		},
	)
	return ioc.GetServices[Setup](c)
}
