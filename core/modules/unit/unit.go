package unit

import (
	"engine/services/ecs"

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

type ClickEvent struct {
	Unit ecs.EntityID
}

func NewClickEvent(unit ecs.EntityID) ClickEvent {
	return ClickEvent{unit}
}

//

type Service interface {
	Coords() ecs.ComponentsArray[CoordsComponent]
	Unit() ecs.ComponentsArray[UnitComponent]
}
