package window

import (
	"engine/modules/loop"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type config struct {
	window  *sdl.Window
	context sdl.GLContext
}

func NewConfig(
	window *sdl.Window,
	context sdl.GLContext,
) config {
	return config{
		window:  window,
		context: context,
	}
}

var Pkg = ioc.NewPkgT(func(b ioc.Builder, config config) {
	ioc.Register(b, func(c ioc.Dic) Api {
		return newApi(
			config.window,
			config.context,
		)
	})

	ioc.Wrap(b, func(c ioc.Dic, b events.Builder) {
		events.Listen(b, func(loop.StopEvent) {
			api := ioc.Get[Api](c)
			sdl.GLDeleteContext(api.Ctx())
			_ = api.Window().Destroy()
			sdl.Quit()
		})
	})
})
