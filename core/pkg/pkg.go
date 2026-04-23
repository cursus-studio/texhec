package corepkg

import (
	definitionspkg "core/modules/definitions/pkg"
	deploypkg "core/modules/deploy/pkg"
	fpsloggerpkg "core/modules/fpslogger/pkg"
	generationpkg "core/modules/generation/pkg"
	loadingpkg "core/modules/loading/pkg"
	pathfindpkg "core/modules/pathfind/pkg"
	playerpkg "core/modules/player/pkg"
	settingspkg "core/modules/settings/pkg"
	tilepkg "core/modules/tile/pkg"
	gamescenes "core/scenes"
	creditsscene "core/scenes/credits"
	gamescene "core/scenes/game"
	menuscene "core/scenes/menu"
	settingsscene "core/scenes/settings"
	enginepkg "engine/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		enginepkg.Pkg,

		definitionspkg.Pkg,
		deploypkg.Pkg,
		fpsloggerpkg.Pkg,
		generationpkg.Pkg,
		loadingpkg.Pkg,
		pathfindpkg.Pkg,
		playerpkg.Pkg,
		settingspkg.Pkg,
		tilepkg.Pkg,

		gamescenes.Pkg,
		creditsscene.Pkg,
		gamescene.Pkg,
		menuscene.Pkg,
		settingsscene.Pkg,
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
})
