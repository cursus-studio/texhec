package gamescenes

import (
	"core/modules/definitions"
	"core/modules/deploy"
	"core/modules/generation"
	"core/modules/pathfind"
	"core/modules/player"
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/modules/audio"
	"engine/modules/scene"

	"github.com/ogiusek/ioc/v2"
)

var (
	MenuID     = scene.NewSceneId("menu")
	GameID     = scene.NewSceneId("game")
	SettingsID = scene.NewSceneId("settings")
	CreditsID  = scene.NewSceneId("credits")
)

const (
	EffectChannel audio.Channel = iota
	MusicChannel
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
		b.SetScene(MenuID, scene.Scene(ioc.Get[MenuBuilder](c)))
		b.SetScene(GameID, scene.Scene(ioc.Get[GameBuilder](c)))
		b.SetScene(SettingsID, scene.Scene(ioc.Get[SettingsBuilder](c)))
		b.SetScene(CreditsID, scene.Scene(ioc.Get[CreditsBuilder](c)))
	})
})
