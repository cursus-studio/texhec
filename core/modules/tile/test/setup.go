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
	c := ioc.NewContainer(corepkg.Pkg)
	setup := ioc.GetServices[Setup](c)
	return setup
}
