package internal

import (
	"engine/modules/focus"
	"engine/services/ecs"
	"fmt"
)

func (s *service) DefaultFocused() ecs.ComponentsArray[focus.DefaultFocusedComponent] {
	return s.defaultFocused
}
func (s *service) Focused() ecs.ComponentsArray[focus.FocusedComponent] {
	return s.focused
}
func (s *service) FocusedEntity() (ecs.EntityID, bool) {
	focusedEntities := s.focused.GetEntities()
	if len(focusedEntities) == 0 {
		focusedEntities = s.defaultFocused.GetEntities()
	}
	if len(focusedEntities) > 1 {
		s.Logger().Log(fmt.Errorf("expected most one focused element"))
		for _, focusedEntity := range focusedEntities {
			s.Focused().Remove(focusedEntity)
		}
	}

	if len(focusedEntities) == 0 {
		return 0, false
	}

	return focusedEntities[0], true
}

func (s *service) NewFocusedBubbleEvent(event any) (focus.BubbleEvent, bool) {
	focusedEntity, ok := s.FocusedEntity()
	if !ok {
		return focus.BubbleEvent{}, false
	}
	return focus.NewBubbleEvent(focusedEntity, event), true
}
