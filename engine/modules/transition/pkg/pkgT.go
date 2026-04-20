package transitionpkg

import (
	prototypepkg "engine/modules/prototype/pkg"
	"engine/modules/transition"
	"engine/modules/transition/internal/transitionimpl"
	"engine/services/codec"

	"github.com/ogiusek/ioc/v2"
)

func PkgT[Component transition.LerpConstraint[Component]]() ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		for _, pkg := range []ioc.Pkg{
			prototypepkg.PkgT[transition.TransitionComponent[Component]](),
		} {
			pkg(b)
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
	})
}
