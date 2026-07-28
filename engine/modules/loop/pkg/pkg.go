package looppkg

import (
	"engine/modules/loop"
	"engine/modules/loop/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) loop.Service {
		return internal.NewService(c)
	})
})
