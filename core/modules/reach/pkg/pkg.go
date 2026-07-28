package reachpkg

import (
	"core/modules/reach"
	"core/modules/reach/internal"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, internal.NewService)
})

func PkgT[FeatureComponent any](b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[reach.Component[FeatureComponent]],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, internal.NewServiceT[FeatureComponent])
}
