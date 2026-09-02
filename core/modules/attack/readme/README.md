# attack
## Architecture
defines attack feature

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               4             31              4            229
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                             5             31              4            230
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/attack.Service`

#### method Service Damage
Type: `func() engine/modules/ecs.ComponentArray[core/modules/attack.DamageComponent]`

#### method Service FullHealth
Type: `func(engine/modules/ecs.EntityID) (core/modules/attack.HealthComponent, bool)`

#### method Service Health
Type: `func() engine/modules/ecs.ComponentArray[core/modules/attack.HealthComponent]`

#### method Service Reach
Type: `func() core/modules/reach.ServiceT[core/modules/attack.TargetComponent]`

#### method Service Register
Type: `func() error`

#### method Service Target
Type: `func() engine/modules/ecs.ComponentArray[core/modules/attack.TargetComponent]`

### type Health
Type: `core/modules/attack.Health`

### type TargetComponent
Type: `core/modules/attack.TargetComponent`

#### property TargetComponent Entity
Type: `engine/modules/ecs.EntityID`

### type HealthComponent
Type: `core/modules/attack.HealthComponent`

#### property HealthComponent Health
Type: `core/modules/attack.Health`

#### method HealthComponent Smooth
Type: `func()`

#### method HealthComponent Lerp
Type: `func(c2 core/modules/attack.HealthComponent, mix32 float32) core/modules/attack.HealthComponent`

### type DamageComponent
Type: `core/modules/attack.DamageComponent`

#### property DamageComponent Damage
Type: `core/modules/attack.Health`

## Variables
### var ErrCannotAttackEnemyOutOfReach
Type: `error`

## Functions
### func NewTarget
Type: `func(target engine/modules/ecs.EntityID) core/modules/attack.TargetComponent`

### func NewHealth
Type: `func(health core/modules/attack.Health) core/modules/attack.HealthComponent`

### func NewDamage
Type: `func(damage core/modules/attack.Health) core/modules/attack.DamageComponent`


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
  - `core/modules/attack.Damage`
  - `core/modules/attack.DamageComponent`
  - `core/modules/attack.Entity`
  - `core/modules/attack.ErrCannotAttackEnemyOutOfReach`
  - `core/modules/attack.Health`
  - `core/modules/attack.HealthComponent`
  - `core/modules/attack.NewDamage`
  - `core/modules/attack.NewHealth`
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
  - `core/modules/tile.Blueprint`
  - `core/modules/tile.NewPos`
  - `core/modules/tile.Pos`
  - `core/modules/tile.Size`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
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
  - `engine/modules/loop.Delta`
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.TickEvent`

`engine/modules/transition`:
  - `engine/modules/transition.LerpInt`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`