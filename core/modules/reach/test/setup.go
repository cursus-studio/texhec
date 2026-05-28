package test

import (
	"core/game"
	"core/modules/reach"
	reachpkg "core/modules/reach/pkg"
	"core/modules/tile"
	corepkg "core/pkg"
	"engine/modules/grid"
	"slices"
	"testing"

	"github.com/ogiusek/ioc/v2"
)

type FeatureComponent struct{}

type Setup struct {
	game.GameWorld `inject:""`
	ReachT         reach.ServiceT[FeatureComponent] `inject:""`
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		corepkg.Pkg,
		reachpkg.PkgT[FeatureComponent](),
	)
	s := ioc.GetServices[Setup](c)
	return s
}

func (s *Setup) ExpectDist(t *testing.T, expected, actual tile.Coord) {
	t.Helper()
	if expected == actual {
		return
	}
	t.Errorf("expected reach %v but got %v", expected, actual)
}

func (s *Setup) ExpectTilesWithinReach(t *testing.T, expected, actual []grid.Coords) {
	t.Helper()
	sorter := func(a, b grid.Coords) int {
		if a.X > b.X {
			return -1
		} else if a.X < b.X {
			return 1
		} else if a.Y > b.Y {
			return -1
		} else if a.Y < b.Y {
			return 1
		}
		return 0
	}
	slices.SortFunc(expected, sorter)
	slices.SortFunc(actual, sorter)
	if slices.Equal(expected, actual) {
		return
	}
	t.Errorf("expected reach %v but got %v", expected, actual)
}
func (s *Setup) getTilesWithinReach(
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
