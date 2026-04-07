package warmuppkg

import (
	"engine/modules/warmup"
	"engine/modules/warmup/internal"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type pkg struct{}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg) Register(b ioc.Builder) {
	ioc.RegisterSingleton(b, func(c ioc.Dic) warmup.Service {
		return internal.NewService(c)
	})
	ioc.WrapService(b, func(c ioc.Dic, b events.Builder) {
		world := ioc.Get[ecs.World](c)
		events.Listen(b, func(warmup.Event) {
			world.WarmUp()
		})
	})
}
