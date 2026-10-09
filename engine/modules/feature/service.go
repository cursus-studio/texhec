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
	// is used when [InteractionsComponent.PreviewInteraction] rules pass
	FieldRenderPreviewRulesPass
	// is used when [InteractionsComponent.PreviewInteraction] rules fail
	FieldRenderPreviewRulesFail
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

type FieldGetter[Field any] struct {
	Index int
}

func NewFieldGetter[Feature, Field any](feature *Feature, field *Field) FieldGetter[Field] {
	return FieldGetter[Field]{GetFieldIndex(feature, field)}
}
func (getter FieldGetter[Field]) From(fields AllFields) (zero Field) {
	if len(fields) <= getter.Index {
		return zero
	}
	field, _ := fields[getter.Index].(Field)
	return field
}

//

// rule should be composed from field getters
type Rule interface {
	AppliedToField() int
}
type RuleRenderComponent[RuleType Rule] struct {
	Rule   RuleType
	Fields AllFields
	Render RuleRender
}

func NewRuleRender[RuleType Rule](rule RuleType, fields AllFields, render RuleRender) RuleRenderComponent[RuleType] {
	return RuleRenderComponent[RuleType]{rule, fields, render}
}

type AnyRuleService interface {
	RenderAny() ecs.AnyComponentArray
	NewAnyRenderComponent(rule any, fields AllFields, render RuleRender) any

	AppliesAny(any, AllFields) error
	Error() error
}
type RuleService[RuleType Rule] interface {
	AnyRuleService
	Render() ecs.ComponentArray[RuleRenderComponent[RuleType]]

	Applies(RuleType, AllFields) error
}

//
// feature

// is deduced from interactions
// if feature isn't locked it is empty
// if feature is locked it contains all [InteractionsComponent.Interactions] with
// populated fields in between and [InteractionsComponent.PreviewInteraction]
type AllFields = []any

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
	AllFields(InteractionsComponent) AllFields
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
