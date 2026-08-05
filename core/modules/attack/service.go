package attack

import "engine/modules/ecs"

type AttackEvent struct {
	By,
	Target ecs.EntityID
}

func NewAttackEvent(
	by,
	target ecs.EntityID,
) AttackEvent {
	return AttackEvent{
		by,
		target,
	}
}

//

type Service interface {
	AttackEvent(AttackEvent)
}
