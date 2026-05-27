package internal

import (
	"engine"
	"engine/modules/focus"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`

	bubbling       ecs.ComponentsArray[focus.BubblingComponent]
	defaultFocused ecs.ComponentsArray[focus.DefaultFocusedComponent]
	focused        ecs.ComponentsArray[focus.FocusedComponent]
}

func NewService(c ioc.Dic) focus.Service {
	s := ioc.GetServices[*service](c)
	s.bubbling = ecs.GetComponentsArray[focus.BubblingComponent](s.World())
	s.defaultFocused = ecs.GetComponentsArray[focus.DefaultFocusedComponent](s.World())
	s.focused = ecs.GetComponentsArray[focus.FocusedComponent](s.World())

	events.Listen(s.EventsBuilder(), s.OnDefaultFocus)
	events.Listen(s.EventsBuilder(), s.OnFocus)
	events.Listen(s.EventsBuilder(), s.OnUnfocus)
	return s
}

func (s *service) OnDefaultFocus(e focus.DefaultFocusEvent) {
	for _, entity := range s.Focus().DefaultFocused().GetEntities() {
		s.Focus().DefaultFocused().Remove(entity)
	}
	s.Focus().DefaultFocused().Set(e.Entity, focus.NewDefaultFocused())
}
func (s *service) OnFocus(e focus.FocusEvent) {
	for _, focusedEntity := range s.Focus().Focused().GetEntities() {
		s.Focus().Focused().Remove(focusedEntity)
	}
	s.Focus().Focused().Set(e.Entity, focus.NewFocused())
}
func (s *service) OnUnfocus(focus.UnfocusEvent) {
	for _, focusedEntity := range s.Focus().Focused().GetEntities() {
		s.Focus().Focused().Remove(focusedEntity)
	}
}
