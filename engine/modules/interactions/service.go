package interactions

import (
	"engine/modules/ecs"
	"reflect"
	"slices"
)

type FeatureKey = reflect.Type
type StepKey = reflect.Type
type InteractionKey = reflect.Type

// interaction
type StatePreviewComponent[State any] struct{ State State }
type MissingPreviewComponent[State any] struct{}

func NewStatePreview[State any](state State) StatePreviewComponent[State] {
	return StatePreviewComponent[State]{state}
}
func NewMissingPreview[State any]() MissingPreviewComponent[State] {
	return MissingPreviewComponent[State]{}
}

type InteractionService[State any] interface {
	StatePreview() ecs.ComponentArray[StatePreviewComponent[State]]
	MissingPreview() ecs.ComponentArray[MissingPreviewComponent[State]]

	// saves state in entity with [IsMissingComponent] or resets interactions
	Save(propertiesEntity ecs.EntityID, state State)
}

// step
type Step[State any] interface {
	State() State
}
type stepT[State any] struct {
	Value State
}

func NewStep[State any](state State) Step[State] {
	return stepT[State]{state}
}
func (step stepT[State]) State() State { return step.Value }

// feature
type Feature interface {
	Event() any
}
type AvailableFeaturesComponent struct {
	Features []FeatureKey
	Selected bool
}
type SelectFeatureEvent struct{ FeatureKey FeatureKey }

func NewSelectedFeature(feature FeatureKey) AvailableFeaturesComponent {
	return AvailableFeaturesComponent{[]FeatureKey{feature}, true}
}
func NewAvailableFeatures(features ...FeatureKey) AvailableFeaturesComponent {
	return AvailableFeaturesComponent{features, false}
}
func NewSelectFeatureEvent(featureKey FeatureKey) SelectFeatureEvent {
	return SelectFeatureEvent{featureKey}
}
func NewDeselectFeatureEvent() SelectFeatureEvent { return SelectFeatureEvent{} }

func (c AvailableFeaturesComponent) Equal(other AvailableFeaturesComponent) bool {
	return slices.Equal(c.Features, other.Features)
}

// service
type Service interface {
	Features() []FeatureKey
	AvailableFeatures() ecs.ComponentArray[AvailableFeaturesComponent]
}
