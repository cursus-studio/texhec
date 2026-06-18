# pathfind
## Architecture
finds path on a grid and moves objects to their target according to their speed

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             71             27            424
Markdown                         1              1              0              5
-------------------------------------------------------------------------------
SUM:                             7             72             27            429
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

#### method Service FindPath
Type: `func(core/modules/pathfind.FindPathEvent)`

#### method Service Register
Type: `func() error`

#### method Service Speed
Type: `func() engine/services/ecs.ComponentsArray[core/modules/pathfind.SpeedComponent]`

#### method Service Step
Type: `func() engine/services/ecs.ComponentsArray[core/modules/pathfind.StepComponent]`

#### method Service Target
Type: `func() engine/services/ecs.ComponentsArray[core/modules/pathfind.TargetComponent]`

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

### type FindPathEvent
Type: `core/modules/pathfind.FindPathEvent`

#### property FindPathEvent Entity
Type: `engine/services/ecs.EntityID`

#### property FindPathEvent Coords
Type: `engine/modules/grid.Coords`

## Variables
### var ErrInvalidPath
Type: `error`

## Functions
### func NewTarget
Type: `func(coords engine/modules/grid.Coords) core/modules/pathfind.TargetComponent`

### func NewSpeed
Type: `func[Number golang.org/x/exp/constraints.Integer](invSpeed Number) core/modules/pathfind.SpeedComponent`

### func NewStep
Type: `func(x engine/modules/grid.Coord, y engine/modules/grid.Coord) core/modules/pathfind.StepComponent`

### func NewFeatureFindPathEvent
Type: `func() engine/modules/interactions.FeatureEvent[core/modules/pathfind.FindPathEvent]`

### func NewFindPathEvent
Type: `func(entity engine/services/ecs.EntityID, coords engine/modules/grid.Coords) core/modules/pathfind.FindPathEvent`


## Dependencies
`core/game`:
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Pathfind`
  - `core/game.Tile`

`core/modules/obstruction`:
  - `core/modules/obstruction.Collisions`
  - `core/modules/obstruction.Component`
  - `core/modules/obstruction.Grid`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.Obstruction`

`core/modules/pathfind`:
  - `core/modules/pathfind.CanStep`
  - `core/modules/pathfind.Coords`
  - `core/modules/pathfind.Entity`
  - `core/modules/pathfind.ErrInvalidPath`
  - `core/modules/pathfind.FindPathEvent`
  - `core/modules/pathfind.InvSpeed`
  - `core/modules/pathfind.NewSpeed`
  - `core/modules/pathfind.NewStep`
  - `core/modules/pathfind.NewTarget`
  - `core/modules/pathfind.Service`
  - `core/modules/pathfind.Speed`
  - `core/modules/pathfind.SpeedComponent`
  - `core/modules/pathfind.Step`
  - `core/modules/pathfind.StepComponent`
  - `core/modules/pathfind.TargetComponent`

`core/modules/tile`:
  - `core/modules/tile.Aligned`
  - `core/modules/tile.Coord`
  - `core/modules/tile.Coords`
  - `core/modules/tile.CoordsInteraction`
  - `core/modules/tile.Entity`
  - `core/modules/tile.ErrInvalidPosition`
  - `core/modules/tile.ErrInvalidStep`
  - `core/modules/tile.ErrPositionAndSpeedIsRequiredToStep`
  - `core/modules/tile.NewPos`
  - `core/modules/tile.NewRot`
  - `core/modules/tile.NewSize`
  - `core/modules/tile.ObjectInteraction`
  - `core/modules/tile.Pos`
  - `core/modules/tile.PosComponent`
  - `core/modules/tile.Rot`
  - `core/modules/tile.RotComponent`
  - `core/modules/tile.Size`
  - `core/modules/tile.SizeComponent`
  - `core/modules/tile.X`
  - `core/modules/tile.Y`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/grid`:
  - `engine/modules/grid.Component`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.CoordsData`
  - `engine/modules/grid.GetTile`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.NewCoord`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/modules/interactions`:
  - `engine/modules/interactions.FeatureEntity`
  - `engine/modules/interactions.FeatureEvent`
  - `engine/modules/interactions.Interaction`
  - `engine/modules/interactions.State`

`engine/modules/interactions/pkg`:
  - `engine/modules/interactions/pkg.FeaturePkg`

`engine/modules/loop`:
  - `engine/modules/loop.TickEvent`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/services/ecs`:
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.Remove`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SystemRegister`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`