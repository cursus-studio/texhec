package deploypkg

import (
	"core/modules/deploy"
	"core/modules/deploy/internal"
	prototypepkg "engine/modules/prototype/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PkgT[deploy.LinkComponent](),
	} {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) deploy.Service {
		return internal.NewService(c)
	})
})
