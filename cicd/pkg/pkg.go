package cicdpkg

import (
	docspkg "cicd/modules/docs/pkg"
	gitpkg "cicd/modules/git/pkg"
	hookspkg "cicd/modules/hooks/pkg"
	"cicd/world"
	enginepkg "engine/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		enginepkg.Pkg,

		gitpkg.Pkg,
		docspkg.Pkg,
		hookspkg.Pkg,
		func(b ioc.Builder) {
			ioc.Register(b, func(c ioc.Dic) world.CICDWorld {
				return ioc.GetServices[world.CICDWorld](c)
			})
		},
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
})
