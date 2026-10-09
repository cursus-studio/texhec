package deploypkg

import (
	"core/game"
	"core/modules/actions"
	"core/modules/deploy"
	"core/modules/deploy/internal"
	"core/modules/reach"
	reachpkg "core/modules/reach/pkg"
	"engine/modules/ecs"
	"engine/modules/entityregistry"
	"engine/modules/feature"
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
		typeregistrypkg.PkgT[deploy.DeployEvent],
		typeregistrypkg.PkgT[deploy.DestroyEvent],
		typeregistrypkg.PkgT[internal.BoughtComponent],

		reachpkg.PkgT[deploy.Component],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) deploy.Service {
		return internal.NewService(c)
	})
	ioc.Wrap(b, func(c ioc.Dic, config featurepkg.Config) {
		{
			f := featurepkg.RegisterFeature[deploy.DeployEvent](config)
			featurepkg.RegisterRuleToField(config, &f, &f.By, actions.OwnRule{
				PlayerContext: feature.NewFieldGetter(&f, &f.PlayerContext),
				Entity:        feature.NewFieldGetter(&f, &f.By),
			})
			featurepkg.RegisterRuleToField(config, &f, &f.By, actions.CanDeployRule{
				Entity: feature.NewFieldGetter(&f, &f.By),
			})
			featurepkg.RegisterRuleToField(config, &f, &f.Blueprint, actions.CanDeployBlueprintRule{
				Entity:    feature.NewFieldGetter(&f, &f.By),
				Blueprint: feature.NewFieldGetter(&f, &f.Blueprint),
			})
			featurepkg.RegisterRuleToField(config, &f, &f.Coords, actions.CanDeployAtRule{
				Blueprint: feature.NewFieldGetter(&f, &f.Blueprint),
				Coords:    feature.NewFieldGetter(&f, &f.Coords),
			})
		}
		{
			f := featurepkg.RegisterFeature[deploy.DestroyEvent](config)
			featurepkg.RegisterRuleToField(config, &f, &f.Entity, actions.OwnRule{
				PlayerContext: feature.NewFieldGetter(&f, &f.PlayerContext),
				Entity:        feature.NewFieldGetter(&f, &f.Entity),
			})
		}
	})
	ioc.Wrap(b, func(c ioc.Dic, config netsyncpkg.Config) {
		world := ioc.GetServices[game.GameWorld](c)
		netsyncpkg.AddEventValidation(config, func(event deploy.DeployEvent) error {
			return world.Player().Own(event.PlayerContext, event.By)
		})
		netsyncpkg.AddEventValidation(config, func(event deploy.DestroyEvent) error {
			return world.Player().Own(event.PlayerContext, event.Entity)
		})
	})

	ioc.Wrap(b, func(c ioc.Dic, b entityregistry.Service) {
		world := ioc.Get[game.GameWorld](c)
		b.Register("deployReach", func(entity ecs.EntityID, structTagValue string) {
			reachVal, err := strconv.Atoi(structTagValue)
			if err != nil {
				world.Logger().Warn(errors.Join(
					fmt.Errorf("couldn't set for entity \"%v\" deployReach", entity),
					err,
				))
				return
			}
			reachVal *= reachVal
			reachComp := reach.NewReach[deploy.Component](grid.Coord(reachVal))
			world.Deploy().Reach().Component().Set(entity, reachComp)
		})
	})
})
