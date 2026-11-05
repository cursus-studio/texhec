package internal

import (
	"core/game"
	"core/modules/generation"
	"engine/modules/ecs"

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
	Config         *Config `inject:""`
}

func NewService(c ioc.Dic) generation.Service {
	s := ioc.GetServices[service](c)
	return &s
}
