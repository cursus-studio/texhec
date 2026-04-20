package loadingpkg

import (
	"core/modules/loading"
	"core/modules/loading/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) loading.System {
		return internal.NewSystem(c)
	})
})
