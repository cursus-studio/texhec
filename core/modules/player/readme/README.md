# player
## Architecture
allowes objects to be owned

## Types
### type Service
Type: `core/modules/player.Service`

#### method Service Owner
Type: `func() engine/services/ecs.ComponentsArray[core/modules/player.OwnerComponent]`

### type OwnerComponent
Type: `core/modules/player.OwnerComponent`

#### property OwnerComponent Owner
Type: `engine/services/ecs.EntityID`

## Functions
### func NewOwner
Type: `func(owner engine/services/ecs.EntityID) core/modules/player.OwnerComponent`


## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             14              2             48
-------------------------------------------------------------------------------
SUM:                             3             14              2             48
-------------------------------------------------------------------------------

```
## Dependencies
`core/game`:
  - `core/game.GameWorld`

`core/modules/player`:
  - `core/modules/player.OwnerComponent`
  - `core/modules/player.Service`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/services/ecs`:
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.GetComponentsArray`

### Third Party
- `github.com/ogiusek/ioc/v2`