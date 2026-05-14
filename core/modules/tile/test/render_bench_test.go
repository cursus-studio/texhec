package test

import (
	"core/modules/generation"
	"engine/modules/camera"
	"engine/modules/entityregistry"
	"engine/modules/grid"
	"engine/modules/loop"
	"engine/modules/seed"
	"engine/services/ecs"
	"fmt"
	"runtime"
	"testing"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/jaypipes/ghw"
	"github.com/ogiusek/events"
)

type Tiles struct {
	Water    ecs.EntityID `path:"tiles/water.biome" tile:"" generate:"25"`
	Sand     ecs.EntityID `path:"tiles/sand.biome" tile:"" generate:"25"`
	Grass    ecs.EntityID `path:"tiles/grass.biome" tile:"" generate:"25"`
	Mountain ecs.EntityID `path:"tiles/mountain.biome" tile:"" generate:"25"`
}

func BenchmarkRendering1MTilesMap(b *testing.B) {
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

	s := NewSetup()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	events.GlobalErrHandler(s.EventsBuilder(), s.Logger().Log)

	errs := ecs.RegisterSystems(
		s.Render(),
		s.Tile().Renderer(),
	)
	for _, err := range errs {
		s.Logger().Log(err)
	}

	// before generating load elements
	_, err := entityregistry.GetRegistry[Tiles](s.EntityRegistry())
	s.Logger().Log(err)

	gridEntity := s.World().NewEntity()

	config := generation.NewConfig(gridEntity, seed.New(21377137), grid.NewCoords(1000, 1000))
	s.Generation().Generate(config).Perform()

	gameCamera := s.World().NewEntity()
	ortho := camera.NewOrtho(-1000, +1000)
	ortho.Zoom = 0.01
	s.Camera().Ortho().Set(gameCamera, ortho)

	events.Emit(s.Events(), loop.FrameEvent{})
	b.ResetTimer()

	for b.Loop() {
		events.Emit(s.Events(), loop.FrameEvent{})
	}
	gl.Finish()
}
