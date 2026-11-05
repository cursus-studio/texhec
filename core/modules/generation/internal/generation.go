package internal

import (
	"core/modules/definitions"
	"core/modules/tile"
	"engine/modules/batcher"
	"engine/modules/collider"
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/groups"
	"engine/modules/noise"
	"fmt"
	"slices"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/ogiusek/events"
)

func MapRange(val, min, max float64) float64 { return min + (val * (max - min)) }

func (s *service) Register() error {
	events.Listen(s.EventsBuilder(), s.GenerateOn)
	s.Grid().Coords().OnUpsert(func(chunkEntity ecs.EntityID) {
		worldGenerationEntity, _, _ := s.Seed().Seed().GetSingleton()
		s.Hierarchy().SetParent(chunkEntity, worldGenerationEntity)
		s.Groups().InheritGroups(chunkEntity)
		s.Collider().Collider().Set(chunkEntity, collider.NewCollider(s.Definitions().Assets().SquareCollider))
	})
	s.Seed().Seed().OnUpsert(func(entity ecs.EntityID) {
		s.Hierarchy().SetParent(entity, s.Scene().Scene())
		s.Groups().Groups().Set(entity, groups.EmptyGroups().Enable(definitions.GameGroup))
	})
	return nil
}

func (s *service) Chances() (*Config, []tile.ID) {
	types := []tile.ID{}

	for _, chance := range s.Config.chances {
		tileComp, ok := s.Tile().Tile().Get(chance.tileType)
		if !ok {
			s.Logger().Log(fmt.Errorf("\"%v\" isn't a tile tile and therefor cannot be used in generation", chance.tileType))
			continue
		}
		types = append(types, slices.Repeat([]tile.ID{tileComp.ID}, chance.chance)...)
	}
	return s.Config, types
}

func (s *service) GenerateOn(event tile.MissingChunkEvent) {
	worldSeedEntity, _, ok := s.Seed().Seed().GetSingleton()
	if !ok {
		return
	}
	if _, ok := s.Tile().Grid().Chunk().Get(worldSeedEntity); ok {
		return
	}
	c, ok := s.Seed().Seed().Get(worldSeedEntity)
	if !ok {
		return
	}
	config, tileTypes := s.Chances()
	gridStateComponent := s.Tile().Grid().NewChunk()
	gridModifiedComponent := s.Tile().Grid().NewChunk()

	obstructGridComponent := s.Obstruction().Grid().NewChunk()

	jobs := int(s.Grid().GetLastIndex()) / config.tilesPerJob

	// apply batch
	applyBatch := batcher.NewBatch(jobs, func(i int) {
		for j := range config.tilesPerJob {
			gridI := grid.Index(i*config.tilesPerJob + j)
			gridValue := gridModifiedComponent.GetTile(gridI)
			gridStateComponent.SetTile(gridI, gridValue)
		}
	})

	// generate batch
	multiplier := 1. / 4

	noise := s.Noise().NewNoise(c.Seed).AddValue(
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
		for j := range config.tilesPerJob {
			gridI := grid.Index(i*config.tilesPerJob + j)
			coords := s.Grid().AbsoluteCoords(event.Coords, s.Grid().IndexCoords(gridI))
			value := noise.Read(mgl64.Vec2{float64(coords.X), float64(coords.Y)})
			value *= float64(len(tileTypes))
			value = min(value, float64(len(tileTypes)-1))
			tileValue := tileTypes[int(value)]
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
		for j := range config.tilesPerJob {
			gridI := grid.Index(i*config.tilesPerJob + j)
			coords := s.Grid().IndexCoords(gridI)
			counts := datastructures.NewSparseArray[tile.ID, int]()
			for _, neighbour := range neighbours {
				coords := grid.NewCoords(coords.X+neighbour.X, coords.Y+neighbour.Y)
				index, ok := s.Grid().CoordsIndex(coords)
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

	obstructBatch := batcher.NewBatch(jobs, func(i int) {
		for j := range config.tilesPerJob {
			gridI := grid.Index(i*config.tilesPerJob + j)
			tileType := gridStateComponent.GetTile(gridI)
			entity, ok := s.Tile().GetTile(tileType)
			if !ok {
				continue
			}
			obstruction, _ := s.Obstruction().Obstrucion().Get(entity)
			obstructGridComponent.SetTile(gridI, obstruction.Obstruction)
		}
	})

	// flush batch
	flushBatch := batcher.NewBatch(1, func(i int) {
		chunkEntity := s.World().NewEntity()

		s.Grid().Coords().Set(chunkEntity, grid.NewChunkCoords(event.Coords.X, event.Coords.Y))
		s.Tile().Grid().Chunk().Set(chunkEntity, gridStateComponent)
		s.Obstruction().Grid().Chunk().Set(chunkEntity, obstructGridComponent)
	})

	// task
	task := s.Batcher().NewTask()
	task.AddConcurrentBatch(generateBatch)
	task.AddConcurrentBatch(applyBatch)
	for range 2 {
		task.AddConcurrentBatch(smoothingBatch)
		task.AddConcurrentBatch(applyBatch)
	}
	task.AddConcurrentBatch(obstructBatch)
	task.AddOrderedBatch(flushBatch)

	s.Batcher().Queue(task.Build())
}
