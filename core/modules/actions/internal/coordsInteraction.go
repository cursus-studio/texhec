package internal

import (
	"core/modules/actions"
	"core/modules/definitions"
	"core/modules/obstruction"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/interactions"
	"engine/modules/render"
	"engine/modules/transform"

	"github.com/go-gl/mathgl/mgl32"
)

func (s *service) CoordsInteraction() interactions.InteractionService[actions.CoordsInteraction] {
	return s.CoordsInteractionService
}

func (s *service) OnTileClick(e grid.ClickEvent) {
	chunkCoords, ok := s.EngineWorld.Grid().Coords().Get(e.Chunk)
	if !ok {
		return
	}
	coords := s.EngineWorld.Grid().AbsoluteCoords(chunkCoords, e.Coords)
	propertiesEntity := s.World().NewEntity()
	s.CoordsInteraction().Save(propertiesEntity, actions.NewCoordsInteraction(coords))
}

//

func (s *service) OnTileHover(e grid.HoverEvent) {
	entities := s.CoordsInteraction().MissingPreview().GetEntities()
	if len(entities) != 1 {
		return
	}
	previewEntity := entities[0]

	chunkCoords, ok := s.EngineWorld.Grid().Coords().Get(e.Chunk)
	if !ok {
		return
	}
	coords := s.EngineWorld.Grid().AbsoluteCoords(chunkCoords, e.Coords)
	pos := tile.NewPos(coords.Coords())
	s.Tile().Pos().Set(previewEntity, pos)

	for _, child := range s.Hierarchy().Children(previewEntity).GetIndices() {
		s.World().RemoveEntity(child)
	}

	canDeploy := true

	if anchor, ok := s.Anchor().Get(previewEntity); ok {
		for _, reachCoords := range s.GameWorld.Deploy().Reach().TilesWithinReach(anchor.Entity) {
			ind := s.Prototype().Clone(s.Definitions().Assets().Blank)
			s.Hierarchy().SetParent(ind, previewEntity)
			s.Transform().Parent().Set(ind, transform.NewParent(transform.Absolute))

			s.Render().Mesh().Set(ind, render.NewMesh(s.Definitions().Assets().SquareMesh))
			s.Render().Texture().Set(ind, render.NewTexture(s.Definitions().Assets().Border))
			s.Groups().InheritGroups(ind)

			s.Tile().Layer().Set(ind, tile.NewLayer(definitions.RangePlaceholderLayer))
			s.Tile().Pos().Set(ind, tile.NewPos(reachCoords.Coords()))
			s.Render().Color().Set(ind, render.NewColor(mgl32.Vec4{0, 0, .5, 1}))
		}

		if !s.GameWorld.Deploy().Reach().Reaches(anchor.Entity, previewEntity) {
			canDeploy = false
		}
	}

	if coordsCursor, ok := s.CoordsCursor().Get(previewEntity); ok {
		blueprintObstruction, _ := s.Obstruction().Component().Get(coordsCursor.PropertiesEntity)
		size, _ := s.Tile().Size().Get(coordsCursor.PropertiesEntity)
		aabb := obstruction.NewAABB(pos, size)
		collisions := s.Obstruction().Collisions(aabb, blueprintObstruction.Obstruction)

		for _, collision := range collisions {
			ind := s.Prototype().Clone(s.Definitions().Assets().Blank)
			s.Hierarchy().SetParent(ind, previewEntity)
			s.Transform().Parent().Set(ind, transform.NewParent(transform.Absolute))

			s.Tile().Layer().Set(ind, tile.NewLayer(definitions.TilePlaceholderLayer))
			s.Render().Mesh().Set(ind, render.NewMesh(s.Definitions().Assets().SquareMesh))
			s.Render().Texture().Set(ind, render.NewTexture(s.Definitions().Assets().Blank))
			s.Groups().InheritGroups(ind)

			s.Tile().Pos().Set(ind, tile.NewPos(collision.Coords()))
			s.Render().Color().Set(ind, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
		}

		if len(collisions) != 0 {
			canDeploy = false
		}
	}

	if canDeploy {
		s.Render().Color().Set(previewEntity, render.NewColor(mgl32.Vec4{0, 1, 0, 1}))
	} else {
		s.Render().Color().Set(previewEntity, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
	}
}

func (s *service) OnCoordsMissingUpsert(entity ecs.EntityID) {
	if _, ok := s.CoordsInteraction().MissingPreview().Get(entity); !ok {
		return
	}
	worldEntity, ok := s.Seed().WorldSeed()
	if !ok {
		return
	}
	coordsCursor, ok := s.CoordsCursor().Get(entity)
	if !ok {
		return
	}

	s.Prototype().CloneTo(coordsCursor.PropertiesEntity, entity)
	s.Obstruction().Deployed().Remove(entity)
	s.Collider().Component().Remove(entity)

	s.Hierarchy().SetParent(entity, worldEntity)
	s.Tile().Layer().Set(entity, tile.NewLayer(definitions.ObjectPlaceholderLayer))

	if !coordsCursor.CustomImage {
		s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Hud().Selected))
	}
}

//

func (s *service) OnCoordsStateUpsert(entity ecs.EntityID) {
	stateComp, ok := s.CoordsInteraction().StatePreview().Get(entity)
	if !ok {
		return
	}

	worldEntity, ok := s.Seed().WorldSeed()
	if !ok {
		return
	}

	targetCoords := stateComp.State.Coords
	s.Hierarchy().SetParent(entity, worldEntity)

	pos := tile.NewPos(targetCoords.Coords())
	s.Tile().Pos().Set(entity, pos)

	s.Transform().Parent().Set(entity, transform.NewParent(transform.Absolute))
	s.Tile().Layer().Set(entity, tile.NewLayer(definitions.ObjectSelectionPlaceholderLayer))

	s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Hud().Selected))
	s.Groups().InheritGroups(entity)

	s.Obstruction().Deployed().Remove(entity)
	s.Collider().Component().Remove(entity)
}
