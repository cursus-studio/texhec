package internal

import (
	"core/game"
	"core/modules/economy"
	"core/modules/player"
	"engine/modules/ecs"
	"engine/modules/uuid"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld `inject:""`
	OwnerLink      uuid.LinkService[player.OwnerLink] `inject:""`

	player       ecs.ComponentArray[player.PlayerComponent]
	actingPlayer ecs.ComponentArray[player.ActingPlayerComponent]
}

func NewService(c ioc.Dic) player.Service {
	s := ioc.GetServices[*service](c)
	s.player = ecs.GetComponentArray[player.PlayerComponent](s.World())
	s.actingPlayer = ecs.GetComponentArray[player.ActingPlayerComponent](s.World())

	s.player.OnUpsert(s.OnPlayerUpsert)
	return s
}

func (s *service) Player() ecs.ComponentArray[player.PlayerComponent] {
	return s.player
}
func (s *service) ActingPlayer() ecs.ComponentArray[player.ActingPlayerComponent] {
	return s.actingPlayer
}
func (s *service) Owner() uuid.LinkService[player.OwnerLink] {
	return s.OwnerLink
}

func (s *service) ControlsObject(entity ecs.EntityID) error {
	owner, ok := s.Owner().Get(entity)
	if !ok {
		return player.ErrRequiresOwner
	}
	if _, ok := s.ActingPlayer().Get(owner); !ok {
		return player.ErrRequiresControl
	}
	return nil
}

func (s *service) OnPlayerUpsert(player ecs.EntityID) {
	worldGenerationEntity, ok := s.Seed().WorldSeed()
	if !ok {
		return
	}
	s.Hierarchy().SetParent(player, worldGenerationEntity)
	s.UUID().Component().Set(player, uuid.New(s.UUID().NewUUID()))
	s.Economy().Wallet().Set(player, economy.NewWallet(0))
}
