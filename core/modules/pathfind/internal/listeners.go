package internal

import (
	"core/modules/definitions"
	"core/modules/pathfind"
	"core/modules/tile"
	"core/modules/ui"
	"engine/modules/collider"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/render"

	"github.com/ogiusek/events"
)

func (s *service) PreviewPath(e pathfind.PreviewPathEvent) {
	worldEntity, ok := s.Tile().GetConfig()
	if !ok {
		return
	}
	events.Emit(s.Events(), ui.NewUnselect[ui.ActionComponent]())

	from, ok := s.Tile().Pos().Get(e.Entity)
	if !ok {
		s.Logger().Log(tile.ErrInvalidPosition)
		return
	}
	to := tile.NewPos(e.Coords.Coords())
	size, _ := s.Tile().Size().Get(e.Entity)
	obstruction, _ := s.Obstruction().Component().Get(e.Entity)
	fromCoords, _ := from.Aligned()
	toCoords, _ := to.Aligned()
	_, ok = s.findPath(fromCoords, toCoords, size, obstruction)
	destination := tile.NewPos(e.Coords.Coords())
	if !ok {
		entity := s.World().NewEntity()
		s.Hierarchy().SetParent(entity, worldEntity)

		s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Hud().Cannot))
		s.Groups().Inherit().Set(entity, groups.InheritGroupsComponent{})

		s.Collider().Component().Set(entity, collider.NewCollider(s.Definitions().Assets().SquareCollider))

		s.Tile().Layer().Set(entity, tile.NewLayer(definitions.PathLayer))
		s.Tile().Pos().Set(entity, destination)
		s.Tile().Size().Set(entity, size)
		s.Ui().Actions().Set(entity, ui.ActionComponent{})
		return
	}
	entity := s.World().NewEntity()
	s.Hierarchy().SetParent(entity, worldEntity)

	s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Hud().Can))
	s.Groups().Inherit().Set(entity, groups.InheritGroupsComponent{})

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
	obstruction, _ := s.Obstruction().Component().Get(e.Entity)
	fromCoords, _ := from.Aligned()
	toCoords, _ := to.Aligned()
	if _, ok := s.findPath(fromCoords, toCoords, size, obstruction); !ok {
		s.Logger().Log(pathfind.ErrInvalidPath)
		return
	}
	s.Target().Set(e.Entity, pathfind.NewTarget(e.Coords))

	events.Emit(s.Events(), ui.NewUnselect[ui.ObjectComponent]())
}

func (s *service) OnObjectSelect(e ui.SelectEvent[ui.ObjectComponent]) {
	worldEntity, ok := s.Tile().GetConfig()
	if !ok {
		return
	}
	for _, entity := range e.Entities {
		target, ok := s.Target().Get(entity)
		if !ok {
			continue
		}
		size, _ := s.Tile().Size().Get(entity)

		marker := s.World().NewEntity()
		s.Hierarchy().SetParent(marker, worldEntity)

		s.Render().Mesh().Set(marker, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(marker, render.NewTexture(s.Definitions().Hud().Target))
		s.Groups().Inherit().Set(marker, groups.InheritGroupsComponent{})

		s.Tile().Layer().Set(marker, tile.NewLayer(definitions.PathLayer))
		s.Tile().Pos().Set(marker, tile.NewPos(target.Coords.Coords()))
		s.Tile().Size().Set(marker, size)
		s.Collider().Component().Set(marker, collider.NewCollider(s.Definitions().Assets().SquareCollider))

		s.Ui().Objects().Set(marker, ui.ObjectComponent{})
	}
}
