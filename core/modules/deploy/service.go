// defines GUI for deploying objects
package deploy

import (
	"core/modules/reach"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/interactions"
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

	Component() ecs.ComponentArray[Component]

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

type DeployFeature struct{}
type DeployEvent struct {
	By,
	Blueprint ecs.EntityID
	Coords grid.Coords
}

func NewDeployFeature() interactions.FeatureEvent[DeployFeature] {
	return interactions.NewFeatureEvent(DeployFeature{})
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
