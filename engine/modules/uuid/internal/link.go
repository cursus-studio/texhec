package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/uuid"

	"github.com/ogiusek/ioc/v2"
)

type linkService[Wrapped any] struct {
	engine.EngineWorld `inject:""`
	linkUUID           ecs.ComponentArray[uuid.LinkUUIDComponent[Wrapped]]
	cache              ecs.ComponentArray[uuid.LinkCacheComponent[Wrapped]]
}

func NewLinkService[Wrapped any](c ioc.Dic) uuid.LinkService[Wrapped] {
	s := ioc.GetServices[*linkService[Wrapped]](c)
	s.linkUUID = ecs.GetComponentArray[uuid.LinkUUIDComponent[Wrapped]](s.World())
	s.cache = ecs.GetComponentArray[uuid.LinkCacheComponent[Wrapped]](s.World())
	return s
}

func (s *linkService[Wrapped]) UUID() ecs.ComponentArray[uuid.LinkUUIDComponent[Wrapped]] {
	return s.linkUUID
}
func (s *linkService[Wrapped]) Cache() ecs.ComponentArray[uuid.LinkCacheComponent[Wrapped]] {
	return s.cache
}
func (s *linkService[Wrapped]) Get(linkSrc ecs.EntityID) (linkDst ecs.EntityID, ok bool) {
	if cache, ok := s.cache.Get(linkSrc); ok {
		return cache.Entity, true
	}
	link, ok := s.linkUUID.Get(linkSrc)
	if !ok {
		return 0, false
	}
	entity, ok := s.EngineWorld.UUID().Entity(link.UUID)
	if !ok {
		return 0, false
	}
	s.cache.Set(linkSrc, uuid.NewLinkCache[Wrapped](entity))
	return entity, true
}

func (s *linkService[Wrapped]) SetUUID(entity ecs.EntityID, id uuid.UUID) {
	s.UUID().Set(entity, uuid.NewLinkUUID[Wrapped](id))
}
