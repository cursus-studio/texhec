package unitpkg

import (
	"core/modules/unit"
	"core/modules/unit/internal"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct {
	layer float32
}

func Package(layer float32) ioc.Pkg {
	return pkg{layer}
}

func (pkg pkg) Register(b ioc.Builder) {
	ioc.RegisterSingleton(b, func(c ioc.Dic) unit.Service {
		return internal.NewService(c, pkg.layer)
	})
}
