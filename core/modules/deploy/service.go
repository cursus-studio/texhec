// defines GUI for deploying objects
package deploy

import (
	"core/modules/actions"
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
	DestroyEvent(DestroyEvent)
}

//

type DeployEvent struct {
	By        actions.FriendlyBuilderObjectStep
	Blueprint actions.BlueprintStep
	Coords    actions.CoordsStep
}
type DestroyEvent struct {
	Object actions.FriendlyObjectStep
}

func NewDeployEvent(
	by,
	blueprint ecs.EntityID,
	coords grid.Coords,
) DeployEvent {
	return DeployEvent{
		interactions.NewStep(actions.NewObjectInteraction(by)),
		interactions.NewStep(actions.NewBlueprintInteraction(blueprint)),
		interactions.NewStep(actions.NewCoordsInteraction(coords)),
	}
}
func NewDestroyEvent(
	object ecs.EntityID,
) DestroyEvent {
	return DestroyEvent{
		interactions.NewStep(actions.NewObjectInteraction(object)),
	}
}
