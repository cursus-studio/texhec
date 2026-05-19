package pipepkg

import (
	"cicd/modules/pipe"
	"cicd/modules/pipe/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) pipe.Service {
		return internal.NewService(c)
	})
})
