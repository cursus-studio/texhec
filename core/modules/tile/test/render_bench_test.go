package test

import (
	"core/game"
	"core/modules/tile"
	corepkg "core/pkg"
	assetspkg "engine/modules/assets/pkg"
	"engine/modules/camera"
	"engine/modules/ecs"
	"engine/modules/entityregistry"
	"engine/modules/grid"
	"engine/modules/loop"
	"engine/modules/seed"
	"engine/modules/transform"
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

func BenchmarkRenderingChunk(b *testing.B) {
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
	tilesInChunk := world.Grid().ChunkSize()
	tilesInChunk *= tilesInChunk
	fmt.Printf("tiles in chunk: %v\n", tilesInChunk)

	events.GlobalErrHandler(world.EventsBuilder(), world.Logger().Log)

	errs := ecs.RegisterSystems(
		world.Render(),
		world.Tile(),
		world.Batcher(),
		world.Tile().Renderer(),
		world.Generation(),
	)
	for _, err := range errs {
		world.Logger().Log(err)
	}

	// before generating load elements
	_, err := entityregistry.GetRegistry[Tiles](world.EntityRegistry())
	world.Logger().Log(err)

	n := float32(world.Grid().ChunkSize())
	gameCamera := world.World().NewEntity()
	ortho := camera.NewOrtho(-1000, +1000)
	ortho.Zoom = 10. / n
	world.Camera().Ortho().Set(gameCamera, ortho)
	tileSize := world.Tile().GetTileSize().Size
	cameraPos := transform.NewPos(tileSize.X()*n, tileSize.Y()*n, 0)
	world.Transform().Pos().Set(gameCamera, cameraPos)

	generationEntity := world.World().NewEntity()
	world.Hierarchy().SetParent(generationEntity, gameCamera)
	world.Seed().Seed().Set(generationEntity, seed.NewSeed(21377137))

	events.Emit(world.Events(), tile.NewMissingChunkEvent(grid.NewChunkCoords(0, 0)))
	for _, task := range world.Batcher().Tasks() {
		task.Perform()
	}
	events.Emit(world.Events(), loop.FrameEvent{})
	b.ResetTimer()
	for b.Loop() {
		events.Emit(world.Events(), loop.FrameEvent{})
	}
	gl.Finish()
}
