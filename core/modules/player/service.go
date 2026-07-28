package player

import (
	"engine/modules/ecs"
	"errors"
)

var (
	ErrRequiresOwner     error = errors.New("player:requires owner")
	ErrRequiresControl   error = errors.New("player:requires control over player")
	ErrRequiresToBeEnemy error = errors.New("player:requires to be enemy")
)

// marks that player is performing a move
type ActingPlayerComponent struct{}
type OwnerComponent struct {
	Owner ecs.EntityID
}

func NewActingPlayer() ActingPlayerComponent     { return ActingPlayerComponent{} }
func NewOwner(owner ecs.EntityID) OwnerComponent { return OwnerComponent{owner} }

//

type Service interface {
	ActingPlayer() ecs.ComponentArray[ActingPlayerComponent]
	Owner() ecs.ComponentArray[OwnerComponent]

	// returns nil if object is controled
	ControlsObject(ecs.EntityID) error
}
