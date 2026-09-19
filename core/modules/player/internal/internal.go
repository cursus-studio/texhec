package internal

import (
	"core/game"
	"core/modules/economy"
	"core/modules/player"
	"engine/modules/ecs"
	"engine/modules/loop"
	"engine/modules/uuid"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld `inject:""`
	OwnerLink      uuid.LinkService[player.OwnerLink] `inject:""`

	player            ecs.ComponentArray[player.PlayerComponent]
	playerUUID        ecs.ComponentArray[player.PlayerUUIDComponent]
	actingPlayer      ecs.ComponentArray[player.ActingPlayerComponent]
	actingConnection  ecs.ComponentArray[player.ActingConnectionComponent]
	playersConnection ecs.ComponentArray[player.PlayersConnectionComponent]
	playerConnection  ecs.ComponentArray[player.PlayerConnectionComponent]
}

func NewService(c ioc.Dic) player.Service {
	s := ioc.GetServices[*service](c)
	s.player = ecs.GetComponentArray[player.PlayerComponent](s.World())
	s.playerUUID = ecs.GetComponentArray[player.PlayerUUIDComponent](s.World())
	s.actingPlayer = ecs.GetComponentArray[player.ActingPlayerComponent](s.World())
	s.actingConnection = ecs.GetComponentArray[player.ActingConnectionComponent](s.World())
	s.playersConnection = ecs.GetComponentArray[player.PlayersConnectionComponent](s.World())
	s.playerConnection = ecs.GetComponentArray[player.PlayerConnectionComponent](s.World())

	s.player.OnUpsert(s.OnPlayerUpsert)
	s.Connection().Component().OnUpsert(s.OnConnectionUpsert)
	s.playerUUID.OnUpsert(s.OnPlayerUUIDUpsert)

	events.Listen(s.EventsBuilder(), s.OnAssignActingPlayerDTO)

	return s
}

func (s *service) Player() ecs.ComponentArray[player.PlayerComponent] {
	return s.player
}
func (s *service) PlayerUUID() ecs.ComponentArray[player.PlayerUUIDComponent] {
	return s.playerUUID
}
func (s *service) ActingPlayer() ecs.ComponentArray[player.ActingPlayerComponent] {
	return s.actingPlayer
}
func (s *service) ActingConnection() ecs.ComponentArray[player.ActingConnectionComponent] {
	return s.actingConnection
}
func (s *service) PlayersConnection() ecs.ComponentArray[player.PlayersConnectionComponent] {
	return s.playersConnection
}
func (s *service) PlayerConnection() ecs.ComponentArray[player.PlayerConnectionComponent] {
	return s.playerConnection
}
func (s *service) Owner() uuid.LinkService[player.OwnerLink] {
	return s.OwnerLink
}

func (s *service) ControlsEntity(entity ecs.EntityID) error {
	owner, ok := s.Owner().Get(entity)
	if !ok {
		return player.ErrRequiresOwner
	}
	if _, ok := s.ActingPlayer().Get(owner); !ok {
		return player.ErrRequiresControl
	}
	return nil
}

func (s *service) ControlsUUID(uuidVal uuid.UUID) error {
	if entity, ok := s.UUID().Entity(uuidVal); ok {
		return s.ControlsEntity(entity)
	}
	return uuid.ErrMissingUUID
}

func (s *service) OnPlayerUpsert(playerEntity ecs.EntityID) {
	worldGenerationEntity, ok := s.Seed().WorldSeed()
	if !ok {
		return
	}
	s.Hierarchy().SetParent(playerEntity, worldGenerationEntity)
	playerUUID := s.UUID().NewUUID()
	s.PlayerUUID().Set(playerEntity, player.NewPlayerUUID(playerUUID))
	if _, ok := s.Economy().Wallet().Get(playerEntity); !ok {
		s.Economy().Wallet().Set(playerEntity, economy.NewWallet(0))
	}
}

func (s *service) OnPlayerUUIDUpsert(playerEntity ecs.EntityID) {
	playerUUID, ok := s.PlayerUUID().Get(playerEntity)
	if !ok {
		return
	}
	s.UUID().Component().Set(playerEntity, uuid.New(playerUUID.UUID))
}

func (s *service) OnConnectionUpsert(entity ecs.EntityID) {
	conn, ok := s.Connection().Component().Get(entity)
	if !ok {
		return
	}
	parent, ok := s.Hierarchy().Parent(entity)
	if !ok {
		return
	}
	if _, ok := s.PlayersConnection().Get(parent); !ok {
		return
	}
	var playerUUID uuid.UUID
	for _, playerEntity := range s.player.GetEntities() {
		if _, ok := s.actingPlayer.Get(playerEntity); ok {
			continue
		}
		if _, ok := s.actingConnection.Get(playerEntity); ok {
			continue
		}
		uuid, ok := s.UUID().Component().Get(playerEntity)
		if !ok {
			continue
		}
		playerUUID = uuid.ID
		break
	}
	var zeroUUID uuid.UUID
	if playerUUID == zeroUUID {
		return
	}
	s.PlayerConnection().Set(entity, player.NewPlayerConnection(playerUUID))
	if err := conn.Conn().Send(player.NewAssignActingPlayerDTO(playerUUID)); err != nil {
		s.Logger().Warn(err)
	}
}

func (s *service) OnAssignActingPlayerDTO(dto player.AssignActingPlayerDTO) {
	if playerEntity, ok := s.UUID().Entity(dto.PlayerUUID); ok {
		s.ActingPlayer().Set(playerEntity, player.NewActingPlayer())
		return
	}
	if _, ok := s.Connection().Component().Get(dto.ConnEntity); !ok {
		return
	}
	events.Emit(s.Events(), loop.NewEmitOnFrameEvent(dto))
}
