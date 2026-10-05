package player

import (
	"core/modules/actions"
	"engine/modules/connection"
	"engine/modules/ecs"
	"engine/modules/uuid"
)

// marks that player is performing a move
type PlayerComponent struct {
	Name string
}
type PlayerUUIDComponent struct {
	UUID uuid.UUID
}
type ActingPlayerComponent struct{}
type ActingConnectionComponent struct{ Connection ecs.EntityID }

func NewPlayer(name string) PlayerComponent            { return PlayerComponent{name} }
func NewPlayerUUID(uuid uuid.UUID) PlayerUUIDComponent { return PlayerUUIDComponent{uuid} }
func NewActingPlayer() ActingPlayerComponent           { return ActingPlayerComponent{} }
func NewActiongConnection(entity ecs.EntityID) ActingConnectionComponent {
	return ActingConnectionComponent{entity}
}

type OwnerLink struct{}

//

type AssignActingPlayerDTO struct {
	connection.MsgCtx
	PlayerUUID uuid.UUID
}

func NewAssignActingPlayerDTO(playerUUID uuid.UUID) AssignActingPlayerDTO {
	return AssignActingPlayerDTO{PlayerUUID: playerUUID}
}

//

type PlayersConnectionComponent struct{}
type PlayerConnectionComponent struct {
	Player uuid.UUID
}

func NewPlayersConnection() PlayersConnectionComponent { return PlayersConnectionComponent{} }
func NewPlayerConnection(player uuid.UUID) PlayerConnectionComponent {
	return PlayerConnectionComponent{player}
}

//

type Service interface {
	// attached to player
	Player() ecs.ComponentArray[PlayerComponent]
	PlayerUUID() ecs.ComponentArray[PlayerUUIDComponent]
	ActingPlayer() ecs.ComponentArray[ActingPlayerComponent]
	ActingConnection() ecs.ComponentArray[ActingConnectionComponent]

	// attached to host and connection
	PlayersConnection() ecs.ComponentArray[PlayersConnectionComponent]
	PlayerConnection() ecs.ComponentArray[PlayerConnectionComponent]

	// points to player entity
	Owner() uuid.LinkService[OwnerLink]

	// returns nil if object is controled
	Own(player actions.PlayerContext, entity actions.Entity) error
	IsEnemy(player actions.PlayerContext, entity actions.Entity) error

	GetContext() actions.PlayerContext
}
