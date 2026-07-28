package test

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/entityregistry"
	enginepkg "engine/pkg"

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
		enginepkg.Pkg,
		func(b ioc.Builder) {
			ioc.Wrap(b, func(c ioc.Dic, registry entityregistry.Service) {
				world := ioc.Get[ecs.World](c)
				arr := ecs.GetComponentArray[TagValueComponent](world)
				registry.Register("tag", func(entity ecs.EntityID, structTagValue string) {
					arr.Set(entity, TagValueComponent{structTagValue})
				})
			})
		},
	)
	return ioc.GetServices[Setup](c)
}
