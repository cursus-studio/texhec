package mock

import (
	corepkg "core/pkg"
	netsyncpkg "engine/modules/netsync/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		corepkg.Pkg,
		netsyncpkg.Pkg(netsyncpkg.NewConfig(0)),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
})
