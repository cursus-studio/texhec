package internal

import (
	"core/modules/definitions"
	"core/modules/deploy"
	"core/modules/generation"
	"core/modules/player"
	"core/modules/tile"
	"engine"
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

func NewConfig() *Config {
	return &Config{
		tilesPerJob: 100,
	}
}

func (c *Config) AddChance(tileType ecs.EntityID, chanceInProcent int) {
	c.chances = append(c.chances, chance{tileType, chanceInProcent})
}

//

type service struct {
	engine.World `inject:"1"`
	Definitions  definitions.Definitions `inject:"1"`
	Tile         tile.Service            `inject:"1"`
	Deploy       deploy.Service          `inject:"1"`
	Player       player.Service          `inject:"1"`
	C            ioc.Dic
}

func NewService(c ioc.Dic) generation.Service {
	s := ioc.GetServices[service](c)
	s.C = c
	return &s
}

func MapRange(val, min, max float64) float64 { return min + (val * (max - min)) }

func (s *service) Chances() (*Config, []tile.ID) {
	config := ioc.Get[*Config](s.C)
	types := []tile.ID{}

	for _, chance := range config.chances {
		tileComp, ok := s.Tile.TileType().Get(chance.tileType)
		if !ok {
			s.Logger.Warn(fmt.Errorf("\"%v\" isn't a tile tile and therefor cannot be used in generation", chance.tileType))
			continue
		}
		types = append(types, slices.Repeat([]tile.ID{tileComp.ID}, chance.chance)...)
	}
	return config, types
}

func (s *service) Generate(c generation.Config) batcher.Task {
	config, tileTypes := s.Chances()
	gridStateComponent := tile.NewTileGrid(c.Size.Coords())
	gridModifiedComponent := tile.NewTileGrid(c.Size.Coords())

	obstructGridComponent := tile.NewObstructGrid(c.Size.Coords())

	jobs := int(gridStateComponent.GetLastIndex()) / config.tilesPerJob

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
		for j := range config.tilesPerJob {
			gridI := grid.Index(i*config.tilesPerJob + j)
			coords := gridModifiedComponent.GetCoords(gridI)
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

	obstructBatch := batcher.NewBatch(jobs, func(i int) {
		for j := range config.tilesPerJob {
			gridI := grid.Index(i*config.tilesPerJob + j)
			tileType := gridStateComponent.GetTile(gridI)
			entity, ok := s.Tile.GetTileType(tileType)
			if !ok {
				continue
			}
			obstruction, _ := s.Tile.Obstruction().Get(entity)
			obstructGridComponent.SetTile(gridI, obstruction.Obstruction)
		}
	})

	// flush batch
	flushBatch := batcher.NewBatch(1, func(i int) {
		size := s.Tile.GetTileSize()
		size.Size[0] *= float32(c.Size.X)
		size.Size[1] *= float32(c.Size.Y)

		s.Transform.Size().Set(c.Entity, size)
		s.Transform.PivotPoint().Set(c.Entity, transform.NewPivotPoint(0, 0, .5))

		s.Collider.Component().Set(c.Entity, collider.NewCollider(s.Definitions.SquareCollider))
		s.Inputs.Stack().Set(c.Entity, inputs.StackComponent{})
		s.Tile.TileGrid().Set(c.Entity, gridStateComponent)
		s.Tile.ObstructionGrid().Set(c.Entity, obstructGridComponent)

		playerEntity := s.NewEntity()
		s.Metadata.Name().Set(playerEntity, metadata.NewName("john"))
		player2Entity := s.NewEntity()
		s.Metadata.Name().Set(player2Entity, metadata.NewName("anna"))

		// generates objects
		type Deployed struct {
			Blueprint,
			Player ecs.EntityID
		}
		toDeploy := []Deployed{
			{s.Definitions.Constructs.Farm, playerEntity},
			{s.Definitions.Units.Tank, player2Entity},
		}
		for index := grid.Index(1); index < gridStateComponent.GetLastIndex(); index++ {
			if len(toDeploy) == 0 {
				break
			}
			coords := gridStateComponent.GetCoords(index)
			deployed := toDeploy[0]
			if entity, err := s.Deploy.Deploy(deployed.Blueprint, deployed.Player, coords); err == nil {
				index += 2
				toDeploy = toDeploy[1:]
				if _, ok := s.Tile.Speed().Get(entity); !ok {
					continue
				}
				step := tile.NewStep(coords.Coords())
				step.X--
				s.Tile.Step().Set(entity, step)
			}
		}
	})

	// task

	task := s.Batcher.NewTask()
	task.AddConcurrentBatch(generateBatch)
	task.AddConcurrentBatch(applyBatch)
	for range 2 {
		task.AddConcurrentBatch(smoothingBatch)
		task.AddConcurrentBatch(applyBatch)
	}
	task.AddConcurrentBatch(obstructBatch)
	task.AddOrderedBatch(flushBatch)

	return task.Build()
}
