package internal

import (
	"core/game"
	"core/modules/attack"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld `inject:""`
}

func NewService(c ioc.Dic) attack.Service {
	s := ioc.GetServices[*service](c)

	events.Listen(s.EventsBuilder(), s.AttackEvent)
	return s
}

func (s *service) AttackEvent(event attack.AttackEvent) {
	s.World().RemoveEntity(event.Target)
}
