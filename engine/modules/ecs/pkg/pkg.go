package ecspkg

import (
	"engine/modules/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) ecs.World {
		return ecs.NewWorld()
	})

	ioc.Wrap(b, func(c ioc.Dic, b events.Builder) {
		world := ioc.Get[ioc.Lazy[ecs.World]](c)
		events.Listen(b, func(event ecs.SetEvent) {
			if arr, ok := world().GetArrByComp(event.Component); ok {
				arr.SetAny(event.Entity, event.Component)
			}
		})
	})
})
