package internal

import (
	"core/modules/deploy"
	"core/modules/tile"
	"engine"
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

// perform verification can you deploy by someone
// if you cannot than do a flip
// if deploy.By ? {
//   log warning (this shouldn't be a button)
// }
// pay and perform everything

func (s *service) Select(e deploy.SelectEvent) {
	events.Emit(s.Events, tile.NewSelectEvent(deploy.NewExecuteEvent(e.Blueprint)))
}
func (s *service) Preview(e deploy.PreviewEvent) {
	// ???
}
func (s *service) Execute(e deploy.ExecuteEvent) {
	deployed := s.Prototype.Clone(e.Blueprint)
	s.Hierarchy.SetParent(deployed, s.Scene.Scene())

	s.Tile.Pos().Set(deployed, tile.NewPos(e.Coords.Coords()))
}
