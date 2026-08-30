package internal

import (
	"core/game"
	"core/modules/player"
	"engine/modules/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld `inject:""`

	player       ecs.ComponentArray[player.PlayerComponent]
	actingPlayer ecs.ComponentArray[player.ActingPlayerComponent]
	owner        ecs.ComponentArray[player.OwnerComponent]
}

func NewService(c ioc.Dic) player.Service {
	s := ioc.GetServices[*service](c)
	s.player = ecs.GetComponentArray[player.PlayerComponent](s.World())
	s.actingPlayer = ecs.GetComponentArray[player.ActingPlayerComponent](s.World())
	s.owner = ecs.GetComponentArray[player.OwnerComponent](s.World())
	return s
}

func (s *service) Player() ecs.ComponentArray[player.PlayerComponent] {
	return s.player
}
func (s *service) ActingPlayer() ecs.ComponentArray[player.ActingPlayerComponent] {
	return s.actingPlayer
}
func (s *service) Owner() ecs.ComponentArray[player.OwnerComponent] {
	return s.owner
}

func (s *service) ControlsObject(entity ecs.EntityID) error {
	owner, ok := s.Owner().Get(entity)
	if !ok {
		return player.ErrRequiresOwner
	}
	if _, ok := s.ActingPlayer().Get(owner.Owner); !ok {
		return player.ErrRequiresControl
	}
	return nil
}
