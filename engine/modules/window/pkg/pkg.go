package windowpkg

import (
	"engine/modules/window"
	"engine/modules/window/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) window.Service {
		s := internal.NewService(c)
		return s
	})
})
