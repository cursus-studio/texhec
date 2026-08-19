package service

import (
	"engine/modules/collider"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/inputs"
	"engine/modules/transform"

	"github.com/ogiusek/events"
)

type HoverEvent struct{ Target inputs.Target }
type ClickEvent struct{ Target inputs.Target }

func (HoverEvent) SetTarget(target inputs.Target) inputs.EventTargetSetter { return HoverEvent{target} }
func (ClickEvent) SetTarget(target inputs.Target) inputs.EventTargetSetter { return ClickEvent{target} }

//

func (s *service) OnUpsert(entity ecs.EntityID) {
	coords, ok := s.Coords().Get(entity)
	if !ok {
		return
	}
	s.Inputs().Hover().Set(entity, inputs.NewHoverComponent(HoverEvent{}))
	s.Inputs().Stack().Set(entity, inputs.StackComponent{})
	s.Inputs().LeftClick().Set(entity, inputs.NewLeftClick(ClickEvent{}))

	size := s.GetTileSize()
	size.Size[0] *= float32(s.Grid().ChunkSize())
	size.Size[1] *= float32(s.Grid().ChunkSize())

	s.Transform().Pos().Set(entity, transform.NewPos(
		float32(coords.X)*size.Size[0],
		float32(coords.Y)*size.Size[1],
		0,
	))
	s.Transform().Size().Set(entity, size)
	s.Transform().PivotPoint().Set(entity, transform.NewPivotPoint(0, 0, .5))
}

func (s *service) getCollisionCoords(collision collider.ObjectRayCollision) grid.Coords {
	size := float32(s.Grid().ChunkSize())
	point := collision.Hit.Point
	return grid.NewCoords(
		grid.Coord(size*(1+point.X())/2),
		grid.Coord(size*(1+point.Y())/2),
	)
}

func (s *service) OnHover(e HoverEvent) {
	coords := s.getCollisionCoords(e.Target.ObjectRayCollision)
	event := grid.NewHoverEvent(e.Target.Entity, coords)
	events.EmitAny(s.Events(), event)
}

func (s *service) OnClick(e ClickEvent) {
	coords := s.getCollisionCoords(e.Target.ObjectRayCollision)
	event := grid.NewClickEvent(e.Target.Entity, coords)
	events.EmitAny(s.Events(), event)
}
