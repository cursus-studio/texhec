package deploypkg

import (
	"core/game"
	"core/modules/deploy"
	"core/modules/deploy/internal"
	"core/modules/reach"
	reachpkg "core/modules/reach/pkg"
	"engine/modules/entityregistry"
	"engine/modules/grid"
	"engine/services/ecs"
	"errors"
	"fmt"
	"strconv"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		reachpkg.PkgT[deploy.Component],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) deploy.Service {
		return internal.NewService(c)
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
