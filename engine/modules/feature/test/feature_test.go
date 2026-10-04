package test

import (
	"engine/modules/feature"
	enginepkg "engine/pkg"
	"reflect"
	"testing"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

func TestFeatures(t *testing.T) {
	emitted := []AttackFeature{}
	world := ioc.GetServices[World](ioc.NewContainer(enginepkg.Pkg, Pkg, func(b ioc.Builder) {
		ioc.Wrap(b, func(c ioc.Dic, b events.Builder) {
			events.Listen(b, func(f AttackFeature) {
				emitted = append(emitted, f)
			})
		})
	}))
	world.Rules = make(map[feature.RuleKey]feature.AnyRuleService)
	world.Rules[reflect.TypeFor[OwnsRule]()] = world.OwnsRule
	world.Rules[reflect.TypeFor[EnemyRule]()] = world.EnemyRule
	world.Rules[reflect.TypeFor[CanAttackRule]()] = world.CanAttackRule
	world.T = t

	// initial state
	world.AssertInteractions(feature.NewInteractions(nil, nil, nil, false))

	// add preview
	invalidAttackerfield := NewEntityField(EnemyPlayerCTX.ID)
	if err := world.World.Feature().PreviewInteraction(invalidAttackerfield); err != nil {
		t.Fatal(err)
	}
	world.AssertInteractions(feature.NewInteractions(nil, invalidAttackerfield, nil, false))

	// submit preview
	world.World.Feature().SubmitPreviewInteraction(reflect.TypeFor[EntityField]())
	world.AssertInteractions(feature.NewInteractions(
		[]any{invalidAttackerfield}, nil,
		[]feature.FeatureKey{reflect.TypeFor[FocusEnemyFeature]()}, false))

	// add preview second time
	attackerfield := NewEntityField(ProvidedPlayerCTX.ID)
	if err := world.World.Feature().PreviewInteraction(attackerfield); err != nil {
		t.Fatal(err)
	}
	world.AssertInteractions(feature.NewInteractions(
		[]any{invalidAttackerfield}, attackerfield,
		[]feature.FeatureKey{reflect.TypeFor[FocusEnemyFeature]()}, false))

	// submit preview and expect only last to be accepted
	world.World.Feature().SubmitPreviewInteraction(reflect.TypeFor[EntityField]())
	world.AssertInteractions(feature.NewInteractions(
		[]any{attackerfield}, nil,
		[]feature.FeatureKey{reflect.TypeFor[AttackFeature](), reflect.TypeFor[MoveToFeature]()}, false))

	// lock feature
	events.Emit(world.World.Events(), feature.NewLockFeatureEvent(reflect.TypeFor[AttackFeature]()))
	world.AssertInteractions(feature.NewInteractions(
		[]any{attackerfield}, nil,
		[]feature.FeatureKey{reflect.TypeFor[AttackFeature]()}, true))

	// add preview for invalid field
	invalidAttackedField := NewEntityField(ProvidedPlayerCTX.ID)
	if err := world.World.Feature().PreviewInteraction(invalidAttackedField); err != nil {
		t.Fatal(err)
	}
	world.AssertInteractions(feature.NewInteractions(
		[]any{attackerfield}, invalidAttackedField,
		[]feature.FeatureKey{reflect.TypeFor[AttackFeature]()}, true))

	// expect invalid field to be rejected
	world.World.Feature().SubmitPreviewInteraction(reflect.TypeFor[EntityField]())
	world.AssertInteractions(feature.NewInteractions(
		[]any{attackerfield}, nil,
		[]feature.FeatureKey{reflect.TypeFor[AttackFeature]()}, true))

	// add preview
	attackedField := NewEntityField(EnemyPlayerCTX.ID)
	if err := world.World.Feature().PreviewInteraction(attackedField); err != nil {
		t.Fatal(err)
	}
	world.AssertInteractions(feature.NewInteractions(
		[]any{attackerfield}, attackedField, []feature.FeatureKey{reflect.TypeFor[AttackFeature]()}, true))

	// expect event to be emitted after submission
	if len(emitted) != 0 {
		t.Fatalf("do not expected emitted event")
	}
	world.World.Feature().SubmitPreviewInteraction(reflect.TypeFor[EntityField]())
	world.AssertInteractions(feature.NewInteractions(nil, nil, nil, false))
	expectedEmitted := AttackFeature{ProvidedPlayerCTX, attackerfield, attackedField}
	if len(emitted) != 1 || emitted[0] != expectedEmitted {
		t.Fatalf("expected [%v] not %v", expectedEmitted, emitted)
	}
}
