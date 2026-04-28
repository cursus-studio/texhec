package colliderpkg

import (
	"engine/modules/collider"
	"engine/modules/collider/internal/collisions"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[collider.Component],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) collider.Service {
		return collisions.NewService(c, 1000)
	})
})
