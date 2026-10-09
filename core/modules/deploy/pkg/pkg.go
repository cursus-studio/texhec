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
			featurepkg.RegisterRuleToField(config, f, actions.NewOwnRule(&f, &f.PlayerContext, &f.By))
			featurepkg.RegisterRuleToField(config, f, actions.NewCanDeployRule(&f, &f.By))
			featurepkg.RegisterRuleToField(config, f, actions.NewCanDeployBlueprintRule(&f, &f.By, &f.Blueprint))
			featurepkg.RegisterRuleToField(config, f, actions.NewCanDeployAtRule(&f, &f.By, &f.Blueprint, &f.Coords))
		}
		{
			f := featurepkg.RegisterFeature[deploy.DestroyEvent](config)
			featurepkg.RegisterRuleToField(config, f, actions.NewOwnRule(&f, &f.PlayerContext, &f.Entity))
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
