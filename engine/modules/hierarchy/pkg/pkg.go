package hierarchypkg

import (
	"engine/modules/hierarchy"
	"engine/modules/hierarchy/internal/hierarchyservice"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/services/codec"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PkgT[hierarchy.Component](),
	} {
		pkg(b)
	}
	ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// components
			Register(hierarchy.Component{})
	})

	ioc.Register(b, func(c ioc.Dic) hierarchy.Service {
		return hierarchyservice.NewService(c)
	})
})
