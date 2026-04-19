package mock

import (
	"engine/mock"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		mock.Pkg,
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
})
