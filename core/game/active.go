package game

import (
	"core/modules/definitions"
	"core/modules/deploy"
	"core/modules/generation"
	"core/modules/pathfind"
	"core/modules/player"
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/modules/scene"

	"github.com/ogiusek/ioc/v2"
)

type GameWorld struct {
	engine.EngineWorld `inject:""`

	// game
	Definitions ioc.Lazy[definitions.Service] `inject:""`
	Deploy      ioc.Lazy[deploy.Service]      `inject:""`
	Generation  ioc.Lazy[generation.Service]  `inject:""`
	Pathfind    ioc.Lazy[pathfind.Service]    `inject:""`
	Player      ioc.Lazy[player.Service]      `inject:""`
	Tile        ioc.Lazy[tile.Service]        `inject:""`
	Ui          ioc.Lazy[ui.Service]          `inject:""`
}

type MenuBuilder scene.Scene
type GameBuilder scene.Scene
type SettingsBuilder scene.Scene
type CreditsBuilder scene.Scene

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Wrap(b, func(c ioc.Dic, b scene.Service) {
		b.SetScene(definitions.MenuID, scene.Scene(ioc.Get[MenuBuilder](c)))
		b.SetScene(definitions.GameID, scene.Scene(ioc.Get[GameBuilder](c)))
		b.SetScene(definitions.SettingsID, scene.Scene(ioc.Get[SettingsBuilder](c)))
		b.SetScene(definitions.CreditsID, scene.Scene(ioc.Get[CreditsBuilder](c)))
	})
})
