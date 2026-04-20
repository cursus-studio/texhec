package uipkg

import (
	"core/modules/ui"
	"core/modules/ui/internal/systems"
	"core/modules/ui/internal/uiservice"
	"engine/modules/prototype/pkg"
	"engine/services/codec"
	"engine/services/ecs"
	"time"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type config struct {
	animationDuration time.Duration
	bgTimePerFrame    time.Duration
}

func NewConfig(
	animationDuration time.Duration,
	bgTimePerFrame time.Duration,
) config {
	return config{
		animationDuration,
		bgTimePerFrame,
	}
}

var Pkg = ioc.NewPkgT(func(b ioc.Builder, config config) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PkgT[ui.AnimatedBackgroundComponent](),
		prototypepkg.PkgT[ui.CursorCameraComponent](),
		prototypepkg.PkgT[ui.UiCameraComponent](),
	} {
		pkg(b)
	}
	ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// components
			Register(ui.AnimatedBackgroundComponent{}).
			Register(ui.CursorCameraComponent{}).
			Register(ui.UiCameraComponent{}).
			// events
			Register(ui.HideUiEvent{})
	})

	ioc.Register(b, func(c ioc.Dic) ui.Service {
		return uiservice.NewService(c, config.animationDuration, config.bgTimePerFrame)
	})
	ioc.Register(b, func(c ioc.Dic) ui.System {
		eventsBuilder := ioc.Get[events.Builder](c)
		return ecs.NewSystemRegister(func() error {
			errs := ecs.RegisterSystems(
				systems.NewSystem(c, config.bgTimePerFrame),
				systems.NewCursorSystem(c),
			)
			if len(errs) != 0 {
				return errs[0]
			}

			events.Listen(eventsBuilder, func(e sdl.MouseButtonEvent) {
				if e.Button != sdl.BUTTON_RIGHT || e.State != sdl.RELEASED {
					return
				}
				events.Emit(eventsBuilder.Events(), ui.HideUiEvent{})
			})
			ioc.Get[ui.Service](c)
			return nil
		})
	})
})
