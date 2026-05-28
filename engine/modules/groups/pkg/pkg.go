package groupspkg

import (
	"engine/modules/groups"
	"engine/modules/groups/internal"
	hierarchypkg "engine/modules/hierarchy/pkg"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[groups.GroupsComponent],
		hierarchypkg.PkgT[groups.GroupsComponent](),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) groups.Service {
		return internal.NewService(c)
	})
})
