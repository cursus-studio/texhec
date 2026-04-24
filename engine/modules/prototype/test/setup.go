package test

import (
	"engine"
	prototypepkg "engine/modules/prototype/pkg"
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
		prototypepkg.PkgT[Cloned1Component],
		prototypepkg.PkgT[Cloned2Component],
	)

	s := ioc.GetServices[Setup](c)

	s.Cloned1 = ecs.GetComponentsArray[Cloned1Component](s.World())
	s.Cloned2 = ecs.GetComponentsArray[Cloned2Component](s.World())
	s.NotCloned = ecs.GetComponentsArray[NotClonedComponent](s.World())

	return s
}
