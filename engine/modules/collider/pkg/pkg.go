package colliderpkg

import (
	"engine/modules/collider"
	"engine/modules/collider/internal/collisions"
	prototypepkg "engine/modules/prototype/pkg"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct{}

func Package() ioc.Pkg { return pkg{} }

func (pkg) Register(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PackageT[collider.Component](),
	} {
		pkg.Register(b)
	}
	ioc.Register(b, func(c ioc.Dic) collider.Service {
		return collisions.NewService(c, 1000)
	})
}
