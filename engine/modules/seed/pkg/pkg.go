package seedpkg

import (
	netsyncpkg "engine/modules/netsync/pkg"
	"engine/modules/seed"
	"engine/modules/seed/internal"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[seed.SeedComponent],

		netsyncpkg.ComponentPkg[seed.SeedComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) seed.Service {
		return internal.NewService(c)
	})
})
