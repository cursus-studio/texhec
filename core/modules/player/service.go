package player

import (
	"engine/modules/ecs"
	"engine/modules/uuid"
	"errors"
)

var (
	ErrRequiresOwner     error = errors.New("player:requires owner")
	ErrRequiresControl   error = errors.New("player:requires control over player")
	ErrRequiresToBeEnemy error = errors.New("player:requires to be enemy")
)

// marks that player is performing a move
type PlayerComponent struct {
	Name string
}
type ActingPlayerComponent struct{}

func NewPlayer(name string) PlayerComponent  { return PlayerComponent{name} }
func NewActingPlayer() ActingPlayerComponent { return ActingPlayerComponent{} }

type OwnerLink struct{}

//

type Service interface {
	Player() ecs.ComponentArray[PlayerComponent]
	ActingPlayer() ecs.ComponentArray[ActingPlayerComponent]
	Owner() uuid.LinkService[OwnerLink]

	// returns nil if object is controled
	ControlsObject(ecs.EntityID) error
}
