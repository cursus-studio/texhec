package test

import (
	corepkg "core/pkg"
	gamescenes "core/scenes"

	"github.com/ogiusek/ioc/v2"
)

type Setup struct {
	gamescenes.GameWorld `inject:""`
}

func NewSetup() Setup {
	c := ioc.NewContainer(corepkg.Pkg)
	setup := ioc.GetServices[Setup](c)
	return setup
}
