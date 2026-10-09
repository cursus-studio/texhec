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

type OwnRule struct {
	PlayerContext feature.FieldGetter[PlayerContext]
	Entity        feature.FieldGetter[Entity]
}
type IsEnemyRule struct {
	PlayerContext feature.FieldGetter[PlayerContext]
	Entity        feature.FieldGetter[Entity]
}

// build
type CanDeployRule struct {
	Entity feature.FieldGetter[Entity]
}
type CanDeployBlueprintRule struct {
	Entity    feature.FieldGetter[Entity]
	Blueprint feature.FieldGetter[Blueprint]
}
type CanDeployAtRule struct {
	Entity    feature.FieldGetter[Entity]
	Blueprint feature.FieldGetter[Blueprint]
	Coords    feature.FieldGetter[Coords]
}

// move
type CanMoveRule struct {
	Entity feature.FieldGetter[Entity]
}
type CanMoveToRule struct {
	Entity feature.FieldGetter[Entity]
	Coords feature.FieldGetter[Coords]
}

// attack
type CanAttackRule struct {
	Entity feature.FieldGetter[Entity]
}
type CanBeAttackedRule struct {
	Attacker feature.FieldGetter[Entity]
	Attacked feature.FieldGetter[Entity]
}

//

type Service interface {
	// fields
	PlayerContext() feature.FieldService[PlayerContext]
	Coords() feature.FieldService[Coords]
	Entity() feature.FieldService[Entity]
	Blueprint() feature.FieldService[Blueprint]

	// rules
	Own() feature.RuleService[OwnRule]
	IsEnemy() feature.RuleService[IsEnemyRule]

	CanDeploy() feature.RuleService[CanDeployRule]
	CanDeployBlueprint() feature.RuleService[CanDeployBlueprintRule]
	CanDeployAt() feature.RuleService[CanDeployAtRule]

	CanMove() feature.RuleService[CanMoveRule]
	CanMoveTo() feature.RuleService[CanMoveToRule]

	CanAttack() feature.RuleService[CanAttackRule]
	CanBeAttacked() feature.RuleService[CanBeAttackedRule]
}
