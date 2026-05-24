package gridcollider

import (
	"engine"
	"engine/modules/collider"
	"engine/modules/grid"
	"engine/modules/inputs"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type HoverEvent[Tile grid.TileConstraint] struct {
	Target inputs.Target
}

func (HoverEvent[Tile]) SetTarget(target inputs.Target) inputs.EventTargetSetter {
	return HoverEvent[Tile]{target}
}

//

type ClickEvent[Tile grid.TileConstraint] struct {
	Target inputs.Target
}

func (ClickEvent[Tile]) SetTarget(target inputs.Target) inputs.EventTargetSetter {
	return ClickEvent[Tile]{target}
}

//

type squareFallThroughPolicy[Tile grid.TileConstraint] struct {
	engine.EngineWorld `inject:""`
	GridT              grid.ServiceT[Tile] `inject:""`

	zero       Tile
	hoverEvent func(ecs.EntityID, grid.Coords) any
}

func NewColliderWithPolicy[Tile grid.TileConstraint](
	c ioc.Dic,
	hoverEvent func(ecs.EntityID, grid.Coords) any,
) collider.FallTroughPolicy {
	s := ioc.GetServices[*squareFallThroughPolicy[Tile]](c)
	s.hoverEvent = hoverEvent

	s.GridT.Chunk().OnUpsert(s.OnUpsert)

	events.Listen(s.EventsBuilder(), s.OnHover)

	return s
}

func (t *squareFallThroughPolicy[Tile]) OnUpsert(entity ecs.EntityID) {
	t.Inputs().Hover().Set(entity, inputs.NewHoverComponent(HoverEvent[Tile]{}))
	t.Inputs().LeftClick().Set(entity, inputs.NewLeftClick(ClickEvent[Tile]{}))
}

func (t *squareFallThroughPolicy[Tile]) getCoords(collision collider.ObjectRayCollision) grid.Coords {
	size := float32(t.Grid().ChunkSize())
	point := collision.Hit.Point
	return grid.NewCoords(
		grid.Coord(size*(1+point.X())/2),
		grid.Coord(size*(1+point.Y())/2),
	)
}

func (t *squareFallThroughPolicy[Tile]) FallThrough(collision collider.ObjectRayCollision) bool {
	gridComponent, ok := t.GridT.Chunk().Get(collision.Entity)
	if !ok {
		return false
	}

	coords := t.getCoords(collision)
	index, ok := t.Grid().CoordsIndex(coords)
	if !ok {
		return true
	}

	tile := gridComponent.GetTile(index)
	return tile == t.zero
}

func (t *squareFallThroughPolicy[Tile]) OnHover(e HoverEvent[Tile]) {
	coords := t.getCoords(e.Target.ObjectRayCollision)
	event := t.hoverEvent(e.Target.Entity, coords)
	events.EmitAny(t.Events(), event)
}
