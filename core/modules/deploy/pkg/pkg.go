package deploypkg

import (
	"core/game"
	"core/modules/actions"
	"core/modules/deploy"
	"core/modules/deploy/internal"
	"core/modules/player"
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
	"unsafe"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type DeployFeature struct {
	player.PlayerContext
	By        actions.FriendlyBuilderEntityStep
	Blueprint actions.BlueprintStep
	Coords    actions.CoordsStep
}

type DestroyFeature struct {
	player.PlayerContext
	Entity actions.FriendlyEntityStep
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[deploy.DeployEvent],
		typeregistrypkg.PkgT[deploy.DestroyEvent],
		typeregistrypkg.PkgT[internal.BoughtComponent],

		reachpkg.PkgT[deploy.Component],
		interactionspkg.FeaturePkg[DeployFeature](
			interactionspkg.NewCopyRelation[actions.CanDeployComponent](
				unsafe.Offsetof(DeployFeature{}.By), unsafe.Offsetof(DeployFeature{}.Blueprint)),
			interactionspkg.NewCopyRelation[actions.CoordsCursorComponent](
				unsafe.Offsetof(DeployFeature{}.Blueprint), unsafe.Offsetof(DeployFeature{}.Coords)),
			interactionspkg.NewCopyRelation[actions.AnchorComponent](
				unsafe.Offsetof(DeployFeature{}.By), unsafe.Offsetof(DeployFeature{}.Coords)),
		),
		interactionspkg.FeaturePkg[DestroyFeature](),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Wrap(b, func(c ioc.Dic, b events.Builder) {
		world := ioc.Get[game.GameWorld](c)
		events.Listen(b, func(f DeployFeature) {
			events.Emit(world.Events(), deploy.NewDeployEvent(
				f.By.State().UUID,
				f.Blueprint.State().UUID,
				f.Coords.State().Coords,
			))
		})
		events.Listen(b, func(f DestroyFeature) {
			events.Emit(world.Events(), deploy.NewDestroyEvent(
				f.Entity.State().UUID,
			))
		})
	})
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
