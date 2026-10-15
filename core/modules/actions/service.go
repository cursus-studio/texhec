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

func NewOwnRule[Feature any](f *Feature, ctx *PlayerContext, entity *Entity) OwnRule {
	return OwnRule{feature.NewFieldGetter(f, ctx), feature.NewFieldGetter(f, entity)}
}
func NewIsEnemyRule[Feature any](f *Feature, ctx *PlayerContext, entity *Entity) IsEnemyRule {
	return IsEnemyRule{feature.NewFieldGetter(f, ctx), feature.NewFieldGetter(f, entity)}
}
func (r OwnRule) AppliedToField() int     { return r.Entity.Index }
func (r IsEnemyRule) AppliedToField() int { return r.Entity.Index }

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

func NewCanDeployRule[Feature any](f *Feature, entity *Entity) CanDeployRule {
	return CanDeployRule{feature.NewFieldGetter(f, entity)}
}
func NewCanDeployBlueprintRule[Feature any](f *Feature, entity *Entity, blueprint *Blueprint) CanDeployBlueprintRule {
	return CanDeployBlueprintRule{feature.NewFieldGetter(f, entity), feature.NewFieldGetter(f, blueprint)}
}
func NewCanDeployAtRule[Feature any](f *Feature, entity *Entity, blueprint *Blueprint, coords *Coords) CanDeployAtRule {
	return CanDeployAtRule{feature.NewFieldGetter(f, entity), feature.NewFieldGetter(f, blueprint), feature.NewFieldGetter(f, coords)}
}
func (r CanDeployRule) AppliedToField() int          { return r.Entity.Index }
func (r CanDeployBlueprintRule) AppliedToField() int { return r.Blueprint.Index }
func (r CanDeployAtRule) AppliedToField() int        { return r.Coords.Index }

// move
type CanMoveRule struct {
	Entity feature.FieldGetter[Entity]
}
type CanMoveToRule struct {
	Entity feature.FieldGetter[Entity]
	Coords feature.FieldGetter[Coords]
}

func NewCanMoveRule[Feature any](f *Feature, entity *Entity) CanMoveRule {
	return CanMoveRule{feature.NewFieldGetter(f, entity)}
}
func NewCanMoveToRule[Feature any](f *Feature, entity *Entity, coords *Coords) CanMoveToRule {
	return CanMoveToRule{feature.NewFieldGetter(f, entity), feature.NewFieldGetter(f, coords)}
}
func (r CanMoveRule) AppliedToField() int   { return r.Entity.Index }
func (r CanMoveToRule) AppliedToField() int { return r.Coords.Index }

// attack
type CanAttackRule struct {
	Entity feature.FieldGetter[Entity]
}
type CanBeAttackedRule struct {
	Attacker feature.FieldGetter[Entity]
	Attacked feature.FieldGetter[Entity]
}

func NewCanAttackRule[Feature any](f *Feature, entity *Entity) CanAttackRule {
	return CanAttackRule{feature.NewFieldGetter(f, entity)}
}
func NewCanBeAttackedRule[Feature any](f *Feature, attacker *Entity, attacked *Entity) CanBeAttackedRule {
	return CanBeAttackedRule{feature.NewFieldGetter(f, attacker), feature.NewFieldGetter(f, attacked)}
}

func (r CanAttackRule) AppliedToField() int     { return r.Entity.Index }
func (r CanBeAttackedRule) AppliedToField() int { return r.Attacked.Index }

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
