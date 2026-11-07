package corepkg

import (
	game "core/game"
	creditsscene "core/game/credits"
	menuscene "core/game/menu"
	playscene "core/game/play"
	settingsscene "core/game/settings"
	actionspkg "core/modules/actions/pkg"
	attackpkg "core/modules/attack/pkg"
	definitionspkg "core/modules/definitions/pkg"
	deploypkg "core/modules/deploy/pkg"
	economypkg "core/modules/economy/pkg"
	fpsloggerpkg "core/modules/fpslogger/pkg"
	generationpkg "core/modules/generation/pkg"
	loadingpkg "core/modules/loading/pkg"
	lobbypkg "core/modules/lobby/pkg"
	obstructionpkg "core/modules/obstruction/pkg"
	pathfindpkg "core/modules/pathfind/pkg"
	playerpkg "core/modules/player/pkg"
	reachpkg "core/modules/reach/pkg"
	settingspkg "core/modules/settings/pkg"
	tilepkg "core/modules/tile/pkg"
	uipkg "core/modules/ui/pkg"
	enginepkg "engine/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		enginepkg.Pkg,

		actionspkg.Pkg,
		attackpkg.Pkg,
		definitionspkg.Pkg,
		deploypkg.Pkg,
		economypkg.Pkg,
		fpsloggerpkg.Pkg,
		generationpkg.Pkg,
		loadingpkg.Pkg,
		lobbypkg.Pkg,
		obstructionpkg.Pkg,
		pathfindpkg.Pkg,
		playerpkg.Pkg,
		reachpkg.Pkg,
		settingspkg.Pkg,
		tilepkg.Pkg,
		uipkg.Pkg,

		game.Pkg,
		creditsscene.Pkg,
		playscene.Pkg,
		menuscene.Pkg,
		settingsscene.Pkg,
		func(b ioc.Builder) {
			ioc.Register(b, func(c ioc.Dic) game.GameWorld {
				return ioc.GetServices[game.GameWorld](c)
			})
		},
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
})
