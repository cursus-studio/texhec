package transition

import (
	"engine/services/ecs"
	"time"

	"golang.org/x/exp/constraints"
)

type System ecs.SystemRegister

type Service interface {
	Easing() ecs.ComponentsArray[EasingComponent]
	EasingFunction() ecs.ComponentsArray[EasingFunctionComponent]
}

//

type LerpConstraint[Component any] interface {
	Lerp(Component, float32) Component
}

func Lerp[Number, T constraints.Float](a, b Number, t T) Number {
	return a + Number(t)*(b-a)
}

//

type Progress float32

type TransitionComponent[Component LerpConstraint[Component]] struct {
	From, To Component
	Progress,
	Duration time.Duration
}

func NewTransition[Component LerpConstraint[Component]](
	from, to Component,
	duration time.Duration,
) TransitionComponent[Component] {
	return TransitionComponent[Component]{
		From:     from,
		To:       to,
		Progress: 0,
		Duration: duration,
	}
}

//

// saves transition component
type TransitionEvent[Component LerpConstraint[Component]] struct {
	Entity    ecs.EntityID
	Component TransitionComponent[Component]
}

func NewTransitionEvent[Component LerpConstraint[Component]](
	entity ecs.EntityID,
	from, to Component,
	duration time.Duration,
) TransitionEvent[Component] {
	return TransitionEvent[Component]{
		Entity: entity,
		Component: NewTransition(
			from, to,
			duration,
		),
	}
}

//

type DelayedEvent struct {
	Event    any
	Duration time.Duration
}

func NewDelayedEvent(
	event any,
	duration time.Duration,
) DelayedEvent {
	return DelayedEvent{
		Event:    event,
		Duration: duration,
	}
}

//

type EasingComponent struct {
	ID ecs.EntityID
}

type EasingFunctionComponent struct {
	EasingFunction func(t Progress) Progress
}

func NewEasing(id ecs.EntityID) EasingComponent {
	return EasingComponent{id}
}

func NewEasingFunction(easingFunction func(t Progress) Progress) EasingFunctionComponent {
	return EasingFunctionComponent{easingFunction}
}
