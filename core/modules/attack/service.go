package attack

import (
	"core/modules/reach"
	"engine/modules/ecs"
	"errors"
)

var (
	ErrCannotAttackEnemyOutOfReach error = errors.New("attack: enemy is out of reach and cannot follow him")
)

type TargetComponent struct {
	Entity ecs.EntityID
}

func NewTarget(target ecs.EntityID) TargetComponent {
	return TargetComponent{target}
}

//

type Service interface {
	ecs.SystemRegister
	Reach() reach.ServiceT[TargetComponent]

	Target() ecs.ComponentArray[TargetComponent]
}
