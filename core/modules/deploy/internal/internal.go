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

	events.Listen(s.EventsBuilder, s.Deploy)

	return s
}

func (s *service) Component() ecs.ComponentsArray[deploy.Component] { return s.component }
func (s *service) Link() ecs.ComponentsArray[deploy.LinkComponent]  { return s.link }

func (s *service) Deploy(deploy deploy.DeployEvent) {
	// perform verification can you deploy by someone
	// if you cannot than do a flip
	// if deploy.By ? {
	//   log warning (this shouldn't be a button)
	// }
	// pay and perform everything
	deployed := s.Prototype.Clone(deploy.Blueprint)
	s.Hierarchy.SetParent(deployed, s.Scene.Scene())

	s.Tile.Pos().Set(deployed, tile.NewPos(deploy.Coords.Coords()))
}
