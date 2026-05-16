// defines generic slice based data structure to store grids
package grid

import "engine/services/ecs"

type Service[Tile TileConstraint] interface {
	Component() ecs.ComponentsArray[SquareGridComponent[Tile]]
}
