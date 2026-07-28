package consolepkg

import (
	"engine/modules/console"
	"engine/modules/console/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) console.Service { return internal.NewConsole() })
})
