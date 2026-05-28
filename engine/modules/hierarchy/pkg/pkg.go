package hierarchypkg

import (
	"engine/modules/hierarchy"
	"engine/modules/hierarchy/internal/hierarchyservice"
	"engine/modules/hierarchy/internal/inheritservice"
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

	ioc.Register(b, hierarchyservice.NewService)
})

func PkgT[Component any]() ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		pkgs := []ioc.Pkg{
			typeregistrypkg.PkgT[hierarchy.InheritComponent[Component]],
		}
		for _, pkg := range pkgs {
			pkg(b)
		}
		ioc.Register(b, inheritservice.NewService[Component])
	})
}
