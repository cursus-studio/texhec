package colliderpkg

import (
	"engine/modules/collider"
	"engine/modules/collider/internal/collisions"
	prototypepkg "engine/modules/prototype/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		prototypepkg.PkgT[collider.Component],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) collider.Service {
		return collisions.NewService(c, 1000)
	})
})
