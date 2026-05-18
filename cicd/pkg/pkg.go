package cicdpkg

import (
	docspkg "cicd/modules/docs/pkg"
	gitpkg "cicd/modules/git/pkg"
	pipepkg "cicd/modules/pipe/pkg"
	projectfspkg "cicd/modules/projectfs/pkg"
	"cicd/world"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		docspkg.Pkg,
		gitpkg.Pkg,
		projectfspkg.Pkg,
		pipepkg.Pkg,

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
