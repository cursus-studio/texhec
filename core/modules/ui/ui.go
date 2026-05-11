package ui

import (
	"engine/services/ecs"
)

type System ecs.SystemRegister

// marker which says module relative to which element to position
type UiCameraComponent struct{}
type AnimatedBackgroundComponent struct{}
type CursorCameraComponent struct{}

type SelectionGroup interface {
	HideOnUnselect(ecs.EntityID)
	Unselect()
	UnselectEvent() any
	OnUnselect(func())
}

type Service interface {
	UiCamera() ecs.ComponentsArray[UiCameraComponent]
	AnimatedBackground() ecs.ComponentsArray[AnimatedBackgroundComponent]
	CursorCamera() ecs.ComponentsArray[CursorCameraComponent]

	Objects() SelectionGroup
	Actions() SelectionGroup

	// returns parent to attach ui elements
	// potentially with enter animation
	ShowMenu() (parents []ecs.EntityID)
	// removes all children
	HideMenu()
}
