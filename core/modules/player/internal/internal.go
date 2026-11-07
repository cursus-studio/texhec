package internal

import (
	"core/game"
	"core/modules/actions"
	"core/modules/economy"
	"core/modules/player"
	"engine/modules/ecs"
	"engine/modules/loop"
	"engine/modules/relation"
	"engine/modules/uuid"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld `inject:""`
	Relation       relation.Service[player.NameComponent] `inject:""`
	OwnerLink      uuid.LinkService[player.OwnerLink]     `inject:""`

	name              ecs.ComponentArray[player.NameComponent]
	color             ecs.ComponentArray[player.ColorComponent]
	playerUUID        ecs.ComponentArray[player.PlayerUUIDComponent]
	actingPlayer      ecs.ComponentArray[player.ActingPlayerComponent]
	actingConnection  ecs.ComponentArray[player.ActingConnectionComponent]
	playersConnection ecs.ComponentArray[player.PlayersConnectionComponent]
	playerConnection  ecs.ComponentArray[player.PlayerConnectionComponent]
}

func NewService(c ioc.Dic) player.Service {
	s := ioc.GetServices[*service](c)
	s.name = ecs.GetComponentArray[player.NameComponent](s.World())
	s.color = ecs.GetComponentArray[player.ColorComponent](s.World())
	s.playerUUID = ecs.GetComponentArray[player.PlayerUUIDComponent](s.World())
	s.actingPlayer = ecs.GetComponentArray[player.ActingPlayerComponent](s.World())
	s.actingConnection = ecs.GetComponentArray[player.ActingConnectionComponent](s.World())
	s.playersConnection = ecs.GetComponentArray[player.PlayersConnectionComponent](s.World())
	s.playerConnection = ecs.GetComponentArray[player.PlayerConnectionComponent](s.World())

	s.Name().OnUpsert(s.OnPlayerUpsert)
	s.playerUUID.OnUpsert(s.OnPlayerUUIDUpsert)

	events.Listen(s.EventsBuilder(), s.OnAssignActingPlayerDTO)

	return s
}

func (s *service) NameRelation() relation.Service[player.NameComponent]       { return s.Relation }
func (s *service) Name() ecs.ComponentArray[player.NameComponent]             { return s.name }
func (s *service) Color() ecs.ComponentArray[player.ColorComponent]           { return s.color }
func (s *service) PlayerUUID() ecs.ComponentArray[player.PlayerUUIDComponent] { return s.playerUUID }
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

//

func (s *service) Own(ctx actions.PlayerContext, entityUUID actions.Entity) error {
	entity, ok := s.UUID().Entity(uuid.UUID(entityUUID))
	if !ok {
		return s.Actions().Own().Error()
	}
	owner, ok := s.Owner().Get(entity)
	if !ok {
		return s.Actions().Own().Error()
	}
	ownerUUID, ok := s.UUID().UUID().Get(owner)
	if !ok || ownerUUID.ID != uuid.UUID(ctx) {
		return s.Actions().Own().Error()
	}
	return nil
}
func (s *service) IsEnemy(ctx actions.PlayerContext, entityUUID actions.Entity) error {
	entity, ok := s.UUID().Entity(uuid.UUID(entityUUID))
	if !ok {
		return s.Actions().IsEnemy().Error()
	}
	owner, ok := s.Owner().Get(entity)
	if !ok {
		return nil
	}
	ownerUUID, _ := s.UUID().UUID().Get(owner)
	if ownerUUID.ID == uuid.UUID(ctx) {
		return s.Actions().IsEnemy().Error()
	}
	return nil
}

//

func (s *service) OnPlayerUpsert(playerEntity ecs.EntityID) {
	worldGenerationEntity, _, ok := s.Seed().Seed().GetSingleton()
	if !ok {
		return
	}
	s.Hierarchy().SetParent(playerEntity, worldGenerationEntity)
	if _, ok := s.PlayerUUID().Get(playerEntity); !ok {
		playerComp, _ := s.Name().Get(playerEntity)
		playerUUID := s.UUID().NewDeterministicUUID(fmt.Appendf(nil, "player:%v", playerComp.Name))
		s.PlayerUUID().Set(playerEntity, player.NewPlayerUUID(playerUUID))
	}
	if _, ok := s.Economy().Wallet().Get(playerEntity); !ok {
		s.Economy().Wallet().Set(playerEntity, economy.NewWallet(0))
	}
}

func (s *service) OnPlayerUUIDUpsert(playerEntity ecs.EntityID) {
	playerUUID, ok := s.PlayerUUID().Get(playerEntity)
	if !ok {
		return
	}
	s.UUID().UUID().Set(playerEntity, uuid.NewUUID(playerUUID.UUID))
}

func (s *service) OnAssignActingPlayerDTO(dto player.AssignActingPlayerDTO) {
	if playerEntity, ok := s.UUID().Entity(dto.PlayerUUID); ok {
		s.ActingPlayer().Set(playerEntity, player.NewActingPlayer())
		return
	}
	if _, ok := s.Connection().Connection().Get(dto.ConnEntity); !ok {
		return
	}
	events.Emit(s.Events(), loop.NewEmitOnFrameEvent(dto))
}

func (s *service) GetContext() actions.PlayerContext {
	entities := s.ActingPlayer().GetEntities()
	if len(entities) != 1 {
		return actions.PlayerContext{}
	}
	entity := entities[0]
	uuid, ok := s.UUID().UUID().Get(entity)
	if !ok {
		return actions.PlayerContext{}
	}
	return actions.PlayerContext(uuid.ID)
}
