// allows us to drag objects and/or camera
package drag

import (
	"engine/modules/ecs"
	"engine/modules/inputs"
)

type DraggableEvent struct {
	Entity ecs.EntityID
	Drag   inputs.DragEvent
}

func NewDraggable(
	entity ecs.EntityID,
) DraggableEvent {
	return DraggableEvent{
		Entity: entity,
	}
}

func (e DraggableEvent) ApplyDrag(dragEvent inputs.DragEvent) any {
	e.Drag = dragEvent
	return e
}

type Service interface {
	ecs.SystemRegister
}
