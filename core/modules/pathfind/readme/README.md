# pathfind
## Architecture
finds path on a grid and moves objects to their target according to their speed

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               7            108             34            760
Markdown                         1              1              0              5
-------------------------------------------------------------------------------
SUM:                             8            109             34            765
-------------------------------------------------------------------------------
```
## TODO
Modify speed to allow moving many tiles per tick.
Modify pathfinding to do not use shortest route and instead use optimal path (optimize paths per chunks).

Research materials:
- `supreme commander`
- `multi agent pathfinding`

## Types
### type Service
Type: `core/modules/pathfind.Service`

#### method Service CanStep
Type: `func(engine/modules/grid.Coords, core/modules/tile.SizeComponent, core/modules/obstruction.Component, core/modules/pathfind.StepComponent) bool`

#### method Service CoordsRegion
Type: `func(engine/modules/grid.Coords, core/modules/obstruction.Obstruction) (core/modules/pathfind.Region, bool)`

#### method Service EntityRegion
Type: `func(engine/modules/ecs.EntityID) (core/modules/pathfind.Region, bool)`

#### method Service FindPath
Type: `func(core/modules/pathfind.FindPathEvent)`

#### method Service RegionObstruction
Type: `func(core/modules/pathfind.Region) (core/modules/obstruction.Obstruction, bool)`
region

#### method Service Register
Type: `func() error`

#### method Service ShareRegion
Type: `func(engine/modules/ecs.EntityID, engine/modules/grid.Coords) bool`

#### method Service Speed
Type: `func() engine/modules/ecs.ComponentArray[core/modules/pathfind.SpeedComponent]`

#### method Service Step
Type: `func() engine/modules/ecs.ComponentArray[core/modules/pathfind.StepComponent]`

#### method Service Target
Type: `func() engine/modules/ecs.ComponentArray[core/modules/pathfind.TargetComponent]`

### type TargetComponent
Type: `core/modules/pathfind.TargetComponent`
all entities without [tile.StepComponent] get one on tick which will move them towards target

#### property TargetComponent Coords
Type: `engine/modules/grid.Coords`

### type SpeedComponent
Type: `core/modules/pathfind.SpeedComponent`

#### property SpeedComponent InvSpeed
Type: `int8`
ticks to move one tile

### type StepComponent
Type: `core/modules/pathfind.StepComponent`
Step coords should be +/- 1 x or y from current target position.
Otherwise step will be removed and warning will be logged.

#### property StepComponent Coords
Type: `engine/modules/grid.Coords`

### type Region
Type: `core/modules/pathfind.Region`
this variable contains region index and is used for region connectivity

### type FindPathEvent
Type: `core/modules/pathfind.FindPathEvent`

#### property FindPathEvent Entity
Type: `engine/modules/ecs.EntityID`

#### property FindPathEvent Coords
Type: `engine/modules/grid.Coords`

## Variables
### var ErrInvalidPath
Type: `error`

### var ErrInvalidServiceOrder
Type: `error`

### var NotARegion
Type: `core/modules/pathfind.Region`

## Functions
### func NewTarget
Type: `func(coords engine/modules/grid.Coords) core/modules/pathfind.TargetComponent`

### func NewSpeed
Type: `func[Number golang.org/x/exp/constraints.Integer](invSpeed Number) core/modules/pathfind.SpeedComponent`

### func NewStep
Type: `func(x engine/modules/grid.Coord, y engine/modules/grid.Coord) core/modules/pathfind.StepComponent`

### func NewFindPathEvent
Type: `func(entity engine/modules/ecs.EntityID, coords engine/modules/grid.Coords) core/modules/pathfind.FindPathEvent`


## Dependencies
`core/game`:
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Pathfind`
  - `core/game.Tile`

`core/modules/actions`:
  - `core/modules/actions.Coords`
  - `core/modules/actions.CoordsCursorComponent`
  - `core/modules/actions.CoordsStep`
  - `core/modules/actions.Entity`
  - `core/modules/actions.FriendlyMobileEntityStep`
  - `core/modules/actions.RegionAnchorComponent`

`core/modules/obstruction`:
  - `core/modules/obstruction.Collisions`
  - `core/modules/obstruction.Component`
  - `core/modules/obstruction.Grid`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.Obstruction`
  - `core/modules/obstruction.Obstructions`

`core/modules/pathfind`:
  - `core/modules/pathfind.CanStep`
  - `core/modules/pathfind.Coords`
  - `core/modules/pathfind.Entity`
  - `core/modules/pathfind.ErrInvalidPath`
  - `core/modules/pathfind.ErrInvalidServiceOrder`
  - `core/modules/pathfind.FindPathEvent`
  - `core/modules/pathfind.InvSpeed`
  - `core/modules/pathfind.NewFindPathEvent`
  - `core/modules/pathfind.NewSpeed`
  - `core/modules/pathfind.NewStep`
  - `core/modules/pathfind.NewTarget`
  - `core/modules/pathfind.Region`
  - `core/modules/pathfind.Service`
  - `core/modules/pathfind.Speed`
  - `core/modules/pathfind.SpeedComponent`
  - `core/modules/pathfind.Step`
  - `core/modules/pathfind.StepComponent`
  - `core/modules/pathfind.Target`
  - `core/modules/pathfind.TargetComponent`

`core/modules/tile`:
  - `core/modules/tile.Aligned`
  - `core/modules/tile.Coord`
  - `core/modules/tile.Coords`
  - `core/modules/tile.ErrInvalidPosition`
  - `core/modules/tile.ErrInvalidStep`
  - `core/modules/tile.ErrPositionAndSpeedIsRequiredToStep`
  - `core/modules/tile.GetTile`
  - `core/modules/tile.Grid`
  - `core/modules/tile.NewPos`
  - `core/modules/tile.NewRot`
  - `core/modules/tile.NewSize`
  - `core/modules/tile.Pos`
  - `core/modules/tile.PosComponent`
  - `core/modules/tile.Rot`
  - `core/modules/tile.RotComponent`
  - `core/modules/tile.Size`
  - `core/modules/tile.SizeComponent`
  - `core/modules/tile.UnloadChunkEvent`
  - `core/modules/tile.X`
  - `core/modules/tile.Y`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.NewSparseSet`
  - `engine/modules/datastructures.SparseArray`
  - `engine/modules/datastructures.SparseSet`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.SystemRegister`
  - `engine/modules/ecs.World`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/grid`:
  - `engine/modules/grid.AbsoluteCoords`
  - `engine/modules/grid.Chunk`
  - `engine/modules/grid.ChunkCoordsComponent`
  - `engine/modules/grid.Component`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.CoordsData`
  - `engine/modules/grid.CoordsIndex`
  - `engine/modules/grid.GetTile`
  - `engine/modules/grid.GetTiles`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.IndexCoords`
  - `engine/modules/grid.NewChunk`
  - `engine/modules/grid.NewCoord`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.RelativeCoords`
  - `engine/modules/grid.ServiceT`
  - `engine/modules/grid.SetTile`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/modules/grid/pkg`:
  - `engine/modules/grid/pkg.PkgT`

`engine/modules/interactions/pkg`:
  - `engine/modules/interactions/pkg.FeaturePkg`
  - `engine/modules/interactions/pkg.NewCopyRelation`

`engine/modules/loop`:
  - `engine/modules/loop.TickEvent`

`engine/modules/relation`:
  - `engine/modules/relation.Get`
  - `engine/modules/relation.Service`

`engine/modules/relation/pkg`:
  - `engine/modules/relation/pkg.MapRelationPkg`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`