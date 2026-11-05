package playscene

import (
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		playPkg,
		playConnectPkg,
		playGeneratePkg,
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
})
