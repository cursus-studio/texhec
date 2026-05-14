package test

import (
	"core/game"
	corepkg "core/pkg"
	assetspkg "engine/modules/assets/pkg"
	"engine/modules/window"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/ogiusek/ioc/v2"
)

type Setup struct {
	game.GameWorld `inject:""`
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		corepkg.Pkg,
		func(b ioc.Builder) {
			ioc.Wrap(b, func(c ioc.Dic, w window.Service) {
				w.Window().SetTitle("tile module benchmark")
				gl.ClearColor(0.2, 0.3, 0.3, 1.0)
			})
			ioc.Wrap(b, func(c ioc.Dic, b assetspkg.Config) { b.SetPath("../../../assets/") })
		},
	)
	s := ioc.GetServices[Setup](c)
	return s
}
