package internal

import (
	"engine/modules/focus"
	"engine/services/ecs"

	"github.com/ogiusek/events"
)

func (s *service) Bubbling() ecs.ComponentsArray[focus.BubblingComponent] {
	return s.bubbling
}

func (s *service) DryRun(bubbleEvent focus.BubbleEvent) (bubbles []ecs.EntityID, captured bool) {
	focusedEntity := bubbleEvent.Entity
	parents := append([]ecs.EntityID{focusedEntity}, s.Hierarchy().GetOrderedParents(focusedEntity)...)

	for _, capturingEntity := range parents {
		comp, ok := s.Bubbling().Get(capturingEntity)
		if !ok {
			continue
		}
		if _, ok := comp.CapturesEvents().GetIndex(bubbleEvent.EventType); !ok {
			continue
		}
		bubbles = append(bubbles, capturingEntity)
		if !comp.Fallthrough() {
			return bubbles, true
		}
	}
	return bubbles, false
}

// emits wrapper events and event if it isn't captured
func (s *service) Emit(bubbleEvent focus.BubbleEvent) {
	bubbles, captured := s.DryRun(bubbleEvent)
	for _, bubble := range bubbles {
		comp, ok := s.bubbling.Get(bubble)
		if !ok {
			continue
		}
		eventToEmit := comp.Capture(bubbleEvent.Event)
		if eventToEmit == nil {
			continue
		}
		if apply, ok := eventToEmit.(ecs.ApplyEntityEvent); ok {
			eventToEmit = apply.ApplyEntity(bubble)
		}
		events.EmitAny(s.Events(), eventToEmit)
	}
	if !captured {
		events.EmitAny(s.Events(), bubbleEvent.Event)
	}
}
