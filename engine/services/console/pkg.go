package console

import "github.com/ogiusek/ioc/v2"

type pkg struct{}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg pkg) Register(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) Console { return newConsole() })
}
