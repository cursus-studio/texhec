package uiservice

import (
	"core/game"
	"core/modules/ui"
	"engine/modules/loop"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type selectionGroup[Component any] struct {
	game.GameWorld `inject:""`
	arr            ecs.ComponentsArray[Component]
	selected       ecs.ComponentsArray[selectedComponent[Component]]
}

type selectedComponent[Component any] struct{}

func NewSelectionGroup[Component any](c ioc.Dic, service *service) ui.SelectionGroup[Component] {
	s := ioc.GetServices[*selectionGroup[Component]](c)
	s.arr = ecs.GetComponentsArray[Component](s.World())
	s.selected = ecs.GetComponentsArray[selectedComponent[Component]](s.World())

	system := ecs.NewSystemRegister(func() error {
		events.Listen(s.EventsBuilder(), s.OnTick)
		events.Listen(s.EventsBuilder(), s.OnSelect)
		events.Listen(s.EventsBuilder(), s.OnUnselect)
		return nil
	})
	service.systems = append(service.systems, system)

	return s.arr
}

func (s *selectionGroup[Component]) OnTick(e loop.TickEvent) {
	for _, entity := range s.selected.GetEntities() {
		events.Emit(s.Events(), ui.NewSelectTick[Component](e, entity))
	}
}
func (s *selectionGroup[Component]) OnSelect(e ui.SelectEvent[Component]) {
	s.selected.Set(e.Entity, selectedComponent[Component]{})
}
func (s *selectionGroup[Component]) OnUnselect(e ui.UnselectEvent[Component]) {
	for _, entity := range s.selected.GetEntities() {
		s.selected.Remove(entity)
	}
	for _, entity := range s.arr.GetEntities() {
		s.World().RemoveEntity(entity)
	}
}
