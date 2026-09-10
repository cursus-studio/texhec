package internal

import (
	"core/modules/pathfind"
	"core/modules/tile"
	"fmt"
)

func (s *service) FindPath(e pathfind.FindPathEvent) {
	s.Logger().Info(fmt.Errorf("find path"))
	entity, ok := s.UUID().Entity(e.UUID)
	if !ok {
		return
	}
	from, ok := s.Tile().Pos().Get(entity)
	if !ok {
		s.Logger().Log(tile.ErrInvalidPosition)
		return
	}
	to := tile.NewPos(e.Coords.Coords())
	size, _ := s.Tile().Size().Get(entity)
	obstruction, _ := s.Obstruction().Component().Get(entity)
	fromCoords, _ := from.Aligned()
	toCoords, _ := to.Aligned()
	if _, ok := s.findPath(fromCoords, toCoords, size, obstruction); !ok {
		s.Logger().Log(pathfind.ErrInvalidPath)
		return
	}
	s.Target().Set(entity, pathfind.NewTarget(e.Coords))
}
