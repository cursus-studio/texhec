package groupspkg

import (
	codecpkg "engine/modules/codec/pkg"
	"engine/modules/groups"
	"engine/modules/groups/internal"
	prototypepkg "engine/modules/prototype/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[groups.GroupsComponent],
		codecpkg.PkgT[groups.InheritGroupsComponent],

		prototypepkg.PkgT[groups.GroupsComponent],
		prototypepkg.PkgT[groups.InheritGroupsComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) groups.Service {
		return internal.NewService(c)
	})
})
