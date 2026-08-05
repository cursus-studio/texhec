package internal

import (
	"core/modules/actions"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/interactions"
	"engine/modules/render"
	"engine/modules/transform"
	"fmt"
)

func (s *service) EntityInteraction() interactions.InteractionService[actions.EntityInteraction] {
	return s.ObjectInteractionService
}

func (s *service) OnClickObject(event tile.ClickEntityEvent) {
	if missingEntities := s.EntityInteraction().MissingPreview().GetEntities(); len(missingEntities) == 1 {
		missingEntity := missingEntities[0]
		if anchor, ok := s.Anchor().Get(missingEntity); ok && !s.GameWorld.Deploy().Reach().Reaches(anchor.Entity, event.Entity) {
			s.Logger().Warn(fmt.Errorf("cannot click entity out of range"))
			return
		}
	}

	region, ok := s.Pathfind().EntityRegion(event.Entity)
	if !ok {
		return
	}

	link, ok := s.Metadata().Link().Get(event.Entity)
	if !ok {
		s.Logger().Warn(fmt.Errorf("cannot click entity which doesn't have original entity"))
		return
	}

	propertiesEntity := s.World().NewEntity()
	s.CanDeploy().Set(propertiesEntity, actions.NewCanDeploy(link.Entity))
	s.CoordsCursor().Set(propertiesEntity, actions.NewCoordsCursor(link.Entity, false))
	s.Anchor().Set(propertiesEntity, actions.NewAnchor(event.Entity))
	s.RegionAnchor().Set(propertiesEntity, actions.NewRegionAnchor(region))

	s.EntityInteraction().Save(propertiesEntity, actions.NewEntityInteraction(event.Entity))
}

func (s *service) OnObjectMissingUpsert(entity ecs.EntityID) {
	// should find objects on which action can be performed and
	// these object should be highlighted or shown in choose menu
}

//

func (s *service) OnObjectStateUpsert(entity ecs.EntityID) {
	stateComp, ok := s.EntityInteraction().StatePreview().Get(entity)
	if !ok {
		return
	}

	targetObject := stateComp.State.Entity
	s.Hierarchy().SetParent(entity, targetObject)
	s.Groups().InheritGroups(entity)
	s.Transform().Inherit().Set(entity, transform.NewInherit(transform.RelativePos|transform.RelativeSizeXY))
	s.Transform().Pos().Set(entity, transform.NewPos(0, 0, -1))

	s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Hud().Can))
}
