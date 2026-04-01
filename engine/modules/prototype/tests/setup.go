package tests

import (
	"engine/modules/prototype"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/services/clock"
	"engine/services/ecs"
	"engine/services/logger"
	"time"

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
	World     ecs.World         `inject:"1"`
	Prototype prototype.Service `inject:"1"`
	Cloned1   ecs.ComponentsArray[Cloned1Component]
	Cloned2   ecs.ComponentsArray[Cloned2Component]
	NotCloned ecs.ComponentsArray[NotClonedComponent]
}

func NewSetup() Setup {
	b := ioc.NewBuilder()

	for _, pkg := range []ioc.Pkg{
		logger.Package(true, func(c ioc.Dic, message string) { print(message) }),
		clock.Package(time.RFC3339Nano),
		ecs.Package(),
		prototypepkg.Package(),
		prototypepkg.PackageT[Cloned1Component](),
		prototypepkg.PackageT[Cloned2Component](),
	} {
		pkg.Register(b)
	}

	c := b.Build()

	s := ioc.GetServices[Setup](c)

	s.Cloned1 = ecs.GetComponentsArray[Cloned1Component](s.World)
	s.Cloned2 = ecs.GetComponentsArray[Cloned2Component](s.World)
	s.NotCloned = ecs.GetComponentsArray[NotClonedComponent](s.World)

	return s
}
