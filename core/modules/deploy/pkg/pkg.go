package deploypkg

import (
	"core/modules/deploy"
	"core/modules/deploy/internal"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct{}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg) Register(b ioc.Builder) {
	ioc.RegisterSingleton(b, func(c ioc.Dic) deploy.Service {
		return internal.NewService(c)
	})
}
