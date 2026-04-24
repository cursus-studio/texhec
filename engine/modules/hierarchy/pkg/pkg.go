package hierarchypkg

import (
	codecpkg "engine/modules/codec/pkg"
	"engine/modules/hierarchy"
	"engine/modules/hierarchy/internal/hierarchyservice"
	prototypepkg "engine/modules/prototype/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[hierarchy.Component],
		prototypepkg.PkgT[hierarchy.Component],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) hierarchy.Service {
		return hierarchyservice.NewService(c)
	})
})
