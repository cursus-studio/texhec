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
	OwnService     feature.RuleService[actions.OwnRule]     `inject:""`
	IsEnemyService feature.RuleService[actions.IsEnemyRule] `inject:""`

	CanDeployService          feature.RuleService[actions.CanDeployRule]          `inject:""`
	CanDeployBlueprintService feature.RuleService[actions.CanDeployBlueprintRule] `inject:""`
	CanDeployAtService        feature.RuleService[actions.CanDeployAtRule]        `inject:""`

	CanMoveService   feature.RuleService[actions.CanMoveRule]   `inject:""`
	CanMoveToService feature.RuleService[actions.CanMoveToRule] `inject:""`

	CanAttackService     feature.RuleService[actions.CanAttackRule]     `inject:""`
	CanBeAttackedService feature.RuleService[actions.CanBeAttackedRule] `inject:""`
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
func (s *service) Own() feature.RuleService[actions.OwnRule] {
	return s.OwnService
}
func (s *service) IsEnemy() feature.RuleService[actions.IsEnemyRule] {
	return s.IsEnemyService
}

func (s *service) CanDeploy() feature.RuleService[actions.CanDeployRule] {
	return s.CanDeployService
}
func (s *service) CanDeployBlueprint() feature.RuleService[actions.CanDeployBlueprintRule] {
	return s.CanDeployBlueprintService
}
func (s *service) CanDeployAt() feature.RuleService[actions.CanDeployAtRule] {
	return s.CanDeployAtService
}

func (s *service) CanMove() feature.RuleService[actions.CanMoveRule] {
	return s.CanMoveService
}
func (s *service) CanMoveTo() feature.RuleService[actions.CanMoveToRule] {
	return s.CanMoveToService
}

func (s *service) CanAttack() feature.RuleService[actions.CanAttackRule] {
	return s.CanAttackService
}
func (s *service) CanBeAttacked() feature.RuleService[actions.CanBeAttackedRule] {
	return s.CanBeAttackedService
}
