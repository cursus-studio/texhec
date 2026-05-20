# generation
## Architecture
generates map using noise functions

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             43             11            262
Markdown                         1              3              0             13
-------------------------------------------------------------------------------
SUM:                             4             46             11            275
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/generation.Service`

#### method Service Generate
Type: `func(core/modules/generation.Config) engine/modules/batcher.Task`
adds to world all grids

### type Config
Type: `core/modules/generation.Config`

#### property Config Entity
Type: `engine/services/ecs.EntityID`

#### property Config Seed
Type: `engine/modules/seed.Seed`

#### property Config Size
Type: `engine/modules/grid.Coords`
will be generated <0, n)

## Functions
### func NewConfig
Type: `func(entity engine/services/ecs.EntityID, seed engine/modules/seed.Seed, size engine/modules/grid.Coords) core/modules/generation.Config`


## Dependencies
`core/game`:
  - `core/game.Definitions`
  - `core/game.Deploy`
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Tile`

`core/modules/generation`:
  - `core/modules/generation.Config`
  - `core/modules/generation.Entity`
  - `core/modules/generation.Seed`
  - `core/modules/generation.Service`
  - `core/modules/generation.Size`

`core/modules/obstruction`:
  - `core/modules/obstruction.Component`
  - `core/modules/obstruction.Grid`
  - `core/modules/obstruction.NewGrid`
  - `core/modules/obstruction.Obstruction`

`core/modules/tile`:
  - `core/modules/tile.Component`
  - `core/modules/tile.GetTile`
  - `core/modules/tile.GetTileSize`
  - `core/modules/tile.Grid`
  - `core/modules/tile.ID`
  - `core/modules/tile.NewGrid`

`engine/modules/batcher`:
  - `engine/modules/batcher.AddConcurrentBatch`
  - `engine/modules/batcher.AddOrderedBatch`
  - `engine/modules/batcher.Build`
  - `engine/modules/batcher.NewBatch`
  - `engine/modules/batcher.NewTask`
  - `engine/modules/batcher.Task`

`engine/modules/collider`:
  - `engine/modules/collider.Component`
  - `engine/modules/collider.NewCollider`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/grid`:
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.GetCoords`
  - `engine/modules/grid.GetIndex`
  - `engine/modules/grid.GetLastIndex`
  - `engine/modules/grid.GetTile`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.SetTile`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/modules/inputs`:
  - `engine/modules/inputs.Stack`
  - `engine/modules/inputs.StackComponent`

`engine/modules/metadata`:
  - `engine/modules/metadata.Name`
  - `engine/modules/metadata.NewName`

`engine/modules/noise`:
  - `engine/modules/noise.AddPerlin`
  - `engine/modules/noise.AddValue`
  - `engine/modules/noise.Build`
  - `engine/modules/noise.NewLayer`
  - `engine/modules/noise.NewNoise`
  - `engine/modules/noise.Read`

`engine/modules/seed`:
  - `engine/modules/seed.Seed`

`engine/modules/transform`:
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.PivotPoint`
  - `engine/modules/transform.Size`

`engine/services/datastructures`:
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.GetIndices`
  - `engine/services/datastructures.NewSparseArray`
  - `engine/services/datastructures.Set`

`engine/services/ecs`:
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.NewEntity`
  - `engine/services/ecs.Set`

### Third Party
- `github.com/go-gl/mathgl/mgl64`
- `github.com/ogiusek/ioc/v2`