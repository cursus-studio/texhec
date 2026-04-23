package window

import (
	"engine/modules/loop"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) Api {
		return newApi(c)
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
