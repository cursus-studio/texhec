# attack
## Architecture
defines attack feature

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             16              1             69
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                             4             16              1             70
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
Type: `engine/modules/ecs.EntityID`

#### property AttackEvent Target
Type: `engine/modules/ecs.EntityID`

## Functions
### func NewAttackEvent
Type: `func(by engine/modules/ecs.EntityID, target engine/modules/ecs.EntityID) core/modules/attack.AttackEvent`


## Dependencies
`core/game`:
  - `core/game.GameWorld`

`core/modules/actions`:
  - `core/modules/actions.AnchorComponent`
  - `core/modules/actions.EnemyEntityStep`
  - `core/modules/actions.Entity`
  - `core/modules/actions.FriendlyBuilderEntityStep`

`core/modules/attack`:
  - `core/modules/attack.AttackEvent`
  - `core/modules/attack.NewAttackEvent`
  - `core/modules/attack.Service`
  - `core/modules/attack.Target`

`engine/modules/ecs`:
  - `engine/modules/ecs.EntityID`

`engine/modules/interactions/pkg`:
  - `engine/modules/interactions/pkg.FeaturePkg`
  - `engine/modules/interactions/pkg.NewCopyRelation`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`