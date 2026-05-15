package ui

import (
	"engine/modules/loop"
	"engine/services/ecs"
)

// marker which says module relative to which element to position
type UiCameraComponent struct{}
type AnimatedBackgroundComponent struct{}
type CursorCameraComponent struct{}

// selection group events
type UnselectEvent[Component any] struct{}
type SelectEvent[Component any] struct{ Entities []ecs.EntityID }

// each tick is emited with currently selected entity
type SelectTickEvent[Component any] struct {
	Tick     loop.TickEvent
	Entities []ecs.EntityID
}

func NewUnselect[Component any]() UnselectEvent[Component] {
	return UnselectEvent[Component]{}
}
func NewSelect[Component any](entities ...ecs.EntityID) SelectEvent[Component] {
	return SelectEvent[Component]{entities}
}
func NewSelectTick[Component any](tick loop.TickEvent, entities []ecs.EntityID) SelectTickEvent[Component] {
	return SelectTickEvent[Component]{tick, entities}
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
