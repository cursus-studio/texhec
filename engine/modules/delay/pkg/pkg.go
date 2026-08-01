package delaypkg

import (
	"engine/modules/delay"
	"engine/modules/delay/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) delay.Service {
		return internal.NewService(c)
	})
})
