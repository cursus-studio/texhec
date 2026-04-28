package smoothpkg

import (
	"engine"
	"engine/modules/loop"
	"engine/modules/smooth"
	"engine/modules/smooth/internal"
	"engine/modules/transition"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) smooth.Service {
		return struct{}{}
	})

	ioc.Register(b, func(c ioc.Dic) smooth.StartSystem {
		return ecs.NewSystemRegister(func() error {
			s := ioc.GetServices[engine.EngineWorld](c)
			events.Listen(s.EventsBuilder(), func(tick loop.TickEvent) {
				events.Emit(s.Events(), internal.FirstEvent(tick))
			})
			return nil
		})
	})

	ioc.Register(b, func(c ioc.Dic) smooth.StopSystem {
		return ecs.NewSystemRegister(func() error {
			s := ioc.GetServices[engine.EngineWorld](c)
			events.Listen(s.EventsBuilder(), func(tick loop.TickEvent) {
				events.Emit(s.Events(), internal.LastEvent(tick))
			})
			return nil
		})
	})
})

// func PkgT[Component transition.LerpConstraint[Component]](b ioc.Builder) {
func PkgT[Component any](b ioc.Builder) {
	var zero Component
	_ = any(zero).(transition.LerpConstraint[Component])
	ioc.Register(b, func(c ioc.Dic) *internal.Service[Component] {
		return internal.NewService[Component](c)
	})
	ioc.Wrap(b, func(c ioc.Dic, _ smooth.Service) {
		internal.NewSystems[Component](c)
	})
}
