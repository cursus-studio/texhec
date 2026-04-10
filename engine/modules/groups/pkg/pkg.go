package groupspkg

import (
	"engine/modules/groups"
	"engine/modules/groups/internal"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/services/codec"

	"github.com/ogiusek/ioc/v2"
)

type pkg struct {
}

func Package() ioc.Pkg {
	return pkg{}
}

func (pkg) Register(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PackageT[groups.GroupsComponent](),
		prototypepkg.PackageT[groups.InheritGroupsComponent](),
	} {
		pkg.Register(b)
	}
	ioc.WrapService(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// components
			Register(groups.GroupsComponent{})
	})
	ioc.RegisterSingleton(b, func(c ioc.Dic) groups.Service {
		return internal.NewService(c)
	})
}
