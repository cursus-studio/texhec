package internal

import (
	"core/modules/actions"
	"core/modules/definitions"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/groups"
	"engine/modules/interactions"
	"engine/modules/render"
	"engine/modules/transform"
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
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

	linkEntity, ok := s.Tile().GetLink(event.Entity)
	if !ok {
		s.Logger().Warn(fmt.Errorf("cannot click entity which doesn't have original entity"))
		return
	}

	propertiesEntity := s.World().NewEntity()
	s.CanDeploy().Set(propertiesEntity, actions.NewCanDeploy(linkEntity))
	s.CoordsCursor().Set(propertiesEntity, actions.NewCoordsCursor(linkEntity, false))
	s.Anchor().Set(propertiesEntity, actions.NewAnchor(event.Entity))
	s.RegionAnchor().Set(propertiesEntity, actions.NewRegionAnchor(region))

	s.EntityInteraction().Save(propertiesEntity, actions.NewEntityInteraction(event.Entity))
}

func (s *service) OnObjectMissingUpsert(entity ecs.EntityID) {
	// should find objects on which action can be performed and
	// these object should be highlighted or shown in choose menu
	anchor, ok := s.Anchor().Get(entity)
	if !ok {
		return
	}
	for _, reachCoords := range s.GameWorld.Deploy().Reach().TilesWithinReach(anchor.Entity) {
		ind := s.World().NewEntity()
		s.Hierarchy().SetParent(ind, entity)
		s.Transform().Inherit().Set(ind, transform.NewInherit(transform.Absolute))

		s.Render().Mesh().Set(ind, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(ind, render.NewTexture(s.Definitions().Assets().Border))
		s.Groups().Component().Set(ind, groups.EmptyGroups().Enable(definitions.GameGroup))

		s.Tile().Layer().Set(ind, tile.NewLayer(definitions.RangePlaceholderLayer))
		s.Tile().Pos().Set(ind, tile.NewPos(reachCoords.Coords()))
		s.Render().Color().Set(ind, render.NewColor(mgl32.Vec4{0, 0, .5, 1}))
	}
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
