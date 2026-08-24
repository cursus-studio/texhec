package reach

import (
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/grid"
	"errors"
)

var (
	ErrOutsideOfReach error = errors.New("reach: outside of reach")
)

// stores reach distance squared (squared to avoid Sqrt)
type Component[FeatureComponent any] struct {
	Reach grid.Coord
}

// takes square of distnace
func NewReach[FeatureComponent any](reach grid.Coord) Component[FeatureComponent] {
	return Component[FeatureComponent]{reach}
}

//

type Service interface {
	// returns rounded up distance between nearest coordinates
	Distance(
		from tile.PosComponent, fromSize tile.SizeComponent,
		to tile.PosComponent, toSize tile.SizeComponent,
	) tile.Coord
}

type ServiceT[FeatureComponent any] interface {
	Component() ecs.ComponentArray[Component[FeatureComponent]]
	Reaches(fromEntity, toEntity ecs.EntityID) bool
	TilesFrom(tile.PosComponent, tile.SizeComponent, Component[FeatureComponent]) []grid.Coords
	TilesWithinReach(entity ecs.EntityID) []grid.Coords
}
