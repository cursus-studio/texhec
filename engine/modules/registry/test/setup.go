package test

import (
	"engine"
	"engine/mock"
	"engine/modules/registry"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type Setup struct {
	engine.EngineWorld `inject:""`
}

type TagValueComponent struct {
	Value string
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		mock.Pkg,
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
