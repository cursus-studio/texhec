package main

import (
	"core/modules/definitions"
	"core/modules/definitions/pkg"
	deploypkg "core/modules/deploy/pkg"
	"core/modules/fpslogger/pkg"
	"core/modules/generation/pkg"
	"core/modules/loading/pkg"
	pathfindpkg "core/modules/pathfind/pkg"
	playerpkg "core/modules/player/pkg"
	"core/modules/settings"
	"core/modules/settings/pkg"
	"core/modules/tile"
	"core/modules/tile/pkg"
	"core/modules/ui"
	"core/modules/ui/pkg"
	gamescenes "core/scenes"
	creditsscene "core/scenes/credits"
	gamescene "core/scenes/game"
	menuscene "core/scenes/menu"
	settingsscene "core/scenes/settings"
	"engine/modules/assets/pkg"
	"engine/modules/audio/pkg"
	"engine/modules/batcher/pkg"
	"engine/modules/camera"
	"engine/modules/camera/pkg"
	"engine/modules/collider/pkg"
	"engine/modules/connection/pkg"
	"engine/modules/drag"
	"engine/modules/drag/pkg"
	"engine/modules/grid"
	"engine/modules/groups/pkg"
	"engine/modules/hierarchy/pkg"
	"engine/modules/inputs"
	"engine/modules/inputs/pkg"
	"engine/modules/layout/pkg"
	"engine/modules/loop/pkg"
	"engine/modules/metadata/pkg"
	"engine/modules/netsync/pkg"
	"engine/modules/noise/pkg"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/modules/record"
	"engine/modules/record/pkg"
	"engine/modules/registry/pkg"
	"engine/modules/render"
	"engine/modules/render/pkg"
	"engine/modules/scene/pkg"
	"engine/modules/smooth/pkg"
	"engine/modules/text"
	"engine/modules/text/pkg"
	"engine/modules/transform"
	"engine/modules/transform/pkg"
	"engine/modules/transition/pkg"
	"engine/modules/uuid/pkg"
	"engine/modules/warmup/pkg"
	"engine/services/clock"
	"engine/services/codec"
	"engine/services/console"
	"engine/services/datastructures"
	"engine/services/ecs"
	"engine/services/graphics/texture"
	"engine/services/graphics/texturearray"
	"engine/services/logger"
	"engine/services/media"
	"engine/services/media/window"
	"time"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

func getDic() ioc.Dic {

	return ioc.NewContainer(
		clock.Pkg,
		ecs.Pkg,
		codec.Pkg,

		assetspkg.Pkg,
		logger.Pkg(logger.NewConfig(
			true,
			func(c ioc.Dic) func(message string) { return ioc.Get[console.Console](c).PrintPermanent },
		)),
		console.Pkg,
		media.Pkg,
		scenepkg.Pkg,

		gtexture.Pkg,
		texturearray.Pkg,
		tilepkg.Pkg,
		generationpkg.Pkg,
		uipkg.Pkg(uipkg.NewConfig(
			time.Millisecond*300, // animation duration
			time.Second/12,       // bgTimePerFrame
		)),
		settingspkg.Pkg,

		//

		// engine packages
		audiopkg.Pkg,
		camerapkg.Pkg,
		colliderpkg.Pkg,
		dragpkg.Pkg,
		groupspkg.Pkg,
		inputspkg.Pkg,
		looppkg.Pkg,
		prototypepkg.Pkg,
		renderpkg.Pkg,
		textpkg.Pkg(textpkg.NewConfig(
			func(c ioc.Dic) text.FontFamilyComponent {
				asset := ioc.Get[definitions.Definitions](c).FontAsset
				return text.FontFamilyComponent{FontFamily: asset}
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
			// 0.8125, // suggested (52/64)
			0.8, // arbitrary number works for some reason
		)),
		transformpkg.Pkg,
		hierarchypkg.Pkg,
		uuidpkg.Pkg,
		batcherpkg.Pkg,
		connectionpkg.Pkg,
		metadatapkg.Pkg,
		netsyncpkg.Pkg(func() netsyncpkg.Config {
			config := netsyncpkg.NewConfig(
				150, // max predictions
			)
			record.AddToConfig[transform.PosComponent](config.RecordConfig())
			record.AddToConfig[camera.OrthoComponent](config.RecordConfig())
			record.AddToConfig[grid.SquareGridComponent[tile.ID]](config.RecordConfig())
			// netsyncpkg.AddComponent[transform.PosComponent](config)
			// netsyncpkg.AddComponent[camera.OrthoComponent](config)
			// netsyncpkg.AddComponent[definition.DefinitionLinkComponent](config)
			// netsyncpkg.AddComponent[tile.PosComponent](config)

			// syncpkg.AddEvent[scenessys.ChangeSceneEvent](config)
			netsyncpkg.AddEvent[drag.DraggableEvent](config)
			netsyncpkg.AddEvent[inputs.DragEvent](config)

			netsyncpkg.AddTransparentEvent[settings.EnterSettingsEvent](config)
			netsyncpkg.AddTransparentEvent[tile.HoverEvent](config)
			netsyncpkg.AddTransparentEvent[ui.HideUiEvent](config)

			// netsyncpkg.AddEventAuthorization(config, func(c inputs.DragEvent) error {
			// 	return errors.New("no")
			// })

			return config
		}()),
		recordpkg.Pkg,
		registrypkg.Pkg,
		smoothpkg.Pkg,
		smoothpkg.PkgT[render.ColorComponent](),
		smoothpkg.PkgT[tile.PosComponent](),
		smoothpkg.PkgT[tile.RotComponent](),
		transitionpkg.Pkg,
		layoutpkg.Pkg,
		loadingpkg.Pkg,
		noisepkg.Pkg,
		warmuppkg.Pkg,

		// game packages
		deploypkg.Pkg,
		pathfindpkg.Pkg,
		fpsloggerpkg.Pkg,
		playerpkg.Pkg,

		gamescenes.Pkg,
		definitionspkg.Pkg,

		creditsscene.Pkg,
		gamescene.Pkg,
		menuscene.Pkg,
		settingsscene.Pkg,
		func(b ioc.Builder) {
			ioc.Wrap(b, func(c ioc.Dic, w window.Api) {
				w.Window().SetTitle("TEXHEC")
			})
			ioc.Wrap(b, func(c ioc.Dic, f gtexture.Factory) {
				f.Wrap(func(t gtexture.Texture) {
					t.Bind()
					defer gl.BindTexture(gl.TEXTURE_2D, 0)

					gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
					gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
					gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
					gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
				})
			})

			ioc.Wrap(b, func(c ioc.Dic, f texturearray.Factory) {
				f.Wrap(func(ta texturearray.TextureArray) {
					ta.Bind()
					defer gl.BindTexture(gl.TEXTURE_2D_ARRAY, 0)

					gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
					gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
					gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
					gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
				})
			})
		},
	)
}
