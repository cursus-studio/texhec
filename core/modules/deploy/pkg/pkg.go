package deploypkg

import (
	"core/game"
	"core/modules/deploy"
	"core/modules/deploy/internal"
	"core/modules/reach"
	reachpkg "core/modules/reach/pkg"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/entityregistry"
	"engine/modules/grid"
	interactionspkg "engine/modules/interactions/pkg"
	"errors"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		reachpkg.PkgT[deploy.Component],
		interactionspkg.FeaturePkg[deploy.DeployEvent](
			interactionspkg.NewCopyRelation[tile.CanDeployComponent](
				unsafe.Offsetof(deploy.DeployEvent{}.By),
				unsafe.Offsetof(deploy.DeployEvent{}.Blueprint),
			),
			interactionspkg.NewCopyRelation[tile.CoordsCursorComponent](
				unsafe.Offsetof(deploy.DeployEvent{}.Blueprint),
				unsafe.Offsetof(deploy.DeployEvent{}.Coords),
			),
			interactionspkg.NewCopyRelation[tile.CoordsAnchorComponent](
				unsafe.Offsetof(deploy.DeployEvent{}.By),
				unsafe.Offsetof(deploy.DeployEvent{}.Coords),
			),
		),
		interactionspkg.FeaturePkg[deploy.DestroyEvent](),
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
