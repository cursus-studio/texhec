// integrates uuid
//
// uuid allows us to create universaly unique entity
package uuid

import (
	"engine/modules/ecs"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrMissingUUID error = errors.New("uuid:missing uuid")
)

// engine interface

type UUIDComponent struct {
	ID UUID
}

func NewUUID(id UUID) UUIDComponent {
	return UUIDComponent{id}
}

type Service interface {
	Factory
	UUID() ecs.ComponentArray[UUIDComponent]
	Entity(UUID) (ecs.EntityID, bool)

	// creates if doesn't have assigned uuid
	CreateGetUUID(ecs.EntityID) UUID
	CreateGetEntity(UUID) ecs.EntityID
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
	NewUUID() UUID // isn't deterministic
	NewDeterministicUUID([]byte) UUID
	NewUUIDFromAny(any) UUID
}
