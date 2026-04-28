package layoutpkg

import (
	"engine/modules/layout"
	"engine/modules/layout/internal/service"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[layout.AlignComponent],
		typeregistrypkg.PkgT[layout.GapComponent],
		typeregistrypkg.PkgT[layout.OrderComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) layout.Service {
		return service.NewLayoutService(c)
	})
})
