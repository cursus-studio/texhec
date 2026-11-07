// defines GUI for deploying objects
package deploy

import (
	"core/modules/actions"
	"core/modules/reach"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/uuid"
)

type DeployCounterComponent struct {
	I uint32
}
type DeployComponent struct {
	Deployable []ecs.EntityID
}

func NewDeployCounter(i uint32) DeployCounterComponent {
	return DeployCounterComponent{i}
}
func NewDeploy(deployable ...ecs.EntityID) DeployComponent {
	return DeployComponent{
		deployable,
	}
}

func (c DeployCounterComponent) Inc() DeployCounterComponent { return NewDeployCounter(c.I + 1) }

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
	Reach() reach.ServiceT[DeployComponent]

	DeployCounter() ecs.ComponentArray[DeployCounterComponent]
	DeployComp() ecs.ComponentArray[DeployComponent]

	CanDeploy(actions.Entity) error
	CanDeployBlueprint(actions.Entity, actions.Blueprint) error
	CanDeployAt(actions.Entity, actions.Blueprint, actions.Coords) error

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
