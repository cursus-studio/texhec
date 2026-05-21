package test

import (
	"core/game"
	"core/modules/generation"
	corepkg "core/pkg"
	assetspkg "engine/modules/assets/pkg"
	"engine/modules/camera"
	"engine/modules/entityregistry"
	"engine/modules/grid"
	"engine/modules/loop"
	"engine/modules/seed"
	"engine/modules/transform"
	"engine/modules/window"
	"engine/services/ecs"
	"fmt"
	"runtime"
	"testing"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/jaypipes/ghw"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type Tiles struct {
	Water    ecs.EntityID `path:"tiles/water.biome" tile:"" generate:"25"`
	Sand     ecs.EntityID `path:"tiles/sand.biome" tile:"" generate:"25"`
	Grass    ecs.EntityID `path:"tiles/grass.biome" tile:"" generate:"25"`
	Mountain ecs.EntityID `path:"tiles/mountain.biome" tile:"" generate:"25"`
}

func benchmarkRenderingXTilesMap(b *testing.B, n int) {
	if gpu, err := ghw.GPU(); err != nil {
		b.Error(err)
		return
	} else {
		for i, gpu := range gpu.GraphicsCards {
			name := gpu.DeviceInfo.Product.Name
			if i == 0 {
				fmt.Printf("gpu: %v\n", name)
			} else {
				fmt.Printf("gpu %v: %v\n", i+1, name)
			}
		}
	}
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	world := ioc.Get[game.GameWorld](ioc.NewContainer(
		corepkg.Pkg,
		func(b ioc.Builder) {
			ioc.Wrap(b, func(c ioc.Dic, w window.Service) {
				w.Window().SetTitle("tile module benchmark")
				w.Window().SetSize(1000, 1000)
				gl.ClearColor(0.2, 0.3, 0.3, 1.0)
			})
			ioc.Wrap(b, func(c ioc.Dic, b assetspkg.Config) { b.SetPath("../../../assets/") })
		},
	))

	events.GlobalErrHandler(world.EventsBuilder(), world.Logger().Log)

	errs := ecs.RegisterSystems(
		world.Render(),
		world.Tile().Renderer(),
	)
	for _, err := range errs {
		world.Logger().Log(err)
	}

	// before generating load elements
	_, err := entityregistry.GetRegistry[Tiles](world.EntityRegistry())
	world.Logger().Log(err)

	gridEntity := world.World().NewEntity()

	config := generation.NewConfig(gridEntity, seed.New(21377137), grid.NewCoords(n, n))
	world.Generation().Generate(config).Perform()

	gameCamera := world.World().NewEntity()
	ortho := camera.NewOrtho(-1000, +1000)
	ortho.Zoom = 10. / float32(n)
	world.Camera().Ortho().Set(gameCamera, ortho)
	tileSize := world.Tile().GetTileSize().Size
	cameraPos := transform.NewPos(tileSize.X()*float32(n)/2, tileSize.Y()*float32(n)/2, 0)
	world.Transform().Pos().Set(gameCamera, cameraPos)

	events.Emit(world.Events(), loop.FrameEvent{})
	b.ResetTimer()

	for b.Loop() {
		events.Emit(world.Events(), loop.FrameEvent{})
	}
	gl.Finish()
}

// const MAP_SIZE
func BenchmarkRendering1MTilesMap(b *testing.B) { benchmarkRenderingXTilesMap(b, 1000) }
func BenchmarkRendering4MTilesMap(b *testing.B) { benchmarkRenderingXTilesMap(b, 2000) }
