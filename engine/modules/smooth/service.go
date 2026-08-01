// catches changes on tick and applies them smoothly between ticks
package smooth

import (
	"engine/modules/ecs"
	"engine/modules/transition"
)

type AddWaypointEvent[StateComponent any] struct {
	Entity ecs.EntityID
	State  StateComponent
}

func NewAddWaypointEvent[StateComponent any](entity ecs.EntityID, state StateComponent) AddWaypointEvent[StateComponent] {
	return AddWaypointEvent[StateComponent]{entity, state}
}

//

type ServiceT[StateComponent any] interface {
	// AddWaypoint appends a state snapshot for the next tick interval.
	// Multiple waypoints within a single tick period are distributed evenly across the frame duration.
	AddWaypoint(ecs.EntityID, StateComponent)
}

type Service interface {
	Start() ecs.SystemRegister
	Stop() ecs.SystemRegister
}

type SmoothConstraint[Component any] interface {
	transition.LerpConstraint[Component]
	// this method is a tag that component is smoothed
	// each lerpable component with this method will automatically be registered to be smoothed
	Smooth()
}
