package internal

import (
	"core/modules/pathfind"
	"core/modules/tile"
)

func (s *service) FindPath(e pathfind.FindPathEvent) {
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
}
