package internal

import (
	"core/game"
	"core/modules/actions"
	"engine/modules/ecs"
	"engine/modules/interactions"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld              `inject:""`
	CoordsInteractionService    interactions.InteractionService[actions.CoordsInteraction]    `inject:""`
	ObjectInteractionService    interactions.InteractionService[actions.ObjectInteraction]    `inject:""`
	BlueprintInteractionService interactions.InteractionService[actions.BlueprintInteraction] `inject:""`

	canDeploy    ecs.ComponentArray[actions.CanDeployComponent]
	coordsCursor ecs.ComponentArray[actions.CoordsCursorComponent]
	coordsAnchor ecs.ComponentArray[actions.CoordsAnchorComponent]
}

func NewService(c ioc.Dic) actions.Service {
	s := ioc.GetServices[*service](c)
	s.canDeploy = ecs.GetComponentArray[actions.CanDeployComponent](s.World())
	s.coordsCursor = ecs.GetComponentArray[actions.CoordsCursorComponent](s.World())
	s.coordsAnchor = ecs.GetComponentArray[actions.CoordsAnchorComponent](s.World())

	s.CoordsInteractionService.MissingPreview().OnUpsert(s.OnCoordsMissingUpsert)
	s.CoordsInteractionService.StatePreview().OnUpsert(s.OnCoordsStateUpsert)
	events.Listen(s.EventsBuilder(), s.OnTileHover)
	events.Listen(s.EventsBuilder(), s.OnTileClick)

	s.ObjectInteractionService.MissingPreview().OnUpsert(s.OnObjectMissingUpsert)
	s.ObjectInteractionService.StatePreview().OnUpsert(s.OnObjectStateUpsert)
	events.Listen(s.EventsBuilder(), s.OnClickObject)

	s.BlueprintInteractionService.MissingPreview().OnUpsert(s.OnBlueprintMissingUpsert)
	s.BlueprintInteractionService.MissingPreview().OnRemove(s.OnBlueprintMissingRemove)
	s.BlueprintInteractionService.StatePreview().OnUpsert(s.OnBlueprintStateUpsert)
	events.Listen(s.EventsBuilder(), s.OnClickBlueprint)

	return s
}

func (s *service) CanDeploy() ecs.ComponentArray[actions.CanDeployComponent] {
	return s.canDeploy
}
func (s *service) CoordsAnchor() ecs.ComponentArray[actions.CoordsAnchorComponent] {
	return s.coordsAnchor
}
func (s *service) CoordsCursor() ecs.ComponentArray[actions.CoordsCursorComponent] {
	return s.coordsCursor
}
