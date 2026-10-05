package actionspkg

import (
	"core/game"
	"core/modules/actions"
	"core/modules/actions/internal"
	featurepkg "engine/modules/feature/pkg"
	"errors"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		featurepkg.FieldPkg[actions.PlayerContext],
		featurepkg.FieldPkg[actions.Coords],
		featurepkg.FieldPkg[actions.Entity],
		featurepkg.FieldPkg[actions.Blueprint],

		featurepkg.RulePkg[actions.OwnRule](func(c ioc.Dic) func(actions.PlayerContext, actions.Entity) error {
			return ioc.Get[game.GameWorld](c).Player().Own
		}, errors.New("player:does not own")),
		featurepkg.RulePkg[actions.IsEnemyRule](func(c ioc.Dic) func(actions.PlayerContext, actions.Entity) error {
			return ioc.Get[game.GameWorld](c).Player().IsEnemy
		}, errors.New("player:is not an enemy")),

		featurepkg.NoInputRulePkg[actions.CanBuildRule](func(c ioc.Dic) func(actions.Entity) error {
			return ioc.Get[game.GameWorld](c).Deploy().CanBuild
		}, errors.New("deploy:cannot build")),
		featurepkg.RulePkg[actions.CanBuildBlueprintRule](func(c ioc.Dic) func(actions.Entity, actions.Blueprint) error {
			return ioc.Get[game.GameWorld](c).Deploy().CanBuildBlueprint
		}, errors.New("deploy:cannot build blueprint")),
		featurepkg.RulePkg[actions.CanPlaceRule](func(c ioc.Dic) func(actions.Blueprint, actions.Coords) error {
			return ioc.Get[game.GameWorld](c).Deploy().CanPlace
		}, errors.New("deploy:cannot place")),

		featurepkg.NoInputRulePkg[actions.CanMoveRule](func(c ioc.Dic) func(actions.Entity) error {
			return ioc.Get[game.GameWorld](c).Pathfind().CanMove
		}, errors.New("pathfind:cannot move")),
		featurepkg.RulePkg[actions.CanMoveToRule](func(c ioc.Dic) func(actions.Entity, actions.Coords) error {
			return ioc.Get[game.GameWorld](c).Pathfind().CanMoveTo
		}, errors.New("pathfind:cannot move to destination")),

		featurepkg.NoInputRulePkg[actions.CanAttackRule](func(c ioc.Dic) func(actions.Entity) error {
			return ioc.Get[game.GameWorld](c).Attack().CanAttack
		}, errors.New("attack:cannot attack")),
		featurepkg.RulePkg[actions.CanBeAttackedRule](func(c ioc.Dic) func(actions.Entity, actions.Entity) error {
			return ioc.Get[game.GameWorld](c).Attack().CanBeAttacked
		}, errors.New("attack:cannot attack target")),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) actions.Service {
		return internal.NewService(c)
	})
})
