package reach

import (
	"core/modules/tile"
	"engine/modules/grid"
	"engine/services/ecs"
	"errors"
)

var (
	ErrOutsideOfReach error = errors.New("reach: outside of reach")
)

// stores reach distance squared (squared to avoid Sqrt)
type Component[FeatureComponent any] struct {
	Reach grid.Coord
}

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
	Component() ecs.ComponentsArray[Component[FeatureComponent]]
	Reaches(fromEntity, toEntity ecs.EntityID) bool
	TilesWithinReach(entity ecs.EntityID) []grid.Coords
}
