package unitpkg

import (
	"core/modules/unit"
	"core/modules/unit/internal"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct {
}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg pkg) Register(b ioc.Builder) {
	ioc.RegisterSingleton(b, func(c ioc.Dic) unit.Service {
		return internal.NewService(c)
	})
	ioc.RegisterSingleton(b, func(c ioc.Dic) unit.System {
		return ecs.NewSystemRegister(func() error {
			return internal.NewSystem(c)
		})
	})
}
