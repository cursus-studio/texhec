package deploy

import (
	"engine/modules/grid"
	"engine/services/ecs"
)

type Component struct {
	Deployable []ecs.EntityID
}

func NewDeploy(deployable ...ecs.EntityID) Component {
	return Component{
		deployable,
	}
}

//

type LinkComponent struct {
	Deploy ecs.EntityID
}

func NewLink(deploy ecs.EntityID) LinkComponent {
	return LinkComponent{
		deploy,
	}
}

//

type Service interface {
	Component() ecs.ComponentsArray[Component]
	Link() ecs.ComponentsArray[LinkComponent]

	// deploy differs from execute event by who deploys.
	// execute adds costs and everything where deploy just deploys without any costs (its deployed by system)
	Deploy(
		blueprint,
		owner ecs.EntityID,
		coords grid.Coords,
	) (ecs.EntityID, error)
	Select(SelectEvent)
	Preview(PreviewEvent)
	Execute(ExecuteEvent)
}

//

// Select unit.
// Add in gui some indicator.
// Change on click event.
type SelectEvent struct {
	By,
	Blueprint ecs.EntityID
}

func NewSelectEvent(
	by,
	blueprint ecs.EntityID,
) SelectEvent {
	return SelectEvent{
		by,
		blueprint,
	}
}

//

// Select unit.
// Add in gui some indicator.
// Perform all checks and costs
type PreviewEvent struct {
	By,
	Blueprint ecs.EntityID
	Coords grid.Coords
}

func NewPreviewEvent(
	by,
	blueprint ecs.EntityID,
) PreviewEvent {
	return PreviewEvent{
		By:        by,
		Blueprint: blueprint,
	}
}

func (e PreviewEvent) ApplyCoords(coords grid.Coords) any {
	e.Coords = coords
	return e
}

//

// Deploys on coords something if it doesn't collide
type ExecuteEvent struct {
	By,
	Blueprint ecs.EntityID
	Coords grid.Coords
}

func NewExecuteEvent(
	by,
	blueprint ecs.EntityID,
) ExecuteEvent {
	return ExecuteEvent{
		By:        by,
		Blueprint: blueprint,
	}
}

func (e ExecuteEvent) ApplyCoords(coords grid.Coords) any {
	e.Coords = coords
	return e
}
