package transitionpkg

import (
	prototypepkg "engine/modules/prototype/pkg"
	"engine/modules/transition"
	"engine/modules/transition/internal/transitionimpl"
	"engine/services/codec"

	"github.com/ogiusek/ioc/v2"
)

type pkgT[Component transition.LerpConstraint[Component]] struct {
}

func PackageT[Component transition.LerpConstraint[Component]]() ioc.Pkg {
	return pkgT[Component]{}
}

func (pkgT[Component]) Register(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PackageT[transition.TransitionComponent[Component]](),
	} {
		pkg.Register(b)
	}
	ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// components
			Register(transition.TransitionComponent[Component]{}).
			// events
			Register(transition.TransitionEvent[Component]{})
	})
	ioc.Wrap(b, func(c ioc.Dic, b transitionimpl.Builder) {
		sys := transitionimpl.NewSysT[Component](c)
		b.Register(sys)
	})
}
