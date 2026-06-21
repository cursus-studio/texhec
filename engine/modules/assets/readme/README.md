# assets
## Architecture
We store all data in `ecs.World` and build on top of it.\
We store path in `PathComponent` and upon retrival we store asset as interface in `CacheComponent`.\
We convert `PathComponent` to `CacheComponent` using dispatchers where each file extension has dedicated dispatcher.\
Using interfaces for `CacheComponent` doesn't affect performance heavily, its nothing in comparison to data stored and its processing.\
To release assets we just remove `CacheComponent` (recommended) or entity with this component.\
`CacheComponent` stores interface and we use\
`func GetAsset[Asset any](assets Service, assetID ecs.EntityID) (Asset, error)`\
to parse is to our asset type.

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/assets/test	0.006s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             60              5            274
Markdown                         1              0              0              8
-------------------------------------------------------------------------------
SUM:                             7             60              5            282
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/assets.Service`

#### method Service Cache
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/assets.CacheComponent]`

#### method Service Get
Type: `func(engine/modules/ecs.EntityID) (engine/modules/assets.Asset, error)`
get also caches asset

#### method Service Path
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/assets.PathComponent]`

#### method Service Register
Type: `func(extension string, dispatcher func(path engine/modules/assets.PathComponent) (engine/modules/assets.Asset, error))`

### type Asset
Type: `engine/modules/assets.Asset`

#### method Asset Release
Type: `func()`

### type PathComponent
Type: `engine/modules/assets.PathComponent`

#### property PathComponent Path
Type: `string`

#### method PathComponent Extension
Type: `func() string`

### type CacheComponent
Type: `engine/modules/assets.CacheComponent`

#### property CacheComponent Cache
Type: `engine/modules/assets.Asset`

## Variables
### var ErrAssetHasDifferentType
Type: `error`

### var ErrAssetNotFound
Type: `error`

## Functions
### func NewPath
Type: `func(path string) engine/modules/assets.PathComponent`

### func NewCache
Type: `func(cache engine/modules/assets.Asset) engine/modules/assets.CacheComponent`

### func GetAsset
Type: `func[Asset any](assets engine/modules/assets.Service, assetID engine/modules/ecs.EntityID) (Asset, error)`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Logger`
  - `engine.World`

`engine/modules/assets`:
  - `engine/modules/assets.Asset`
  - `engine/modules/assets.Cache`
  - `engine/modules/assets.CacheComponent`
  - `engine/modules/assets.ErrAssetNotFound`
  - `engine/modules/assets.Extension`
  - `engine/modules/assets.NewCache`
  - `engine/modules/assets.NewPath`
  - `engine/modules/assets.Path`
  - `engine/modules/assets.PathComponent`
  - `engine/modules/assets.Release`
  - `engine/modules/assets.Service`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.SparseArray`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/ogiusek/ioc/v2`