package layoutpkg

import (
	"engine/modules/layout"
	"engine/modules/layout/internal/service"
	prototypepkg "engine/modules/prototype/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PkgT[layout.AlignComponent](),
		prototypepkg.PkgT[layout.GapComponent](),
		prototypepkg.PkgT[layout.OrderComponent](),
	} {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) layout.Service {
		return service.NewLayoutService(c)
	})
})
