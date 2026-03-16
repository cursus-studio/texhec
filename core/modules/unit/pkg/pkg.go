package unitpkg

import (
	"core/modules/unit"
	"core/modules/unit/internal"
	transitionpkg "engine/modules/transition/pkg"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct {
	layer float32
}

func Package(layer float32) ioc.Pkg {
	return pkg{layer}
}

func (pkg pkg) Register(b ioc.Builder) {
	ioc.RegisterSingleton(b, func(c ioc.Dic) unit.Service {
		return internal.NewService(c)
	})
	ioc.RegisterSingleton(b, func(c ioc.Dic) unit.System {
		return ecs.NewSystemRegister(func() error {
			return internal.NewSystem(c, pkg.layer)
		})
	})

	for _, pkg := range []ioc.Pkg{
		transitionpkg.PackageT[unit.RotationComponent](),
		transitionpkg.PackageT[unit.CoordsComponent](),
	} {
		pkg.Register(b)
	}
}
