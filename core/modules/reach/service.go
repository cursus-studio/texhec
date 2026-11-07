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
type ReachComponent[FeatureComponent any] struct {
	Reach grid.Coord
}

// takes square of distnace
func NewReach[FeatureComponent any](reach grid.Coord) ReachComponent[FeatureComponent] {
	return ReachComponent[FeatureComponent]{reach}
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
	Reach() ecs.ComponentArray[ReachComponent[FeatureComponent]]
	Reaches(fromEntity, toEntity ecs.EntityID) bool
	ReachesCoords(fromEntity ecs.EntityID, toPos tile.PosComponent, toSize tile.SizeComponent) bool
	TilesFrom(tile.PosComponent, tile.SizeComponent, ReachComponent[FeatureComponent]) []grid.Coords
	TilesWithinReach(entity ecs.EntityID) []grid.Coords
}
