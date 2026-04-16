package loadingpkg

import (
	"core/modules/loading"
	"core/modules/loading/internal"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct{}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg) Register(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) loading.System {
		return internal.NewSystem(c)
	})
}
