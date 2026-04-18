package colliderpkg

import (
	"engine/modules/collider"
	"engine/modules/collider/internal/collisions"
	prototypepkg "engine/modules/prototype/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PkgT[collider.Component](),
	} {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) collider.Service {
		return collisions.NewService(c, 1000)
	})
})
