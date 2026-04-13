package tile

import (
	"engine/modules/grid"
	"engine/modules/transform"
	"engine/modules/transition"
	"engine/services/ecs"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/exp/constraints"
)

type System ecs.SystemRegister
type SystemRenderer ecs.SystemRegister

//

type ID uint8

func NewTileGrid(w, h grid.Coord) grid.SquareGridComponent[ID] {
	return grid.NewSquareGrid[ID](w, h)
}

//

type TypeComponent struct {
	ID ID
}

func NewTileType(id ID) TypeComponent {
	return TypeComponent{id}
}

//

type Coord float64

type PosComponent struct {
	X, Y Coord
}

func NewPos[Number constraints.Integer | constraints.Float](x, y Number) PosComponent {
	return PosComponent{Coord(x), Coord(y)}
}

func (c1 PosComponent) Lerp(c2 PosComponent, mix32 float32) PosComponent {
	return PosComponent{
		transition.Lerp(c1.X, c2.X, mix32),
		transition.Lerp(c1.Y, c2.Y, mix32),
	}
}

//

type LayerComponent struct {
	Z Coord
}

func NewLayer[Number constraints.Integer | constraints.Float](z Number) LayerComponent {
	return LayerComponent{Coord(z)}
}

//

type SizeComponent struct {
	X, Y Coord
}

func NewSize[Number constraints.Integer | constraints.Float](x, y Number) SizeComponent {
	return SizeComponent{Coord(x), Coord(y)}
}

func (c1 SizeComponent) Lerp(c2 SizeComponent, mix32 float32) SizeComponent {
	return SizeComponent{
		transition.Lerp(c1.X, c2.X, mix32),
		transition.Lerp(c1.Y, c2.Y, mix32),
	}
}

func (c *SizeComponent) Size() (Coord, Coord) {
	return c.X, c.Y
}

//

type RotComponent struct {
	Radians float32
}

func NewRot(radians float32) RotComponent {
	return RotComponent{radians}
}

func (c1 RotComponent) Lerp(c2 RotComponent, mix32 float32) RotComponent {
	return RotComponent{transition.Lerp(c1.Radians, c2.Radians, mix32)}
}

func (e *RotComponent) Quat() mgl32.Quat {
	return mgl32.QuatRotate(e.Radians, mgl32.Vec3{0, 0, 1})
}

//

// mask of ways in which tile is obstructed
type Obstruction uint8

func NewObstructGrid(w, h grid.Coord) grid.SquareGridComponent[Obstruction] {
	return grid.NewSquareGrid[Obstruction](w, h)
}

// defines how entity or tile obstruct
type ObstructionComponent struct {
	Obstruction Obstruction
}

func NewObstruction(obstruction Obstruction) ObstructionComponent {
	return ObstructionComponent{obstruction}
}

//

// adding and removing deployed component modified obstruction component
type DeployedComponent struct{}

func NewDeployed() DeployedComponent {
	return DeployedComponent{}
}

//

// aabb on grid
type AABB struct {
	Coords PosComponent
	Size   SizeComponent
	Tiles  []grid.Coords
}

func NewAABB(coords PosComponent, size SizeComponent) AABB {
	posX := grid.Coord(coords.X)
	posY := grid.Coord(coords.Y)
	if Coord(posX) != coords.X {
		size.X++
	}
	if Coord(posY) != coords.Y {
		size.Y++
	}
	sizeX := grid.Coord(math.Ceil(float64(size.X)))
	sizeY := grid.Coord(math.Ceil(float64(size.Y)))
	tiles := make([]grid.Coords, 0, sizeX*sizeY)
	for x := posX; x < posX+sizeX; x++ {
		for y := posY; y < posY+sizeY; y++ {
			tiles = append(tiles, grid.NewCoords(x, y))
		}
	}
	return AABB{coords, size, tiles}
}

//

type Service interface {
	TileType() ecs.ComponentsArray[TypeComponent]
	TileGrid() ecs.ComponentsArray[grid.SquareGridComponent[ID]]
	ObstructionGrid() ecs.ComponentsArray[grid.SquareGridComponent[Obstruction]]
	GetTileType(ID) (ecs.EntityID, bool)

	Pos() ecs.ComponentsArray[PosComponent]
	Size() ecs.ComponentsArray[SizeComponent]
	Rot() ecs.ComponentsArray[RotComponent]
	Layer() ecs.ComponentsArray[LayerComponent]
	Obstruction() ecs.ComponentsArray[ObstructionComponent]
	Deployed() ecs.ComponentsArray[DeployedComponent]

	GetTileSize() transform.SizeComponent
}

//

type ApplyCoordsEvent interface {
	ApplyCoords(grid.Coords) any
}

type SelectEvent struct {
	HoverEvent any
}

func NewSelectEvent(hoverEvent any) SelectEvent {
	return SelectEvent{hoverEvent}
}

//

type HoverEvent struct {
	Grid ecs.EntityID
	Tile grid.Index
}

func NewHoverEvent(
	grid ecs.EntityID,
	tile grid.Index,
) any {
	return HoverEvent{grid, tile}
}

//

type ClickEntityEvent struct {
	Entity ecs.EntityID
}

func NewClickEntityEvent() ClickEntityEvent {
	return ClickEntityEvent{}
}

func (e ClickEntityEvent) ApplyEntity(entity ecs.EntityID) any {
	e.Entity = entity
	return e
}
