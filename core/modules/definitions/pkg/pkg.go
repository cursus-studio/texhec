package definitionspkg

import (
	"core/modules/definitions"
	"core/modules/definitions/internal"
	_ "image/png"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) definitions.Service {
		return internal.NewService(c)
	})
})
