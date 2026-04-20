package smoothpkg

import (
	"engine"
	"engine/modules/smooth"
	"engine/modules/smooth/internal"
	"engine/modules/transition"
	"engine/services/ecs"
	"engine/services/frames"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

func PkgT[Component transition.LerpConstraint[Component]]() ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		ioc.Register(b, func(c ioc.Dic) *internal.Service[Component] {
			return internal.NewService[Component](c)
		})
		ioc.Wrap(b, func(c ioc.Dic, _ smooth.Service) {
			internal.NewSystems[Component](c)
		})
	})
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) smooth.Service {
		return struct{}{}
	})

	ioc.Register(b, func(c ioc.Dic) smooth.StartSystem {
		return ecs.NewSystemRegister(func() error {
			s := ioc.GetServices[engine.EngineWorld](c)
			events.Listen(s.EventsBuilder(), func(tick frames.TickEvent) {
				events.Emit(s.Events(), internal.FirstEvent(tick))
			})
			return nil
		})
	})

	ioc.Register(b, func(c ioc.Dic) smooth.StopSystem {
		return ecs.NewSystemRegister(func() error {
			s := ioc.GetServices[engine.EngineWorld](c)
			events.Listen(s.EventsBuilder(), func(tick frames.TickEvent) {
				events.Emit(s.Events(), internal.LastEvent(tick))
			})
			return nil
		})
	})
})
