// this module is responsible for create foundations for creating in game GUI
package ui

import (
	"engine/modules/ecs"
)

// marker which says module relative to which element to position
type UiCameraComponent struct{}
type AnimatedBackgroundComponent struct{}
type CursorCameraComponent struct{}

//

type MaxChunksZoomOut = uint32
type GameCameraComponent struct {
	MaxChunksZoomOut MaxChunksZoomOut
}

func NewGameCamera(maxChunksZoomOut MaxChunksZoomOut) GameCameraComponent {
	return GameCameraComponent{maxChunksZoomOut}
}

//

type Service interface {
	ecs.SystemRegister

	UiCamera() ecs.ComponentArray[UiCameraComponent]
	AnimatedBackground() ecs.ComponentArray[AnimatedBackgroundComponent]
	CursorCamera() ecs.ComponentArray[CursorCameraComponent]

	GameCamera() ecs.ComponentArray[GameCameraComponent]

	AddUICamera(sceneParent ecs.EntityID) ecs.EntityID
	AddBgCamera(sceneParent ecs.EntityID) ecs.EntityID

	// returns parent to attach ui elements
	// potentially with enter animation
	ShowMenu() (parents []ecs.EntityID)
	// removes all children
	HideMenu()
}
