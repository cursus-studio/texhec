# pathfind
## Architecture
finds path on a grid and moves objects to their target according to their speed

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             97             94            523
-------------------------------------------------------------------------------
SUM:                             6             97             94            523
-------------------------------------------------------------------------------

```
## Types
### type Service
Type: `core/modules/pathfind.Service`

#### method Service CanStep
Type: `func(engine/modules/grid.Coords, core/modules/tile.SizeComponent, core/modules/obstruction.Component, core/modules/pathfind.StepComponent) bool`

#### method Service FindPath
Type: `func(core/modules/pathfind.FindPathEvent)`

#### method Service PreviewPath
Type: `func(core/modules/pathfind.PreviewPathEvent)`

#### method Service Register
Type: `func() error`

#### method Service Select
Type: `func(core/modules/pathfind.SelectEvent)`

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

### type SelectEvent
Type: `core/modules/pathfind.SelectEvent`
Select object.
Add in gui some indicator.
Change on click event.

#### property SelectEvent Entity
Type: `engine/services/ecs.EntityID`

### type PreviewPathEvent
Type: `core/modules/pathfind.PreviewPathEvent`
Select object.
Add in gui some indicator.
Perform all checks and costs

#### property PreviewPathEvent Entity
Type: `engine/services/ecs.EntityID`

#### property PreviewPathEvent Coords
Type: `engine/modules/grid.Coords`

#### method PreviewPathEvent ApplyCoords
Type: `func(coords engine/modules/grid.Coords) any`

### type FindPathEvent
Type: `core/modules/pathfind.FindPathEvent`
Adds [TargetComponent] to entity

#### property FindPathEvent Entity
Type: `engine/services/ecs.EntityID`

#### property FindPathEvent Coords
Type: `engine/modules/grid.Coords`

#### method FindPathEvent ApplyCoords
Type: `func(coords engine/modules/grid.Coords) any`

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

### func NewSelectEvent
Type: `func(entity engine/services/ecs.EntityID) core/modules/pathfind.SelectEvent`

### func NewPreviewPathEvent
Type: `func(entity engine/services/ecs.EntityID) core/modules/pathfind.PreviewPathEvent`

### func NewFindPathEvent
Type: `func(entity engine/services/ecs.EntityID) core/modules/pathfind.FindPathEvent`


## Dependencies
`core/game`:
  - `core/game.Definitions`
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Pathfind`
  - `core/game.Tile`
  - `core/game.Ui`

`core/modules/definitions`:
  - `core/modules/definitions.Assets`
  - `core/modules/definitions.Can`
  - `core/modules/definitions.Cannot`
  - `core/modules/definitions.GameGroup`
  - `core/modules/definitions.Hud`
  - `core/modules/definitions.PathLayer`
  - `core/modules/definitions.SquareCollider`
  - `core/modules/definitions.SquareMesh`
  - `core/modules/definitions.Target`

`core/modules/obstruction`:
  - `core/modules/obstruction.Collisions`
  - `core/modules/obstruction.Component`
  - `core/modules/obstruction.Grid`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.Obstruction`

`core/modules/pathfind`:
  - `core/modules/pathfind.ApplyCoords`
  - `core/modules/pathfind.CanStep`
  - `core/modules/pathfind.Coords`
  - `core/modules/pathfind.Entity`
  - `core/modules/pathfind.ErrInvalidPath`
  - `core/modules/pathfind.FindPathEvent`
  - `core/modules/pathfind.InvSpeed`
  - `core/modules/pathfind.NewFindPathEvent`
  - `core/modules/pathfind.NewPreviewPathEvent`
  - `core/modules/pathfind.NewSpeed`
  - `core/modules/pathfind.NewStep`
  - `core/modules/pathfind.NewTarget`
  - `core/modules/pathfind.PreviewPathEvent`
  - `core/modules/pathfind.SelectEvent`
  - `core/modules/pathfind.Service`
  - `core/modules/pathfind.Speed`
  - `core/modules/pathfind.SpeedComponent`
  - `core/modules/pathfind.Step`
  - `core/modules/pathfind.StepComponent`
  - `core/modules/pathfind.TargetComponent`

`core/modules/tile`:
  - `core/modules/tile.Aligned`
  - `core/modules/tile.Coord`
  - `core/modules/tile.ErrInvalidPosition`
  - `core/modules/tile.ErrInvalidStep`
  - `core/modules/tile.ErrPositionAndSpeedIsRequiredToStep`
  - `core/modules/tile.Layer`
  - `core/modules/tile.NewLayer`
  - `core/modules/tile.NewPos`
  - `core/modules/tile.NewRot`
  - `core/modules/tile.NewSelectEvent`
  - `core/modules/tile.NewSize`
  - `core/modules/tile.Pos`
  - `core/modules/tile.PosComponent`
  - `core/modules/tile.Rot`
  - `core/modules/tile.RotComponent`
  - `core/modules/tile.Size`
  - `core/modules/tile.SizeComponent`
  - `core/modules/tile.X`
  - `core/modules/tile.Y`

`core/modules/ui`:
  - `core/modules/ui.ActionComponent`
  - `core/modules/ui.Actions`
  - `core/modules/ui.Entities`
  - `core/modules/ui.NewUnselect`
  - `core/modules/ui.ObjectComponent`
  - `core/modules/ui.Objects`
  - `core/modules/ui.SelectEvent`

`engine/modules/collider`:
  - `engine/modules/collider.Component`
  - `engine/modules/collider.NewCollider`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/grid`:
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.GetIndex`
  - `engine/modules/grid.GetTile`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/modules/groups`:
  - `engine/modules/groups.Component`
  - `engine/modules/groups.EmptyGroups`
  - `engine/modules/groups.Enable`
  - `engine/modules/groups.Ptr`
  - `engine/modules/groups.Val`

`engine/modules/inputs`:
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewLeftClick`

`engine/modules/loop`:
  - `engine/modules/loop.TickEvent`

`engine/modules/render`:
  - `engine/modules/render.Mesh`
  - `engine/modules/render.NewMesh`
  - `engine/modules/render.NewTexture`
  - `engine/modules/render.Texture`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/services/datastructures`:
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.NewSparseArray`
  - `engine/services/datastructures.Set`
  - `engine/services/datastructures.SparseArray`

`engine/services/ecs`:
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.NewEntity`
  - `engine/services/ecs.Remove`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SystemRegister`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`