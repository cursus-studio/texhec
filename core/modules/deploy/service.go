// defines GUI for deploying objects
package deploy

import (
	"core/modules/player"
	"core/modules/reach"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/uuid"
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

type DeployEvent struct {
	player.PlayerContext
	By, Blueprint uuid.UUID
	Coords        grid.Coords
}
type DestroyEvent struct {
	player.PlayerContext
	UUID uuid.UUID
}

func NewDeployEvent(
	ctx player.PlayerContext,
	by, blueprint uuid.UUID,
	coords grid.Coords,
) DeployEvent {
	return DeployEvent{
		ctx,
		by, blueprint,
		coords,
	}
}
func NewDestroyEvent(ctx player.PlayerContext, uuid uuid.UUID) DestroyEvent {
	return DestroyEvent{ctx, uuid}
}

//

type Service interface {
	ecs.SystemRegister
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
