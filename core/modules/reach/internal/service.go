package internal

import (
	"core/game"
	"core/modules/reach"
	"core/modules/tile"

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
	from tile.PosComponent, fromSize tile.SizeComponent,
	to tile.PosComponent, toSize tile.SizeComponent,
) tile.Coord {
	var dx, dy tile.Coord

	if fromRight := from.X + tile.Coord(fromSize.X) - 1; to.X >= fromRight {
		dx = to.X - fromRight
	} else if toRight := to.X + tile.Coord(toSize.X) - 1; from.X >= toRight {
		dx = from.X - toRight
	}

	if fromBottom := from.Y + tile.Coord(fromSize.Y) - 1; to.Y >= fromBottom {
		dy = to.Y - fromBottom
	} else if toBottom := to.Y + tile.Coord(toSize.Y) - 1; from.Y >= toBottom {
		dy = from.Y - toBottom
	}

	return tile.Coord((dx * dx) + (dy * dy))
}
