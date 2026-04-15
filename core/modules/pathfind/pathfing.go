package pathfind

import (
	"engine/modules/grid"
	"engine/services/ecs"
)

// TODO:
// - add pathfind component
// - each tick check each component and add tile.Step for each without tile.Step
// - abstract away can move in tile module and use it in here
// - follow patterns in deploy module how to add hover effects

// Select unit.
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

// Select unit.
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

// Deploys on coords something if it doesn't collide
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
