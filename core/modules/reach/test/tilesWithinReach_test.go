package test

import (
	"core/modules/tile"
	"engine/modules/grid"
	"testing"
)

func TestTilesWithinReach(t *testing.T) {
	s := NewSetup()

	s.ExpectTilesWithinReach(t,
		[]grid.Coords{grid.NewCoords(0, 0), grid.NewCoords(1, 1)},
		[]grid.Coords{grid.NewCoords(1, 1), grid.NewCoords(0, 0)},
	)

	// Legend:
	// - means outside reach
	// + means within reach
	// X means object

	s.ExpectTilesWithinReach(t,
		[]grid.Coords{
			// -+-
			// +X+
			// -+-
			grid.NewCoords(-1, 0),
			grid.NewCoords(0, -1), grid.NewCoords(1, 0),
			grid.NewCoords(0, 1),
		},
		s.getTilesWithinReach(tile.NewPos(0, 0), tile.NewSize(1, 1), 1),
	)
	s.ExpectTilesWithinReach(t,
		// no result for not fixed position
		[]grid.Coords{},
		s.getTilesWithinReach(tile.NewPos(.5, .5), tile.NewSize(1, 1), 1),
	)
	s.ExpectTilesWithinReach(t,
		[]grid.Coords{
			// -+-
			// +X+
			// -+-
			grid.NewCoords(0, 1),
			grid.NewCoords(1, 0), grid.NewCoords(2, 1),
			grid.NewCoords(1, 2),
		},
		s.getTilesWithinReach(tile.NewPos(1, 1), tile.NewSize(1, 1), 1),
	)

	s.ExpectTilesWithinReach(t,
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
		s.getTilesWithinReach(tile.NewPos(2, 2), tile.NewSize(1, 1), 2),
	)

	s.ExpectTilesWithinReach(t,
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
		s.getTilesWithinReach(tile.NewPos(1, 1), tile.NewSize(2, 2), 1),
	)

	s.ExpectTilesWithinReach(t,
		[]grid.Coords{
			// -++-
			// +XX+
			// -++-
			grid.NewCoords(1, 0), grid.NewCoords(2, 0),
			grid.NewCoords(0, 1), grid.NewCoords(3, 1),
			grid.NewCoords(1, 2), grid.NewCoords(2, 2),
		},
		s.getTilesWithinReach(tile.NewPos(1, 1), tile.NewSize(2, 1), 1),
	)
}
