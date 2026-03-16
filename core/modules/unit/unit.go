package unit

import (
	"engine/services/ecs"

	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/exp/constraints"
)

type System ecs.SystemRegister

type UnitComponent struct {
	Unit ecs.EntityID
}

func NewUnit(unit ecs.EntityID) UnitComponent { return UnitComponent{unit} }

//

type Coord float64

type CoordsComponent struct {
	X, Y Coord
}

func NewCoords[Number constraints.Integer | constraints.Float](x, y Number) CoordsComponent {
	return CoordsComponent{Coord(x), Coord(y)}
}

func (c1 CoordsComponent) Lerp(c2 CoordsComponent, mix32 float32) CoordsComponent {
	r := CoordsComponent{
		c1.X*Coord(1-mix32) + c2.X*Coord(mix32),
		c1.Y*Coord(1-mix32) + c2.Y*Coord(mix32),
	}
	return r
}

func (c *CoordsComponent) Coords() (Coord, Coord) {
	return c.X, c.Y
}

//

type RotationComponent struct {
	Radians float32
}

func NewRotation(radians float32) RotationComponent {
	return RotationComponent{radians}
}

func (c1 RotationComponent) Lerp(c2 RotationComponent, mix32 float32) RotationComponent {
	return RotationComponent{c1.Radians*(1-mix32) + c2.Radians*(mix32)}
}

func (e *RotationComponent) Quat() mgl32.Quat {
	return mgl32.QuatRotate(e.Radians, mgl32.Vec3{0, 0, 1})
}

//

type Service interface {
	Coords() ecs.ComponentsArray[CoordsComponent]
	Unit() ecs.ComponentsArray[UnitComponent]
	Rotation() ecs.ComponentsArray[RotationComponent]
}

//

type ClickEvent struct {
	Unit ecs.EntityID
}

func NewClickEvent(unit ecs.EntityID) ClickEvent {
	return ClickEvent{unit}
}
