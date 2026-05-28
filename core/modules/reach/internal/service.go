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
	if fromPos.X != tile.Coord(int(fromPos.X)) ||
		fromPos.Y != tile.Coord(int(fromPos.Y)) {
		return tile.Coord(math.MaxFloat32)
	}
	var dx, dy grid.Coord

	from := grid.NewCoords(
		grid.Coord(math.Round(float64(fromPos.X))),
		grid.Coord(math.Round(float64(fromPos.Y))),
	)
	to := grid.NewCoords(
		grid.Coord(math.Round(float64(toPos.X))),
		grid.Coord(math.Round(float64(toPos.Y))),
	)

	fromRight := from.X + fromSize.X - 1
	toRight := to.X + toSize.X - 1

	if to.X >= fromRight {
		dx = to.X - fromRight
	} else if from.X >= toRight {
		dx = from.X - toRight
	}

	fromBottom := from.Y + fromSize.Y - 1
	toBottom := to.Y + toSize.Y - 1

	if to.Y >= fromBottom {
		dy = to.Y - fromBottom
	} else if from.Y >= toBottom {
		dy = from.Y - toBottom
	}

	totalDist := float64((dx * dx) + (dy * dy))
	return tile.Coord(totalDist)
}
