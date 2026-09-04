# obstruction
## Architecture
defines how obstruction map is stored and accessed

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	core/modules/obstruction/test	0.013s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             57             11            315
-------------------------------------------------------------------------------
SUM:                             6             57             11            315
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/obstruction.Service`

#### method Service Collisions
Type: `func(core/modules/obstruction.AABB, core/modules/obstruction.Obstruction) []engine/modules/grid.Coords`

#### method Service Component
Type: `func() engine/modules/ecs.ComponentArray[core/modules/obstruction.Component]`

#### method Service Deployed
Type: `func() engine/modules/ecs.ComponentArray[core/modules/obstruction.DeployedComponent]`

#### method Service Grid
Type: `func() engine/modules/grid.ServiceT[core/modules/obstruction.Obstruction]`

#### method Service Obstructions
Type: `func() engine/modules/datastructures.SparseSet[core/modules/obstruction.Obstruction]`

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

### Third Party
- `github.com/ogiusek/ioc/v2`