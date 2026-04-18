package internal

import (
	"core/modules/player"
	gamescenes "core/scenes"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	gamescenes.GameWorld `inject:""`

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
