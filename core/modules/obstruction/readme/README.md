# obstruction
## Architecture
defines how obstruction map is stored and accessed

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	core/modules/obstruction/test	0.005s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             54             11            319
-------------------------------------------------------------------------------
SUM:                             6             54             11            319
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/obstruction.Service`

#### method Service Collisions
Type: `func(core/modules/obstruction.AABB, core/modules/obstruction.Obstruction) []engine/modules/grid.Coords`

#### method Service Component
Type: `func() engine/services/ecs.ComponentsArray[core/modules/obstruction.Component]`

#### method Service Deployed
Type: `func() engine/services/ecs.ComponentsArray[core/modules/obstruction.DeployedComponent]`

#### method Service Grid
Type: `func() engine/modules/grid.ServiceT[core/modules/obstruction.Obstruction]`

#### method Service Register
Type: `func() error`

### type Obstruction
Type: `core/modules/obstruction.Obstruction`
mask of ways in which tile is obstructed

### type Component
Type: `core/modules/obstruction.Component`
Defines how entity or tile obstruct
On obstruction collision new entity is removed and warning is logged

#### property Component Obstruction
Type: `core/modules/obstruction.Obstruction`

### type AABB
Type: `core/modules/obstruction.AABB`
aabb on grid

#### property AABB Coords
Type: `core/modules/tile.PosComponent`

#### property AABB Size
Type: `core/modules/tile.SizeComponent`

#### property AABB Tiles
Type: `[]engine/modules/grid.Coords`

### type DeployedComponent
Type: `core/modules/obstruction.DeployedComponent`
adding and removing deployed component modifies obstruction component

## Variables
### var ErrPositionIsOccupied
Type: `error`

## Functions
### func NewObstruction
Type: `func(obstruction core/modules/obstruction.Obstruction) core/modules/obstruction.Component`

### func NewAABB
Type: `func(coords core/modules/tile.PosComponent, size core/modules/tile.SizeComponent) core/modules/obstruction.AABB`

### func NewDeployed
Type: `func() core/modules/obstruction.DeployedComponent`


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
  - `core/modules/obstruction.Component`
  - `core/modules/obstruction.Deployed`
  - `core/modules/obstruction.DeployedComponent`
  - `core/modules/obstruction.ErrPositionIsOccupied`
  - `core/modules/obstruction.Grid`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.NewObstruction`
  - `core/modules/obstruction.Obstruction`
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

`engine/modules/inputs`:
  - `engine/modules/inputs.Stack`
  - `engine/modules/inputs.StackComponent`

`engine/modules/record`:
  - `engine/modules/record.AddToConfig`
  - `engine/modules/record.ComponentGetter`
  - `engine/modules/record.Config`
  - `engine/modules/record.Entities`
  - `engine/modules/record.Entity`
  - `engine/modules/record.NewConfig`
  - `engine/modules/record.RecordingID`
  - `engine/modules/record.StartBackwardsRecording`
  - `engine/modules/record.Stop`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/services/ecs`:
  - `engine/services/ecs.AddDirtySet`
  - `engine/services/ecs.BeforeGet`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.DirtySet`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.NewDirtySet`
  - `engine/services/ecs.OnUpsert`
  - `engine/services/ecs.RemoveEntity`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SetEmpty`
  - `engine/services/ecs.SystemRegister`

### Third Party
- `github.com/ogiusek/ioc/v2`