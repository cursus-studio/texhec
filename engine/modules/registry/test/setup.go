package test

import (
	"engine/modules/registry"
	registrypkg "engine/modules/registry/pkg"
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
	b := ioc.NewBuilder()
	pkgs := []ioc.Pkg{
		clock.Package(time.RFC3339Nano),
		logger.Package(true, func(c ioc.Dic, message string) { print(message) }),
		ecs.Package(),
		registrypkg.Package(),
	}
	for _, pkg := range pkgs {
		pkg.Register(b)
	}
	ioc.WrapService(b, func(c ioc.Dic, registry registry.Service) {
		world := ioc.Get[ecs.World](c)
		registry.Register("tag", func(structTagValue string) ecs.EntityID {
			entity := world.NewEntity()
			ecs.SaveComponent(world, entity, TagValueComponent{structTagValue})
			return entity
		})
	})
	c := b.Build()
	return ioc.GetServices[Setup](c)
}
