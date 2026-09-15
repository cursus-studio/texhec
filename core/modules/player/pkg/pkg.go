package playerpkg

import (
	"core/modules/player"
	"core/modules/player/internal"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	uuidpkg "engine/modules/uuid/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		uuidpkg.LinkPkgT[player.OwnerLink],
		typeregistrypkg.PkgT[player.PlayerComponent],
		typeregistrypkg.PkgT[player.ActingPlayerComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) player.Service {
		return internal.NewService(c)
	})
})
