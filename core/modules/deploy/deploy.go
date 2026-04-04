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
	Deploy(DeployEvent)
}

//

type DeployEvent struct {
	// By,
	Blueprint ecs.EntityID
	Coords    grid.Coords
}

func NewDeployEvent(
	// by,
	blueprint ecs.EntityID,
	coords grid.Coords,
) DeployEvent {
	return DeployEvent{
		// by,
		blueprint,
		coords,
	}
}
