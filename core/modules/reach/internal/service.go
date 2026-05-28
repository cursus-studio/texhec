package internal

import (
	"core/game"
	"core/modules/reach"
	"core/modules/tile"
	"engine/modules/grid"
	"math"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld `inject:""`
}

func NewService(c ioc.Dic) reach.Service {
	s := ioc.GetServices[*service](c)
	return s
}

func (s *service) Distance(
	fromPos tile.PosComponent, fromSize tile.SizeComponent,
	toPos tile.PosComponent, toSize tile.SizeComponent,
) tile.Coord {
	if fromPos.X != tile.Coord(grid.Coord(fromPos.X)) ||
		fromPos.Y != tile.Coord(grid.Coord(fromPos.Y)) {
		return tile.Coord(math.MaxFloat32)
	}
	var dx, dy grid.Coord

	from := grid.NewCoords(grid.Coord(fromPos.X), grid.Coord(fromPos.Y))
	to := grid.NewCoords(grid.Coord(toPos.X), grid.Coord(toPos.Y))

	if fromRight := from.X + fromSize.X - 1; to.X >= fromRight {
		dx = to.X - fromRight
	} else if toRight := to.X + toSize.X - 1; from.X >= toRight {
		dx = from.X - toRight
	}

	if fromBottom := from.Y + fromSize.Y - 1; to.Y >= fromBottom {
		dy = to.Y - fromBottom
	} else if toBottom := to.Y + toSize.Y - 1; from.Y >= toBottom {
		dy = from.Y - toBottom
	}

	return tile.Coord((dx * dx) + (dy * dy))
}
