package fpsloggerpkg

import (
	"core/modules/fpslogger"
	"core/modules/fpslogger/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) fpslogger.System {
		return internal.NewFpsLoggerSystem(c)
	})
})
