# generation
## Architecture
generates map using noise functions

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             43             15            282
Markdown                         1              3              0             13
-------------------------------------------------------------------------------
SUM:                             4             46             15            295
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/generation.Service`

#### method Service Register
Type: `func() error`


## Dependencies
`core/game`:
  - `core/game.Definitions`
  - `core/game.Deploy`
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Tile`

`core/modules/generation`:
  - `core/modules/generation.Service`

`core/modules/tile`:
  - `core/modules/tile.Component`
  - `core/modules/tile.Config`
  - `core/modules/tile.Coords`
  - `core/modules/tile.GetConfig`
  - `core/modules/tile.GetTile`
  - `core/modules/tile.GetTileSize`
  - `core/modules/tile.Grid`
  - `core/modules/tile.ID`
  - `core/modules/tile.MissingChunkEvent`
  - `core/modules/tile.Seed`

`engine/modules/batcher`:
  - `engine/modules/batcher.AddConcurrentBatch`
  - `engine/modules/batcher.AddOrderedBatch`
  - `engine/modules/batcher.Build`
  - `engine/modules/batcher.NewBatch`
  - `engine/modules/batcher.NewTask`
  - `engine/modules/batcher.Queue`

`engine/modules/collider`:
  - `engine/modules/collider.Component`
  - `engine/modules/collider.NewCollider`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/grid`:
  - `engine/modules/grid.AbsoluteCoords`
  - `engine/modules/grid.Chunk`
  - `engine/modules/grid.ChunkSize`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.CoordsIndex`
  - `engine/modules/grid.GetLastIndex`
  - `engine/modules/grid.GetTile`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.IndexCoords`
  - `engine/modules/grid.NewChunk`
  - `engine/modules/grid.NewChunkCoords`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.SetTile`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/modules/groups`:
  - `engine/modules/groups.Inherit`
  - `engine/modules/groups.InheritGroupsComponent`

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

`engine/modules/transform`:
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.PivotPoint`
  - `engine/modules/transform.Pos`
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
  - `engine/services/ecs.SystemRegister`

### Third Party
- `github.com/go-gl/mathgl/mgl64`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`