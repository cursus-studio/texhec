# attack
## Architecture
defines attack feature

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             21              3            142
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                             4             21              3            143
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/attack.Service`

#### method Service Reach
Type: `func() core/modules/reach.ServiceT[core/modules/attack.TargetComponent]`

#### method Service Register
Type: `func() error`

#### method Service Target
Type: `func() engine/modules/ecs.ComponentArray[core/modules/attack.TargetComponent]`

### type TargetComponent
Type: `core/modules/attack.TargetComponent`

#### property TargetComponent Entity
Type: `engine/modules/ecs.EntityID`

## Variables
### var ErrCannotAttackEnemyOutOfReach
Type: `error`

## Functions
### func NewTarget
Type: `func(target engine/modules/ecs.EntityID) core/modules/attack.TargetComponent`


## Dependencies
`core/game`:
  - `core/game.Attack`
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Pathfind`
  - `core/game.Tile`

`core/modules/actions`:
  - `core/modules/actions.EnemyEntityStep`
  - `core/modules/actions.Entity`
  - `core/modules/actions.FriendlyOffensiveEntityStep`

`core/modules/attack`:
  - `core/modules/attack.Entity`
  - `core/modules/attack.ErrCannotAttackEnemyOutOfReach`
  - `core/modules/attack.NewTarget`
  - `core/modules/attack.Reach`
  - `core/modules/attack.Service`
  - `core/modules/attack.Target`
  - `core/modules/attack.TargetComponent`

`core/modules/obstruction`:
  - `core/modules/obstruction.Collisions`
  - `core/modules/obstruction.Component`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.Obstruction`

`core/modules/pathfind`:
  - `core/modules/pathfind.NewTarget`
  - `core/modules/pathfind.Speed`
  - `core/modules/pathfind.Target`

`core/modules/reach`:
  - `core/modules/reach.Component`
  - `core/modules/reach.NewReach`
  - `core/modules/reach.Reaches`
  - `core/modules/reach.ServiceT`
  - `core/modules/reach.TilesFrom`

`core/modules/reach/pkg`:
  - `core/modules/reach/pkg.PkgT`

`core/modules/tile`:
  - `core/modules/tile.NewPos`
  - `core/modules/tile.Pos`
  - `core/modules/tile.Size`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewSetEvent`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/grid`:
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`

`engine/modules/interactions/pkg`:
  - `engine/modules/interactions/pkg.FeaturePkg`

`engine/modules/loop`:
  - `engine/modules/loop.TickEvent`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`