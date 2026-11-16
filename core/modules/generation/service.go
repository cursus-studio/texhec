// generates map using noise functions
package generation

import (
	"engine/modules/ecs"
	"engine/modules/grid"
)

type ChunkGenerator uint8

// service used identifier used as a singleton
type GeneratorComponent struct {
	Generator ChunkGenerator
}

func NewGenerator(generator ChunkGenerator) GeneratorComponent { return GeneratorComponent{generator} }

type ChunkGenerateService interface {
	IsSpecial(grid.ChunkCoordsComponent) bool
}

//

type Service interface {
	// generates all missing chunks
	ecs.SystemRegister
	Generator() ecs.ComponentArray[GeneratorComponent]

	IsSpecial(grid.ChunkCoordsComponent) bool
}
