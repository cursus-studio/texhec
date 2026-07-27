package actionspkg

import (
	"core/game"
	"core/modules/actions"
	"core/modules/actions/internal"
	"core/modules/player"
	interactionspkg "engine/modules/interactions/pkg"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		// coords interaction
		interactionspkg.InteractionPkg[actions.CoordsInteraction](),
		interactionspkg.StepPkg[actions.CoordsStep](func(c ioc.Dic) func(state actions.CoordsInteraction) error {
			return func(state actions.CoordsInteraction) error { return nil }
		}),

		// object interaction
		interactionspkg.InteractionPkg[actions.ObjectInteraction](),
		interactionspkg.StepPkg[actions.ObjectStep](func(c ioc.Dic) func(state actions.ObjectInteraction) error {
			return func(state actions.ObjectInteraction) error { return nil }
		}),
		interactionspkg.StepPkg[actions.FriendlyObjectStep](func(c ioc.Dic) func(state actions.ObjectInteraction) error {
			world := ioc.Get[game.GameWorld](c)
			return func(state actions.ObjectInteraction) error {
				return world.Player().ControlsObject(state.Entity)
			}
		}),
		interactionspkg.StepPkg[actions.FriendlyMobileObjectStep](func(c ioc.Dic) func(state actions.ObjectInteraction) error {
			world := ioc.Get[game.GameWorld](c)
			return func(state actions.ObjectInteraction) error {
				if err := world.Player().ControlsObject(state.Entity); err != nil {
					return err
				}
				if _, ok := world.Pathfind().Speed().Get(state.Entity); !ok {
					return actions.ErrRequiresSpeed
				}
				return nil
			}
		}),
		interactionspkg.StepPkg[actions.FriendlyBuilderObjectStep](func(c ioc.Dic) func(state actions.ObjectInteraction) error {
			world := ioc.Get[game.GameWorld](c)
			return func(state actions.ObjectInteraction) error {
				if err := world.Player().ControlsObject(state.Entity); err != nil {
					return err
				}
				link, _ := world.Metadata().Link().Get(state.Entity)
				if _, ok := world.Deploy().Component().Get(link.Entity); !ok {
					return actions.ErrRequiresDeploy
				}
				return nil
			}
		}),
		interactionspkg.StepPkg[actions.EnemyObjectStep](func(c ioc.Dic) func(state actions.ObjectInteraction) error {
			world := ioc.Get[game.GameWorld](c)
			return func(state actions.ObjectInteraction) error {
				if err := world.Player().ControlsObject(state.Entity); err != nil {
					return nil
				}
				return player.ErrRequiresToBeEnemy
			}
		}),

		// bluepring interaction
		interactionspkg.InteractionPkg[actions.BlueprintInteraction](),
		interactionspkg.StepPkg[actions.BlueprintStep](func(c ioc.Dic) func(state actions.BlueprintInteraction) error {
			return func(state actions.BlueprintInteraction) error { return nil }
		}),

		typeregistrypkg.PkgT[actions.CoordsCursorComponent],
		typeregistrypkg.PkgT[actions.CoordsAnchorComponent],
		typeregistrypkg.PkgT[actions.CanDeployComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) actions.Service {
		return internal.NewService(c)
	})
})
