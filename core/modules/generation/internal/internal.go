package internal

import (
	"core/game"
	"core/modules/generation"
	"core/modules/tile"
	"engine/modules/batcher"
	"engine/modules/collider"
	"engine/modules/grid"
	"engine/modules/inputs"
	"engine/modules/metadata"
	"engine/modules/noise"
	"engine/modules/transform"
	"engine/services/datastructures"
	"engine/services/ecs"
	"fmt"
	"slices"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type chance struct {
	tileType ecs.EntityID
	chance   int
}

type Config struct {
	chances     []chance
	tilesPerJob int
}

func NewConfig(tilesPerJob int) *Config {
	return &Config{
		tilesPerJob: tilesPerJob,
	}
}

func (c *Config) AddChance(tileType ecs.EntityID, chanceInProcent int) {
	c.chances = append(c.chances, chance{tileType, chanceInProcent})
}

//

type service struct {
	game.GameWorld `inject:""`
	C              ioc.Dic
}

func NewService(c ioc.Dic) generation.Service {
	s := ioc.GetServices[service](c)
	s.C = c
	return &s
}

func MapRange(val, min, max float64) float64 { return min + (val * (max - min)) }

func (s *service) Register() error {
	events.Listen(s.EventsBuilder(), s.GenerateOn)
	return nil
}

func (s *service) Chances() (*Config, []tile.ID) {
	config := ioc.Get[*Config](s.C)
	types := []tile.ID{}

	for _, chance := range config.chances {
		tileComp, ok := s.Tile().Component().Get(chance.tileType)
		if !ok {
			s.Logger().Log(fmt.Errorf("\"%v\" isn't a tile tile and therefor cannot be used in generation", chance.tileType))
			continue
		}
		types = append(types, slices.Repeat([]tile.ID{tileComp.ID}, chance.chance)...)
	}
	return config, types
}

func (s *service) GenerateOn(event tile.MissingChunkEvent) {
	// this shouldn't override parrent component instead create additional child
	// this isn't a comment to purposfully throw error
	worldGenerationEntity, ok := s.Tile().GetConfig()
	if !ok {
		return
	}
	if _, ok := s.Tile().Grid().Chunk().Get(worldGenerationEntity); ok {
		return
	}
	c, ok := s.Tile().Config().Get(worldGenerationEntity)
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
			obstruction, _ := s.Obstruction().Component().Get(entity)
			obstructGridComponent.SetTile(gridI, obstruction.Obstruction)
		}
	})

	// flush batch
	flushBatch := batcher.NewBatch(1, func(i int) {
		chunkEntity := s.World().NewEntity()

		s.Hierarchy().SetParent(chunkEntity, worldGenerationEntity)
		s.Groups().InheritGroups(chunkEntity)
		size := s.Tile().GetTileSize()
		size.Size[0] *= float32(s.Grid().ChunkSize())
		size.Size[1] *= float32(s.Grid().ChunkSize())

		s.Transform().Pos().Set(chunkEntity, transform.NewPos(
			float32(event.Coords.X)*size.Size[0],
			float32(event.Coords.Y)*size.Size[1],
			0,
		))
		s.Transform().Size().Set(chunkEntity, size)
		s.Transform().PivotPoint().Set(chunkEntity, transform.NewPivotPoint(0, 0, .5))

		s.Collider().Component().Set(chunkEntity, collider.NewCollider(s.Definitions().Assets().SquareCollider))
		s.Inputs().Stack().Set(chunkEntity, inputs.StackComponent{})
		s.Grid().Coords().Set(chunkEntity, grid.NewChunkCoords(event.Coords.X, event.Coords.Y))
		s.Tile().Grid().Chunk().Set(chunkEntity, gridStateComponent)
		s.Obstruction().Grid().Chunk().Set(chunkEntity, obstructGridComponent)

		playerEntity := s.World().NewEntity()
		s.Hierarchy().SetParent(playerEntity, worldGenerationEntity)
		s.Metadata().Name().Set(playerEntity, metadata.NewName("john"))
		player2Entity := s.World().NewEntity()
		s.Hierarchy().SetParent(player2Entity, worldGenerationEntity)
		s.Metadata().Name().Set(player2Entity, metadata.NewName("anna"))

		// generates objects
		type Deployed struct {
			Blueprint,
			Player ecs.EntityID
		}
		toDeploy := []Deployed{
			{s.Definitions().Objects().Farm, playerEntity},
			{s.Definitions().Objects().Tank, player2Entity},
		}
	loop:
		for y := range s.Grid().ChunkSize() {
			for x := range s.Grid().ChunkSize() {
				coords := grid.NewCoords(x, y)
				coords = s.Grid().AbsoluteCoords(event.Coords, coords)
				deployed := toDeploy[0]
				if _, err := s.Deploy().Deploy(deployed.Blueprint, deployed.Player, coords); err == nil {
					toDeploy = toDeploy[1:]
				}
				if len(toDeploy) == 0 {
					break loop
				}
			}
		}
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
