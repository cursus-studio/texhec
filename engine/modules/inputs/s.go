package inputs

import (
	"engine/services/ecs"
	"engine/services/media/window"
)

type System ecs.SystemRegister

// this event is called when nothing is dragged
type DragEvent struct {
	Camera   ecs.EntityID
	From, To window.MousePos // from and to is normalized
}

//

// interfaces which can be implemented by events

type ApplyDragEvent interface {
	ApplyDrag(DragEvent) (event any)
}

type ApplyEntityEvent interface {
	ApplyEntity(entityEmitting ecs.EntityID) (event any)
}

//

type SynchronizePositionEvent DragEvent

func (SynchronizePositionEvent) ApplyDrag(dragEvent DragEvent) any {
	return SynchronizePositionEvent(dragEvent)
}
