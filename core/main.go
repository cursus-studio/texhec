package main

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/fpslogger"
	"core/modules/loading"
	"core/modules/settings"
	"core/modules/tile"
	"core/modules/ui"
	_ "embed"
	"engine/modules/audio"
	"engine/modules/batcher"
	"engine/modules/camera"
	"engine/modules/connection"
	"engine/modules/drag"
	"engine/modules/inputs"
	"engine/modules/logger"
	"engine/modules/loop"
	"engine/modules/netsync"
	"engine/modules/render"
	"engine/modules/scene"
	"engine/modules/smooth"
	"engine/modules/text"
	"engine/modules/transition"
	"engine/services/ecs"
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

// golangci-lint run --fix
func main() {
	print("started main\n")

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

	c := getDic()
	world := ioc.GetServices[game.GameWorld](c)

	// load world before starting timer
	events.Emit(world.Events(), scene.NewChangeSceneEvent(definitions.GameID))

	// before start
	events.GlobalErrHandler(world.EventsBuilder(), func(err error) {
		world.Logger().Log(err)
	})

	temporaryInlineSystems := ecs.NewSystemRegister(func() error {
		events.Listen(world.EventsBuilder(), func(e sdl.KeyboardEvent) {
			if e.Keysym.Sym == sdl.K_q {
				world.Logger().Log(errors.Join(logger.ErrInfo, errors.New("quiting program due to pressing 'Q'")))
				events.Emit(world.Events(), loop.NewStopEvent())
			}
			if e.Keysym.Sym == sdl.K_ESCAPE {
				world.Logger().Log(errors.Join(logger.ErrInfo, errors.New("quiting program due to pressing 'ESC'")))
				events.Emit(world.Events(), loop.NewStopEvent())
			}
			if e.State == sdl.PRESSED && e.Keysym.Sym == sdl.K_f {
				world.Logger().Log(errors.Join(logger.ErrInfo, errors.New("toggling screen size due to pressing 'F'")))
				flags := world.Window().Window().GetFlags()
				if flags&sdl.WINDOW_FULLSCREEN_DESKTOP == sdl.WINDOW_FULLSCREEN_DESKTOP {
					_ = world.Window().Window().SetFullscreen(0)
				} else {
					_ = world.Window().Window().SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
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

		// render
		ioc.Get[render.System](c),
		ioc.Get[tile.SystemRenderer](c),
		ioc.Get[render.SystemRenderer](c),
		ioc.Get[text.SystemRenderer](c),
		ioc.Get[fpslogger.System](c),
	)
	for _, err := range errs {
		world.Logger().Log(err)
	}

	game := ioc.GetServices[game.GameWorld](c)
	game.Logger().Log(errors.Join(logger.ErrInfo, errors.New("initialized engine")))
	runtime.LockOSThread()
	game.Loop().Run(loop.NewConfigureEvent(60, 1))
}
