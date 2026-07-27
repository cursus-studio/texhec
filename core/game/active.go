package game

import (
	"core/modules/actions"
	"core/modules/definitions"
	"core/modules/deploy"
	"core/modules/fpslogger"
	"core/modules/generation"
	"core/modules/loading"
	"core/modules/obstruction"
	"core/modules/pathfind"
	"core/modules/player"
	"core/modules/reach"
	"core/modules/settings"
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/modules/scene"

	"github.com/ogiusek/ioc/v2"
)

type GameWorld struct {
	engine.EngineWorld `inject:""`

	// game
	Actions     ioc.Lazy[actions.Service]     `inject:""`
	Definitions ioc.Lazy[definitions.Service] `inject:""`
	Deploy      ioc.Lazy[deploy.Service]      `inject:""`
	FpsLogger   ioc.Lazy[fpslogger.Service]   `inject:""`
	Generation  ioc.Lazy[generation.Service]  `inject:""`
	Loading     ioc.Lazy[loading.Service]     `inject:""`
	Obstruction ioc.Lazy[obstruction.Service] `inject:""`
	Pathfind    ioc.Lazy[pathfind.Service]    `inject:""`
	Player      ioc.Lazy[player.Service]      `inject:""`
	Reach       ioc.Lazy[reach.Service]       `inject:""`
	Settings    ioc.Lazy[settings.Service]    `inject:""`
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
