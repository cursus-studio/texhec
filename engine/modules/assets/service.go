// We store all data in `ecs.World` and build on top of it.\
// We store path in `PathComponent` and upon retrival we store asset as interface in `CacheComponent`.\
// We convert `PathComponent` to `CacheComponent` using dispatchers where each file extension has dedicated dispatcher.\
// Using interfaces for `CacheComponent` doesn't affect performance heavily, its nothing in comparison to data stored and its processing.\
// To release assets we just remove `CacheComponent` (recommended) or entity with this component.\
// `CacheComponent` stores interface and we use\
// `func GetAsset[Asset any](assets Service, assetID ecs.EntityID) (Asset, error)`\
// to parse is to our asset type.
package assets

import (
	"engine/services/ecs"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Asset interface{ Release() }

// components

type PathComponent struct{ Path string }
type CacheComponent struct{ Cache Asset }

func NewPath(path string) PathComponent   { return PathComponent{path} }
func NewCache(cache Asset) CacheComponent { return CacheComponent{cache} }

func (s *PathComponent) Extension() string {
	parts := strings.Split(s.Path, ".")
	parts = strings.Split(parts[len(parts)-1], "/")
	return parts[len(parts)-1]
}

// add asset struct

type Service interface {
	Path() ecs.ComponentsArray[PathComponent]
	Cache() ecs.ComponentsArray[CacheComponent]

	Register(
		/* shouldn't have dots and be after dots in asset */ extension string,
		dispatcher func(path PathComponent) (Asset, error),
	)

	// get also caches asset
	Get(ecs.EntityID) (Asset, error)
}

var (
	ErrAssetHasDifferentType error = errors.New("asset is not of requested type")
	ErrAssetNotFound         error = errors.New("asset not found")
)

func GetAsset[Asset any](assets Service, assetID ecs.EntityID) (Asset, error) {
	rawAsset, err := assets.Get(assetID)
	if err != nil {
		var a Asset
		return a, err
	}
	asset, ok := rawAsset.(Asset)
	if !ok {
		var a Asset
		err := errors.Join(
			ErrAssetHasDifferentType,
			fmt.Errorf(
				"asset is of type \"%s\" and expected to be \"%s\"",
				reflect.TypeOf(rawAsset).String(),
				reflect.TypeFor[Asset]().String(),
			),
		)
		return a, err
	}
	return asset, nil
}
