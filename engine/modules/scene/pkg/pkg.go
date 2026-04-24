package scenepkg

import (
	codecpkg "engine/modules/codec/pkg"
	"engine/modules/scene"
	"engine/modules/scene/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[scene.ChangeSceneEvent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) scene.Service {
		return internal.NewService(c)
	})
})
