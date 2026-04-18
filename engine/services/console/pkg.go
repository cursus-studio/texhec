package console

import "github.com/ogiusek/ioc/v2"

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) Console { return newConsole() })
})
