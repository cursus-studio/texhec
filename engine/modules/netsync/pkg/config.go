package netsyncpkg

import (
	"engine/modules/ecs"
	"engine/modules/netsync/internal/config"
	"engine/modules/record"
)

type Config struct {
	config config.Config
}

func newConfig() Config {
	return Config{
		config: config.NewConfig(),
	}
}

func (c *Config) RecordConfig() record.Config {
	return c.config.RecordConfig
}
func (c *Config) SetMaxPredictions(maxPredictions int) {
	c.config.MaxPredictions = maxPredictions
}

func AddEvent[EventType any](c Config) {
	config.AddEvent[EventType](c.config)
}

func AddTickEvent[EventType any](c Config) {
	config.AddTickEvent[EventType](c.config)
}

func AddGenericEventValidation(c Config, handler func(client ecs.EntityID, event any) error) {
	config.AddGenericEventValidation(c.config, handler)
}

func AddEventValidation[EventType any](c Config, handler func(event EventType) error) {
	config.AddEventValidation(c.config, handler)
}
