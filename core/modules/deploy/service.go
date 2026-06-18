// defines GUI for deploying objects
package deploy

import (
	"core/modules/reach"
	"engine/modules/grid"
	"engine/modules/interactions"
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

type Service interface {
	Reach() reach.ServiceT[Component]

	Component() ecs.ComponentsArray[Component]

	// deploy differs from execute event by who deploys.
	// execute adds costs and everything where deploy just deploys without any costs (its deployed by system)
	Deploy(
		blueprint,
		owner ecs.EntityID,
		coords grid.Coords,
	) (ecs.EntityID, error)
	DeployEvent(DeployEvent)
}

//

type DeployEvent struct {
	By,
	Blueprint ecs.EntityID
	Coords grid.Coords
}

func NewFeatureDeployEvent() interactions.FeatureEvent[DeployEvent] {
	return interactions.NewFeatureEvent(DeployEvent{})
}
func NewDeployEvent(
	by,
	blueprint ecs.EntityID,
	coords grid.Coords,
) DeployEvent {
	return DeployEvent{
		By:        by,
		Blueprint: blueprint,
		Coords:    coords,
	}
}
