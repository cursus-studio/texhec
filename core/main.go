package main

import (
	"core/game"
	"core/modules/definitions"
	"engine/modules/inputs"
	"engine/modules/loop"
	"engine/modules/scene"
	"engine/services/ecs"
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
		f, err := os.Create(filepath.Base(fmt.Sprintf("ignore.cpu.pprof%v", name)))
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

	// before start
	events.GlobalErrHandler(world.EventsBuilder(), func(err error) {
		world.Logger().Log(err)
	})

	temporaryInlineSystems := ecs.NewSystemRegister(func() error {
		events.Listen(world.EventsBuilder(), func(e inputs.KeyboardEvent) {
			if e.Keysym.Sym == sdl.K_q {
				world.Logger().Info(errors.New("quiting program due to pressing 'Q'"))
				events.Emit(world.Events(), loop.NewStopEvent())
			}
			if e.Keysym.Sym == sdl.K_ESCAPE {
				world.Logger().Info(errors.New("quiting program due to pressing 'ESC'"))
				events.Emit(world.Events(), loop.NewStopEvent())
			}
			if e.State == sdl.PRESSED && e.Keysym.Sym == sdl.K_f {
				world.Logger().Info(errors.New("toggling screen size due to pressing 'F'"))
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
		world.NetSync().Start(),
		world.Smooth().Start(),
		// update {
		world.Connection(),

		// inputs
		world.Inputs(),
		world.Audio(),

		// update
		world.Camera(),
		world.Drag(),
		world.Transition(),
		temporaryInlineSystems,

		world.Tile(),
		world.Pathfind(),

		// ui update
		world.Ui(),
		world.Settings(),
		world.Loading(),
		world.Batcher(),
		// } (update)

		world.Smooth().Stop(),
		world.NetSync().Stop(),

		// render
		world.Render(),
		world.Tile().Renderer(),
		world.Render().Renderer(),
		world.Text().Renderer(),
		world.FpsLogger(),
	)
	for _, err := range errs {
		world.Logger().Log(err)
	}

	events.Emit(world.Events(), scene.NewChangeSceneEvent(definitions.MenuID))

	world.Logger().Info(errors.New("initialized engine"))
	runtime.LockOSThread()
	world.Loop().Run(loop.NewConfigureEvent(60, 1))
}
