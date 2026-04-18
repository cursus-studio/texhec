package registrypkg

import (
	"engine/modules/registry"
	"engine/modules/registry/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) registry.Service {
		return internal.NewService(c)
	})
})
