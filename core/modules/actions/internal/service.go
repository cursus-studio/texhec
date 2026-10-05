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

	CanBuildService          feature.NoInputRuleService[actions.CanBuildRule, actions.Entity]                      `inject:""`
	CanBuildBlueprintService feature.RuleService[actions.CanBuildBlueprintRule, actions.Entity, actions.Blueprint] `inject:""`
	CanPlaceService          feature.RuleService[actions.CanPlaceRule, actions.Blueprint, actions.Coords]          `inject:""`

	CanMoveService   feature.NoInputRuleService[actions.CanMoveRule, actions.Entity]            `inject:""`
	CanMoveToService feature.RuleService[actions.CanMoveToRule, actions.Entity, actions.Coords] `inject:""`

	CanAttackService     feature.NoInputRuleService[actions.CanAttackRule, actions.Entity]              `inject:""`
	CanBeAttackedService feature.RuleService[actions.CanBeAttackedRule, actions.Entity, actions.Entity] `inject:""`
}

func NewService(c ioc.Dic) actions.Service {
	s := ioc.GetServices[*service](c)

	s.CanBuildBlueprint().Render().OnUpsert(s.OnCanBuildBlueprintUpsert)
	events.Listen(s.EventsBuilder(), s.OnBlueprintClick)

	events.Listen(s.EventsBuilder(), s.OnCoordsHover)
	events.Listen(s.EventsBuilder(), s.OnCoordsClick)

	s.Entity().Render().OnUpsert(s.OnEntityUpsert)
	events.Listen(s.EventsBuilder(), s.OnEntityClick)
	events.Listen(s.EventsBuilder(), s.OnEntityHover)

	// s.CoordsInteractionService.MissingPreview().OnUpsert(s.OnCoordsMissingUpsert)
	// s.CoordsInteractionService.StatePreview().OnUpsert(s.OnCoordsStateUpsert)
	// events.Listen(s.EventsBuilder(), s.OnTileHover)
	// events.Listen(s.EventsBuilder(), s.OnTileClick)
	//
	// s.ObjectInteractionService.MissingPreview().OnUpsert(s.OnObjectMissingUpsert)
	// s.ObjectInteractionService.StatePreview().OnUpsert(s.OnObjectStateUpsert)
	// events.Listen(s.EventsBuilder(), s.OnClickObject)
	//
	// s.BlueprintInteractionService.MissingPreview().OnUpsert(s.OnBlueprintMissingUpsert)
	// s.BlueprintInteractionService.MissingPreview().OnRemove(s.OnBlueprintMissingRemove)
	// s.BlueprintInteractionService.StatePreview().OnUpsert(s.OnBlueprintStateUpsert)
	// events.Listen(s.EventsBuilder(), s.OnClickBlueprint)

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

func (s *service) CanBuild() feature.NoInputRuleService[actions.CanBuildRule, actions.Entity] {
	return s.CanBuildService
}
func (s *service) CanBuildBlueprint() feature.RuleService[actions.CanBuildBlueprintRule, actions.Entity, actions.Blueprint] {
	return s.CanBuildBlueprintService
}
func (s *service) CanPlace() feature.RuleService[actions.CanPlaceRule, actions.Blueprint, actions.Coords] {
	return s.CanPlaceService
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
