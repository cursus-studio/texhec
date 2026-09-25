package player

import (
	"engine/modules/connection"
	"engine/modules/ecs"
	"engine/modules/uuid"
	"errors"
)

var (
	ErrRequiresOwner     error = errors.New("player:requires owner")
	ErrRequiresControl   error = errors.New("player:requires control over player")
	ErrRequiresToBeEnemy error = errors.New("player:requires to be enemy")
)

//

// event context
type PlayerContext struct {
	PlayerUUID uuid.UUID
}

// marks that player is performing a move
type PlayerComponent struct {
	Name string
}
type PlayerUUIDComponent struct {
	UUID uuid.UUID
}
type ActingPlayerComponent struct{}
type ActingConnectionComponent struct{ Connection uuid.UUID }

func NewPlayer(name string) PlayerComponent            { return PlayerComponent{name} }
func NewPlayerUUID(uuid uuid.UUID) PlayerUUIDComponent { return PlayerUUIDComponent{uuid} }
func NewActingPlayer() ActingPlayerComponent           { return ActingPlayerComponent{} }
func NewActiongConnection(uuid uuid.UUID) ActingConnectionComponent {
	return ActingConnectionComponent{uuid}
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
	Player() ecs.ComponentArray[PlayerComponent]
	PlayerUUID() ecs.ComponentArray[PlayerUUIDComponent]
	ActingPlayer() ecs.ComponentArray[ActingPlayerComponent]
	ActingConnection() ecs.ComponentArray[ActingConnectionComponent]

	PlayersConnection() ecs.ComponentArray[PlayersConnectionComponent]
	PlayerConnection() ecs.ComponentArray[PlayerConnectionComponent]

	Owner() uuid.LinkService[OwnerLink]

	// returns nil if object is controled
	ControlsEntity(ecs.EntityID) error
	ControlsUUID(uuid.UUID) error
}
