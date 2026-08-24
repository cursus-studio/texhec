package smoothpkg

import (
	"engine/modules/smooth"
	"engine/modules/smooth/internal"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) smooth.Service { return internal.NewService(c) })
})

// func PkgT[Component smooth.SmoothConstraint[Component]](b ioc.Builder) {
func PkgT[Component any](b ioc.Builder) {
	var zero Component
	_ = any(zero).(smooth.SmoothConstraint[Component])
	ioc.Register(b, func(c ioc.Dic) *internal.ServiceT[Component] {
		return internal.NewServiceT[Component](c)
	})
	ioc.Wrap(b, func(c ioc.Dic, _ smooth.Service) {
		internal.NewSystems[Component](c)
	})
	ioc.Register(b, func(c ioc.Dic) smooth.ServiceT[Component] {
		return ioc.Get[*internal.ServiceT[Component]](c)
	})
}
