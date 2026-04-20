package groupspkg

import (
	"engine/modules/groups"
	"engine/modules/groups/internal"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/services/codec"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PkgT[groups.GroupsComponent](),
		prototypepkg.PkgT[groups.InheritGroupsComponent](),
	} {
		pkg(b)
	}
	ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// components
			Register(groups.GroupsComponent{})
	})
	ioc.Register(b, func(c ioc.Dic) groups.Service {
		return internal.NewService(c)
	})
})
