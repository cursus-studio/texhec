// this module is responsible for create foundations for creating in game GUI
package ui

import (
	"engine/services/ecs"
)

// marker which says module relative to which element to position
type UiCameraComponent struct{}
type AnimatedBackgroundComponent struct{}
type CursorCameraComponent struct{}

// selection group events
type UnselectEvent[Component any] struct{}
type SelectEvent[Component any] struct{ Entities []ecs.EntityID }

func NewUnselect[Component any]() UnselectEvent[Component] {
	return UnselectEvent[Component]{}
}
func NewSelect[Component any](entities ...ecs.EntityID) SelectEvent[Component] {
	return SelectEvent[Component]{entities}
}

//

// groups selected elements with component and allows to remove all of them at once
// [SelectionGroup] differs from [ecs.ComponentsArray] that it listens to extra events
type SelectionGroup[Component any] ecs.ComponentsArray[Component]

type ObjectComponent struct{}
type ActionComponent struct{}

//

type Service interface {
	ecs.SystemRegister

	UiCamera() ecs.ComponentsArray[UiCameraComponent]
	AnimatedBackground() ecs.ComponentsArray[AnimatedBackgroundComponent]
	CursorCamera() ecs.ComponentsArray[CursorCameraComponent]

	Objects() SelectionGroup[ObjectComponent]
	Actions() SelectionGroup[ActionComponent]

	// returns parent to attach ui elements
	// potentially with enter animation
	ShowMenu() (parents []ecs.EntityID)
	// removes all children
	HideMenu()
}
