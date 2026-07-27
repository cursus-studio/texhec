# attack
## Architecture
defines attack feature

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             14              1             65
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                             4             14              1             66
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/attack.Service`

#### method Service AttackEvent
Type: `func(core/modules/attack.AttackEvent)`

### type AttackEvent
Type: `core/modules/attack.AttackEvent`

#### property AttackEvent By
Type: `core/modules/actions.FriendlyBuilderObjectStep`

#### property AttackEvent Target
Type: `core/modules/actions.EnemyObjectStep`

## Functions
### func NewDeployEvent
Type: `func(by engine/modules/ecs.EntityID, target engine/modules/ecs.EntityID) core/modules/attack.AttackEvent`


## Dependencies
`core/game`:
  - `core/game.GameWorld`

`core/modules/actions`:
  - `core/modules/actions.CoordsAnchorComponent`
  - `core/modules/actions.EnemyObjectStep`
  - `core/modules/actions.Entity`
  - `core/modules/actions.FriendlyBuilderObjectStep`
  - `core/modules/actions.NewObjectInteraction`

`core/modules/attack`:
  - `core/modules/attack.AttackEvent`
  - `core/modules/attack.By`
  - `core/modules/attack.Service`
  - `core/modules/attack.Target`

`engine/modules/ecs`:
  - `engine/modules/ecs.EntityID`

`engine/modules/interactions`:
  - `engine/modules/interactions.NewStep`
  - `engine/modules/interactions.State`

`engine/modules/interactions/pkg`:
  - `engine/modules/interactions/pkg.FeaturePkg`
  - `engine/modules/interactions/pkg.NewCopyRelation`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`