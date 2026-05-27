package focus

import (
	"engine/services/datastructures"
	"engine/services/ecs"
	"reflect"
)

// bubbling

// captures events that are emitted
// from child with [FocusedComponent]
// to uppermost parent with [BubblingComponent] with fallthrough(T) == false
// if none [BubblingComponent] stops further emission then [Event] is emited
type BubblingConstraint interface {
	// stores a list of events which can be passed to capture
	// this should be a global variable it never should be stored in component
	CapturesEvents() datastructures.SetReader[reflect.Type]

	// should wrap event and return wrapping event.
	Capture(any) any

	// should return a constant
	Fallthrough() bool
}

// This implements bubbling
type BubblingComponent struct {
	BubblingConstraint
}

func NewBubbling(event BubblingConstraint) BubblingComponent {
	return BubblingComponent{event}
}

//

type BubbleEvent struct {
	Entity ecs.EntityID
	// golang generics are to restrictive to use them.
	// this has to use any because propagating it everywhere would require to granural configuration everywhere
	Event     any
	EventType reflect.Type
}

func NewBubbleEvent(entity ecs.EntityID, event any) BubbleEvent {
	return BubbleEvent{entity, event, reflect.TypeOf(event)}
}

//
// focus

// focuses default entity like scene or camera
type UnfocusEvent struct{}

// unfocuses all elements and only focuses specific one
type FocusEvent struct{ Entity ecs.EntityID }
type DefaultFocusEvent struct{ Entity ecs.EntityID }

func NewUnfocusEvent() UnfocusEvent                              { return UnfocusEvent{} }
func NewFocusEvent(entity ecs.EntityID) FocusEvent               { return FocusEvent{entity} }
func NewDefaultFocusEvent(entity ecs.EntityID) DefaultFocusEvent { return DefaultFocusEvent{entity} }

func (FocusEvent) ApplyEntity(entity ecs.EntityID) any { return FocusEvent{entity} }

// element should be focused on click for example
// on right click or escape element should get unfocused
type FocusedComponent struct{}
type DefaultFocusedComponent struct{}

func NewFocused() FocusedComponent               { return FocusedComponent{} }
func NewDefaultFocused() DefaultFocusedComponent { return DefaultFocusedComponent{} }

type Service interface {
	// bubbling
	Bubbling() ecs.ComponentsArray[BubblingComponent]

	DryRun(BubbleEvent) (bubbles []ecs.EntityID, captured bool)
	Emit(BubbleEvent)

	// focus
	DefaultFocused() ecs.ComponentsArray[DefaultFocusedComponent]
	Focused() ecs.ComponentsArray[FocusedComponent]

	FocusedEntity() (ecs.EntityID, bool)
	NewFocusedBubbleEvent(event any) (BubbleEvent, bool)
}
