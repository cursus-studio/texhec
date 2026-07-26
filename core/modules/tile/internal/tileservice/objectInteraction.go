package tileservice

import (
	"core/modules/definitions"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/interactions"
	"engine/modules/render"
	"engine/modules/transform"
	"fmt"
)

func (s *service) ObjectInteraction() interactions.InteractionService[tile.ObjectInteraction] {
	return s.ObjectInteractionService
}

func (s *service) OnClickObject(event tile.ClickEntityEvent) {
	link, ok := s.Metadata().Link().Get(event.Entity)
	if !ok {
		s.Logger().Warn(fmt.Errorf("cannot click entity which doesn't have original entity"))
		return
	}

	propertiesEntity := s.World().NewEntity()
	s.CanDeploy().Set(propertiesEntity, tile.NewCanDeploy(link.Entity))
	s.CoordsAnchor().Set(propertiesEntity, tile.NewCoordsAnchor(event.Entity))
	s.CoordsCursor().Set(propertiesEntity, tile.NewCoordsCursor(link.Entity, true))

	s.ObjectInteraction().Save(propertiesEntity, tile.NewObjectInteraction(event.Entity))
}

func (s *service) OnObjectMissingUpsert(entity ecs.EntityID) {
	// should find objects on which action can be performed and
	// these object should be highlighted or shown in choose menu
}

//

func (s *service) OnObjectStateUpsert(entity ecs.EntityID) {
	stateComp, ok := s.ObjectInteraction().StatePreview().Get(entity)
	if !ok {
		return
	}

	worldEntity, ok := s.Seed().WorldSeed()
	if !ok {
		return
	}

	targetObject := stateComp.State.Entity
	s.Hierarchy().SetParent(entity, worldEntity)

	if pos, ok := s.Tile().Pos().Get(targetObject); ok {
		s.Tile().Pos().Set(entity, pos)
	}
	if size, ok := s.Tile().Size().Get(targetObject); ok {
		s.Tile().Size().Set(entity, size)
	}
	s.Tile().Rot().Set(entity, tile.NewRot(0))

	s.Transform().Parent().Set(entity, transform.NewParent(transform.Absolute))
	s.Tile().Layer().Set(entity, tile.NewLayer(definitions.ObjectSelectionPlaceholderLayer))
	s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Hud().Can))
	s.Groups().InheritGroups(entity)
}
