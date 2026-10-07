package feature

import (
	"engine/modules/ecs"
	"errors"
	"reflect"
	"slices"
)

var (
	ErrInvalidType error = errors.New("feature:invalid type")
)

type FieldKey = reflect.Type
type RuleKey = reflect.Type
type FeatureKey = reflect.Type

//
// field

type FieldRender uint8

const (
	_ FieldRender = iota
	// is used for each Interaction in [InteractionsComponent.Interactions]
	FieldRenderInteraction
	// is used when value is empty and we need to render it
	FieldRenderMissing
	// is used when [InteractionsComponent.PreviewInteraction] rules fail
	FieldRenderPreviewRulesFail
	// is used when [InteractionsComponent.PreviewInteraction] rules pass
	FieldRenderPreviewRulesPass
)

type ProvideFieldService[FieldType any] func() FieldType

type FieldRenderComponent[FieldType any] struct {
	Value  FieldType
	Render FieldRender
}

func NewFieldRender[FieldType any](value FieldType, render FieldRender) FieldRenderComponent[FieldType] {
	return FieldRenderComponent[FieldType]{value, render}
}

type AnyFieldService interface {
	RenderAny() ecs.AnyComponentArray
	NewAnyRenderComponent(value any, fieldRender FieldRender) any
}
type FieldService[FieldType any] interface {
	AnyFieldService
	Render() ecs.ComponentArray[FieldRenderComponent[FieldType]]
}

//
// rule

type RuleRender uint8

const (
	_ RuleRender = iota
	// is used for rules for [InteractionsComponent.Interactions]
	RuleRenderCompletePreview
	// is used for rules for [InteractionsComponent.PreviewInteraction] if all rules pass
	RuleRenderPreviewRulesPass
	// is used for rules for [InteractionsComponent.PreviewInteraction] if any rules fail
	RuleRenderPreviewRulesFail
)

type NoInputRule[AppliedTo any] = Rule[AppliedTo, AppliedTo]
type Rule[DataIn any, AppliedTo any] any
type RuleRenderComponent[RuleType Rule[DataIn, AppliedTo], DataIn any, AppliedTo any] struct {
	In        DataIn
	AppliedTo AppliedTo
	Render    RuleRender
}

func NewRuleRender[RuleType Rule[DataIn, AppliedTo], DataIn any, AppliedTo any](
	in DataIn, appliedTo AppliedTo, render RuleRender) RuleRenderComponent[RuleType, DataIn, AppliedTo] {
	return RuleRenderComponent[RuleType, DataIn, AppliedTo]{in, appliedTo, render}
}

type AnyRuleService interface {
	RenderAny() ecs.AnyComponentArray
	NewAnyRenderComponent(in, appliedTo any, render RuleRender) any

	AppliesToAny(dataIn, appliedTo any) error
	Error() error
}
type RuleService[RuleType Rule[DataIn, AppliedTo], DataIn any, AppliedTo any] interface {
	AnyRuleService
	Render() ecs.ComponentArray[RuleRenderComponent[RuleType, DataIn, AppliedTo]]

	Applies(DataIn, AppliedTo) error
}
type NoInputRuleService[RuleType NoInputRule[AppliedTo], AppliedTo any] = RuleService[RuleType, AppliedTo, AppliedTo]

//
// feature

type InteractionsComponent struct {
	Interactions       []any
	PreviewInteraction any
	Features           []FeatureKey
	LockedFeature      bool
}
type RenderComponent struct{}
type LockFeatureEvent struct {
	// to unlock use empty value
	Feature FeatureKey
}

func NewInteractions(interactions []any, previewInteraction any, features []FeatureKey, lockFeature bool) InteractionsComponent {
	return InteractionsComponent{interactions, previewInteraction, features, lockFeature}
}
func NewRender() RenderComponent                                 { return RenderComponent{} }
func NewLockFeatureEvent(featureKey FeatureKey) LockFeatureEvent { return LockFeatureEvent{featureKey} }
func NewUnlockFeatureEvent() LockFeatureEvent                    { return LockFeatureEvent{} }

func (c InteractionsComponent) Equal(other InteractionsComponent) bool {
	return slices.Equal(c.Interactions, other.Interactions) &&
		c.PreviewInteraction == other.PreviewInteraction &&
		slices.Equal(c.Features, other.Features) &&
		c.LockedFeature == other.LockedFeature
}

type InteractionsService interface {
	Interactions() InteractionsComponent
	OnInteractionsMod(func(ecs.EntityID))

	SubmitInteraction(field any)
	PreviewInteraction(field any)
	SubmitPreviewInteraction(fieldType FieldKey)
}

type Service interface {
	InteractionsService

	ecs.SystemRegister
	Features() []FeatureKey
}
