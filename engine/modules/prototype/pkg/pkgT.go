package prototypepkg

import (
	"engine/modules/prototype/internal"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type clonerPkg[Component any] struct{}

func PackageT[Component any]() ioc.Pkg {
	return clonerPkg[Component]{}
}

func (clonerPkg[Component]) Register(b ioc.Builder) {
	ioc.WrapService(b, func(c ioc.Dic, b internal.Service) {
		b.Add(ecs.GetComponentsArray[Component](ioc.Get[ecs.World](c)))
	})
}
