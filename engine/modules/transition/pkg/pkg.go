package transitionpkg

import (
	codecpkg "engine/modules/codec/pkg"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/modules/transition"
	"engine/modules/transition/internal/service"
	"engine/modules/transition/internal/system"
	"engine/modules/transition/internal/transitionimpl"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[transition.EasingComponent],
		codecpkg.PkgT[transition.Progress],

		prototypepkg.PkgT[transition.EasingComponent],
		prototypepkg.PkgT[transition.EasingFunctionComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) transitionimpl.Builder {
		b := transitionimpl.NewBuilder()
		b.Register(system.NewSystem(c)) // delayedEvent system
		return b
	})

	ioc.Register(b, func(c ioc.Dic) transition.System {
		return ioc.Get[transitionimpl.Builder](c).Build()
	})
	ioc.Register(b, func(c ioc.Dic) transition.Service {
		return service.NewService(c)
	})
})
