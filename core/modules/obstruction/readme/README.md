# obstruction
## Architecture
defines how obstruction map is stored and accessed

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	core/modules/obstruction/test	0.037s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             56             11            317
-------------------------------------------------------------------------------
SUM:                             6             56             11            317
-------------------------------------------------------------------------------
```
## Dependencies
`core/game`:
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Tile`

`core/modules/definitions`:
  - `core/modules/definitions.AirspaceObstruction`
  - `core/modules/definitions.LowlandObstruction`
  - `core/modules/definitions.WaterObstruction`

`core/modules/obstruction`:
  - `core/modules/obstruction.AABB`
  - `core/modules/obstruction.Deployed`
  - `core/modules/obstruction.DeployedComponent`
  - `core/modules/obstruction.Grid`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.NewObstruction`
  - `core/modules/obstruction.Obstrucion`
  - `core/modules/obstruction.Obstruction`
  - `core/modules/obstruction.ObstructionComponent`
  - `core/modules/obstruction.Service`
  - `core/modules/obstruction.Tiles`

`core/modules/tile`:
  - `core/modules/tile.Coord`
  - `core/modules/tile.ErrInvalidPosition`
  - `core/modules/tile.Pos`
  - `core/modules/tile.PosComponent`
  - `core/modules/tile.Size`
  - `core/modules/tile.SizeComponent`
  - `core/modules/tile.X`
  - `core/modules/tile.Y`

`core/pkg`:
  - `core/pkg.Pkg`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSparseSet`
  - `engine/modules/datastructures.SparseSet`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/grid`:
  - `engine/modules/grid.Chunk`
  - `engine/modules/grid.Component`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.CoordsData`
  - `engine/modules/grid.Entity`
  - `engine/modules/grid.GetTile`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.ServiceT`
  - `engine/modules/grid.SetTile`

`engine/modules/grid/pkg`:
  - `engine/modules/grid/pkg.PkgT`

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.ComponentPkg`

`engine/modules/record`:
  - `engine/modules/record.AddToConfig`
  - `engine/modules/record.ComponentGetter`
  - `engine/modules/record.Config`
  - `engine/modules/record.Entities`
  - `engine/modules/record.NewConfig`
  - `engine/modules/record.RecordingID`
  - `engine/modules/record.StartBackwardsRecording`
  - `engine/modules/record.Stop`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/ogiusek/ioc/v2`