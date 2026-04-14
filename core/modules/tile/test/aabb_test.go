package test

import (
	"core/modules/tile"
	"engine/modules/grid"
	"testing"
)

func TestAABB(t *testing.T) {
	unorderedEqual := func(a, b []grid.Coords) bool {
		if len(a) != len(b) {
			return false
		}
		counts := make(map[grid.Coords]int)
		for _, x := range a {
			counts[x]++
		}
		for _, x := range b {
			if counts[x] == 0 {
				return false
			}
			counts[x]--
		}
		return true
	}

	expectEqual := func(pos tile.PosComponent, size tile.SizeComponent, expected []grid.Coords) {
		t.Helper()
		had := tile.NewAABB(pos, size).Tiles
		if !unorderedEqual(expected, had) {
			t.Errorf("expected %v but got %v", expected, had)
		}
	}
	expectEqual(tile.NewPos(0, 0), tile.NewSize(0, 0), []grid.Coords{})
	expectEqual(tile.NewPos(0, 0), tile.NewSize(1, 1), []grid.Coords{
		grid.NewCoords(0, 0),
	})
	expectEqual(tile.NewPos(.1, 0), tile.NewSize(1, 1), []grid.Coords{
		grid.NewCoords(0, 0),
		grid.NewCoords(1, 0),
	})
	expectEqual(tile.NewPos(2, 0.17164845126343442), tile.NewSize(1, 1), []grid.Coords{
		grid.NewCoords(2, 0),
		grid.NewCoords(2, 1),
	})
	expectEqual(tile.NewPos(.02, 0), tile.NewSize(2, 2), []grid.Coords{
		grid.NewCoords(0, 0),
		grid.NewCoords(1, 0),
		grid.NewCoords(2, 0),

		grid.NewCoords(0, 1),
		grid.NewCoords(1, 1),
		grid.NewCoords(2, 1),
	})

	expectEqual(tile.NewPos(0, 0.053360091944934296), tile.NewSize(2, 2), []grid.Coords{
		grid.NewCoords(0, 0),
		grid.NewCoords(0, 1),
		grid.NewCoords(0, 2),

		grid.NewCoords(1, 0),
		grid.NewCoords(1, 1),
		grid.NewCoords(1, 2),
	})
	expectEqual(tile.NewPos(0, 0.06915298505355394), tile.NewSize(2, 2), []grid.Coords{
		grid.NewCoords(0, 0),
		grid.NewCoords(0, 1),
		grid.NewCoords(0, 2),

		grid.NewCoords(1, 0),
		grid.NewCoords(1, 1),
		grid.NewCoords(1, 2),
	})

}
