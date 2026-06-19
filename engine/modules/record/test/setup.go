package test

import (
	"engine/modules/codec"
	"engine/modules/ecs"
	"engine/modules/record"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/modules/uuid"
	enginepkg "engine/pkg"

	"github.com/ogiusek/ioc/v2"
)

type Component struct {
	Counter int
}

type Setup struct {
	Config record.Config
	Codec  codec.Service

	World          ecs.World
	UUID           uuid.Service
	Record         record.Service
	ComponentArray ecs.ComponentArray[Component]
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		enginepkg.Pkg,
		typeregistrypkg.PkgT[Component],
	)

	s := Setup{
		Codec:  ioc.Get[codec.Service](c),
		Config: record.NewConfig(),

		World:  ioc.Get[ecs.World](c),
		UUID:   ioc.Get[uuid.Service](c),
		Record: ioc.Get[record.Service](c),
	}

	s.ComponentArray = ecs.GetComponentArray[Component](s.World)

	record.AddToConfig[Component](s.Config)

	return s
}
