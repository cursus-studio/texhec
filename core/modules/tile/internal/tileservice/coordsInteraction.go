package tileservice

import (
	"core/modules/definitions"
	"core/modules/obstruction"
	"core/modules/tile"
	"engine/modules/grid"
	"engine/modules/interactions"
	"engine/modules/render"
	"engine/modules/transform"
	"engine/services/ecs"

	"github.com/go-gl/mathgl/mgl32"
)

func (s *service) CoordsInteraction() interactions.InteractionService[tile.CoordsInteraction] {
	return s.CoordsInteractionService
}

// on grid hover everything
func (s *service) OnTileHover(e grid.HoverEvent) {
	featureEntity := s.Interactions().FeatureEntity()
	if _, ok := s.CoordsInteractionService.MissingInteraction().Get(featureEntity); !ok {
		return
	}

	chunkCoords, ok := s.EngineWorld.Grid().Coords().Get(e.Chunk)
	if !ok {
		return
	}
	coords := s.EngineWorld.Grid().AbsoluteCoords(chunkCoords, e.Coords)
	interaction := interactions.NewInteraction(tile.CoordsInteraction{Coords: coords})

	s.CoordsInteractionService.Interaction().Set(featureEntity, interaction)
}

// on tile click
func (s *service) OnTileClick(e grid.ClickEvent) {
	featureEntity := s.Interactions().FeatureEntity()
	if _, ok := s.CoordsInteractionService.Interaction().Get(featureEntity); !ok {
		return
	}
	s.CoordsInteractionService.MissingInteraction().Remove(featureEntity)
}

// on interactions.CoordsInteraction mod
func (s *service) OnCoordsInteractionUpsert(entity ecs.EntityID) {
	worldEntity, ok := s.Seed().WorldSeed()
	if !ok {
		return
	}

	placeholders := s.CoordsInteractionService.InteractionGUI().GetEntities()
	if len(placeholders) > 1 {
		for _, entity := range s.CoordsInteractionService.InteractionGUI().GetEntities()[1:] {
			s.World().RemoveEntity(entity)
		}
	}
	featureEntity := s.Interactions().FeatureEntity()
	coordsCursor, ok := s.coordsCursor.Get(featureEntity)
	if !ok {
		return
	}
	coords, ok := s.CoordsInteractionService.Interaction().Get(featureEntity)
	if !ok {
		return
	}

	var placeholderEntity ecs.EntityID
	if len(placeholders) == 0 {
		placeholderEntity = s.World().NewEntity()
	} else {
		placeholderEntity = placeholders[0]
	}
	s.Prototype().CloneTo(coordsCursor.PropertiesEntity, placeholderEntity)
	s.Obstruction().Deployed().Remove(placeholderEntity)
	s.Collider().Component().Remove(placeholderEntity)

	s.Hierarchy().SetParent(placeholderEntity, worldEntity)
	s.CoordsInteractionService.InteractionGUI().Set(placeholderEntity, interactions.InteractionGUIComponent[tile.CoordsInteraction]{})
	s.Tile().Layer().Set(placeholderEntity, tile.NewLayer(definitions.ObjectPlaceholderLayer))

	pos := tile.NewPos(coords.State.Coords.Coords())
	s.Tile().Pos().Set(placeholderEntity, pos)

	if !coordsCursor.CustomImage {
		s.Render().Texture().Set(placeholderEntity, render.NewTexture(s.Definitions().Hud().Selected))
		// set image to default
	}

	//

	canDeploy := true
	for _, child := range s.Hierarchy().Children(placeholderEntity).GetIndices() {
		s.World().RemoveEntity(child)
	}

	// range

	if coordsCursorRange, ok := s.coordsCursorRange.Get(featureEntity); ok && coordsCursorRange.Entity != 0 {
		for _, coords := range s.GameWorld.Deploy().Reach().TilesWithinReach(coordsCursorRange.Entity) {
			entity := s.Prototype().Clone(s.Definitions().Assets().Blank)
			s.Hierarchy().SetParent(entity, placeholderEntity)
			s.Transform().Parent().Set(entity, transform.NewParent(transform.Absolute))

			s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
			s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Assets().Border))
			s.Groups().InheritGroups(entity)

			s.Tile().Layer().Set(entity, tile.NewLayer(definitions.RangePlaceholderLayer))
			s.Tile().Pos().Set(entity, tile.NewPos(coords.Coords()))
			s.Render().Color().Set(entity, render.NewColor(mgl32.Vec4{0, 0, .5, 1}))
		}

		if !s.GameWorld.Deploy().Reach().Reaches(coordsCursorRange.Entity, placeholderEntity) {
			canDeploy = false
		}
	}

	// collisions
	blueprintObstruction, _ := s.Obstruction().Component().Get(coordsCursor.PropertiesEntity)
	size, _ := s.Tile().Size().Get(coordsCursor.PropertiesEntity)
	aabb := obstruction.NewAABB(pos, size)
	collisions := s.Obstruction().Collisions(aabb, blueprintObstruction.Obstruction)
	for _, collision := range collisions {
		entity := s.Prototype().Clone(s.Definitions().Assets().Blank)
		s.Hierarchy().SetParent(entity, placeholderEntity)
		s.Transform().Parent().Set(entity, transform.NewParent(transform.Absolute))

		s.Tile().Layer().Set(entity, tile.NewLayer(definitions.TilePlaceholderLayer))
		s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Assets().Blank))
		s.Groups().InheritGroups(entity)

		s.Tile().Layer().Set(entity, tile.NewLayer(definitions.TilePlaceholderLayer))
		s.Tile().Pos().Set(entity, tile.NewPos(collision.Coords()))
		s.Render().Color().Set(entity, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
	}
	if len(collisions) != 0 {
		canDeploy = false
		goto showCanDeploy
	}

showCanDeploy:
	if canDeploy {
		s.Render().Color().Set(placeholderEntity, render.NewColor(mgl32.Vec4{0, 1, 0, 1}))
		return
	}
	s.Render().Color().Set(placeholderEntity, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
}
