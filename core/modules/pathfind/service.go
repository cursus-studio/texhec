// finds path on a grid and moves objects to their target according to their speed
package pathfind

import (
	"core/modules/obstruction"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/grid"
	"errors"

	"golang.org/x/exp/constraints"
)

var (
	ErrInvalidPath         error = errors.New("pathfind:invalid path")
	ErrInvalidServiceOrder       = errors.New("invalid services order. Chunk region map cannot be generated before chunk")
)

// all entities without [tile.StepComponent] get one on tick which will move them towards target
type TargetComponent struct {
	grid.Coords
}

func NewTarget(coords grid.Coords) TargetComponent {
	return TargetComponent{coords}
}

//

type SpeedComponent struct {
	InvSpeed int8 // ticks to move one tile
}

func NewSpeed[Number constraints.Integer](invSpeed Number) SpeedComponent {
	return SpeedComponent{int8(invSpeed)}
}

//

// Step coords should be +/- 1 x or y from current target position.
// Otherwise step will be removed and warning will be logged.
type StepComponent struct{ grid.Coords }

func NewStep(x, y grid.Coord) StepComponent { return StepComponent{grid.NewCoords(x, y)} }

//

// this variable contains region index and is used for region connectivity
type Region uint16

var NotARegion = ^Region(0)

type Service interface {
	ecs.SystemRegister
	Target() ecs.ComponentArray[TargetComponent]
	Speed() ecs.ComponentArray[SpeedComponent]
	Step() ecs.ComponentArray[StepComponent]

	// region
	RegionObstruction(Region) (obstruction.Obstruction, bool)
	CoordsRegion(grid.Coords, obstruction.Obstruction) (Region, bool)
	EntityRegion(ecs.EntityID) (Region, bool)
	ShareRegion(ecs.EntityID, grid.Coords) bool

	//
	FindPath(FindPathEvent)

	CanStep(
		grid.Coords,
		tile.SizeComponent,
		obstruction.Component,
		StepComponent,
	) bool
}

// TODO:
// - add path caching (best would be improving its performance)
// - add more security to ensure 2 objects cannot move to the same tile in one tick (this causes error currently)
// - improve errors
// - look on `HPA*` and `JPS`

type FindPathEvent struct {
	Entity ecs.EntityID
	Coords grid.Coords
}

func NewFindPathEvent(entity ecs.EntityID, coords grid.Coords) FindPathEvent {
	return FindPathEvent{entity, coords}
}
