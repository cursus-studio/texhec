package loggerpkg

import (
	"engine/modules/logger"
	"engine/modules/logger/internal"

	"github.com/ogiusek/ioc/v2"
)

type Config internal.Config

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) internal.Service {
		return internal.NewService()
	})
	ioc.Register(b, func(c ioc.Dic) Config {
		return ioc.Get[internal.Service](c)
	})
	ioc.Register(b, func(c ioc.Dic) logger.Service {
		return ioc.Get[internal.Service](c)
	})
})
