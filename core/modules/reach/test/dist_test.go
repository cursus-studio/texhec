package test

import (
	"core/modules/tile"
	"testing"
)

func pow(num tile.Coord) tile.Coord {
	return num * num
}

func TestDistance(t *testing.T) {
	s := NewSetup()

	s.ExpectDist(t, .5, s.Reach().Distance(
		tile.NewPos(.5, .5), tile.NewSize(1, 1),
		tile.NewPos(1, 0), tile.NewSize(1, 1),
	))
	s.ExpectDist(t, 1, s.Reach().Distance(
		tile.NewPos(0, 0), tile.NewSize(1, 1),
		tile.NewPos(1, 0), tile.NewSize(1, 1),
	))
	s.ExpectDist(t, 2, s.Reach().Distance(
		tile.NewPos(0, 0), tile.NewSize(1, 1),
		tile.NewPos(1, 1), tile.NewSize(1, 1),
	))
	s.ExpectDist(t, pow(2), s.Reach().Distance(
		tile.NewPos(0, 0), tile.NewSize(1, 1),
		tile.NewPos(2, 0), tile.NewSize(1, 1),
	))
	s.ExpectDist(t, pow(3), s.Reach().Distance(
		tile.NewPos(0, 0), tile.NewSize(1, 1),
		tile.NewPos(3, 0), tile.NewSize(1, 1),
	))
}
