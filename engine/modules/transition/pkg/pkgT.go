package transitionpkg

import (
	codecpkg "engine/modules/codec/pkg"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/modules/transition"
	"engine/modules/transition/internal/transitionimpl"

	"github.com/ogiusek/ioc/v2"
)

func PkgT[Component transition.LerpConstraint[Component]](b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[transition.TransitionComponent[Component]],
		codecpkg.PkgT[transition.TransitionEvent[Component]],

		prototypepkg.PkgT[transition.TransitionComponent[Component]],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Wrap(b, func(c ioc.Dic, b transitionimpl.Builder) {
		sys := transitionimpl.NewSysT[Component](c)
		b.Register(sys)
	})
}
