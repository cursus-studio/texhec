package player

import (
	"core/modules/actions"
	"engine/modules/connection"
	"engine/modules/ecs"
	"engine/modules/relation"
	"engine/modules/uuid"
	"errors"

	"github.com/go-gl/mathgl/mgl32"
)

var (
	ErrNameCollision = errors.New("player:name collision")
)

// marks that player is performing a move
type NameComponent struct{ Name string }
type ColorComponent struct{ Color mgl32.Vec3 }
type PlayerUUIDComponent struct {
	UUID uuid.UUID
}
type ActingPlayerComponent struct{}
type ActingConnectionComponent struct{ Connection ecs.EntityID }

func NewName(name string) NameComponent                { return NameComponent{name} }
func NewColor(color mgl32.Vec3) ColorComponent         { return ColorComponent{color} }
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

type PlayersConnectionComponent struct {
	AcceptNewConnections bool
}
type PlayerConnectionComponent struct {
	Player uuid.UUID
}

func NewPlayersConnection(acceptnewConnections bool) PlayersConnectionComponent {
	return PlayersConnectionComponent{acceptnewConnections}
}
func NewPlayerConnection(player uuid.UUID) PlayerConnectionComponent {
	return PlayerConnectionComponent{player}
}

//

type Service interface {
	NameRelation() relation.Service[NameComponent]
	// attached to player
	Name() ecs.ComponentArray[NameComponent]
	Color() ecs.ComponentArray[ColorComponent]
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
