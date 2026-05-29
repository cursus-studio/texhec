package internal

import (
	"core/game"
	"core/modules/reach"
	"core/modules/tile"
	"engine/modules/grid"
	"engine/services/ecs"
	"math"

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
	return dist <= tile.Coord(reach.Reach)
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

	//

	start := grid.NewCoords(grid.Coord(pos.X), grid.Coord(pos.Y))
	r := grid.Coord(math.Sqrt(float64(reachComp.Reach)))

	min := grid.NewCoords(start.X-r, start.Y-r)
	max := grid.NewCoords(start.X+size.X+r-1, start.Y+size.Y+r-1)

	tiles := make([]grid.Coords, 0, (max.X-min.X)*(max.Y-min.Y))

	for y := min.Y; y <= max.Y || (min.Y > max.Y && y >= min.Y); y++ {
		for x := min.X; x <= max.X || (min.X > max.X && x >= min.X); x++ {
			if x >= start.X && x < start.X+size.X &&
				y >= start.Y && y < start.Y+size.Y {
				continue
			}

			var dx, dy grid.Coord

			fromRight := start.X + size.X
			if x >= fromRight {
				dx = x - fromRight + 1
			} else if x < start.X {
				dx = start.X - x
			}

			fromBottom := start.Y + size.Y
			if y >= fromBottom {
				dy = y - fromBottom + 1
			} else if y < start.Y {
				dy = start.Y - y
			}

			if (dx*dx)+(dy*dy) <= grid.Coord(reachComp.Reach) {
				tiles = append(tiles, grid.NewCoords(grid.Coord(x), grid.Coord(y)))
			}
		}
	}
	return tiles
}
