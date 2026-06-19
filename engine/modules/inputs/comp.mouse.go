package inputs

import (
	"engine/modules/ecs"
	"engine/modules/window"
)

// many elements can be hovered at once
type HoveredComponent struct {
	Camera ecs.EntityID
}

func NewHovered(camera ecs.EntityID) HoveredComponent {
	return HoveredComponent{camera}
}

//

type DraggedComponent struct {
	Camera ecs.EntityID
}

func NewDragged(camera ecs.EntityID) DraggedComponent {
	return DraggedComponent{camera}
}

//

// keeps element selected even if user drags outside
type KeepSelectedComponent struct{}

// this is special component stating that on click it enables clicking elements below
type StackComponent struct{}

// this means that element got pressed and if not next that isn't going to be pressed
type StackedComponent struct{}

//

type LeftClickComponent struct{ Event any }
type DoubleLeftClickComponent struct{ Event any }

type RightClickComponent struct{ Event any }
type DoubleRightClickComponent struct{ Event any }

type MouseEnterComponent struct{ Event any }
type MouseLeaveComponent struct{ Event any }

type HoverComponent struct{ Event any }
type DragComponent struct{ Event any }

func NewLeftClick(e any) LeftClickComponent { return LeftClickComponent{e} }
func NewDoubleLeftClick(e any) DoubleLeftClickComponent {
	return DoubleLeftClickComponent{e}
}

func NewRightClick(e any) RightClickComponent { return RightClickComponent{e} }
func NewDoubleRightClick(e any) DoubleRightClickComponent {
	return DoubleRightClickComponent{e}
}

func NewMouseEnterComponent(event any) MouseEnterComponent { return MouseEnterComponent{event} }
func NewMouseLeaveComponent(event any) MouseLeaveComponent { return MouseLeaveComponent{event} }

func NewHoverComponent(event any) HoverComponent { return HoverComponent{event} }
func NewDragComponent(event any) DragComponent   { return DragComponent{event} }

//

// this event is called when nothing is dragged
type DragEvent struct {
	Camera   ecs.EntityID
	From, To window.MousePos // from and to is normalized
}

// interfaces which can be implemented by events {
type ApplyDragEvent interface {
	ApplyDrag(DragEvent) (event any)
}

type SynchronizePositionEvent DragEvent

func (SynchronizePositionEvent) ApplyDrag(dragEvent DragEvent) any {
	return SynchronizePositionEvent(dragEvent)
}
