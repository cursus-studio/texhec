package internal

import (
	"core/game"
	"core/modules/generation"
	"core/modules/tile"
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"engine/modules/grid"
	"fmt"
	"slices"

	"github.com/ogiusek/ioc/v2"
)

type chance struct {
	tileType ecs.EntityID
	chance   int
}
type chunkGenerate struct {
	tileType ecs.EntityID
}

type Config struct {
	chances       []chance
	chunkGenerate chunkGenerate
	tilesPerJob   int
}

func NewConfig(tilesPerJob int) *Config {
	return &Config{
		tilesPerJob: tilesPerJob,
	}
}

func (c *Config) AddChance(tileType ecs.EntityID, chanceInProcent int) {
	c.chances = append(c.chances, chance{tileType, chanceInProcent})
}
func (c *Config) AddChunkGenerate(tileType ecs.EntityID) {
	c.chunkGenerate = chunkGenerate{tileType}
}

//

type service struct {
	game.GameWorld `inject:""`
	Config         *Config `inject:""`
	generator      ecs.ComponentArray[generation.GeneratorComponent]
	generators     datastructures.SparseArray[generation.ChunkGenerator, generation.ChunkGenerateService]

	types   []tile.ID
	special tile.ID
}

type Service interface {
	RegisterGenerator(generation.ChunkGenerator, generation.ChunkGenerateService)
	generation.Service
}

func NewService(c ioc.Dic) Service {
	s := ioc.GetServices[*service](c)
	s.generator = ecs.GetComponentArray[generation.GeneratorComponent](s.World())
	s.generators = datastructures.NewSparseArray[generation.ChunkGenerator, generation.ChunkGenerateService]()

	for _, chance := range s.Config.chances {
		tileComp, ok := s.Tile().Tile().Get(chance.tileType)
		if !ok {
			s.Logger().Fatal(fmt.Errorf("\"%v\" isn't a tile tile and therefor cannot be used in generation", chance.tileType))
			continue
		}
		s.types = append(s.types, slices.Repeat([]tile.ID{tileComp.ID}, chance.chance)...)
	}

	if tileComp, ok := s.Tile().Tile().Get(s.Config.chunkGenerate.tileType); ok {
		s.special = tileComp.ID
	}

	return s
}

func (s *service) Generator() ecs.ComponentArray[generation.GeneratorComponent] {
	return s.generator
}

func (s *service) IsSpecial(coords grid.ChunkCoordsComponent) bool {
	_, generator, ok := s.generator.GetSingleton()
	if !ok {
		return false
	}
	service, ok := s.generators.Get(generator.Generator)
	if !ok {
		return false
	}
	return service.IsSpecial(coords)
}

func (s *service) RegisterGenerator(kind generation.ChunkGenerator, service generation.ChunkGenerateService) {
	if added := s.generators.Set(kind, service); !added {
		panic("cannot register generator twice")
	}
}
