package batcherpkg

import (
	"engine/modules/batcher"
	"engine/modules/batcher/internal"
	"runtime"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) *internal.Service {
		return internal.NewService(
			c,
			max(1, runtime.NumCPU()-1),
		)
	})

	ioc.Register(b, func(c ioc.Dic) batcher.Service {
		return ioc.Get[*internal.Service](c)
	})

	ioc.Register(b, func(c ioc.Dic) batcher.System {
		return ioc.Get[*internal.Service](c).System()
	})
})
