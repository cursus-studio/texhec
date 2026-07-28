package seedpkg

import (
	"engine/modules/seed"
	"engine/modules/seed/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) seed.Service {
		return internal.NewService(c)
	})
})
