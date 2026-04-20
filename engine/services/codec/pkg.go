package codec

import (
	"engine/services/logger"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) Builder {
		return NewBuilder(ioc.Get[logger.Logger](c))
	})
	ioc.Register(b, func(c ioc.Dic) Codec {
		return ioc.Get[Builder](c).Build()
	})
})
