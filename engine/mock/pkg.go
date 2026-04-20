package mock

import (
	assetspkg "engine/modules/assets/pkg"
	audiopkg "engine/modules/audio/pkg"
	batcherpkg "engine/modules/batcher/pkg"
	camerapkg "engine/modules/camera/pkg"
	colliderpkg "engine/modules/collider/pkg"
	connectionpkg "engine/modules/connection/pkg"
	groupspkg "engine/modules/groups/pkg"
	hierarchypkg "engine/modules/hierarchy/pkg"
	inputspkg "engine/modules/inputs/pkg"
	layoutpkg "engine/modules/layout/pkg"
	metadatapkg "engine/modules/metadata/pkg"
	netsyncpkg "engine/modules/netsync/pkg"
	noisepkg "engine/modules/noise/pkg"
	prototypepkg "engine/modules/prototype/pkg"
	recordpkg "engine/modules/record/pkg"
	registrypkg "engine/modules/registry/pkg"
	renderpkg "engine/modules/render/pkg"
	scenepkg "engine/modules/scene/pkg"
	"engine/modules/text"
	textpkg "engine/modules/text/pkg"
	transformpkg "engine/modules/transform/pkg"
	transitionpkg "engine/modules/transition/pkg"
	uuidpkg "engine/modules/uuid/pkg"
	warmuppkg "engine/modules/warmup/pkg"
	"engine/services/clock"
	"engine/services/codec"
	"engine/services/console"
	"engine/services/datastructures"
	"engine/services/ecs"
	"engine/services/graphics/texturearray"
	"engine/services/logger"
	"engine/services/media/window"
	"engine/services/runtime"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		ecs.Pkg,
		assetspkg.Pkg,
		audiopkg.Pkg,
		batcherpkg.Pkg(batcherpkg.NewConfig(1, time.Second)),
		camerapkg.Pkg,
		colliderpkg.Pkg,
		connectionpkg.Pkg,
		groupspkg.Pkg,
		hierarchypkg.Pkg,
		inputspkg.Pkg,
		layoutpkg.Pkg,
		metadatapkg.Pkg,
		netsyncpkg.Pkg(netsyncpkg.NewConfig(0)),
		noisepkg.Pkg,
		prototypepkg.Pkg,
		recordpkg.Pkg,
		registrypkg.Pkg,
		renderpkg.Pkg,
		scenepkg.Pkg,
		textpkg.Pkg(textpkg.NewConfig(
			func(c ioc.Dic) text.FontFamilyComponent {
				return text.FontFamilyComponent{}
			},
			text.FontSizeComponent{FontSize: 16},
			text.BreakComponent{Break: text.BreakWord},
			text.TextAlignComponent{Vertical: 0, Horizontal: 0},
			text.TextColorComponent{Color: mgl32.Vec4{1, 1, 1, 1}},
			func() datastructures.SparseSet[rune] {
				set := datastructures.NewSparseSet[rune]()
				for i := int32('a'); i <= int32('z'); i++ {
					set.Add(rune(i))
				}
				for i := int32('A'); i <= int32('Z'); i++ {
					set.Add(rune(i))
				}
				for i := int32('0'); i <= int32('9'); i++ {
					set.Add(rune(i))
				}
				for i := int32('!'); i <= int32('/'); i++ {
					set.Add(rune(i))
				}
				for i := int32(':'); i <= int32('@'); i++ {
					set.Add(rune(i))
				}
				for i := int32('['); i <= int32('`'); i++ {
					set.Add(rune(i))
				}
				for i := int32('{'); i <= int32('~'); i++ {
					set.Add(rune(i))
				}
				set.Add(' ')

				return set
			}(),
			64,
			0.8, // arbitrary number works for some reason

		)),
		transformpkg.Pkg,
		transitionpkg.Pkg,
		uuidpkg.Pkg,
		warmuppkg.Pkg,

		clock.Pkg,
		codec.Pkg,
		console.Pkg,
		logger.Pkg(logger.NewConfig(
			true,
			func(c ioc.Dic) func(message string) { return func(message string) { print(message) } },
		)),
		runtime.Pkg,

		window.Pkg(window.NewConfig(nil, nil)),
		texturearray.Pkg,
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
})
