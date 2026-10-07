// defines GUI for deploying objects
package deploy

import (
	"core/modules/actions"
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
	actions.PlayerContext
	By        actions.Entity
	Blueprint actions.Blueprint
	Coords    actions.Coords
}
type DestroyEvent struct {
	actions.PlayerContext
	Entity actions.Entity
}

func NewDeployEvent(
	ctx actions.PlayerContext,
	by, blueprint uuid.UUID,
	coords grid.Coords,
) DeployEvent {
	return DeployEvent{
		ctx, actions.Entity(by), actions.Blueprint(blueprint),
		actions.Coords(coords),
	}
}
func NewDestroyEvent(ctx actions.PlayerContext, uuid uuid.UUID) DestroyEvent {
	return DestroyEvent{ctx, actions.Entity(uuid)}
}

//

type Service interface {
	ecs.SystemRegister
	Reach() reach.ServiceT[Component]

	Component() ecs.ComponentArray[Component]

	CanDeploy(actions.Entity) error
	CanDeployBlueprint(actions.Entity, actions.Blueprint) error
	CanDeployAt(actions.Blueprint, actions.Coords) error

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
