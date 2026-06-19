package internal

import (
	"core/game"
	"core/modules/player"
	"engine/modules/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld `inject:""`

	owner ecs.ComponentArray[player.OwnerComponent]
}

func NewService(c ioc.Dic) player.Service {
	s := ioc.GetServices[*service](c)
	s.owner = ecs.GetComponentArray[player.OwnerComponent](s.World())
	return s
}

func (s *service) Owner() ecs.ComponentArray[player.OwnerComponent] {
	return s.owner
}
