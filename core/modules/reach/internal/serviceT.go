package internal

import (
	"core/game"
	"core/modules/reach"
	"core/modules/tile"
	"engine/modules/grid"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type serviceT[Component any] struct {
	game.GameWorld `inject:""`
	component      ecs.ComponentsArray[reach.Component[Component]]
}

func NewServiceT[Component any](c ioc.Dic) reach.ServiceT[Component] {
	s := ioc.GetServices[*serviceT[Component]](c)
	s.component = ecs.GetComponentsArray[reach.Component[Component]](s.World())
	return s
}

func (s *serviceT[Component]) Component() ecs.ComponentsArray[reach.Component[Component]] {
	return s.component
}
func (s *serviceT[Component]) Reaches(fromEntity, toEntity ecs.EntityID) bool {
	reach, _ := s.component.Get(fromEntity)
	if reach.Reach == 0 {
		return false
	}

	fromPos, ok := s.Tile().Pos().Get(fromEntity)
	if !ok {
		return false
	}
	fromSize, _ := s.Tile().Size().Get(fromEntity)

	toPos, ok := s.Tile().Pos().Get(toEntity)
	if !ok {
		return false
	}
	toSize, _ := s.Tile().Size().Get(toEntity)

	dist := s.Reach().Distance(fromPos, fromSize, toPos, toSize)
	return dist <= reach.Reach
}
func (s *serviceT[Component]) TilesWithinReach(entity ecs.EntityID) []grid.Coords {
	reachComp, _ := s.component.Get(entity)
	if reachComp.Reach == 0 {
		return nil
	}
	pos, ok := s.Tile().Pos().Get(entity)
	if !ok {
		return nil
	}
	size, _ := s.Tile().Size().Get(entity)

	if pos.X != tile.Coord(int(pos.X)) || pos.Y != tile.Coord(int(pos.Y)) {
		return nil
	}
	startX, startY := grid.Coord(pos.X), grid.Coord(pos.Y)
	sizeX := size.X
	sizeY := size.Y
	r := grid.Coord(reachComp.Reach)

	minX := startX - r
	minY := startY - r
	maxX := startX + sizeX + r - 1
	maxY := startY + sizeY + r - 1

	var tiles []grid.Coords

	for y := minY; y <= maxY || (minY > maxY && y >= minY); y++ {
		for x := minX; x <= maxX || (minX > maxX && x >= minX); x++ {
			if x >= startX && x < startX+sizeX &&
				y >= startY && y < startY+sizeY {
				continue
			}

			var dx, dy grid.Coord

			fromRight := startX + sizeX
			if x >= fromRight {
				dx = x - fromRight + 1
			} else if x < startX {
				dx = startX - x
			}

			fromBottom := startY + sizeY
			if y >= fromBottom {
				dy = y - fromBottom + 1
			} else if y < startY {
				dy = startY - y
			}

			if (dx*dx)+(dy*dy) <= r*r {
				tiles = append(tiles, grid.NewCoords(grid.Coord(x), grid.Coord(y)))
			}
		}
	}
	return tiles
}
