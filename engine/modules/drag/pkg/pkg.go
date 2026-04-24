package dragpkg

import (
	codecpkg "engine/modules/codec/pkg"
	"engine/modules/drag"
	"engine/modules/drag/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[drag.DraggableEvent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) drag.System {
		return internal.NewSystem(c)
	})
})
