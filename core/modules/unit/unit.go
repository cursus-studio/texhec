package unit

import "engine/services/ecs"

type System ecs.SystemRegister

type UnitComponent struct {
	Unit ecs.EntityID
}

func NewUnit(unit ecs.EntityID) UnitComponent { return UnitComponent{unit} }

//

type Service interface {
	Unit() ecs.ComponentsArray[UnitComponent]
}

//

type ClickEvent struct {
	Unit ecs.EntityID
}

func NewClickEvent(unit ecs.EntityID) ClickEvent {
	return ClickEvent{unit}
}
