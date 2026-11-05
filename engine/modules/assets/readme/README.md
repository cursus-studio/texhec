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
ok  	engine/modules/assets/test	0.013s
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