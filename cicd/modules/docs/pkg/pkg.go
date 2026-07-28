package docspkg

import (
	"cicd/modules/docs"
	"cicd/modules/docs/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) docs.Service {
		return internal.NewService(c)
	})
})
