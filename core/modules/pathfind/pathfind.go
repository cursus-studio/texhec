package pathfind

import (
	"core/modules/obstruction"
	"core/modules/tile"
	"engine/modules/grid"
	"engine/services/ecs"
	"errors"

	"golang.org/x/exp/constraints"
)

var (
	ErrInvalidPath error = errors.New("pathfind:invalid path")
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

type Service interface {
	ecs.SystemRegister
	Target() ecs.ComponentsArray[TargetComponent]
	Speed() ecs.ComponentsArray[SpeedComponent]
	Step() ecs.ComponentsArray[StepComponent]

	Select(SelectEvent)
	PreviewPath(PreviewPathEvent)
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

// Select object.
// Add in gui some indicator.
// Change on click event.
type SelectEvent struct {
	Entity ecs.EntityID
}

func NewSelectEvent(entity ecs.EntityID) SelectEvent {
	return SelectEvent{
		entity,
	}
}

//

// Select object.
// Add in gui some indicator.
// Perform all checks and costs
type PreviewPathEvent struct {
	Entity ecs.EntityID
	Coords grid.Coords
}

func NewPreviewPathEvent(
	entity ecs.EntityID,
) PreviewPathEvent {
	return PreviewPathEvent{
		Entity: entity,
	}
}

func (e PreviewPathEvent) ApplyCoords(coords grid.Coords) any {
	e.Coords = coords
	return e
}

//

// Adds [TargetComponent] to entity
type FindPathEvent struct {
	Entity ecs.EntityID
	Coords grid.Coords
}

func NewFindPathEvent(
	entity ecs.EntityID,
) FindPathEvent {
	return FindPathEvent{
		Entity: entity,
	}
}

func (e FindPathEvent) ApplyCoords(coords grid.Coords) any {
	e.Coords = coords
	return e
}
