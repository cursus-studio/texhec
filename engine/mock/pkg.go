package mock

import (
	netsyncpkg "engine/modules/netsync/pkg"
	enginepkg "engine/pkg"
	"engine/services/logger"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		enginepkg.Pkg,
		netsyncpkg.Pkg(netsyncpkg.NewConfig(0)),

		logger.Pkg(logger.NewConfig(
			true,
			func(c ioc.Dic) func(message string) { return func(message string) { print(message) } },
		)),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
})
