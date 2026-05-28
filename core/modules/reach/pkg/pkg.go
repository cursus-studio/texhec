package reachpkg

import (
	"core/modules/reach/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, internal.NewService)
})

func PkgT[FeatureComponent any]() ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		ioc.Register(b, internal.NewServiceT[FeatureComponent])
	})
}
