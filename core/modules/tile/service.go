package tile

import (
	"engine/modules/grid"
	"engine/modules/transform"
	"engine/modules/transition"
	"engine/services/ecs"

	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/exp/constraints"
)

type System ecs.SystemRegister
type SystemRenderer ecs.SystemRegister

//

type ID uint8

func NewGrid(w, h grid.Coord) grid.SquareGridComponent[ID] {
	return grid.NewSquareGrid[ID](w, h)
}

type Component struct {
	ID ID
}

func NewTile(id ID) Component {
	return Component{id}
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

type Service interface {
	Tile() ecs.ComponentsArray[Component]
	Grid() ecs.ComponentsArray[grid.SquareGridComponent[ID]]

	Pos() ecs.ComponentsArray[PosComponent]
	Size() ecs.ComponentsArray[SizeComponent]
	Rot() ecs.ComponentsArray[RotComponent]
	Layer() ecs.ComponentsArray[LayerComponent]

	GetTileSize() transform.SizeComponent

	Unit(entity, blueprint ecs.EntityID)
	Construct(entity, blueprint ecs.EntityID)
}

//

type ClickEvent struct {
	Grid ecs.EntityID
	Tile grid.Index
}

func NewClickEvent(
	grid ecs.EntityID,
	tile grid.Index,
) any {
	return ClickEvent{grid, tile}
}

//

type ClickUnitEvent struct {
	Unit ecs.EntityID
}

func NewClickUnitEvent() ClickUnitEvent {
	return ClickUnitEvent{}
}

func (e ClickUnitEvent) ApplyEntity(entity ecs.EntityID) any {
	e.Unit = entity
	return e
}

//

type ClickConstructEvent struct {
	Construct ecs.EntityID
}

func NewClickConstructEvent() ClickConstructEvent {
	return ClickConstructEvent{}
}

func (e ClickConstructEvent) ApplyEntity(entity ecs.EntityID) any {
	e.Construct = entity
	return e
}
