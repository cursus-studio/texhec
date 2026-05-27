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
	"engine/services/ecs"
)

type Service interface {
	ecs.SystemRegister

	TextInput() ecs.ComponentsArray[TextInputComponent]

	Hovered() ecs.ComponentsArray[HoveredComponent]
	Dragged() ecs.ComponentsArray[DraggedComponent]
	Stacked() ecs.ComponentsArray[StackedComponent]

	KeepSelected() ecs.ComponentsArray[KeepSelectedComponent]

	LeftClick() ecs.ComponentsArray[LeftClickComponent]
	DoubleLeftClick() ecs.ComponentsArray[DoubleLeftClickComponent]

	RightClick() ecs.ComponentsArray[RightClickComponent]
	DoubleRightClick() ecs.ComponentsArray[DoubleRightClickComponent]

	MouseEnter() ecs.ComponentsArray[MouseEnterComponent]
	MouseLeave() ecs.ComponentsArray[MouseLeaveComponent]

	Hover() ecs.ComponentsArray[HoverComponent]
	Drag() ecs.ComponentsArray[DragComponent]

	Stack() ecs.ComponentsArray[StackComponent]

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
