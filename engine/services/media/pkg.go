package media

import (
	"engine/services/media/audio"
	"engine/services/media/inputs"
	"engine/services/media/window"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		audio.Pkg,
		inputs.Pkg,
		window.Pkg,
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
