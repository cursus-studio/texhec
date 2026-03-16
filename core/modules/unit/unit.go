package unit

import (
	"engine/services/ecs"

	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/exp/constraints"
)

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

func (e *RotationComponent) Quat() mgl32.Quat {
	return mgl32.QuatRotate(e.Radians, mgl32.Vec3{0, 0, -1})
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
