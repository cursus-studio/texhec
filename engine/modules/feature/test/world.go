package test

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/feature"
	"engine/modules/feature/internal"
	featurepkg "engine/modules/feature/pkg"
	"errors"
	"fmt"
	"maps"
	"reflect"
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

type OwnsRule struct {
	PlayerCTX feature.FieldGetter[PlayerCTX]
	Entity    feature.FieldGetter[EntityField]
}
type CanAttackRule struct {
	Entity feature.FieldGetter[EntityField]
	Enemy  feature.FieldGetter[EntityField]
}
type EnemyRule struct {
	PlayerCTX feature.FieldGetter[PlayerCTX]
	Enemy     feature.FieldGetter[EntityField]
}

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
		featurepkg.RulePkg(func(c ioc.Dic) func(r OwnsRule, fields feature.AllFields) error {
			return func(rule OwnsRule, fields feature.AllFields) error {
				if rule.PlayerCTX.From(fields).ID == rule.Entity.From(fields).OwnerID {
					return nil
				}
				return ErrNotAnOwner
			}
		}, errors.New("does not own")),
		featurepkg.RulePkg(func(c ioc.Dic) func(rule CanAttackRule, fields feature.AllFields) error {
			return func(rule CanAttackRule, fields feature.AllFields) error {
				if rule.Entity.From(fields).OwnerID != rule.Enemy.From(fields).OwnerID {
					return nil
				}
				return ErrNotAnEnemy
			}
		}, errors.New("cannot attack")),
		featurepkg.RulePkg(func(c ioc.Dic) func(rule EnemyRule, fields feature.AllFields) error {
			return func(rule EnemyRule, fields feature.AllFields) error {
				if rule.PlayerCTX.From(fields).ID != rule.Enemy.From(fields).OwnerID {
					return nil
				}
				return ErrNotAnEnemy
			}
		}, errors.New("is not an enemy")),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	// features
	ioc.Wrap(b, func(c ioc.Dic, config featurepkg.Config) {
		featurepkg.RegisterFieldProvider(config, func() PlayerCTX { return ProvidedPlayerCTX })

		{
			f := featurepkg.RegisterFeature[AttackFeature](config)
			featurepkg.RegisterRuleToField(config, &f, &f.Attacker, OwnsRule{
				feature.NewFieldGetter(&f, &f.PlayerCTX),
				feature.NewFieldGetter(&f, &f.Attacker),
			})
			featurepkg.RegisterRuleToField(config, &f, &f.Attacked, CanAttackRule{
				feature.NewFieldGetter(&f, &f.Attacker),
				feature.NewFieldGetter(&f, &f.Attacked),
			})
		}
		{
			f := featurepkg.RegisterFeature[MoveToFeature](config)
			featurepkg.RegisterRuleToField(config, &f, &f.From, OwnsRule{
				feature.NewFieldGetter(&f, &f.PlayerCTX),
				feature.NewFieldGetter(&f, &f.From),
			})
		}
		{
			f := featurepkg.RegisterFeature[FocusEnemyFeature](config)
			featurepkg.RegisterRuleToField(config, &f, &f.Enemy, EnemyRule{
				feature.NewFieldGetter(&f, &f.PlayerCTX),
				feature.NewFieldGetter(&f, &f.Enemy),
			})
		}
	})
})

type World struct {
	World         engine.EngineWorld                 `inject:""`
	EntityField   feature.FieldService[EntityField]  `inject:""`
	OwnsRule      feature.RuleService[OwnsRule]      `inject:""`
	CanAttackRule feature.RuleService[CanAttackRule] `inject:""`
	EnemyRule     feature.RuleService[EnemyRule]     `inject:""`
	Config        featurepkg.Config                  `inject:""`

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
	allFields := w.World.Feature().AllFields(expected)
	if !expected.LockedFeature {
		if len(allFields) != 0 {
			w.T.Fatalf("expected [] from AllFields and got %v", allFields)
		}
	} else if expected.LockedFeature {
		featureKey := expected.Features[0]
		featureConfig := w.Config.GetFeatureConfig(featureKey)
		expectedAllFields := []any{}
		interactions := slices.Clone(expected.Interactions)
		if expected.PreviewInteraction != nil {
			interactions = append(interactions, expected.PreviewInteraction)
		}
		for len(interactions) != 0 {
			fieldConfig := featureConfig.Fields[len(expectedAllFields)]
			if fieldConfig.Provide != nil {
				expectedAllFields = append(expectedAllFields, fieldConfig.Provide())
				continue
			}
			expectedAllFields = append(expectedAllFields, interactions[0])
			interactions = interactions[1:]
		}
		if !reflect.DeepEqual(allFields, expectedAllFields) {
			w.T.Fatalf("expected %v from AllFields and got %v", allFields, expectedAllFields)
		}
	}

	// verify world state
	expectedComponents := make(map[any]struct{})

	var featureConfig internal.FeatureConfig
	if expected.LockedFeature {
		featureKey := expected.Features[0]
		featureConfig = w.Config.GetFeatureConfig(featureKey)
	}
	// for every field we want to render [feature.FieldRenderInteraction] with
	// [feature.RuleRenderCompletePreview] for every rule
	previewFieldIndex := len(allFields)
	if expected.LockedFeature && expected.PreviewInteraction != nil {
		previewFieldIndex--
	}
	for i, field := range allFields[:previewFieldIndex] {
		fieldConfig := featureConfig.Fields[i]
		if fieldConfig.Provide != nil {
			continue
		}
		comp := feature.NewFieldRender(field.(EntityField), feature.FieldRenderInteraction)
		expectedComponents[comp] = struct{}{}
		for _, ruleConfig := range fieldConfig.Rules {
			rule := w.Rules[ruleConfig.RuleType]
			if err := rule.AppliesAny(ruleConfig.Rule, allFields); err != nil {
				w.T.Fatal(err.Error())
			}
			comp := rule.NewAnyRenderComponent(ruleConfig.Rule, allFields, feature.RuleRenderCompletePreview)
			expectedComponents[fmt.Sprintf("%v", comp)] = struct{}{}
		}
	}
	if !expected.LockedFeature {
		for _, field := range expected.Interactions {
			comp := feature.NewFieldRender(field.(EntityField), feature.FieldRenderInteraction)
			expectedComponents[comp] = struct{}{}
		}
	}

	// for preview field we want [feature.FieldRenderPreviewRulesPass] or [feature.FieldRenderPreviewRulesFail]
	// with [feature.RuleRenderPreviewRulesPass] or [feature.RuleRenderPreviewRulesFail]
	if expected.LockedFeature && expected.PreviewInteraction != nil {
		ruleRender := feature.RuleRenderPreviewRulesPass
		fieldRender := feature.FieldRenderPreviewRulesPass
		for _, ruleConfig := range featureConfig.Fields[previewFieldIndex].Rules {
			rule := w.Rules[ruleConfig.RuleType]
			if err := rule.AppliesAny(ruleConfig.Rule, allFields); err != nil {
				ruleRender = feature.RuleRenderPreviewRulesFail
				fieldRender = feature.FieldRenderPreviewRulesFail
				break
			}
		}
		comp := feature.NewFieldRender(expected.PreviewInteraction.(EntityField), fieldRender)
		expectedComponents[comp] = struct{}{}
		for _, ruleConfig := range featureConfig.Fields[previewFieldIndex].Rules {
			rule := w.Rules[ruleConfig.RuleType]
			comp := rule.NewAnyRenderComponent(ruleConfig.Rule, allFields, ruleRender)
			expectedComponents[fmt.Sprintf("%v", comp)] = struct{}{}
		}
	} else if expected.LockedFeature && expected.PreviewInteraction == nil {
		comp := feature.NewFieldRender(EntityField{}, feature.FieldRenderMissing)
		expectedComponents[comp] = struct{}{}
		for _, ruleConfig := range featureConfig.Fields[previewFieldIndex].Rules {
			rule := w.Rules[ruleConfig.RuleType]
			comp := rule.NewAnyRenderComponent(ruleConfig.Rule, allFields, feature.RuleRenderPreviewRulesFail)
			expectedComponents[fmt.Sprintf("%v", comp)] = struct{}{}
		}
	}

	// if there is no preview we want to render [feature.FieldRenderMissing]
	if expected.LockedFeature && expected.PreviewInteraction == nil {
		comp := feature.NewFieldRender(EntityField{}, feature.FieldRenderMissing)
		expectedComponents[comp] = struct{}{}
	}

	// verify is this actual
	arrays := []ecs.AnyComponentArray{
		w.EntityField.RenderAny(),
	}
	stringArrays := []ecs.AnyComponentArray{
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
	for _, array := range stringArrays {
		for _, entity := range array.GetEntities() {
			comp, _ := array.GetAny(entity)
			foundComponents[fmt.Sprintf("%v", comp)] = struct{}{}
		}
	}

	if !maps.Equal(expectedComponents, foundComponents) {
		w.T.Fatalf("for %v expected\n%v\nand found\n%v",
			expected, expectedComponents, foundComponents)
	}
}
