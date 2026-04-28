package internal

import (
	"core/game"
	"core/modules/player"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld `inject:""`

	owner ecs.ComponentsArray[player.OwnerComponent]
}

func NewService(c ioc.Dic) player.Service {
	s := ioc.GetServices[*service](c)
	s.owner = ecs.GetComponentsArray[player.OwnerComponent](s.World())
	return s
}

func (s *service) Owner() ecs.ComponentsArray[player.OwnerComponent] {
	return s.owner
}
