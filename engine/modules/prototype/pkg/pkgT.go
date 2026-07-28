package prototypepkg

import (
	"engine/modules/ecs"
	"engine/modules/prototype/internal"

	"github.com/ogiusek/ioc/v2"
)

func PkgT[Component any](b ioc.Builder) {
	ioc.Wrap(b, func(c ioc.Dic, b internal.Service) {
		b.Add(ecs.GetComponentArray[Component](ioc.Get[ecs.World](c)))
	})
}
