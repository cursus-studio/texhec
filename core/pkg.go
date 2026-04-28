package main

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/settings"
	"core/modules/tile"
	"core/modules/ui"
	corepkg "core/pkg"
	"engine/modules/camera"
	"engine/modules/drag"
	"engine/modules/graphics"
	"engine/modules/grid"
	"engine/modules/inputs"
	"engine/modules/logger"
	loggerpkg "engine/modules/logger/pkg"
	netsyncpkg "engine/modules/netsync/pkg"
	"engine/modules/record"
	"engine/modules/text"
	textpkg "engine/modules/text/pkg"
	"engine/modules/transform"
	"engine/modules/window"
	"errors"
	"fmt"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/ogiusek/ioc/v2"
)

func getDic() ioc.Dic {
	pkgs := []ioc.Pkg{
		corepkg.Pkg,
		func(b ioc.Builder) {
			ioc.Wrap(b, func(c ioc.Dic, def definitions.Service) {
				// definitions have to be loaded explicitly
				// they aren't loaded by default so tests won't look for files
				def.Load()
			})
			ioc.Wrap(b, func(c ioc.Dic, w window.Service) {
				w.Window().SetTitle("TEXHEC")
				gl.ClearColor(0.2, 0.3, 0.3, 1.0)
			})
			ioc.Wrap(b, func(c ioc.Dic, f graphics.Service) {
				f.TextureArray().Wrap(func(ta graphics.TextureArray) {
					ta.Bind()
					defer gl.BindTexture(gl.TEXTURE_2D_ARRAY, 0)

					gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
					gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
					gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
					gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
				})
				f.Texture().Wrap(func(t graphics.Texture) {
					t.Bind()
					defer gl.BindTexture(gl.TEXTURE_2D, 0)

					gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
					gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
					gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
					gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
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
			ioc.Wrap(b, func(c ioc.Dic, config loggerpkg.Config) {
				world := ioc.GetServices[game.GameWorld](c)
				config.AddFormatHandler(func(meta, msg error) error {
					typeMsg, color := "LOG", "37"
					if errors.Is(meta, logger.ErrInfo) {
						typeMsg, color = "Info", "34"
					}
					if logger.IsWarning(meta) {
						typeMsg, color = "Warn", "33"
					}
					if errors.Is(meta, logger.ErrFatal) {
						typeMsg, color = "Fatal", "31"
					}
					return fmt.Errorf(
						"\033[%sm[ %s ]\033[0m %s %s",
						color,
						typeMsg,
						world.Clock().Now().Format("15:04:05.000000"),
						msg.Error(),
					)
				})
				config.AddDeliverHandler(func(meta, msg error) {
					world.Console().PrintPermanent(msg.Error())

					if errors.Is(meta, logger.ErrFatal) {
						world.Console().Flush()
					}
					if logger.IsWarning(meta) {
						world.Console().Flush()
						panic("Debug warning")
					}
				})
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
				world := ioc.GetServices[game.GameWorld](c)
				s.FontFamily().SetEmpty(text.NewFontFamily(world.Definitions().Assets().FontAsset))
			})
		},
	}

	return ioc.NewContainer(pkgs...)
}
