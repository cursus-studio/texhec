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
	// is used when value is empty and we need to render it
	MissingField
	// is used when [InteractionsComponent].PreviewInteraction rules fail
	PreviewRulesFail
	// is used when [InteractionsComponent].PreviewInteraction rules pass
	PreviewRulesPass
	// is used for each Interaction in [InteractionsComponent].Interactions
	Interaction
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

type Rule[DataIn any, AppliedTo any] any
type RuleRenderComponent[RuleType Rule[DataIn, AppliedTo], DataIn any, AppliedTo any] struct {
	In        DataIn
	AppliedTo AppliedTo
}

func NewRuleRender[RuleType Rule[DataIn, AppliedTo], DataIn any, AppliedTo any](
	in DataIn, appliedTo AppliedTo) RuleRenderComponent[RuleType, DataIn, AppliedTo] {
	return RuleRenderComponent[RuleType, DataIn, AppliedTo]{in, appliedTo}
}

type AnyRuleService interface {
	RenderAny() ecs.AnyComponentArray
	NewAnyRenderComponent(in, appliedTo any) any

	AppliesToAny(dataIn, appliedTo any) error
}

type RuleService[RuleType Rule[DataIn, AppliedTo], DataIn any, AppliedTo any] interface {
	AnyRuleService
	Render() ecs.ComponentArray[RuleRenderComponent[RuleType, DataIn, AppliedTo]]

	Applies(DataIn, AppliedTo) error
}

//
// feature

type InteractionsComponent struct {
	Interactions       []any
	PreviewInteraction any
	Features           []FeatureKey
	LockedFeature      bool
}
type LockFeatureEvent struct {
	// to unlock use empty value
	Feature FeatureKey
}

func NewInteractions(interactions []any, previewInteraction any, features []FeatureKey, lockFeature bool) InteractionsComponent {
	return InteractionsComponent{interactions, previewInteraction, features, lockFeature}
}
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

	PreviewInteraction(field any) error // error is returned when field type isn't registered
	SubmitPreviewInteraction(fieldType FieldKey)
}

type Service interface {
	InteractionsService

	ecs.SystemRegister
	Features() []FeatureKey
}
