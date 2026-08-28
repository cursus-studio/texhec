// integrates uuid
//
// uuid allows us to create universaly unique entity
package uuid

import (
	"engine/modules/ecs"

	"github.com/google/uuid"
)

// engine interface

type Component struct {
	ID UUID
}

func New(id UUID) Component {
	return Component{id}
}

type Service interface {
	Factory
	Component() ecs.ComponentArray[Component]
	Entity(UUID) (ecs.EntityID, bool)
}

// raw interface

type UUID uuid.UUID

func (id *UUID) String() string  { return uuid.UUID(*id).String() }
func (uuid *UUID) Bytes() []byte { return uuid[:] }

type Factory interface {
	NewUUID() UUID
	NewUUIDFromString(string) UUID
}
