package interactions

import (
	"engine/modules/ecs"
	"reflect"
)

type Name = string
type Event = any

//

type MissingInteractionComponent[State any] struct{}
type InteractionGUIComponent[State any] struct{}
type InteractionComponent[State any] struct{ State State }
type FinishMeasurementEvent[State any] struct{ State State }

func NewInteraction[State any](state State) InteractionComponent[State] {
	return InteractionComponent[State]{State: state}
}
func NewFinishMeasurementEvent[State any](state State) FinishMeasurementEvent[State] {
	return FinishMeasurementEvent[State]{state}
}

type AnyInteractionService interface {
	MissingInteractionAny() ecs.AnyComponentArray
	InteractionAny() ecs.AnyComponentArray

	Name() Name
	StateType() reflect.Type
	// saves [MissingInteractionComponent] if it there is no [InteractionComponent]
	Measure() (alreadyMeasured bool)
}
type InteractionService[State any] interface {
	// elements are removed when interaction is removed.
	// they can be used to indicate that element is used.
	InteractionGUI() ecs.ComponentArray[InteractionGUIComponent[State]]

	MissingInteraction() ecs.ComponentArray[MissingInteractionComponent[State]]
	Interaction() ecs.ComponentArray[InteractionComponent[State]]
	AnyInteractionService
	FinishMeasurement(FinishMeasurementEvent[State])
}

//

type FeatureEvent[Event any] struct{}
type FeatureComponent struct{ Event Event }
type InstanceComponent struct{}

func NewFeature(event Event) FeatureComponent         { return FeatureComponent{event} }
func NewFeatureEvent[Event any]() FeatureEvent[Event] { return FeatureEvent[Event]{} }

type AnyFeatureService interface {
	Name() Name
	EventType() reflect.Type
	Interactions() []AnyInteractionService
}

// listens to [FeatureEvent]
type FeatureService[Event any] interface{ AnyFeatureService }
type Service interface {
	// entity with this stores all components with interactions and selected feature
	Instance() ecs.ComponentArray[InstanceComponent]
	Feature() ecs.ComponentArray[FeatureComponent]

	FeatureEntity() ecs.EntityID

	// entity here isn't used
	// it's here just to make calling it easier by OnUpsert because only there it should be used
	Proceed(ecs.EntityID)

	// export features matching any interaction already filled
}
