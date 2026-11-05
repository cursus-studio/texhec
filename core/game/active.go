package game

import (
	"core/modules/actions"
	"core/modules/attack"
	"core/modules/definitions"
	"core/modules/deploy"
	"core/modules/economy"
	"core/modules/fpslogger"
	"core/modules/generation"
	"core/modules/loading"
	"core/modules/lobby"
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
	Attack      ioc.Lazy[attack.Service]      `inject:""`
	Definitions ioc.Lazy[definitions.Service] `inject:""`
	Deploy      ioc.Lazy[deploy.Service]      `inject:""`
	Economy     ioc.Lazy[economy.Service]     `inject:""`
	FpsLogger   ioc.Lazy[fpslogger.Service]   `inject:""`
	Generation  ioc.Lazy[generation.Service]  `inject:""`
	Loading     ioc.Lazy[loading.Service]     `inject:""`
	Lobby       ioc.Lazy[lobby.Service]       `inject:""`
	Obstruction ioc.Lazy[obstruction.Service] `inject:""`
	Pathfind    ioc.Lazy[pathfind.Service]    `inject:""`
	Player      ioc.Lazy[player.Service]      `inject:""`
	Reach       ioc.Lazy[reach.Service]       `inject:""`
	Settings    ioc.Lazy[settings.Service]    `inject:""`
	Tile        ioc.Lazy[tile.Service]        `inject:""`
	Ui          ioc.Lazy[ui.Service]          `inject:""`
}

// [PlayBuilder] [SettingsBuilder] [CreditsBuilder]
type MainMenuBuilder scene.Scene

//

// [PlayConnectBuilder] [PlayGenerateBuilder] [PlayLoadBuilder]
type PlayBuilder scene.Scene
type SettingsBuilder scene.Scene
type CreditsBuilder scene.Scene

//

// 1. set ip (set by scene)
// 2. preview [HostForm] and [LobbyForm] (synced by connection, creates a view and removed set ip view)
// 3. [StartGameEvent] (event creates a view)
type PlayConnectBuilder scene.Scene

// 1. [HostComponent]. preview [HostForm] and [LobbyForm]
// 2. [StartGameEvent]
type PlayGenerateBuilder scene.Scene

// 1. [HostComponent]. preview [HostForm] and [LobbyForm]
// 2. [StartGameEvent]
type PlayLoadBuilder scene.Scene

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Wrap(b, func(c ioc.Dic, b scene.Service) {
		b.SetScene(definitions.MainMenuID, scene.Scene(ioc.Get[MainMenuBuilder](c)))

		b.SetScene(definitions.PlayID, scene.Scene(ioc.Get[PlayBuilder](c)))
		b.SetScene(definitions.SettingsID, scene.Scene(ioc.Get[SettingsBuilder](c)))
		b.SetScene(definitions.CreditsID, scene.Scene(ioc.Get[CreditsBuilder](c)))

		b.SetScene(definitions.PlayConnectID, scene.Scene(ioc.Get[PlayConnectBuilder](c)))
		b.SetScene(definitions.PlayGenerateID, scene.Scene(ioc.Get[PlayGenerateBuilder](c)))
		// b.SetScene(definitions.PlayLoadID, scene.Scene(ioc.Get[PlayLoadBuilder](c)))
	})
})
