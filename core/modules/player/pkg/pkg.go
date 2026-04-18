package playerpkg

import (
	"core/modules/player"
	"core/modules/player/internal"
	prototypepkg "engine/modules/prototype/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PkgT[player.OwnerComponent](),
	} {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) player.Service {
		return internal.NewService(c)
	})
})
