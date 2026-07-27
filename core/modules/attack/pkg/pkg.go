package attackpkg

import (
	"core/modules/actions"
	"core/modules/attack"
	"core/modules/attack/internal"
	interactionspkg "engine/modules/interactions/pkg"
	"unsafe"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		interactionspkg.FeaturePkg[attack.AttackEvent](
			interactionspkg.NewCopyRelation[actions.CoordsAnchorComponent](
				unsafe.Offsetof(attack.AttackEvent{}.By),
				unsafe.Offsetof(attack.AttackEvent{}.Target),
			),
		),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) attack.Service {
		return internal.NewService(c)
	})
})
