package internal

import (
	"core/game"
	"core/modules/actions"
	"engine/modules/feature"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld       `inject:""`
	PlayerContextService feature.FieldService[actions.PlayerContext] `inject:""`
	CoordsService        feature.FieldService[actions.Coords]        `inject:""`
	EntityService        feature.FieldService[actions.Entity]        `inject:""`
	BlueprintService     feature.FieldService[actions.Blueprint]     `inject:""`

	// rules
	OwnService     feature.RuleService[actions.OwnRule, actions.PlayerContext, actions.Entity]     `inject:""`
	IsEnemyService feature.RuleService[actions.IsEnemyRule, actions.PlayerContext, actions.Entity] `inject:""`

	CanDeployService          feature.NoInputRuleService[actions.CanDeployRule, actions.Entity]                      `inject:""`
	CanDeployBlueprintService feature.RuleService[actions.CanDeployBlueprintRule, actions.Entity, actions.Blueprint] `inject:""`
	CanDeployAtService        feature.RuleService[actions.CanDeployAtRule, actions.Blueprint, actions.Coords]        `inject:""`

	CanMoveService   feature.NoInputRuleService[actions.CanMoveRule, actions.Entity]            `inject:""`
	CanMoveToService feature.RuleService[actions.CanMoveToRule, actions.Entity, actions.Coords] `inject:""`

	CanAttackService     feature.NoInputRuleService[actions.CanAttackRule, actions.Entity]              `inject:""`
	CanBeAttackedService feature.RuleService[actions.CanBeAttackedRule, actions.Entity, actions.Entity] `inject:""`
}

func NewService(c ioc.Dic) actions.Service {
	s := ioc.GetServices[*service](c)

	s.Entity().Render().OnUpsert(s.OnOwnUpsert)
	s.Own().Render().OnUpsert(s.OnOwnUpsert)

	s.IsEnemy().Render().OnUpsert(s.OnEnemyUpsert)
	s.CanBeAttacked().Render().OnUpsert(s.OnEnemyUpsert)

	s.CanMoveTo().Render().OnUpsert(s.OnCanMoveToUpsert)
	s.CanDeployAt().Render().OnUpsert(s.OnCanDeployAtUpsert)

	s.CanDeploy().Render().OnUpsert(s.OnCanDeployUpsert)
	s.CanAttack().Render().OnUpsert(s.OnCanAttackUpsert)

	s.CanDeployBlueprint().Render().OnUpsert(s.OnCanDeployBlueprintUpsert)
	events.Listen(s.EventsBuilder(), s.OnBlueprintClick)

	events.Listen(s.EventsBuilder(), s.OnCoordsHover)
	events.Listen(s.EventsBuilder(), s.OnCoordsClick)

	events.Listen(s.EventsBuilder(), s.OnEntityClick)
	events.Listen(s.EventsBuilder(), s.OnEntityHover)

	return s
}

// fields
func (s *service) PlayerContext() feature.FieldService[actions.PlayerContext] {
	return s.PlayerContextService
}
func (s *service) Coords() feature.FieldService[actions.Coords] {
	return s.CoordsService
}
func (s *service) Entity() feature.FieldService[actions.Entity] {
	return s.EntityService
}
func (s *service) Blueprint() feature.FieldService[actions.Blueprint] {
	return s.BlueprintService
}

// rules
func (s *service) Own() feature.RuleService[actions.OwnRule, actions.PlayerContext, actions.Entity] {
	return s.OwnService
}
func (s *service) IsEnemy() feature.RuleService[actions.IsEnemyRule, actions.PlayerContext, actions.Entity] {
	return s.IsEnemyService
}

func (s *service) CanDeploy() feature.NoInputRuleService[actions.CanDeployRule, actions.Entity] {
	return s.CanDeployService
}
func (s *service) CanDeployBlueprint() feature.RuleService[actions.CanDeployBlueprintRule, actions.Entity, actions.Blueprint] {
	return s.CanDeployBlueprintService
}
func (s *service) CanDeployAt() feature.RuleService[actions.CanDeployAtRule, actions.Blueprint, actions.Coords] {
	return s.CanDeployAtService
}

func (s *service) CanMove() feature.NoInputRuleService[actions.CanMoveRule, actions.Entity] {
	return s.CanMoveService
}
func (s *service) CanMoveTo() feature.RuleService[actions.CanMoveToRule, actions.Entity, actions.Coords] {
	return s.CanMoveToService
}

func (s *service) CanAttack() feature.NoInputRuleService[actions.CanAttackRule, actions.Entity] {
	return s.CanAttackService
}
func (s *service) CanBeAttacked() feature.RuleService[actions.CanBeAttackedRule, actions.Entity, actions.Entity] {
	return s.CanBeAttackedService
}
