package test

import (
	"core/game"
	corepkg "core/pkg"

	"github.com/ogiusek/ioc/v2"
)

type Setup struct {
	game.GameWorld `inject:""`
}

func NewSetup() Setup {
	return ioc.GetServices[Setup](ioc.NewContainer(corepkg.Pkg))
}
