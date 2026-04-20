package noisepkg

import (
	"engine/modules/noise"
	"engine/modules/noise/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) noise.Service {
		return internal.NewService(c)
	})
})
