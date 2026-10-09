package actionspkg

import (
	"core/game"
	"core/modules/actions"
	"core/modules/actions/internal"
	"engine/modules/ecs"
	"engine/modules/feature"
	featurepkg "engine/modules/feature/pkg"
	netsyncpkg "engine/modules/netsync/pkg"
	"errors"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		featurepkg.FieldPkg[actions.PlayerContext],
		featurepkg.FieldPkg[actions.Coords],
		featurepkg.FieldPkg[actions.Entity],
		featurepkg.FieldPkg[actions.Blueprint],

		featurepkg.RulePkg(func(c ioc.Dic) func(actions.OwnRule, feature.AllFields) error {
			world := ioc.Get[game.GameWorld](c)
			return func(rule actions.OwnRule, fields feature.AllFields) error {
				return world.Player().Own(rule.PlayerContext.From(fields), rule.Entity.From(fields))
			}
		}, errors.New("player:does not own")),
		featurepkg.RulePkg(func(c ioc.Dic) func(actions.IsEnemyRule, feature.AllFields) error {
			world := ioc.Get[game.GameWorld](c)
			return func(rule actions.IsEnemyRule, fields feature.AllFields) error {
				return world.Player().IsEnemy(rule.PlayerContext.From(fields), rule.Entity.From(fields))
			}
		}, errors.New("player:is not an enemy")),

		// build
		featurepkg.RulePkg(func(c ioc.Dic) func(actions.CanDeployRule, feature.AllFields) error {
			world := ioc.Get[game.GameWorld](c)
			return func(rule actions.CanDeployRule, fields feature.AllFields) error {
				return world.Deploy().CanDeploy(rule.Entity.From(fields))
			}
		}, errors.New("deploy:cannot build")),
		featurepkg.RulePkg(func(c ioc.Dic) func(actions.CanDeployBlueprintRule, feature.AllFields) error {
			world := ioc.Get[game.GameWorld](c)
			return func(rule actions.CanDeployBlueprintRule, fields feature.AllFields) error {
				return world.Deploy().CanDeployBlueprint(rule.Entity.From(fields), rule.Blueprint.From(fields))
			}
		}, errors.New("deploy:cannot build blueprint")),
		featurepkg.RulePkg(func(c ioc.Dic) func(actions.CanDeployAtRule, feature.AllFields) error {
			world := ioc.Get[game.GameWorld](c)
			return func(rule actions.CanDeployAtRule, fields feature.AllFields) error {
				return world.Deploy().CanDeployAt(
					rule.Entity.From(fields), rule.Blueprint.From(fields), rule.Coords.From(fields))
			}
		}, errors.New("deploy:cannot place")),

		// move
		featurepkg.RulePkg(func(c ioc.Dic) func(actions.CanMoveRule, feature.AllFields) error {
			world := ioc.Get[game.GameWorld](c)
			return func(rule actions.CanMoveRule, fields feature.AllFields) error {
				return world.Pathfind().CanMove(rule.Entity.From(fields))
			}
		}, errors.New("pathfind:cannot move")),
		featurepkg.RulePkg(func(c ioc.Dic) func(actions.CanMoveToRule, feature.AllFields) error {
			world := ioc.Get[game.GameWorld](c)
			return func(rule actions.CanMoveToRule, fields feature.AllFields) error {
				return world.Pathfind().CanMoveTo(rule.Entity.From(fields), rule.Coords.From(fields))
			}
		}, errors.New("pathfind:cannot move to destination")),

		// attack
		featurepkg.RulePkg(func(c ioc.Dic) func(actions.CanAttackRule, feature.AllFields) error {
			world := ioc.Get[game.GameWorld](c)
			return func(rule actions.CanAttackRule, fields feature.AllFields) error {
				return world.Attack().CanAttack(rule.Entity.From(fields))
			}
		}, errors.New("attack:cannot attack")),
		featurepkg.RulePkg(func(c ioc.Dic) func(actions.CanBeAttackedRule, feature.AllFields) error {
			world := ioc.Get[game.GameWorld](c)
			return func(rule actions.CanBeAttackedRule, fields feature.AllFields) error {
				return world.Attack().CanBeAttacked(rule.Attacker.From(fields), rule.Attacked.From(fields))
			}
		}, errors.New("attack:cannot attack target")),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) actions.Service {
		return internal.NewService(c)
	})

	ioc.Wrap(b, func(c ioc.Dic, config netsyncpkg.Config) {
		world := ioc.GetServices[game.GameWorld](c)
		netsyncpkg.AddGenericEventValidation(config, func(client ecs.EntityID, event any) error {
			err := world.Feature().Validate(event)
			if err == feature.ErrInvalidType {
				return nil
			}
			return err
		})
	})
})
