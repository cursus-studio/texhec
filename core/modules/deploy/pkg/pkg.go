package deploypkg

import (
	"core/modules/deploy"
	"core/modules/deploy/internal"
	prototypepkg "engine/modules/prototype/pkg"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct{}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg) Register(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PackageT[deploy.LinkComponent](),
	} {
		pkg.Register(b)
	}
	ioc.RegisterSingleton(b, func(c ioc.Dic) deploy.Service {
		return internal.NewService(c)
	})
}
