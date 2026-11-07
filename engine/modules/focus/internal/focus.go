package internal

import (
	"engine/modules/ecs"
	"engine/modules/focus"
)

func (s *service) DefaultFocused() ecs.ComponentArray[focus.DefaultFocusedComponent] {
	return s.defaultFocused
}
func (s *service) Focused() ecs.ComponentArray[focus.FocusedComponent] {
	return s.focused
}

func (s *service) NewFocusedBubbleEvent(event any) (focus.BubbleEvent, bool) {
	focusedEntity, _, ok := s.Focused().GetSingleton()
	if !ok {
		focusedEntity, _, ok = s.DefaultFocused().GetSingleton()
	}
	if !ok {
		return focus.BubbleEvent{}, false
	}
	return focus.NewBubbleEvent(focusedEntity, event), true
}
