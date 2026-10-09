package example

import (
	"engine/modules/ecs"
	"engine/modules/feature"
	featurepkg "engine/modules/feature/pkg"
	"errors"

	"github.com/ogiusek/ioc/v2"
)

// Notes:
// - in real implementation listeners should be services methods (func(ecs.EntityID))
// - no implementations or world is here
// - features should be structs without any interfaces to be composable. so [EntityField] should be a struct type wrapper to

//
// field types

type PlayerCTX any

type EntityField any
type CoordsField any
type BlueprintField any

//
// field implementation

type entityFieldService struct {
	Field ioc.Lazy[feature.FieldService[EntityField]] `inject:""`
	// World
}

func NewEntityFieldSystem(c ioc.Dic) ecs.SystemRegister {
	return ioc.GetServices[*entityFieldService](c)
}

func (s *entityFieldService) Register() error {
	// register field renderers
	s.Field().Render().OnUpsert(func(entity ecs.EntityID) {
		// add something to show that field exists

		// _ /*this is of type below and what we can use to render*/, ok := s.Field().Render().Get(entity)
		// type FieldRenderComponent[FieldType any] struct {
		// 	Value       FieldType // we're supposed to render Value with certain modifications depending on flags like PassedTests or IsPreview
		// 	PassedTests bool
		// 	IsPreview   bool
		// }
	})
	s.Field().Render().OnRemove(func(ecs.EntityID) {
		// remove that something
	})
	return nil
}

//
// field rules

type OwnsRule struct {
	PlayerCTX feature.FieldGetter[PlayerCTX]
	Entity    feature.FieldGetter[EntityField]
}
type CanAttackRule feature.Rule
type ReachesRule feature.Rule
type CanDeployRule feature.Rule
type CanMoveToRule feature.Rule
type CanPlaceOnRule feature.Rule

//
// rule implementation

type ownsRuleService struct {
	Rule ioc.Lazy[feature.RuleService[OwnsRule]] `inject:""`
	// World
}

func NewOwnsRuleSystem(c ioc.Dic) ecs.SystemRegister {
	return ioc.GetServices[*ownsRuleService](c)
}
func (s *ownsRuleService) Register() error {
	// register rule renderers
	s.Rule().Render().OnUpsert(func(entity ecs.EntityID) {
		// add something to show that rule exists

		// _/*this is of type below and what we can use to render*/, ok := s.Rule().Render().Get(entity)
		// type RuleRenderComponent[RuleType Rule[DataIn, AppliedTo], DataIn any, AppliedTo any] struct {
		// 	In        DataIn
		// 	AppliedTo AppliedTo // we're meant to show how appliedTo applied to dataIn fails
		// }
	})
	s.Rule().Render().OnRemove(func(ecs.EntityID) {
		// remove that something
	})
	return nil
}

//
// features

type AttackFeature struct {
	PlayerCTX
	Attacker EntityField // is owned by PlayerCTX
	Attacked EntityField // attacker can attack attacker
}
type MoveFeature struct {
	PlayerCTX
	Entity EntityField // is owned by PlayerCTX
	Coords CoordsField // can entity obstruction be on coords
}
type DeployFeature struct {
	PlayerCTX
	Entity   EntityField    // is owned by PlayerCTX
	Deployed BlueprintField // entity can deploy, PlayerCTX wallet is big enough
	Coords   CoordsField    // is in range of entity, can obstruction blueprint obstruction be on coords
}
type DestroyFeature struct {
	PlayerCTX
	Entity EntityField // is owned by PlayerCTX
}

//
// wire services

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		// fields
		featurepkg.FieldPkg[EntityField],
		featurepkg.FieldPkg[CoordsField],
		featurepkg.FieldPkg[BlueprintField],

		// rules
		featurepkg.RulePkg(func(c ioc.Dic) func(OwnsRule, feature.AllFields) error {
			return func(rule OwnsRule, fields feature.AllFields) error {
				playerCtx := rule.PlayerCTX.From(fields)
				entity := rule.Entity.From(fields)
				if playerCtx == entity {
					return nil
				}
				return nil // return rule error
			}
		}, errors.New("does not own")),
		// RulePkg[CanAttackRule](),
		// RulePkg[ReachesRule](),
		// RulePkg[CanDeployRule](),
		// RulePkg[CanMoveToRule](),
		// RulePkg[CanPlaceOnRule](),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	// features
	// each register/wrap can be in different module
	ioc.Wrap(b, func(c ioc.Dic, config featurepkg.Config) {
		// RegisterProvidedField(config, ioc.Get[ProvideFieldService[PlayerCTX]](c))
		featurepkg.RegisterFieldProvider(config, func() PlayerCTX { return nil /**/ })

		// renderers are optional but recomended
		featurepkg.RegisterSystems(config,
			// for each rendered field
			NewEntityFieldSystem(c),
			// ...

			// for each rendered rule
			NewOwnsRuleSystem(c),
			// ...
		)
	})
	ioc.Wrap(b, func(c ioc.Dic, config featurepkg.Config) {
		f := featurepkg.RegisterFeature[AttackFeature](config)
		featurepkg.RegisterRuleToField(config, &f, &f.Attacker, OwnsRule{
			feature.NewFieldGetter(&f, &f.PlayerCTX),
			feature.NewFieldGetter(&f, &f.Attacker),
		})
		featurepkg.RegisterRuleToField[CanAttackRule](config, &f, &f.Attacked, nil) // should replace these nils
	})
	ioc.Wrap(b, func(c ioc.Dic, config featurepkg.Config) {
		f := featurepkg.RegisterFeature[MoveFeature](config)
		featurepkg.RegisterRuleToField(config, &f, &f.Entity, OwnsRule{
			feature.NewFieldGetter(&f, &f.PlayerCTX),
			feature.NewFieldGetter(&f, &f.Entity),
		})
		featurepkg.RegisterRuleToField[CanMoveToRule](config, &f, &f.Coords, nil)
	})
	ioc.Wrap(b, func(c ioc.Dic, config featurepkg.Config) {
		f := featurepkg.RegisterFeature[DeployFeature](config)
		featurepkg.RegisterRuleToField(config, &f, &f.Entity, OwnsRule{
			feature.NewFieldGetter(&f, &f.PlayerCTX),
			feature.NewFieldGetter(&f, &f.Entity),
		})
		featurepkg.RegisterRuleToField[CanDeployRule](config, &f, &f.Deployed, nil)
		featurepkg.RegisterRuleToField[CanPlaceOnRule](config, &f, &f.Coords, nil)
	})
	ioc.Wrap(b, func(c ioc.Dic, config featurepkg.Config) {
		f := featurepkg.RegisterFeature[DestroyFeature](config)
		featurepkg.RegisterRuleToField(config, &f, &f.Entity, OwnsRule{
			feature.NewFieldGetter(&f, &f.PlayerCTX),
			feature.NewFieldGetter(&f, &f.Entity),
		})
	})
})
