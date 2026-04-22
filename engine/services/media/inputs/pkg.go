package inputs

import (
	"engine/services/clock"
	"engine/services/frames"
	"engine/services/logger"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) Api {
		return newInputsApi(
			ioc.Get[logger.Logger](c),
			ioc.Get[clock.Clock](c),
			ioc.Get[events.Events](c),
		)
	})

	ioc.Wrap(b, func(c ioc.Dic, b events.Builder) {
		events.Listen(b, func(qe sdl.QuitEvent) {
			ioc.Get[frames.Frames](c).Stop()
		})
	})
})
