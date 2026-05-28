package test

import (
	"core/modules/tile"
	"math"
	"testing"
)

func pow(num tile.Coord) tile.Coord {
	return num * num
}

func TestDistance(t *testing.T) {
	s := NewSetup(t)

	s.ExpectDist(math.MaxFloat32, s.Reach().Distance(
		tile.NewPos(.5, .5), tile.NewSize(1, 1),
		tile.NewPos(1, 0), tile.NewSize(1, 1),
	))
	s.ExpectDist(1, s.Reach().Distance(
		tile.NewPos(0, 0), tile.NewSize(1, 1),
		tile.NewPos(1, 0), tile.NewSize(1, 1),
	))
	s.ExpectDist(2, s.Reach().Distance(
		tile.NewPos(0, 0), tile.NewSize(1, 1),
		tile.NewPos(1, 1), tile.NewSize(1, 1),
	))
	s.ExpectDist(pow(2), s.Reach().Distance(
		tile.NewPos(0, 0), tile.NewSize(1, 1),
		tile.NewPos(2, 0), tile.NewSize(1, 1),
	))
	s.ExpectDist(pow(3), s.Reach().Distance(
		tile.NewPos(0, 0), tile.NewSize(1, 1),
		tile.NewPos(3, 0), tile.NewSize(1, 1),
	))
}
