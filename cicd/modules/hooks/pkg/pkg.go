package hookspkg

import (
	"cicd/modules/hooks"
	"cicd/modules/hooks/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) hooks.Service {
		return internal.NewHooks(c)
	})
})
