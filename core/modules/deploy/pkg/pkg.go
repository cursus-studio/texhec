package deploypkg

import (
	"core/modules/deploy"
	"core/modules/deploy/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) deploy.Service {
		return internal.NewService(c)
	})
})
