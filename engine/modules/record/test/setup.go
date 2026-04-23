package test

import (
	"engine/modules/record"
	"engine/modules/uuid"
	enginepkg "engine/pkg"
	"engine/services/codec"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type Component struct {
	Counter int
}

type Setup struct {
	Config record.Config
	Codec  codec.Codec

	World          ecs.World
	UUID           uuid.Service
	Record         record.Service
	ComponentArray ecs.ComponentsArray[Component]
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		enginepkg.Pkg,
		func(b ioc.Builder) {
			ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
				b.
					Register(Component{})
			})
		},
	)

	s := Setup{
		Codec:  ioc.Get[codec.Codec](c),
		Config: record.NewConfig(),

		World:  ioc.Get[ecs.World](c),
		UUID:   ioc.Get[uuid.Service](c),
		Record: ioc.Get[record.Service](c),
	}

	s.ComponentArray = ecs.GetComponentsArray[Component](s.World)

	record.AddToConfig[Component](s.Config)

	return s
}
