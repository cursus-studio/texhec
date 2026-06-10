package tileservice

import (
	"core/modules/tile"
	"engine/modules/grid"
	"engine/modules/interactions"
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
