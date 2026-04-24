package hierarchypkg

import (
	"engine/modules/hierarchy"
	"engine/modules/hierarchy/internal/hierarchyservice"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[hierarchy.Component],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) hierarchy.Service {
		return hierarchyservice.NewService(c)
	})
})
