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
	featurepkg "engine/modules/feature/pkg"
	"engine/modules/grid"
	netsyncpkg "engine/modules/netsync/pkg"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"errors"
	"fmt"
	"strconv"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[attack.TargetComponent],
		typeregistrypkg.PkgT[attack.HealthComponent],
		typeregistrypkg.PkgT[attack.DamageComponent],

		typeregistrypkg.PkgT[attack.AttackEvent],

		reachpkg.PkgT[attack.TargetComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) attack.Service {
		return internal.NewService(c)
	})
	ioc.Wrap(b, func(c ioc.Dic, config featurepkg.Config) {
		f := featurepkg.RegisterFeature[attack.AttackEvent](config)
		featurepkg.RegisterRuleToField(config, f, actions.NewOwnRule(&f, &f.PlayerContext, &f.Attacker))
		featurepkg.RegisterRuleToField(config, f, actions.NewCanAttackRule(&f, &f.Attacker))
		featurepkg.RegisterRuleToField(config, f, actions.NewCanBeAttackedRule(&f, &f.Attacker, &f.Target))
	})
	ioc.Wrap(b, func(c ioc.Dic, config netsyncpkg.Config) {
		world := ioc.GetServices[game.GameWorld](c)
		netsyncpkg.AddEventValidation(config, func(event attack.AttackEvent) error {
			return world.Player().Own(event.PlayerContext, event.Attacker)
		})
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
