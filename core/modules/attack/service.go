package attack

import (
	"core/modules/actions"
	"engine/modules/ecs"
	"engine/modules/interactions"
)

type AttackEvent struct {
	By     actions.FriendlyBuilderObjectStep
	Target actions.EnemyObjectStep
}

func NewDeployEvent(
	by,
	target ecs.EntityID,
) AttackEvent {
	return AttackEvent{
		interactions.NewStep(actions.NewObjectInteraction(by)),
		interactions.NewStep(actions.NewObjectInteraction(target)),
	}
}

//

type Service interface {
	AttackEvent(AttackEvent)
}
