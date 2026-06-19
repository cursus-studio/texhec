// generates map using noise functions
package generation

import (
	"engine/modules/ecs"
)

type Service interface {
	// generates all missing chunks
	ecs.SystemRegister
}
