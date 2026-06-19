package ecspkg

import (
	"engine/modules/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) events.Builder {
		return events.NewBuilder()
	})
	ioc.Register(b, func(c ioc.Dic) events.Events {
		return ioc.Get[events.Builder](c).Build()
	})
	ioc.Wrap(b, func(c ioc.Dic, b events.Builder) {
		world := ioc.Get[ioc.Lazy[ecs.World]](c)
		events.Listen(b, func(e ecs.RemoveEntityEvent) {
			world().RemoveEntity(e.Entity)
		})
	})
	ioc.Register(b, func(c ioc.Dic) ecs.World {
		return ecs.NewWorld()
	})
})
