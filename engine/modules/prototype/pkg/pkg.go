package prototypepkg

import (
	"engine/modules/prototype"
	"engine/modules/prototype/internal"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct {
}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg) Register(b ioc.Builder) {
	ioc.RegisterSingleton(b, func(c ioc.Dic) internal.Service {
		return internal.NewService(c)
	})
	ioc.RegisterSingleton(b, func(c ioc.Dic) prototype.Service {
		return ioc.Get[internal.Service](c)
	})
}
