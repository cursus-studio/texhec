package main

import (
	"core/modules/settings"
	"core/modules/tile"
	"core/modules/ui"
	corepkg "core/pkg"
	gamescenes "core/scenes"
	"engine/modules/camera"
	"engine/modules/drag"
	"engine/modules/grid"
	"engine/modules/inputs"
	netsyncpkg "engine/modules/netsync/pkg"
	"engine/modules/record"
	"engine/modules/render"
	smoothpkg "engine/modules/smooth/pkg"
	"engine/modules/text"
	textpkg "engine/modules/text/pkg"
	"engine/modules/transform"
	"engine/modules/window"
	"engine/services/console"
	gtexture "engine/services/graphics/texture"
	"engine/services/graphics/texturearray"
	"engine/services/logger"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/ogiusek/ioc/v2"
)

func getDic() ioc.Dic {
	pkgs := []ioc.Pkg{
		corepkg.Pkg,
		// game
		smoothpkg.PkgT[render.ColorComponent],

		func(b ioc.Builder) {
			ioc.Wrap(b, func(c ioc.Dic, w window.Service) {
				w.Window().SetTitle("TEXHEC")
				gl.ClearColor(0.2, 0.3, 0.3, 1.0)
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
			ioc.Wrap(b, func(_ ioc.Dic, c textpkg.Config) {
				for i := int32('a'); i <= int32('z'); i++ {
					c.UsedGlyphs().Add(rune(i))
				}
				for i := int32('A'); i <= int32('Z'); i++ {
					c.UsedGlyphs().Add(rune(i))
				}
				for i := int32('0'); i <= int32('9'); i++ {
					c.UsedGlyphs().Add(rune(i))
				}
				for i := int32('!'); i <= int32('/'); i++ {
					c.UsedGlyphs().Add(rune(i))
				}
				for i := int32(':'); i <= int32('@'); i++ {
					c.UsedGlyphs().Add(rune(i))
				}
				for i := int32('['); i <= int32('`'); i++ {
					c.UsedGlyphs().Add(rune(i))
				}
				for i := int32('{'); i <= int32('~'); i++ {
					c.UsedGlyphs().Add(rune(i))
				}
				c.UsedGlyphs().Add(' ')
			})
			ioc.Wrap(b, func(c ioc.Dic, config logger.Config) {
				config.PanicOnWarn(true)
				config.Flush(ioc.Get[console.Console](c).PrintPermanent)
			})
			ioc.Wrap(b, func(c ioc.Dic, config netsyncpkg.Config) {
				config.SetMaxPredictions(150)
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
			})
			ioc.Wrap(b, func(c ioc.Dic, s text.Service) {
				world := ioc.GetServices[*gamescenes.GameWorld](c)
				s.FontFamily().SetEmpty(text.NewFontFamily(world.Definitions().FontAsset))
			})
		},
	}

	return ioc.NewContainer(pkgs...)
}
