package graphicspkg

import (
	"engine/modules/graphics"
	"engine/modules/graphics/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) graphics.Service {
		return internal.NewService()
	})
})
