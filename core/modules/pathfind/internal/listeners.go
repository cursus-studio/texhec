package internal

import (
	"core/modules/pathfind"
	"core/modules/tile"
)

func (s *service) FindPath(e pathfind.FindPathEvent) {
	object := e.Object.State().Entity
	coords := e.Coords.State().Coords

	from, ok := s.Tile().Pos().Get(object)
	if !ok {
		s.Logger().Log(tile.ErrInvalidPosition)
		return
	}
	to := tile.NewPos(coords.Coords())
	size, _ := s.Tile().Size().Get(object)
	obstruction, _ := s.Obstruction().Component().Get(object)
	fromCoords, _ := from.Aligned()
	toCoords, _ := to.Aligned()
	if _, ok := s.findPath(fromCoords, toCoords, size, obstruction); !ok {
		s.Logger().Log(pathfind.ErrInvalidPath)
		return
	}
	s.Target().Set(object, pathfind.NewTarget(coords))
}
