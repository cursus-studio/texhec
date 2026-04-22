package main

import (
	"core/modules/fpslogger"
	"core/modules/loading"
	"core/modules/settings"
	"core/modules/tile"
	"core/modules/ui"
	gamescenes "core/scenes"
	_ "embed"
	"engine/modules/audio"
	"engine/modules/batcher"
	"engine/modules/camera"
	"engine/modules/connection"
	"engine/modules/drag"
	"engine/modules/inputs"
	"engine/modules/netsync"
	"engine/modules/render"
	"engine/modules/scene"
	"engine/modules/smooth"
	"engine/modules/text"
	"engine/modules/transition"
	"engine/services/ecs"
	"engine/services/logger"
	"engine/services/media/window"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

// golangci-lint run --fix
func main() {
	print("started\n")

	{ // go tool pprof -http=:8080 ignore.cpu.pprof
		name := ""
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
		f, err := os.Create(fmt.Sprintf("ignore.cpu.pprof%v", name))
		if err != nil {
			panic(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	runtime.LockOSThread()

	c := getDic()

	gl.ClearColor(0.2, 0.3, 0.3, 1.0)

	// load world before starting timer
	events.Emit(ioc.Get[events.Events](c), scene.NewChangeSceneEvent(gamescenes.GameID))

	{ // before start
		logger := ioc.Get[logger.Logger](c)
		eventsBuilder := ioc.Get[events.Builder](c)
		events.GlobalErrHandler(eventsBuilder, func(err error) {
			logger.Warn(err)
		})

		temporaryInlineSystems := ecs.NewSystemRegister(func() error {
			events.Listen(eventsBuilder, func(e sdl.KeyboardEvent) {
				if e.Keysym.Sym == sdl.K_q {
					logger.Info("quiting program due to pressing 'Q'")
					events.Emit(eventsBuilder.Events(), inputs.NewQuitEvent())
				}
				if e.Keysym.Sym == sdl.K_ESCAPE {
					logger.Info("quiting program due to pressing 'ESC'")
					events.Emit(eventsBuilder.Events(), inputs.NewQuitEvent())
				}
				if e.State == sdl.PRESSED && e.Keysym.Sym == sdl.K_f {
					logger.Info("toggling screen size due to pressing 'F'")
					window := ioc.Get[window.Api](c)
					flags := window.Window().GetFlags()
					if flags&sdl.WINDOW_FULLSCREEN_DESKTOP == sdl.WINDOW_FULLSCREEN_DESKTOP {
						_ = window.Window().SetFullscreen(0)
					} else {
						_ = window.Window().SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
					}
				}
			})

			return nil
		})

		errs := ecs.RegisterSystems(
			ioc.Get[netsync.StartSystem](c),
			ioc.Get[smooth.StartSystem](c),
			// update {
			ioc.Get[connection.System](c),

			// inputs
			ioc.Get[inputs.System](c),
			ioc.Get[audio.System](c),

			// update
			ioc.Get[camera.System](c),
			ioc.Get[drag.System](c),
			ioc.Get[transition.System](c),
			temporaryInlineSystems,

			ioc.Get[tile.System](c),

			// ui update
			ioc.Get[ui.System](c),
			ioc.Get[settings.System](c),
			ioc.Get[loading.System](c),
			ioc.Get[batcher.System](c),
			// } (update)

			ioc.Get[smooth.StopSystem](c),
			ioc.Get[netsync.StopSystem](c),

			ioc.Get[inputs.ShutdownSystem](c), // after batcher and before render system

			// render
			ioc.Get[render.System](c),
			ioc.Get[tile.SystemRenderer](c),
			ioc.Get[render.SystemRenderer](c),
			ioc.Get[text.SystemRenderer](c),
			ioc.Get[fpslogger.System](c),
		)
		for _, err := range errs {
			logger.Warn(err)
		}
	}
	game := ioc.GetServices[gamescenes.GameWorld](c)
	err := game.Frames().Run()
	game.Logger().Warn(err)
}
