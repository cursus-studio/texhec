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
	"engine/modules/grid"
	interactionspkg "engine/modules/interactions/pkg"
	"errors"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/ogiusek/ioc/v2"
)

type DeployFeature struct {
	By        actions.FriendlyBuilderEntityStep
	Blueprint actions.BlueprintStep
	Coords    actions.CoordsStep
}

func (f DeployFeature) Event() any {
	return deploy.NewDeployEvent(
		f.By.State().Entity,
		f.Blueprint.State().Entity,
		f.Coords.State().Coords,
	)
}

type DestroyFeature struct {
	Entity actions.FriendlyEntityStep
}

func (f DestroyFeature) Event() any {
	return deploy.NewDestroyEvent(f.Entity.State().Entity)
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
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
