// integrates inputs into the engine
//
// defines mouse and keyboard high-level abstractions.
//
// ### mouse
// directly parses mouse inputs to specific entity click events
//
// ### keyboard
// takes keyboard inputs and spreads them using signals which can be captured by entities in focused element hierarchy
package inputs

import (
	"engine/modules/collider"
	"engine/modules/ecs"
)

type Service interface {
	ecs.SystemRegister

	TextInput() ecs.ComponentArray[TextInputComponent]

	Hovered() ecs.ComponentArray[HoveredComponent]
	Dragged() ecs.ComponentArray[DraggedComponent]
	Stacked() ecs.ComponentArray[StackedComponent]

	KeepSelected() ecs.ComponentArray[KeepSelectedComponent]

	LeftClick() ecs.ComponentArray[LeftClickComponent]
	DoubleLeftClick() ecs.ComponentArray[DoubleLeftClickComponent]

	RightClick() ecs.ComponentArray[RightClickComponent]
	DoubleRightClick() ecs.ComponentArray[DoubleRightClickComponent]

	MouseEnter() ecs.ComponentArray[MouseEnterComponent]
	MouseLeave() ecs.ComponentArray[MouseLeaveComponent]

	Hover() ecs.ComponentArray[HoverComponent]
	Drag() ecs.ComponentArray[DragComponent]

	Stack() ecs.ComponentArray[StackComponent]

	// returns ordered targets with additional data
	StackedData() []Target
}

type EventTargetSetter interface {
	SetTarget(Target) EventTargetSetter
}

type Target struct {
	collider.ObjectRayCollision
	Camera ecs.EntityID
}
