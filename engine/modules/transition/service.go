// changes components from stateA to stateB across many frames
package transition

import (
	"engine/modules/ecs"
	"math"
	"time"

	"golang.org/x/exp/constraints"
)

type Service interface {
	ecs.SystemRegister
	Easing() ecs.ComponentArray[EasingComponent]
	EasingFunction() ecs.ComponentArray[EasingFunctionComponent]
}

//

type LerpConstraint[Component any] interface {
	Lerp(Component, float32) Component
}

func Lerp[Number, T constraints.Float](a, b Number, t T) Number {
	return a + Number(t)*(b-a)
}
func LerpInt[Number constraints.Integer, T constraints.Float](a, b Number, t T) Number {
	return Number(math.Round(Lerp(float64(a), float64(b), t)))
}

//

type Progress float32

// type TransitionComponent[Component LerpConstraint[Component]] struct {
type TransitionComponent[Component any] struct {
	From, To Component
	Progress,
	Duration time.Duration
}

// func NewTransition[Component LerpConstraint[Component]](
func NewTransition[Component any](
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
// type TransitionEvent[Component LerpConstraint[Component]] struct {
type TransitionEvent[Component any] struct {
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
