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
	"reflect"
	"strconv"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		reachpkg.PkgT[deploy.Component],
		interactionspkg.FeaturePkg("deploy", []reflect.Type{
			reflect.TypeFor[tile.ObjectInteraction](),
			reflect.TypeFor[tile.SourceObjectInteraction](),
			reflect.TypeFor[tile.CoordsInteraction](),
		}, func(c ioc.Dic) func() deploy.DeployEvent {
			s := ioc.Get[game.GameWorld](c)
			return func() deploy.DeployEvent {
				e := deploy.DeployEvent{}
				featureEntity := s.Interactions().FeatureEntity()
				if comp, ok := s.Tile().CoordsInteraction().Interaction().Get(featureEntity); ok {
					e.Coords = comp.State.Coords
				}
				if comp, ok := s.Tile().ObjectInteraction().Interaction().Get(featureEntity); ok {
					e.By = comp.State.Entity
				}
				if comp, ok := s.Tile().SourceObjectInteraction().Interaction().Get(featureEntity); ok {
					e.Blueprint = comp.State.Entity
				}
				return e
			}
		}),
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
