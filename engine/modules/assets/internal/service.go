package internal

import (
	"engine/modules/assets"
	"engine/services/datastructures"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

//

type assetsService struct {
	*extensions
	path  ecs.ComponentsArray[assets.PathComponent]
	cache ecs.ComponentsArray[assets.CacheComponent]

	cached datastructures.SparseArray[ecs.EntityID, assets.Asset]
}

func NewService(c ioc.Dic) assets.Service {
	s := &assetsService{}
	s.extensions = NewExtensions(c)

	s.extensions = NewExtensions(c)
	s.path = ecs.GetComponentsArray[assets.PathComponent](s.World())
	s.cache = ecs.GetComponentsArray[assets.CacheComponent](s.World())

	s.cached = datastructures.NewSparseArray[ecs.EntityID, assets.Asset]()

	s.cache.OnUpsert(s.OnUpsert)
	s.cache.OnRemove(s.OnRemove)

	return s
}

func (s *assetsService) OnUpsert(e ecs.EntityID) {
	if asset, ok := s.cached.Get(e); ok {
		asset.Release()
	}
	if cached, ok := s.cache.Get(e); ok {
		s.cached.Set(e, cached.Cache)
	}
}
func (s *assetsService) OnRemove(e ecs.EntityID) {
	if asset, ok := s.cached.Get(e); ok {
		asset.Release()
		s.cached.Remove(e)
	}
}

func (s *assetsService) Path() ecs.ComponentsArray[assets.PathComponent]   { return s.path }
func (s *assetsService) Cache() ecs.ComponentsArray[assets.CacheComponent] { return s.cache }

func (s *assetsService) Get(entity ecs.EntityID) (assets.Asset, error) {
	if cache, ok := s.cache.Get(entity); ok {
		return cache.Cache, nil
	}

	path, ok := s.path.Get(entity)
	if !ok {
		return nil, assets.ErrAssetNotFound
	}

	extension := path.Extension()
	dispatcher, ok := s.ExtensionDispatcher(extension)
	if !ok {
		return nil, assets.ErrAssetNotFound
	}
	asset, err := dispatcher(path)
	if err != nil {
		return nil, err
	}
	s.cache.Set(entity, assets.NewCache(asset))
	return asset, nil
}
