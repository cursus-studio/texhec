package mock

import (
	netsyncpkg "engine/modules/netsync/pkg"
	enginepkg "engine/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		enginepkg.Pkg,
		netsyncpkg.Pkg(netsyncpkg.NewConfig(0)),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
})
