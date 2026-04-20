package media

import (
	"engine/services/media/audio"
	"engine/services/media/inputs"
	"engine/services/media/window"

	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type config struct {
	w   *sdl.Window
	ctx sdl.GLContext
}

func NewConfig(
	w *sdl.Window,
	ctx sdl.GLContext,
) config {
	return config{
		w,
		ctx,
	}
}

var Pkg = ioc.NewPkgT(func(b ioc.Builder, config config) {
	for _, pkg := range []ioc.Pkg{
		audio.Pkg,
		inputs.Pkg,
		window.Pkg(window.NewConfig(config.w, config.ctx)),
	} {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) Api {
		return newApi(
			ioc.Get[inputs.Api](c),
			ioc.Get[window.Api](c),
			ioc.Get[audio.Api](c),
		)
	})
})
