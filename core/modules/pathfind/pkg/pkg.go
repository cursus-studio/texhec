package pathfindpkg

import (
	"core/modules/pathfind"
	"core/modules/pathfind/internal"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct {
}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg) Register(b ioc.Builder) {
	ioc.RegisterSingleton(b, func(c ioc.Dic) pathfind.Service {
		return internal.NewService(c)
	})
}
