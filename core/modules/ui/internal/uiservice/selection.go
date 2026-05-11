package uiservice

import (
	"core/game"
	"core/modules/ui"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type UnselectEvent[Component any] struct{}

type selectionGroup[Component any] struct {
	game.GameWorld `inject:""`
	componentArray ecs.ComponentsArray[Component]

	zero Component
}

func NewSelectionGroup[Component any](c ioc.Dic) ui.SelectionGroup {
	s := ioc.GetServices[*selectionGroup[Component]](c)
	s.componentArray = ecs.GetComponentsArray[Component](s.World())
	events.Listen(s.EventsBuilder(), s.unselectListener)
	return s
}

func (s *selectionGroup[Component]) unselectListener(UnselectEvent[Component]) {
	for _, entity := range s.componentArray.GetEntities() {
		s.World().RemoveEntity(entity)
	}
}

func (s *selectionGroup[Component]) HideOnUnselect(entity ecs.EntityID) {
	s.componentArray.Set(entity, s.zero)
}
func (s *selectionGroup[Component]) Unselect() { events.Emit(s.Events(), UnselectEvent[Component]{}) }

func (s *selectionGroup[Component]) UnselectEvent() any { return UnselectEvent[Component]{} }
func (s *selectionGroup[Component]) OnUnselect(listener func()) {
	events.Listen(s.EventsBuilder(), func(UnselectEvent[Component]) {
		listener()
	})
}
