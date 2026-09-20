package attackpkg

import (
	"core/game"
	"core/modules/actions"
	"core/modules/attack"
	"core/modules/attack/internal"
	"core/modules/reach"
	reachpkg "core/modules/reach/pkg"
	"engine/modules/ecs"
	"engine/modules/entityregistry"
	"engine/modules/grid"
	interactionspkg "engine/modules/interactions/pkg"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"errors"
	"fmt"
	"strconv"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type AttackFeature struct {
	By     actions.FriendlyOffensiveEntityStep
	Target actions.EnemyEntityStep
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[attack.TargetComponent],
		typeregistrypkg.PkgT[attack.HealthComponent],
		typeregistrypkg.PkgT[attack.DamageComponent],

		typeregistrypkg.PkgT[attack.AttackEvent],

		reachpkg.PkgT[attack.TargetComponent],
		interactionspkg.FeaturePkg[AttackFeature](
		// interactionspkg.NewCopyRelation[actions.AnchorComponent](
		// 	unsafe.Offsetof(AttackFeature{}.By), unsafe.Offsetof(AttackFeature{}.Target)),
		),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Wrap(b, func(c ioc.Dic, b events.Builder) {
		world := ioc.Get[game.GameWorld](c)
		events.Listen(b, func(f AttackFeature) {
			events.Emit(world.Events(), attack.NewAttackEvent(f.By.State().UUID, f.Target.State().UUID))
		})
	})
	ioc.Register(b, func(c ioc.Dic) attack.Service {
		return internal.NewService(c)
	})

	ioc.Wrap(b, func(c ioc.Dic, b entityregistry.Service) {
		world := ioc.Get[game.GameWorld](c)
		b.Register("attackReach", func(entity ecs.EntityID, structTagValue string) {
			reachVal, err := strconv.Atoi(structTagValue)
			if err != nil {
				world.Logger().Warn(errors.Join(
					fmt.Errorf("couldn't set for entity \"%v\" attackReach", entity),
					err,
				))
				return
			}
			reachVal *= reachVal
			reachComp := reach.NewReach[attack.TargetComponent](grid.Coord(reachVal))
			world.Attack().Reach().Component().Set(entity, reachComp)
		})
		b.Register("health", func(entity ecs.EntityID, structTagValue string) {
			val, err := strconv.Atoi(structTagValue)
			if err != nil {
				world.Logger().Warn(errors.Join(
					fmt.Errorf("couldn't set for entity \"%v\" health", entity),
					err,
				))
				return
			}
			comp := attack.NewHealth(attack.Health(val))
			world.Attack().Health().Set(entity, comp)
		})
		b.Register("damage", func(entity ecs.EntityID, structTagValue string) {
			val, err := strconv.Atoi(structTagValue)
			if err != nil {
				world.Logger().Warn(errors.Join(
					fmt.Errorf("couldn't set for entity \"%v\" damage", entity),
					err,
				))
				return
			}
			comp := attack.NewDamage(attack.Health(val))
			world.Attack().Damage().Set(entity, comp)
		})
	})
})
