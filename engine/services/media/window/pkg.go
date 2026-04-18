package window

import (
	runtimeservice "engine/services/runtime"

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

	ioc.Wrap(b, func(c ioc.Dic, b runtimeservice.Builder) {
		b.OnCleanUp(func(r runtimeservice.Runtime) {
			api := ioc.Get[Api](c)
			sdl.GLDeleteContext(api.Ctx())
			_ = api.Window().Destroy()
			sdl.Quit()
		})
	})
})
