package internal

import (
	"core/modules/definitions"
	"core/modules/generation"
	"core/modules/tile"
	"engine"
	"engine/modules/batcher"
	"engine/modules/collider"
	"engine/modules/grid"
	"engine/modules/inputs"
	"engine/modules/noise"
	"engine/services/datastructures"
	"engine/services/ecs"
	"fmt"
	"slices"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/ogiusek/ioc/v2"
)

type config struct {
	types       []tile.ID
	tilesPerJob int
}

type service struct {
	engine.World `inject:"1"`
	Definitions  definitions.Definitions `inject:"1"`
	Tile         tile.Service            `inject:"1"`

	config
}

func NewService(c ioc.Dic) generation.Service {
	s := ioc.GetServices[service](c)
	s.types = []tile.ID{}
	s.addChance(s.Definitions.Tiles.Water, 35)
	s.addChance(s.Definitions.Tiles.Sand, 15)
	s.addChance(s.Definitions.Tiles.Grass, 45)
	s.addChance(s.Definitions.Tiles.Mountain, 5)

	s.tilesPerJob = 100
	return &s
}

func (s *service) addChance(tileType ecs.EntityID, chance int) {
	tileComp, ok := s.Tile.Tile().Get(tileType)
	if !ok {
		s.Logger.Warn(fmt.Errorf("\"%v\" isn't a tile tile and therefor cannot be used in generation", tileType))
		return
	}
	s.types = append(s.types, slices.Repeat([]tile.ID{tileComp.ID}, chance)...)
}

func MapRange(val, min, max float64) float64 { return min + (val * (max - min)) }

func (s *service) Generate(c generation.Config) batcher.Task {
	gridStateComponent := tile.NewGrid(c.Size.Coords())
	gridModifiedComponent := tile.NewGrid(c.Size.Coords())

	jobs := int(gridStateComponent.GetLastIndex()) / s.tilesPerJob

	// apply batch
	applyBatch := batcher.NewBatch(jobs, func(i int) {
		for j := range s.tilesPerJob {
			gridI := grid.Index(i*s.tilesPerJob + j)
			gridValue := gridModifiedComponent.GetTile(gridI)
			gridStateComponent.SetTile(gridI, gridValue)
		}
	})

	// generate batch
	multiplier := 1. / 4

	noise := s.Noise.NewNoise(c.Seed).AddValue(
		noise.NewLayer(100*multiplier, .10),
		noise.NewLayer(100*multiplier, .10),
		noise.NewLayer(040*multiplier, .10),
		noise.NewLayer(040*multiplier, .05),
		noise.NewLayer(040*multiplier, .05),
	).AddPerlin(
		noise.NewLayer(500*multiplier, .50),
		noise.NewLayer(500*multiplier, .50),
		noise.NewLayer(500*multiplier, .50),
		noise.NewLayer(500*multiplier, .50),
		noise.NewLayer(500*multiplier, .50),
		noise.NewLayer(500*multiplier, .50),
		noise.NewLayer(500*multiplier, .50),
		noise.NewLayer(100*multiplier, .20),
		//
		noise.NewLayer(040*multiplier, .05),
		noise.NewLayer(020*multiplier, .05),
	).Build()

	generateBatch := batcher.NewBatch(jobs, func(i int) {
		for j := range s.tilesPerJob {
			gridI := grid.Index(i*s.tilesPerJob + j)
			coords := gridModifiedComponent.GetCoords(gridI)
			value := noise.Read(mgl64.Vec2{float64(coords.X), float64(coords.Y)})
			value *= float64(len(s.types))
			value = min(value, float64(len(s.types)-1))
			tileValue := s.types[int(value)]
			gridModifiedComponent.SetTile(gridI, tileValue)
		}
	})

	// smoothing batch
	neighbours := []grid.Coords{}
	neighbourDistance := grid.Coord(3)
	for x := -neighbourDistance; x <= neighbourDistance; x++ {
		for y := -neighbourDistance; y <= neighbourDistance; y++ {
			if x == 0 && y == 0 {
				continue
			}
			neighbours = append(neighbours, grid.NewCoords(x, y))
		}
	}

	sensitivity := 1.5

	smoothingBatch := batcher.NewBatch(jobs, func(i int) {
		for j := range s.tilesPerJob {
			gridI := grid.Index(i*s.tilesPerJob + j)
			coords := gridStateComponent.GetCoords(gridI)
			counts := datastructures.NewSparseArray[tile.ID, int]()
			for _, neighbour := range neighbours {
				coords := grid.NewCoords(coords.X+neighbour.X, coords.Y+neighbour.Y)
				index, ok := gridStateComponent.GetIndex(coords.Coords())
				if !ok {
					continue
				}
				value := gridStateComponent.GetTile(index)
				count, _ := counts.Get(value)
				counts.Set(value, count+1)
			}

			var dominantTile tile.ID
			maxCount := 0
			for _, tileType := range counts.GetIndices() {
				count, _ := counts.Get(tileType)
				if count > maxCount {
					maxCount = count
					dominantTile = tileType
				}
			}

			currentTile := gridStateComponent.GetTile(gridI)
			currentTypeCount, _ := counts.Get(currentTile)

			newTile := currentTile
			if maxCount > int(float64(currentTypeCount)*sensitivity) {
				newTile = dominantTile
			}

			gridModifiedComponent.SetTile(gridI, newTile)
		}
	})

	// flush batch
	flushBatch := batcher.NewBatch(1, func(i int) {
		size := s.Tile.GetTileSize()
		size.Size[0] *= float32(c.Size.X)
		size.Size[1] *= float32(c.Size.Y)

		s.Transform.Size().Set(c.Entity, size)

		s.Collider.Component().Set(c.Entity, collider.NewCollider(s.Definitions.SquareCollider))
		s.Inputs.Stack().Set(c.Entity, inputs.StackComponent{})
		s.Tile.Grid().Set(c.Entity, gridStateComponent)
	})

	// task

	task := s.Batcher.NewTask()
	task.AddConcurrentBatch(generateBatch)
	task.AddConcurrentBatch(applyBatch)
	for range 2 {
		task.AddConcurrentBatch(smoothingBatch)
		task.AddConcurrentBatch(applyBatch)
	}
	task.AddOrderedBatch(flushBatch)

	return task.Build()
}
