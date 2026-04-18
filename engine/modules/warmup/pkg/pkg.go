package warmuppkg

import (
	"engine/modules/warmup"
	"engine/modules/warmup/internal"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) warmup.Service {
		return internal.NewService(c)
	})
	ioc.Wrap(b, func(c ioc.Dic, b events.Builder) {
		world := ioc.Get[ecs.World](c)
		events.Listen(b, func(warmup.Event) {
			world.WarmUp()
		})
	})
})
