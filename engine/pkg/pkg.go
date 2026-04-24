package enginepkg

import (
	assetspkg "engine/modules/assets/pkg"
	audiopkg "engine/modules/audio/pkg"
	batcherpkg "engine/modules/batcher/pkg"
	camerapkg "engine/modules/camera/pkg"
	codecpkg "engine/modules/codec/pkg"
	colliderpkg "engine/modules/collider/pkg"
	connectionpkg "engine/modules/connection/pkg"
	dragpkg "engine/modules/drag/pkg"
	groupspkg "engine/modules/groups/pkg"
	hierarchypkg "engine/modules/hierarchy/pkg"
	inputspkg "engine/modules/inputs/pkg"
	layoutpkg "engine/modules/layout/pkg"
	looppkg "engine/modules/loop/pkg"
	metadatapkg "engine/modules/metadata/pkg"
	netsyncpkg "engine/modules/netsync/pkg"
	noisepkg "engine/modules/noise/pkg"
	prototypepkg "engine/modules/prototype/pkg"
	recordpkg "engine/modules/record/pkg"
	registrypkg "engine/modules/registry/pkg"
	renderpkg "engine/modules/render/pkg"
	scenepkg "engine/modules/scene/pkg"
	smoothpkg "engine/modules/smooth/pkg"
	textpkg "engine/modules/text/pkg"
	transformpkg "engine/modules/transform/pkg"
	transitionpkg "engine/modules/transition/pkg"
	uuidpkg "engine/modules/uuid/pkg"
	warmuppkg "engine/modules/warmup/pkg"
	windowpkg "engine/modules/window/pkg"
	"engine/services/clock"
	"engine/services/console"
	"engine/services/ecs"
	"engine/services/graphics/texture"
	"engine/services/graphics/texturearray"
	"engine/services/logger"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		assetspkg.Pkg,
		audiopkg.Pkg,
		batcherpkg.Pkg,
		camerapkg.Pkg,
		codecpkg.Pkg,
		colliderpkg.Pkg,
		connectionpkg.Pkg,
		dragpkg.Pkg,
		groupspkg.Pkg,
		hierarchypkg.Pkg,
		inputspkg.Pkg,
		layoutpkg.Pkg,
		looppkg.Pkg,
		metadatapkg.Pkg,
		netsyncpkg.Pkg,
		noisepkg.Pkg,
		prototypepkg.Pkg,
		recordpkg.Pkg,
		registrypkg.Pkg,
		renderpkg.Pkg,
		scenepkg.Pkg,
		smoothpkg.Pkg,
		textpkg.Pkg,
		transformpkg.Pkg,
		transitionpkg.Pkg,
		uuidpkg.Pkg,
		warmuppkg.Pkg,
		windowpkg.Pkg,

		clock.Pkg,
		console.Pkg,
		ecs.Pkg,
		gtexture.Pkg,
		texturearray.Pkg,
		logger.Pkg,
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
})
