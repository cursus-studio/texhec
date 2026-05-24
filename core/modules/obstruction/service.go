// defines how obstruction map is stored and accessed
package obstruction

import (
	"core/modules/tile"
	"engine/modules/grid"
	"engine/services/ecs"
	"errors"
)

var (
	ErrPositionIsOccupied error = errors.New("obstruction:position is occupied")
)

// mask of ways in which tile is obstructed
type Obstruction uint8

// Defines how entity or tile obstruct
// On obstruction collision new entity is removed and warning is logged
type Component struct {
	Obstruction Obstruction
}

func NewObstruction(obstruction Obstruction) Component {
	return Component{obstruction}
}

// aabb on grid
type AABB struct {
	Coords tile.PosComponent
	Size   tile.SizeComponent
	Tiles  []grid.Coords
}

func NewAABB(coords tile.PosComponent, size tile.SizeComponent) AABB {
	posX := grid.Coord(coords.X)
	posY := grid.Coord(coords.Y)
	if tile.Coord(posX) != coords.X {
		size.X++
	}
	if tile.Coord(posY) != coords.Y {
		size.Y++
	}
	tiles := make([]grid.Coords, 0, size.X*size.Y)
	for x := posX; x < posX+size.X; x++ {
		for y := posY; y < posY+size.Y; y++ {
			tiles = append(tiles, grid.NewCoords(x, y))
		}
	}
	return AABB{coords, size, tiles}
}

// adding and removing deployed component modifies obstruction component
type DeployedComponent struct{}

func NewDeployed() DeployedComponent {
	return DeployedComponent{}
}

//

type Service interface {
	ecs.SystemRegister
	Grid() grid.ServiceT[Obstruction]
	Component() ecs.ComponentsArray[Component]
	Deployed() ecs.ComponentsArray[DeployedComponent]

	Collisions(AABB, Obstruction) []grid.Coords
}
