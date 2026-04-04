package internal

import (
	"core/modules/deploy"
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/modules/grid"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.World `inject:"1"`
	Tile         tile.Service `inject:"1"`

	component ecs.ComponentsArray[deploy.Component]
	link      ecs.ComponentsArray[deploy.LinkComponent]
}

func NewService(c ioc.Dic) deploy.Service {
	s := ioc.GetServices[*service](c)

	s.component = ecs.GetComponentsArray[deploy.Component](s.World)
	s.link = ecs.GetComponentsArray[deploy.LinkComponent](s.World)

	events.Listen(s.EventsBuilder, s.Execute)
	events.Listen(s.EventsBuilder, s.Preview)
	events.Listen(s.EventsBuilder, s.Select)

	return s
}

func (s *service) Component() ecs.ComponentsArray[deploy.Component] { return s.component }
func (s *service) Link() ecs.ComponentsArray[deploy.LinkComponent]  { return s.link }

func (s *service) Deploy(blueprint ecs.EntityID, coords grid.Coords) {
	// perform verification can you deploy by someone
	// if you cannot than do a flip
	// if deploy.By ? {
	//   log warning (this shouldn't be a button)
	// }
	// do not pay because this is performed by system

	deployed := s.Prototype.Clone(blueprint)
	s.Hierarchy.SetParent(deployed, s.Scene.Scene())

	s.Tile.Pos().Set(deployed, tile.NewPos(coords.Coords()))
	events.Emit(s.Events, ui.HideUiEvent{})
}

func (s *service) Select(e deploy.SelectEvent) {
	events.Emit(s.Events, tile.NewSelectEvent(
		deploy.NewPreviewEvent(e.By, e.Blueprint),
		deploy.NewExecuteEvent(e.By, e.Blueprint),
	))
}
func (s *service) Preview(e deploy.PreviewEvent) {
	byName, ok := s.Metadata.Name().Get(e.By)
	if !ok {
		return
	}
	blueprintName, ok := s.Metadata.Name().Get(e.Blueprint)
	if !ok {
		return
	}
	s.Logger.Info("can %v deploy %v on %v ? show it in gui.", byName.Name, blueprintName.Name, e.Coords)
}
func (s *service) Execute(e deploy.ExecuteEvent) {
	// perform verification can you deploy by someone
	// if you cannot than do a flip
	// if deploy.By ? {
	//   log warning (this shouldn't be a button)
	// }
	// pay and perform everything

	deployed := s.Prototype.Clone(e.Blueprint)
	s.Hierarchy.SetParent(deployed, s.Scene.Scene())

	s.Tile.Pos().Set(deployed, tile.NewPos(e.Coords.Coords()))
	events.Emit(s.Events, ui.HideUiEvent{})
}
