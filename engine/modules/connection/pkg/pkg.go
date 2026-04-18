package connectionpkg

import (
	"engine/modules/connection"
	"engine/modules/connection/internal"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) connection.System {
		return ecs.NewSystemRegister(func() error {
			ioc.Get[connection.Service](c)
			return nil
		})
	})
	ioc.Register(b, func(c ioc.Dic) connection.Service {
		return internal.NewService(c)
	})
})
