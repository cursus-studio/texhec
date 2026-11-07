package test

import (
	"core/game"
	"core/modules/tile"
	"core/modules/ui"
	corepkg "core/pkg"
	assetspkg "engine/modules/assets/pkg"
	"engine/modules/camera"
	"engine/modules/ecs"
	"engine/modules/entityregistry"
	"engine/modules/graphics"
	"engine/modules/grid"
	"engine/modules/loop"
	"engine/modules/seed"
	"engine/modules/window"
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

// modify this to render all these damn chunks
func benchmarkXbyX(b *testing.B, side grid.Coord) {
	b.Helper()

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

			ioc.Wrap(b, func(c ioc.Dic, f graphics.Service) {
				f.TextureArray().Wrap(func(ta graphics.TextureArray) {
					ta.Bind()
					defer gl.BindTexture(gl.TEXTURE_2D_ARRAY, 0)

					gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
					gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
					gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
					gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
				})
			})
			ioc.Wrap(b, func(c ioc.Dic, b assetspkg.Config) { b.SetPath("../../../assets/") })
		},
	))
	tilesInChunk := world.Grid().ChunkSize()
	tilesInChunk *= tilesInChunk
	fmt.Printf("tiles rendered: %v\n", tilesInChunk*side*side)

	events.GlobalErrHandler(world.EventsBuilder(), world.Logger().Log)

	errs := ecs.RegisterSystems(
		// update
		world.Camera(),
		world.Transition(),

		world.Generation(),
		world.Tile(),

		// ui update
		world.Ui(),
		world.Settings(),
		world.Batcher(),
		// } (update)

		// render
		world.Render(),
		world.Tile().Renderer(),
		world.Render().Renderer(),
		world.Text().Renderer(),
	)
	for _, err := range errs {
		world.Logger().Log(err)
	}

	// before generating load elements
	_, err := entityregistry.GetRegistry[Tiles](world.EntityRegistry())
	world.Logger().Log(err)

	//

	gameCameraEntity := world.World().NewEntity()
	gameCamera := ui.NewGameCamera(ui.MaxChunksZoomOut(side))
	world.Ui().GameCamera().Set(gameCameraEntity, gameCamera)

	generationEntity := world.World().NewEntity()
	world.Hierarchy().SetParent(generationEntity, gameCameraEntity)
	world.Seed().Seed().Set(generationEntity, seed.NewSeed(21377137))

	for x := range gameCamera.MaxChunksZoomOut {
		for y := range gameCamera.MaxChunksZoomOut {
			events.Emit(world.Events(), tile.NewMissingChunkEvent(grid.NewChunkCoords(grid.Coord(x), grid.Coord(y))))
		}
	}
	for _, task := range world.Batcher().Tasks() {
		task.Perform()
	}

	world.Ui().GameCamera().Remove(gameCameraEntity)

	{ // set camera to max zoom out for preview
		cameraLimitsEntity := world.Camera().Limits().GetEntities()[0]
		limitsComp, _ := world.Camera().Limits().Get(cameraLimitsEntity)
		world.Camera().Limits().Set(cameraLimitsEntity, camera.NewCameraLimits(
			limitsComp.MinZoom, limitsComp.MinZoom, limitsComp.Min, limitsComp.Max))
	}

	events.Emit(world.Events(), loop.FrameEvent{})
	b.ResetTimer()
	for b.Loop() {
		events.Emit(world.Events(), loop.FrameEvent{})
	}
	gl.Finish()
}

// func BenchmarkRendering6144x6144(b *testing.B) { benchmarkXbyX(b, 6144/32) }
func BenchmarkRendering1024x1024(b *testing.B) { benchmarkXbyX(b, 1024/32) }
func BenchmarkRendering32x32(b *testing.B)     { benchmarkXbyX(b, 1) }
