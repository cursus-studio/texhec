package interactions

import (
	"engine/services/ecs"
	"reflect"
)

type Name = string
type Event = any

//

type MissingInteractionComponent[State any] struct{}
type InteractionComponent[State any] struct{ State State }

func NewInteraction[State any](state State) InteractionComponent[State] {
	return InteractionComponent[State]{State: state}
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
	MissingInteraction() ecs.ComponentsArray[MissingInteractionComponent[State]]
	Interaction() ecs.ComponentsArray[InteractionComponent[State]]
	AnyInteractionService
	FinishMeasurement(State) error
}

//

type FeatureEvent[Event any] struct{ Event Event }
type FeatureEventComponent struct{ Event Event }
type FeatureComponent struct{}

func NewFeatureEvent[Event any](event Event) FeatureEvent[Event] { return FeatureEvent[Event]{event} }
func (e *FeatureEvent[Event]) Component() FeatureEventComponent {
	return FeatureEventComponent{e.Event}
}

type AnyFeatureService interface {
	Name() Name
	EventType() reflect.Type
	Interactions() []AnyInteractionService
}

// listens to [FeatureEvent]
type FeatureService[Event any] interface{ AnyFeatureService }
type Service interface {
	// entity with this stores all components with interactions and selected feature
	Feature() ecs.ComponentsArray[FeatureComponent]
	FeatureEvent() ecs.ComponentsArray[FeatureEventComponent]

	FeatureEntity() ecs.EntityID

	// entity here isn't used
	// it's here just to make calling it easier by OnUpsert because only there it should be used
	Proceed(ecs.EntityID)

	// export features matching any interaction already filled
}
