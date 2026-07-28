package test

import (
	"engine"
	"engine/modules/ecs"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	enginepkg "engine/pkg"

	"github.com/ogiusek/ioc/v2"
)

type Cloned1Component struct {
	Parametr int
}
type Cloned2Component struct {
	Parametr int
}

type NotClonedComponent struct {
	Parametr int
}

type Setup struct {
	engine.EngineWorld `inject:""`
	Cloned1            ecs.ComponentArray[Cloned1Component]
	Cloned2            ecs.ComponentArray[Cloned2Component]
	NotCloned          ecs.ComponentArray[NotClonedComponent]
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		enginepkg.Pkg,
		typeregistrypkg.PkgT[Cloned1Component],
		typeregistrypkg.PkgT[Cloned2Component],
	)

	s := ioc.GetServices[Setup](c)

	s.Cloned1 = ecs.GetComponentArray[Cloned1Component](s.World())
	s.Cloned2 = ecs.GetComponentArray[Cloned2Component](s.World())
	s.NotCloned = ecs.GetComponentArray[NotClonedComponent](s.World())

	return s
}
