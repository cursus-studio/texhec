package dragpkg

import (
	"engine/modules/drag"
	"engine/modules/drag/internal"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[drag.DraggableEvent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) drag.System {
		return internal.NewSystem(c)
	})
})
