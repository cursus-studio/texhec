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

	T *testing.T
}

func NewSetup(t *testing.T) Setup {
	c := ioc.NewContainer(
		corepkg.Pkg,
		reachpkg.PkgT[FeatureComponent](),
	)
	s := ioc.GetServices[Setup](c)
	s.T = t
	return s
}

func (s *Setup) ExpectDist(expected, actual tile.Coord) {
	s.T.Helper()
	if expected == actual {
		return
	}
	s.T.Errorf("expected reach %v but got %v", expected, actual)
}

func (s *Setup) ExpectTilesWithinReach(expected, actual []grid.Coords) {
	s.T.Helper()
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
	s.T.Errorf("expected reach %v but got %v", expected, actual)
}
