package example

import (
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/interactions"
	interactionspkg "engine/modules/interactions/pkg"
	"unsafe"

	"github.com/ogiusek/ioc/v2"
)

// Feature
// Step
// Interaction

// One way to use:
// - Make an interaction (many steps are applied)
// - Select Feature (also apply step of first interaction)
// - Proceed with filling other Steps

// Second way to use:
// - Select Feature
// - Proceed with filling other Steps

// Needed usage mechanics:
// - Open Interaction
// - Narrowing Feature selection
// - Feature selection
// - Fixed Step

// Edge cases to handle:
// - interaction doesn't meet required step criteria -> show inconsistency and reject interaction

// Interaction mechanics:
// - notify that interaction is missing and which entities match step requirements

// On interaction:
// - select step according to feature
// - preview interaction
// - start and preview next step or execute feature

// example usage:
// interactions: simple user interactions
type UnitInteraction struct {
	Entity ecs.EntityID
}
type CoordInteraction struct {
	Coords grid.Coords
}

type AnchorRuleComponent struct{ Coords grid.Coords }
type RangeRuleComponent struct{ Range int }
type CursorRuleComponent struct{ Cursor ecs.EntityID }

// steps: these are validated interactions
// other examples of steps: friendly builder, friendly army
type BlueprintUnit interactions.Step[UnitInteraction]
type EnemyUnit interactions.Step[UnitInteraction]
type FriendlyUnit interactions.Step[UnitInteraction]
type Coord interactions.Step[CoordInteraction]

// features: these are emited as events.
// Here struct tags have component:"field names"
type MoveFeature struct {
	FriendlyUnit
	Coord
}
type StopFeature struct {
	FriendlyUnit
}
type BuildFeature struct {
	FriendlyUnit
	BlueprintUnit
	Coord
}
type AttackFeature struct {
	FriendlyUnit
	EnemyUnit
}
type HealFeature struct {
	Healer FriendlyUnit
	Healed FriendlyUnit
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		// require most. Do not need configuration but whole implementation and integration with GUI
		interactionspkg.InteractionPkg[UnitInteraction](),
		interactionspkg.InteractionPkg[CoordInteraction](),

		// require argument configuration
		interactionspkg.StepPkg[FriendlyUnit](func(c ioc.Dic) func(state UnitInteraction) error {
			return func(state UnitInteraction) error { return nil /* is friendly */ }
		}),
		interactionspkg.StepPkg[EnemyUnit](func(c ioc.Dic) func(state UnitInteraction) error {
			return func(state UnitInteraction) error { return nil /* !is friendly */ }
		}),
		interactionspkg.StepPkg[Coord](func(c ioc.Dic) func(state CoordInteraction) error {
			return func(state CoordInteraction) error { return nil }
		}),

		// requires relation configuration
		interactionspkg.FeaturePkg[MoveFeature](
			interactionspkg.NewRelation(unsafe.Offsetof(MoveFeature{}.FriendlyUnit), unsafe.Offsetof(MoveFeature{}.Coord),
				func(c ioc.Dic) func(srcEntity ecs.EntityID, tgtEntity ecs.EntityID) {
					return func(sourceEntity, targetEntity ecs.EntityID) {
						// set cursor
					}
				}),
		),
		interactionspkg.FeaturePkg[StopFeature](),
		interactionspkg.FeaturePkg[BuildFeature](
			interactionspkg.NewRelation(unsafe.Offsetof(BuildFeature{}.FriendlyUnit), unsafe.Offsetof(BuildFeature{}.BlueprintUnit),
				func(c ioc.Dic) func(srcEntity ecs.EntityID, tgtEntity ecs.EntityID) {
					return func(sourceEntity, targetEntity ecs.EntityID) {
						// set can_deploy
					}
				}),
			interactionspkg.NewRelation(unsafe.Offsetof(BuildFeature{}.FriendlyUnit), unsafe.Offsetof(BuildFeature{}.Coord),
				func(c ioc.Dic) func(srcEntity ecs.EntityID, tgtEntity ecs.EntityID) {
					return func(sourceEntity, targetEntity ecs.EntityID) {
						// set cursor & anchor & range
					}
				}),
		),
		interactionspkg.FeaturePkg[AttackFeature](
			interactionspkg.NewRelation(unsafe.Offsetof(AttackFeature{}.FriendlyUnit), unsafe.Offsetof(AttackFeature{}.EnemyUnit),
				func(c ioc.Dic) func(srcEntity ecs.EntityID, tgtEntity ecs.EntityID) {
					return func(sourceEntity, targetEntity ecs.EntityID) {
						// set anchor&range
					}
				}),
		),
		interactionspkg.FeaturePkg[HealFeature](
			interactionspkg.NewRelation(unsafe.Offsetof(HealFeature{}.Healer), unsafe.Offsetof(HealFeature{}.Healed),
				func(c ioc.Dic) func(srcEntity ecs.EntityID, tgtEntity ecs.EntityID) {
					return func(sourceEntity, targetEntity ecs.EntityID) {
						// set anchor&range
					}
				}),
		),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	// interactions:
	// all interactions previews will be heavily coupled to ECS framework so just imagine:
	// - if interaction is selected show circle
	// - if interaction is suggested to select show other colored circle
	// - highlighting for certain rules
})
