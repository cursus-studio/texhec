package internal

import (
	"core/modules/deploy"
	"core/modules/tile"
	"engine"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.World `inject:"1"`
	Tile         tile.Service `inject:"1"`
}

func NewService(c ioc.Dic) deploy.Service {
	s := ioc.GetServices[*service](c)

	events.Listen(s.EventsBuilder, s.Deploy)

	return s
}

func (s *service) Deploy(deploy deploy.DeployEvent) {
	// perform verification can you deploy by someone
	// if you cannot than do a flip
	// if deploy.By ? {
	//   log warning (this shouldn't be a button)
	// }
	// pay and perform everything
	deployed := s.Prototype.Clone(deploy.Blueprint)
	s.Tile.Pos().Set(deployed, tile.NewPos(deploy.Coords.Coords()))
}
