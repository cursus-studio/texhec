package uipkg

import (
	"core/modules/ui"
	"core/modules/ui/internal/systems"
	"core/modules/ui/internal/uiservice"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/services/ecs"
	"time"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type Config interface {
	AnimationDuration(time.Duration)
	BgFrameDuration(time.Duration)
}

type config struct {
	animationDuration time.Duration
	bgTimePerFrame    time.Duration
}

func newConfig() Config {
	return &config{
		animationDuration: time.Millisecond * 300, // animation duration
		bgTimePerFrame:    time.Second / 12,       // bgTimePerFrame
	}
}
func (c *config) AnimationDuration(d time.Duration) { c.animationDuration = d }
func (c *config) BgFrameDuration(d time.Duration)   { c.bgTimePerFrame = d }

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[ui.AnimatedBackgroundComponent],
		typeregistrypkg.PkgT[ui.CursorCameraComponent],
		typeregistrypkg.PkgT[ui.UiCameraComponent],

		typeregistrypkg.PkgT[ui.HideUiEvent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) Config { return newConfig() })

	ioc.Register(b, func(c ioc.Dic) ui.Service {
		config := ioc.Get[Config](c).(*config)
		return uiservice.NewService(c, config.animationDuration)
	})
	ioc.Register(b, func(c ioc.Dic) ui.System {
		eventsBuilder := ioc.Get[events.Builder](c)
		config := ioc.Get[Config](c).(*config)
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
