package actions

import (
	"engine/modules/feature"
	"engine/modules/grid"
	"engine/modules/uuid"
)

//
// fields

type PlayerContext uuid.UUID
type Coords grid.Coords
type Entity uuid.UUID
type Blueprint uuid.UUID

//

type PlayerContextGetter interface {
	Context() PlayerContext
}

func (ctx PlayerContext) Context() PlayerContext { return ctx }

//
// rules

type OwnRule feature.Rule[PlayerContext, Entity]
type IsEnemyRule feature.Rule[PlayerContext, Entity]

// build
type CanDeployRule feature.NoInputRule[Entity]
type CanDeployBlueprintRule feature.Rule[Entity, Blueprint]
type CanDeployAtRule feature.Rule[Blueprint, Coords]

// move
type CanMoveRule feature.NoInputRule[Entity]
type CanMoveToRule feature.Rule[Entity, Coords]

// attack
type CanAttackRule feature.NoInputRule[Entity]
type CanBeAttackedRule feature.Rule[Entity, Entity]

//

type Service interface {
	// fields
	PlayerContext() feature.FieldService[PlayerContext]
	Coords() feature.FieldService[Coords]
	Entity() feature.FieldService[Entity]
	Blueprint() feature.FieldService[Blueprint]

	// rules
	Own() feature.RuleService[OwnRule, PlayerContext, Entity]
	IsEnemy() feature.RuleService[IsEnemyRule, PlayerContext, Entity]

	CanDeploy() feature.NoInputRuleService[CanDeployRule, Entity]
	CanDeployBlueprint() feature.RuleService[CanDeployBlueprintRule, Entity, Blueprint]
	CanDeployAt() feature.RuleService[CanDeployAtRule, Blueprint, Coords]

	CanMove() feature.NoInputRuleService[CanMoveRule, Entity]
	CanMoveTo() feature.RuleService[CanMoveToRule, Entity, Coords]

	CanAttack() feature.NoInputRuleService[CanAttackRule, Entity]
	CanBeAttacked() feature.RuleService[CanBeAttackedRule, Entity, Entity]
}
