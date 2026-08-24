package main

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/obstruction"
	"core/modules/tile"
	corepkg "core/pkg"
	colliderpkg "engine/modules/collider/pkg"
	"engine/modules/graphics"
	"engine/modules/grid"
	gridpkg "engine/modules/grid/pkg"
	"engine/modules/logger"
	loggerpkg "engine/modules/logger/pkg"
	netsyncpkg "engine/modules/netsync/pkg"
	"engine/modules/record"
	"engine/modules/seed"
	"engine/modules/text"
	textpkg "engine/modules/text/pkg"
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
				for i := rune(32); i <= 126; i++ {
					c.UsedGlyphs().Add(i)
				}

				// 2. Polish Characters (Latin Extended-A)
				polishRunes := []rune{'ą', 'ć', 'ę', 'ł', 'ń', 'ó', 'ś', 'ź', 'ż', 'Ą', 'Ć', 'Ę', 'Ł', 'Ń', 'Ó', 'Ś', 'Ź', 'Ż'}
				for _, r := range polishRunes {
					c.UsedGlyphs().Add(r)
				}
			})
			ioc.Wrap(b, func(c ioc.Dic, s text.Service) {
				world := ioc.Get[game.GameWorld](c)
				s.FontFamily().SetEmpty(text.NewFontFamily(world.Definitions().Assets().FontAsset))
			})
			ioc.Wrap(b, func(c ioc.Dic, config loggerpkg.Config) {
				world := ioc.Get[game.GameWorld](c)
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
				})
			})

			ioc.Wrap(b, func(c ioc.Dic, config netsyncpkg.Config) {
				config.SetMaxPredictions(150)
				// i want to sync:
				// - world: seed
				record.AddToConfig[seed.SeedComponent](config.RecordConfig())
				// - chunks: chukn coords, grid content
				record.AddToConfig[grid.ChunkCoordsComponent](config.RecordConfig())
				record.AddToConfig[grid.ChunkComponent[tile.ID]](config.RecordConfig())
				record.AddToConfig[grid.ChunkComponent[obstruction.Obstruction]](config.RecordConfig())
				// - objects: coords, blueprint, owner, deployed mark
				// - players: name, wallet
			})
			ioc.Wrap(b, func(c ioc.Dic, config colliderpkg.Config) {
				tileSize := ioc.Get[gridpkg.Config](c).GetTileSize()
				chunkSize := float32(ioc.Get[gridpkg.Config](c).GetChunkSize().Val())
				config.SetChunkSize(tileSize * chunkSize / 2)
			})
		},
	}

	return ioc.NewContainer(pkgs...)
}
