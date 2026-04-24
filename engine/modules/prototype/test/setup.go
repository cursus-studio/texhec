package test

import (
	"engine"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	enginepkg "engine/pkg"
	"engine/services/ecs"

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
	Cloned1            ecs.ComponentsArray[Cloned1Component]
	Cloned2            ecs.ComponentsArray[Cloned2Component]
	NotCloned          ecs.ComponentsArray[NotClonedComponent]
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		enginepkg.Pkg,
		typeregistrypkg.PkgT[Cloned1Component],
		typeregistrypkg.PkgT[Cloned2Component],
	)

	s := ioc.GetServices[Setup](c)

	s.Cloned1 = ecs.GetComponentsArray[Cloned1Component](s.World())
	s.Cloned2 = ecs.GetComponentsArray[Cloned2Component](s.World())
	s.NotCloned = ecs.GetComponentsArray[NotClonedComponent](s.World())

	return s
}
