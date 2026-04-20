package prototypepkg

import (
	"engine/modules/prototype"
	"engine/modules/prototype/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) internal.Service {
		return internal.NewService(c)
	})
	ioc.Register(b, func(c ioc.Dic) prototype.Service {
		return ioc.Get[internal.Service](c)
	})
})
