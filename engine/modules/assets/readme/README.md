# Assets
## Architecture
We store all data in `ecs.World` and build on top of it.\
We store path in `PathComponent` and upon retrival we store asset as interface in `CacheComponent`.\
We convert `PathComponent` to `CacheComponent` using dispatchers where each file extension has dedicated dispatcher.\
Using interfaces for `CacheComponent` doesn't affect performance heavily, its nothing in comparison to data stored and its processing.\
To release assets we just remove `CacheComponent` (recommended) or entity with this component.\
`CacheComponent` stores interface and we use\
`func GetAsset[Asset any](assets Service, assetID ecs.EntityID) (Asset, error)`\
to parse is to our asset type.

```go
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

// Can return ErrAssetHasDifferentType, ErrAssetNotFound or error returned by dispatcher
func GetAsset[Asset any](assets Service, assetID ecs.EntityID) (Asset, error)
```

This is main service interface.\
On `Get` we either return content of `CacheComponent` if exists or\
use dispatchers for `PathComponent` extension to read file contents.

### Life cycle
Currently assets are stale and hot reloading isn't supported.\
To hot reload you'll need to remove asset cache and replace path component and you'll need to support in in consumers.

## Usage examples
### Asset
```go
type ExampleAsset interface {
	Data() string
	Release()
}

type exampleAsset struct {
    data string
}

func NewExampleAsset(data string) ExampleAsset {
	return &exampleAsset{data: data}
}

func (a *exampleAsset) Data() string { return a.data }
func (a *exampleAsset) Release()     {}
```

### Extension registration
```go
func (pkg) Register(b ioc.Builder) {
	ioc.Wrap(b, func(c ioc.Dic, b assets.Service) {
		b.Register("extension", func(path assets.PathComponent) (assets.Asset, error) {
			return NewExampleAsset(path.Path), nil
		})
	})
}
```

### Defining assets manually
This is only shows how everything works under the hood.\
Doing so to define assets **isn't recommended**.

```go
type Retrieved struct {
	Btn      ecs.EntityID
	Square   ecs.EntityID
	Settings ecs.EntityID
	Bg       ecs.EntityID
}

func (pkg) Register(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) Retrieved {
		world := ioc.Get[ecs.World](c)
		assetsService := ioc.Get[assets.Service](c)
		retrieved := Retrieved{
			Btn:      world.NewEntity(),
			Square:   world.NewEntity(),
			Settings: world.NewEntity(),
			Bg:       world.NewEntity(),
		}
		assetsService.Path().Set(retrieved.Btn, assets.NewPath("hud/btn.extension"))
		assetsService.Path().Set(retrieved.Square, assets.NewPath("hud/square.extension"))
		assetsService.Path().Set(retrieved.Settings, assets.NewPath("hud/settings.extension"))
		assetsService.Path().Set(retrieved.Bg, assets.NewPath("hud/bg.extension"))
		return retrieved
	})
}
```

### Initializing assets entities using registry (recommended)
We use `registry` and defined `path` struct tag to define asset path.

```go
type Retrieved struct {
	Btn      ecs.EntityID `path:"hud/btn.extension"`
	Square   ecs.EntityID `path:"hud/square.extension"`
	Settings ecs.EntityID `path:"hud/settings.extension"`
	Bg       ecs.EntityID `path:"hud/bg.extension"`
}

func (pkg) Register(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) Retrieved {
		logger := ioc.Get[logger.Logger](c)
		registryService := ioc.Get[registry.Service](c)

        // We define them using
		retrieved, err := registry.GetRegistry[Retrieved](registryService)
		logger.Warn(err)
		return retrieved
	})
}
```

### Retrieving assets
There is dedicated method to retrieve assets.\
We pass to it entity with `PathComponent`.\
After retrival was a success then `CacheComponent` is added and used during next retrivals.

```go
func AssetRetrieval(assetsService assets.Service, assetEntity ecs.EntityID) (ExampleAsset, error) {
	asset, err := assets.GetAsset[ExampleAsset](assetsService, assetEntity)
    return asset, err
}
```

### Releasing assets by destroying entity
```go
func AssetRelease(world ecs.World, assetEntity ecs.EntityID) {
    world.RemoveEntity(assetEntity)
}
```

### Releasing assets by destroying component
```go
func AssetRelease(assetsService assets.Service, assetEntity ecs.EntityID) {
	assetsService.Cache().Remove(assetEntity)
}
```

## Dependencies
- [datastructures](/engine/services/datastructures/readme/README.md)
- [ecs](/engine/services/ecs/readme/README.md)
- [logger](/engine/services/logger/readme/README.md)
- [registry](/engine/modules/registry/readme/README.md)
