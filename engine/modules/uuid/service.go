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

//

type LinkUUIDComponent[Wrappd any] struct{ UUID UUID }
type LinkCacheComponent[Wrappd any] struct{ Entity ecs.EntityID }

func NewLinkUUID[Wrapped any](uuid UUID) LinkUUIDComponent[Wrapped] {
	return LinkUUIDComponent[Wrapped]{uuid}
}
func NewLinkCache[Wrapped any](entity ecs.EntityID) LinkCacheComponent[Wrapped] {
	return LinkCacheComponent[Wrapped]{entity}
}

type LinkService[Wrapped any] interface {
	UUID() ecs.ComponentArray[LinkUUIDComponent[Wrapped]]
	Cache() ecs.ComponentArray[LinkCacheComponent[Wrapped]]
	Get(linkSrc ecs.EntityID) (linkDst ecs.EntityID, ok bool)

	SetUUID(ecs.EntityID, UUID)
}

// raw interface

type UUID uuid.UUID

func (id *UUID) String() string  { return uuid.UUID(*id).String() }
func (uuid *UUID) Bytes() []byte { return uuid[:] }

type Factory interface {
	NewUUID() UUID
	NewUUIDFromString(string) UUID
}
