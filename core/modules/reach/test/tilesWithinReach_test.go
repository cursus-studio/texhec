package test

import (
	"core/modules/reach"
	"core/modules/tile"
	"engine/modules/grid"
	"testing"
)

// 	Component() ecs.ComponentsArray[Component[FeatureComponent]]
// 	Reaches(fromEntity, toEntity ecs.EntityID) bool

func TestTilesWithinReach(t *testing.T) {
	s := NewSetup(t)

	getTilesWithinReach := func(
		pos tile.PosComponent,
		size tile.SizeComponent,
		reachRange tile.Coord,
	) []grid.Coords {
		entity := s.World().NewEntity()
		s.Tile().Pos().Set(entity, pos)
		s.Tile().Size().Set(entity, size)
		s.ReachT.Component().Set(entity, reach.NewReach[FeatureComponent](reachRange))
		coords := s.ReachT.TilesWithinReach(entity)
		s.World().RemoveEntity(entity)
		return coords
	}

	s.ExpectTilesWithinReach(
		[]grid.Coords{grid.NewCoords(0, 0), grid.NewCoords(1, 1)},
		[]grid.Coords{grid.NewCoords(1, 1), grid.NewCoords(0, 0)},
	)

	// Legend:
	// - means outside reach
	// + means within reach
	// X means object

	s.ExpectTilesWithinReach(
		[]grid.Coords{
			// -+-
			// +X+
			// -+-
			grid.NewCoords(-1, 0),
			grid.NewCoords(0, -1), grid.NewCoords(1, 0),
			grid.NewCoords(0, 1),
		},
		getTilesWithinReach(tile.NewPos(0, 0), tile.NewSize(1, 1), 1),
	)
	s.ExpectTilesWithinReach(
		// no result for not fixed position
		[]grid.Coords{},
		getTilesWithinReach(tile.NewPos(.5, .5), tile.NewSize(1, 1), 1),
	)
	s.ExpectTilesWithinReach(
		[]grid.Coords{
			// -+-
			// +X+
			// -+-
			grid.NewCoords(0, 1),
			grid.NewCoords(1, 0), grid.NewCoords(2, 1),
			grid.NewCoords(1, 2),
		},
		getTilesWithinReach(tile.NewPos(1, 1), tile.NewSize(1, 1), 1),
	)

	s.ExpectTilesWithinReach(
		[]grid.Coords{
			// --+--
			// -+++-
			// ++X++
			// -+++-
			// --+--
			grid.NewCoords(2, 0),
			grid.NewCoords(1, 1), grid.NewCoords(2, 1), grid.NewCoords(3, 1),
			grid.NewCoords(0, 2), grid.NewCoords(1, 2), grid.NewCoords(3, 2), grid.NewCoords(4, 2),
			grid.NewCoords(1, 3), grid.NewCoords(2, 3), grid.NewCoords(3, 3),
			grid.NewCoords(2, 4),
		},
		getTilesWithinReach(tile.NewPos(2, 2), tile.NewSize(1, 1), 2),
	)

	s.ExpectTilesWithinReach(
		[]grid.Coords{
			// -++-
			// +XX+
			// +XX+
			// -++-
			grid.NewCoords(1, 0), grid.NewCoords(2, 0),
			grid.NewCoords(0, 1), grid.NewCoords(3, 1),
			grid.NewCoords(0, 2), grid.NewCoords(3, 2),
			grid.NewCoords(1, 3), grid.NewCoords(2, 3),
		},
		getTilesWithinReach(tile.NewPos(1, 1), tile.NewSize(2, 2), 1),
	)

	s.ExpectTilesWithinReach(
		[]grid.Coords{
			// -++-
			// +XX+
			// -++-
			grid.NewCoords(1, 0), grid.NewCoords(2, 0),
			grid.NewCoords(0, 1), grid.NewCoords(3, 1),
			grid.NewCoords(1, 2), grid.NewCoords(2, 2),
		},
		getTilesWithinReach(tile.NewPos(1, 1), tile.NewSize(2, 1), 1),
	)
}
