package hierarchypkg

import (
	"engine/modules/hierarchy"
	"engine/modules/hierarchy/internal/hierarchyservice"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/services/codec"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct{}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg pkg) Register(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PackageT[hierarchy.Component](),
	} {
		pkg.Register(b)
	}
	ioc.WrapService(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// components
			Register(hierarchy.Component{})
	})

	ioc.RegisterSingleton(b, func(c ioc.Dic) hierarchy.Service {
		return hierarchyservice.NewService(c)
	})
}
