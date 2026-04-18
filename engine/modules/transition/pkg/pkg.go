package transitionpkg

import (
	prototypepkg "engine/modules/prototype/pkg"
	"engine/modules/transition"
	"engine/modules/transition/internal/service"
	"engine/modules/transition/internal/system"
	"engine/modules/transition/internal/transitionimpl"
	"engine/services/codec"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PkgT[transition.EasingComponent](),
		prototypepkg.PkgT[transition.EasingFunctionComponent](),
	} {
		pkg(b)
	}
	ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// components
			Register(transition.EasingComponent{}).
			Register(transition.Progress(0))
	})
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
