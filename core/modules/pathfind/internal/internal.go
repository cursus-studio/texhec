package internal

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/pathfind"
	"core/modules/tile"
	"core/modules/ui"
	"engine/modules/collider"
	"engine/modules/grid"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/loop"
	"engine/modules/render"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld `inject:""`

	target ecs.ComponentsArray[pathfind.TargetComponent]
}

func NewService(c ioc.Dic) pathfind.Service {
	s := ioc.GetServices[*service](c)
	s.target = ecs.GetComponentsArray[pathfind.TargetComponent](s.World())

	events.Listen(s.EventsBuilder(), s.Select)
	events.Listen(s.EventsBuilder(), s.PreviewPath)
	events.Listen(s.EventsBuilder(), s.FindPath)
	events.Listen(s.EventsBuilder(), s.OnTick)
	return s
}

func (s *service) Target() ecs.ComponentsArray[pathfind.TargetComponent] { return s.target }

func (s *service) Select(e pathfind.SelectEvent) {
	events.Emit(s.Events(), tile.NewSelectEvent(pathfind.NewPreviewPathEvent(e.Entity)))
}

func (s *service) PreviewPath(e pathfind.PreviewPathEvent) {
	events.Emit(s.Events(), ui.NewUnselect[ui.ActionComponent]())

	from, ok := s.Tile().Pos().Get(e.Entity)
	if !ok {
		s.Logger().Log(tile.ErrInvalidPosition)
		return
	}
	to := tile.NewPos(e.Coords.Coords())
	size, _ := s.Tile().Size().Get(e.Entity)
	obstruction, _ := s.Tile().Obstruction().Get(e.Entity)
	fromCoords, _ := from.Aligned()
	toCoords, _ := to.Aligned()
	_, ok = s.findPath(fromCoords, toCoords, size, obstruction)
	destination := tile.NewPos(e.Coords.Coords())
	if !ok {
		entity := s.World().NewEntity()
		s.Hierarchy().SetParent(entity, s.Scene().Scene())

		s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Hud().Cannot))
		s.Groups().Component().Set(entity, groups.EmptyGroups().Ptr().Enable(definitions.GameGroup).Val())

		s.Collider().Component().Set(entity, collider.NewCollider(s.Definitions().Assets().SquareCollider))

		s.Tile().Layer().Set(entity, tile.NewLayer(definitions.PathLayer))
		s.Tile().Pos().Set(entity, destination)
		s.Tile().Size().Set(entity, size)
		s.Ui().Actions().Set(entity, ui.ActionComponent{})
		return
	}
	entity := s.World().NewEntity()
	s.Hierarchy().SetParent(entity, s.Scene().Scene())

	s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Hud().Can))
	s.Groups().Component().Set(entity, groups.EmptyGroups().Ptr().Enable(definitions.GameGroup).Val())

	s.Collider().Component().Set(entity, collider.NewCollider(s.Definitions().Assets().SquareCollider))
	if destination.X == tile.Coord(e.Coords.X) && destination.Y == tile.Coord(e.Coords.Y) {
		s.Inputs().LeftClick().Set(entity, inputs.NewLeftClick(pathfind.NewFindPathEvent(e.Entity).ApplyCoords(e.Coords)))
	}

	s.Tile().Layer().Set(entity, tile.NewLayer(definitions.PathLayer))
	s.Tile().Pos().Set(entity, destination)
	s.Tile().Size().Set(entity, size)
	s.Ui().Actions().Set(entity, ui.ActionComponent{})
}
func (s *service) FindPath(e pathfind.FindPathEvent) {
	events.Emit(s.Events(), ui.NewUnselect[ui.ActionComponent]())

	from, ok := s.Tile().Pos().Get(e.Entity)
	if !ok {
		s.Logger().Log(tile.ErrInvalidPosition)
		return
	}
	to := tile.NewPos(e.Coords.Coords())
	size, _ := s.Tile().Size().Get(e.Entity)
	obstruction, _ := s.Tile().Obstruction().Get(e.Entity)
	fromCoords, _ := from.Aligned()
	toCoords, _ := to.Aligned()
	if _, ok := s.findPath(fromCoords, toCoords, size, obstruction); !ok {
		s.Logger().Log(pathfind.ErrInvalidPath)
		return
	}
	s.Target().Set(e.Entity, pathfind.NewTarget(e.Coords))

	events.Emit(s.Events(), ui.NewUnselect[ui.ObjectComponent]())
}

func (s *service) OnTick(e loop.TickEvent) {
	for _, entity := range s.Target().GetEntities() {
		if _, ok := s.Tile().Step().Get(entity); ok {
			continue
		}

		from, ok := s.Tile().Pos().Get(entity)
		if !ok {
			s.Logger().Log(tile.ErrInvalidPosition)
			return
		}
		target, _ := s.Target().Get(entity)
		fromAligned, _ := from.Aligned()
		to := tile.NewPos(target.Coords.Coords())
		toAligned, _ := to.Aligned()
		if toAligned == fromAligned {
			s.Target().Remove(entity)
			continue
		}
		size, _ := s.Tile().Size().Get(entity)
		obstruction, _ := s.Tile().Obstruction().Get(entity)
		fromCoords, _ := from.Aligned()
		toCoords, _ := to.Aligned()
		path, ok := s.findPath(fromCoords, toCoords, size, obstruction)
		if !ok {
			s.Logger().Log(pathfind.ErrInvalidPath)
			continue
		}
		step := tile.NewStep(grid.Coord(path[0].X), grid.Coord(path[0].Y))
		for !s.Tile().CanStep(fromCoords, size, obstruction, step) {
			path, _ = s.findPath(fromCoords, toCoords, size, obstruction)
			step = tile.NewStep(grid.Coord(path[0].X), grid.Coord(path[0].Y))
		}
		s.Tile().Step().Set(entity, step)
	}
}
