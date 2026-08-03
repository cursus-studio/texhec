package attackpkg

import (
	"core/modules/actions"
	"core/modules/attack"
	"core/modules/attack/internal"
	interactionspkg "engine/modules/interactions/pkg"
	"unsafe"

	"github.com/ogiusek/ioc/v2"
)

type AttackFeature struct {
	By     actions.FriendlyBuilderEntityStep
	Target actions.EnemyEntityStep
}

func (f AttackFeature) Event() any {
	return attack.NewAttackEvent(
		f.By.State().Entity,
		f.Target.State().Entity,
	)
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		interactionspkg.FeaturePkg[AttackFeature](
			interactionspkg.NewCopyRelation[actions.AnchorComponent](
				unsafe.Offsetof(AttackFeature{}.By), unsafe.Offsetof(AttackFeature{}.Target)),
		),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) attack.Service {
		return internal.NewService(c)
	})
})
