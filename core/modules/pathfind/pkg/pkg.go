package pathfindpkg

import (
	"core/modules/pathfind"
	"core/modules/pathfind/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) pathfind.Service {
		return internal.NewService(c)
	})
})
