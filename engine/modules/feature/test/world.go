package test

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/feature"
	featurepkg "engine/modules/feature/pkg"
	"errors"
	"maps"
	"slices"
	"testing"

	"github.com/ogiusek/ioc/v2"
)

var (
	ErrNotAnOwner error = errors.New("not an owner")
	ErrNotAnEnemy error = errors.New("not an enemy")

	ProvidedPlayerCTX = NewPlayerCTX(12)
	EnemyPlayerCTX    = NewPlayerCTX(13)
)

type PlayerCTX struct{ ID int }
type EntityField struct{ OwnerID int }

func NewPlayerCTX(id int) PlayerCTX          { return PlayerCTX{id} }
func NewEntityField(ownerID int) EntityField { return EntityField{ownerID} }

type OwnsRule feature.Rule[PlayerCTX, EntityField]        // [PlayerCTX] owns [EntityField]
type CanAttackRule feature.Rule[EntityField, EntityField] // [EntityField] canAttack [EntityField]
type EnemyRule feature.Rule[PlayerCTX, EntityField]       // [PlayerCTX] isEnemy [EntityField]

type AttackFeature struct {
	PlayerCTX
	Attacker EntityField
	Attacked EntityField
}
type MoveToFeature struct {
	PlayerCTX
	From EntityField
	To   EntityField
}
type FocusEnemyFeature struct {
	PlayerCTX
	Enemy EntityField
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		// fields
		featurepkg.FieldPkg[EntityField],

		// rules
		featurepkg.RulePkg[OwnsRule](func(c ioc.Dic) func(ctx PlayerCTX, entity EntityField) error {
			return func(ctx PlayerCTX, entity EntityField) error {
				if ctx.ID == entity.OwnerID {
					return nil
				}
				return ErrNotAnOwner
			}
		}),
		featurepkg.RulePkg[CanAttackRule](func(c ioc.Dic) func(ctx EntityField, entity EntityField) error {
			return func(ctx EntityField, entity EntityField) error {
				if ctx.OwnerID != entity.OwnerID {
					return nil
				}
				return ErrNotAnEnemy
			}
		}),
		featurepkg.RulePkg[EnemyRule](func(c ioc.Dic) func(ctx PlayerCTX, entity EntityField) error {
			return func(ctx PlayerCTX, entity EntityField) error {
				if ctx.ID != entity.OwnerID {
					return nil
				}
				return ErrNotAnEnemy
			}
		}),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	// features
	ioc.Wrap(b, func(c ioc.Dic, config featurepkg.Config) {
		featurepkg.RegisterFieldProvider(config, func() PlayerCTX { return ProvidedPlayerCTX })

		{
			f := featurepkg.RegisterFeature[AttackFeature](config)
			featurepkg.RegisterRuleToField[OwnsRule](config, &f, &f.PlayerCTX, &f.Attacker)
			featurepkg.RegisterRuleToField[CanAttackRule](config, &f, &f.Attacker, &f.Attacked)
		}
		{
			f := featurepkg.RegisterFeature[MoveToFeature](config)
			featurepkg.RegisterRuleToField[OwnsRule](config, &f, &f.PlayerCTX, &f.From)
		}
		{
			f := featurepkg.RegisterFeature[FocusEnemyFeature](config)
			featurepkg.RegisterRuleToField[EnemyRule](config, &f, &f.PlayerCTX, &f.Enemy)
		}
	})
})

type World struct {
	World         engine.EngineWorld                                           `inject:""`
	EntityField   feature.FieldService[EntityField]                            `inject:""`
	OwnsRule      feature.RuleService[OwnsRule, PlayerCTX, EntityField]        `inject:""`
	CanAttackRule feature.RuleService[CanAttackRule, EntityField, EntityField] `inject:""`
	EnemyRule     feature.RuleService[EnemyRule, PlayerCTX, EntityField]       `inject:""`
	Config        featurepkg.Config                                            `inject:""`

	Rules map[feature.RuleKey]feature.AnyRuleService
	T     *testing.T
}

func (w World) AssertInteractions(expected feature.InteractionsComponent) {
	w.T.Helper()
	{ // verify struct state
		actual := w.World.Feature().Interactions()
		if !actual.Equal(expected) {
			w.T.Fatalf("expected\n%v\nnot\n%v", expected, actual)
		}
		if expected.LockedFeature && len(expected.Features) != 1 {
			w.T.Fatalf("if feature is locked there should be a single feature")
		}
	}

	// verify world state
	expectedComponents := make(map[any]struct{})

	// populate expected components
	for _, interaction := range expected.Interactions {
		comp := feature.NewFieldRender(interaction.(EntityField), feature.Interaction)
		expectedComponents[comp] = struct{}{}
	}
	if expected.LockedFeature && expected.PreviewInteraction == nil {
		comp := feature.NewFieldRender(EntityField{}, feature.MissingField)
		expectedComponents[comp] = struct{}{}
	}
	if expected.LockedFeature {
		featureKey := expected.Features[0]
		featureConfig := w.Config.GetFeatureConfig(featureKey)
		allFields := []any{}
		interactions := slices.Clone(expected.Interactions)
		for len(interactions) != 0 {
			fieldConfig := featureConfig.Fields[len(allFields)]
			if fieldConfig.Provide != nil {
				allFields = append(allFields, fieldConfig.Provide())
				continue
			}
			allFields = append(allFields, interactions[0])
			interactions = interactions[1:]
		}
		fieldIndex := len(allFields)
		fieldConfig := featureConfig.Fields[fieldIndex]

		rulesPass := feature.PreviewRulesPass
		for _, ruleConfig := range fieldConfig.Rules {
			rule := w.Rules[ruleConfig.RuleType]
			dataIn := allFields[ruleConfig.DataInFieldIndex]
			comp := rule.NewAnyRenderComponent(dataIn, expected.PreviewInteraction)
			expectedComponents[comp] = struct{}{}
			if err := rule.AppliesToAny(dataIn, expected.PreviewInteraction); err != nil {
				rulesPass = feature.PreviewRulesFail
			}
		}
		if expected.PreviewInteraction != nil {
			comp := feature.NewFieldRender(expected.PreviewInteraction.(EntityField), rulesPass)
			expectedComponents[comp] = struct{}{}
		}
	}

	// verify is this actual
	arrays := []ecs.AnyComponentArray{
		w.EntityField.RenderAny(),
		w.OwnsRule.RenderAny(),
		w.CanAttackRule.RenderAny(),
		w.EnemyRule.RenderAny(),
	}
	foundComponents := make(map[any]struct{})
	for _, array := range arrays {
		for _, entity := range array.GetEntities() {
			comp, _ := array.GetAny(entity)
			foundComponents[comp] = struct{}{}
		}
	}

	if !maps.Equal(expectedComponents, foundComponents) {
		w.T.Fatalf("for %v expected\n%v\nand found\n%v",
			expected, expectedComponents, foundComponents)
	}
}
