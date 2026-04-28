package codecpkg

import (
	"engine/modules/codec"
	"engine/modules/codec/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) internal.Builder {
		return internal.NewBuilder(c)
	})
	ioc.Register(b, func(c ioc.Dic) codec.Service {
		return ioc.Get[internal.Builder](c).Build()
	})
})

func PkgT[Component any](b ioc.Builder) {
	var zero Component
	ioc.Wrap(b, func(c ioc.Dic, b internal.Builder) {
		b.Register(zero)
	})
}
