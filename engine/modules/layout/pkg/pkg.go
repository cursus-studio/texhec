package layoutpkg

import (
	"engine/modules/layout"
	"engine/modules/layout/internal/service"
	prototypepkg "engine/modules/prototype/pkg"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct{}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg) Register(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PackageT[layout.AlignComponent](),
		prototypepkg.PackageT[layout.GapComponent](),
		prototypepkg.PackageT[layout.OrderComponent](),
	} {
		pkg.Register(b)
	}
	ioc.RegisterSingleton(b, func(c ioc.Dic) layout.Service {
		return service.NewLayoutService(c)
	})
}
