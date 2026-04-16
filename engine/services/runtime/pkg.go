package runtime

import (
	"github.com/ogiusek/ioc/v2"
)

type pkg struct{}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg) Register(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) Builder {
		return newBuilder()
	})
	ioc.Register(b, func(c ioc.Dic) Runtime {
		return ioc.Get[Builder](c).Build()
	})
}
