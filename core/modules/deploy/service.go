// defines GUI for deploying objects
package deploy

import (
	"core/modules/reach"
	"core/modules/tile"
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
	By        tile.FriendlyBuilderObjectStep
	Blueprint tile.BlueprintStep
	Coords    tile.CoordsStep
}
type DestroyEvent struct {
	Object tile.FriendlyObjectStep
}

func NewDeployEvent(
	by,
	blueprint ecs.EntityID,
	coords grid.Coords,
) DeployEvent {
	return DeployEvent{
		interactions.NewStepT[tile.FriendlyBuilderObjectStep](tile.NewObjectInteraction(by)),
		interactions.NewStepT[tile.BlueprintStep](tile.NewBlueprintInteraction(blueprint)),
		interactions.NewStepT[tile.CoordsStep](tile.NewCoordsInteraction(coords)),
	}
}
func NewDestroyEvent(
	object ecs.EntityID,
) DestroyEvent {
	return DestroyEvent{
		interactions.NewStepT[tile.ObjectStep](tile.NewObjectInteraction(object)),
	}
}
