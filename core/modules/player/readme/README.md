# player
## Architecture
allowes objects to be owned

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             14              2             48
Markdown                         1              0              0              2
-------------------------------------------------------------------------------
SUM:                             4             14              2             50
-------------------------------------------------------------------------------
```
## TODO
Restrict actions to allow only user to perform his actions.
Perhaps attach `PlayerComponent` to camera.

## Types
### type Service
Type: `core/modules/player.Service`

#### method Service Owner
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.OwnerComponent]`

### type OwnerComponent
Type: `core/modules/player.OwnerComponent`

#### property OwnerComponent Owner
Type: `engine/modules/ecs.EntityID`

## Functions
### func NewOwner
Type: `func(owner engine/modules/ecs.EntityID) core/modules/player.OwnerComponent`


## Dependencies
`core/game`:
  - `core/game.GameWorld`

`core/modules/player`:
  - `core/modules/player.OwnerComponent`
  - `core/modules/player.Service`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/ioc/v2`