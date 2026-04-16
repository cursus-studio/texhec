package pathfind

import (
	"engine/modules/grid"
	"engine/services/ecs"
	"errors"
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

type Service interface {
	Target() ecs.ComponentsArray[TargetComponent]

	Select(SelectEvent)
	PreviewPath(PreviewPathEvent)
	FindPath(FindPathEvent)
}

// TODO:
// - add path caching (best would be improving its performance)
// - add more security to ensure 2 objects cannot move to the same tile in one tick (this causes error currently)
// - improve errors

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
