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
type stepT[StepT Step[State], State any] struct {
	Value State
}

func NewStepT[StepT Step[State], State any](state State) StepT {
	return any(stepT[StepT, State]{state}).(StepT)
}
func (step stepT[StepT, State]) State() State { return step.Value }

// feature
type Event = any
type AvailableFeaturesComponent struct{ Features []FeatureKey }
type SelectFeatureEvent struct{ FeatureKey FeatureKey }

func NewAvailableFeatures(features ...FeatureKey) AvailableFeaturesComponent {
	return AvailableFeaturesComponent{features}
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
