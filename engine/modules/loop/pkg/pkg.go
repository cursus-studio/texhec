package looppkg

import (
	"engine/modules/loop"
	"engine/modules/loop/internal"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[loop.EmitOnTickComponent],
		typeregistrypkg.PkgT[loop.EmitOnTickEvent],

		typeregistrypkg.PkgT[loop.FrameEvent],
		typeregistrypkg.PkgT[loop.TickEvent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) loop.Service {
		return internal.NewService(c)
	})
})
