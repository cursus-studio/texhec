package internal

import (
	"core/modules/definitions"
	"core/modules/pathfind"
	"core/modules/tile"
	"core/modules/ui"
	gamescenes "core/scenes"
	"engine/modules/collider"
	"engine/modules/grid"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/render"
	"engine/services/ecs"
	"engine/services/frames"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	gamescenes.GameWorld `inject:""`

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
	for _, entity := range s.Tile().Placeholder().GetEntities() {
		s.World().RemoveEntity(entity)
	}

	from, ok := s.Tile().Pos().Get(e.Entity)
	if !ok {
		s.Logger().Warn(tile.ErrInvalidPosition)
		return
	}
	to := tile.NewPos(e.Coords.Coords())
	size, _ := s.Tile().Size().Get(e.Entity)
	obstruction, _ := s.Tile().Obstruction().Get(e.Entity)
	fromCoords, _ := from.Aligned()
	toCoords, _ := to.Aligned()
	path, ok := s.findPath(fromCoords, toCoords, size, obstruction)
	if !ok {
		for _, pos := range []tile.PosComponent{
			tile.NewPos(e.Coords.Coords()),
			from,
		} {
			entity := s.Prototype().Clone(s.Definitions().Blank)
			s.Hierarchy().SetParent(entity, s.Scene().Scene())

			s.Tile().Layer().Set(entity, tile.NewLayer(definitions.PlaceholderTileLayer))
			s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().SquareMesh))
			s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Blank))
			s.Groups().Component().Set(entity, groups.EmptyGroups().Ptr().Enable(definitions.GameGroup).Val())

			s.Collider().Component().Set(entity, collider.NewCollider(s.Definitions().SquareCollider))

			s.Tile().Layer().Set(entity, tile.NewLayer(definitions.PlaceholderTileLayer))
			s.Tile().Pos().Set(entity, pos)
			s.Tile().Placeholder().Set(entity, tile.NewPlaceholder())
			s.Render().Color().Set(entity, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
		}
		return
	}
	for _, pos := range path {
		entity := s.Prototype().Clone(s.Definitions().Blank)
		s.Hierarchy().SetParent(entity, s.Scene().Scene())

		s.Tile().Layer().Set(entity, tile.NewLayer(definitions.PlaceholderTileLayer))
		s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().SquareMesh))
		s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Blank))
		s.Groups().Component().Set(entity, groups.EmptyGroups().Ptr().Enable(definitions.GameGroup).Val())

		s.Collider().Component().Set(entity, collider.NewCollider(s.Definitions().SquareCollider))
		if pos.X == tile.Coord(e.Coords.X) && pos.Y == tile.Coord(e.Coords.Y) {
			s.Inputs().LeftClick().Set(entity, inputs.NewLeftClick(pathfind.NewFindPathEvent(e.Entity).ApplyCoords(e.Coords)))
		}

		s.Tile().Pos().Set(entity, pos)
		s.Tile().Placeholder().Set(entity, tile.NewPlaceholder())
		s.Render().Color().Set(entity, render.NewColor(mgl32.Vec4{0, 1, 0, 1}))
	}
}
func (s *service) FindPath(e pathfind.FindPathEvent) {
	for _, entity := range s.Tile().Placeholder().GetEntities() {
		s.World().RemoveEntity(entity)
	}

	from, ok := s.Tile().Pos().Get(e.Entity)
	if !ok {
		s.Logger().Warn(tile.ErrInvalidPosition)
		return
	}
	to := tile.NewPos(e.Coords.Coords())
	size, _ := s.Tile().Size().Get(e.Entity)
	obstruction, _ := s.Tile().Obstruction().Get(e.Entity)
	fromCoords, _ := from.Aligned()
	toCoords, _ := to.Aligned()
	if _, ok := s.findPath(fromCoords, toCoords, size, obstruction); !ok {
		s.Logger().Warn(pathfind.ErrInvalidPath)
		return
	}
	s.Target().Set(e.Entity, pathfind.NewTarget(e.Coords))

	events.Emit(s.Events(), ui.HideUiEvent{})
}

func (s *service) OnTick(e frames.TickEvent) {
	for _, entity := range s.Target().GetEntities() {
		if _, ok := s.Tile().Step().Get(entity); ok {
			continue
		}

		from, ok := s.Tile().Pos().Get(entity)
		if !ok {
			s.Logger().Warn(tile.ErrInvalidPosition)
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
			s.Logger().Warn(pathfind.ErrInvalidPath)
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
