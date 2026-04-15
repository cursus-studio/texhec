package tile

import (
	"engine/modules/grid"
	"engine/modules/transform"
	"engine/modules/transition"
	"engine/services/ecs"
	"errors"

	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/exp/constraints"
)

type System ecs.SystemRegister
type SystemRenderer ecs.SystemRegister

//

var (
	// error logged when grid.GetIndex returns !ok
	ErrInvalidPosition                  error = errors.New("tile:position not found on the grid")
	ErrPositionIsOccupied               error = errors.New("tile:position is occupied")
	ErrInvalidStep                      error = errors.New("tile:invalid step")
	ErrPositionAndSpeedIsRequiredToStep error = errors.New("tile:to step you need to have speed and position")
)

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

type Coord float32

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

func (p *PosComponent) IsAligned() bool {
	return Coord(grid.Coord(p.X)) == p.X && Coord(grid.Coord(p.Y)) == p.Y
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
	X, Y grid.Coord
}

func NewSize[Number constraints.Integer](x, y Number) SizeComponent {
	return SizeComponent{grid.Coord(x), grid.Coord(y)}
}

func (c *SizeComponent) Size() (grid.Coord, grid.Coord) {
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
//
//

type PlaceholderComponent struct{}

func NewPlaceholder() PlaceholderComponent {
	return PlaceholderComponent{}
}

//
//
//

// obstruction

// mask of ways in which tile is obstructed
type Obstruction uint8

func NewObstructGrid(w, h grid.Coord) grid.SquareGridComponent[Obstruction] {
	return grid.NewSquareGrid[Obstruction](w, h)
}

// Defines how entity or tile obstruct
// On obstruction collision new entity is removed and warning is logged
type ObstructionComponent struct {
	Obstruction Obstruction
}

func NewObstruction(obstruction Obstruction) ObstructionComponent {
	return ObstructionComponent{obstruction}
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
//
//

type SpeedComponent struct {
	InvSpeed int8 // ticks to move one tile
}

func NewSpeed[Number constraints.Integer](invSpeed Number) SpeedComponent {
	return SpeedComponent{int8(invSpeed)}
}

//

// Step coords should be +/- 1 x or y from current target position.
// Otherwise step will be removed and warning will be logged.
type StepComponent struct {
	grid.Coords
}

func NewStep(x, y grid.Coord) StepComponent {
	return StepComponent{grid.NewCoords(x, y)}
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

	// elemenets with this components are gui indicators
	Placeholder() ecs.ComponentsArray[PlaceholderComponent]

	Obstruction() ecs.ComponentsArray[ObstructionComponent]
	Deployed() ecs.ComponentsArray[DeployedComponent]

	Speed() ecs.ComponentsArray[SpeedComponent]
	Step() ecs.ComponentsArray[StepComponent]

	//

	// 1x1 size to transform
	GetTileSize() transform.SizeComponent
	Collisions(AABB, Obstruction) []grid.Coords
	CanStep(
		PosComponent,
		SizeComponent,
		ObstructionComponent,
		StepComponent,
	) bool
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
