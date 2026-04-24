package prototypepkg

import (
	"engine/modules/prototype/internal"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

func PkgT[Component any](b ioc.Builder) {
	ioc.Wrap(b, func(c ioc.Dic, b internal.Service) {
		b.Add(ecs.GetComponentsArray[Component](ioc.Get[ecs.World](c)))
	})
}
