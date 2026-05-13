package inputs

import (
	"engine/services/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardEvent struct {
	sdl.KeyboardEvent
}

func NewKeyboardEvent(e sdl.KeyboardEvent) KeyboardEvent {
	return KeyboardEvent{e}
}

// if you want to get on click always despite what is focused then listen to [sdl.KeyboardEvent]

// focuses default entity like scene or camera
type UnfocusEvent struct{}

// unfocuses all elements and only focuses specific one
type FocusEvent struct{ Entity ecs.EntityID }
type DefaultFocusEvent struct{ Entity ecs.EntityID }

func NewUnfocusEvent() UnfocusEvent                              { return UnfocusEvent{} }
func NewFocusEvent(entity ecs.EntityID) FocusEvent               { return FocusEvent{entity} }
func NewDefaultFocusEvent(entity ecs.EntityID) DefaultFocusEvent { return DefaultFocusEvent{entity} }

// on keyboardEvent capture events are emitted
// from child with [FocusedComponent]
// to uppermost parent with [CaptureKeyboardComponent] with fallthrough == false
// if none [CaptureKeyboardComponent] stops further emission then [KeyboardEvent] is emited
type CaptureKeyboardConstraint interface {
	Capture(KeyboardEvent) any
	Fallthrough() bool
}
type CaptureKeyboardComponent struct {
	CaptureKeyboardConstraint
}

func NewCaptureKeyboard(event CaptureKeyboardConstraint) CaptureKeyboardComponent {
	return CaptureKeyboardComponent{event}
}

//

type DefaultFocusedComponent struct{}

func NewDefaultFocused() DefaultFocusedComponent { return DefaultFocusedComponent{} }

// element should be focused on click for example
// on right click or escape element should get unfocused
type FocusedComponent struct{}

func NewFocused() FocusedComponent { return FocusedComponent{} }
