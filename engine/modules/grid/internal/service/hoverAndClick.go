package service

import (
	"engine/modules/collider"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/inputs"

	"github.com/ogiusek/events"
)

type HoverEvent struct{ Target inputs.Target }
type ClickEvent struct{ Target inputs.Target }

func (HoverEvent) SetTarget(target inputs.Target) inputs.EventTargetSetter { return HoverEvent{target} }
func (ClickEvent) SetTarget(target inputs.Target) inputs.EventTargetSetter { return ClickEvent{target} }

//

func (t *service) OnUpsert(entity ecs.EntityID) {
	t.Inputs().Hover().Set(entity, inputs.NewHoverComponent(HoverEvent{}))
	t.Inputs().LeftClick().Set(entity, inputs.NewLeftClick(ClickEvent{}))
}

func (t *service) getCollisionCoords(collision collider.ObjectRayCollision) grid.Coords {
	size := float32(t.Grid().ChunkSize())
	point := collision.Hit.Point
	return grid.NewCoords(
		grid.Coord(size*(1+point.X())/2),
		grid.Coord(size*(1+point.Y())/2),
	)
}

func (t *service) OnHover(e HoverEvent) {
	coords := t.getCollisionCoords(e.Target.ObjectRayCollision)
	event := grid.NewHoverEvent(e.Target.Entity, coords)
	events.EmitAny(t.Events(), event)
}

func (t *service) OnClick(e ClickEvent) {
	coords := t.getCollisionCoords(e.Target.ObjectRayCollision)
	event := grid.NewClickEvent(e.Target.Entity, coords)
	events.EmitAny(t.Events(), event)
}
