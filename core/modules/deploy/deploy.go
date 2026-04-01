package deploy

import (
	"engine/modules/grid"
	"engine/services/ecs"
)

type Service interface {
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
