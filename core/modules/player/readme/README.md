# player
## Architecture
allowes objects to be owned

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             20              3             95
Markdown                         2              0              0              3
-------------------------------------------------------------------------------
SUM:                             5             20              3             98
-------------------------------------------------------------------------------
```
## TODO
Restrict actions to allow only user to perform his actions.
Perhaps attach `PlayerComponent` to camera.

## Types
### type Service
Type: `core/modules/player.Service`

#### method Service ActingPlayer
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.ActingPlayerComponent]`

#### method Service ControlsObject
Type: `func(engine/modules/ecs.EntityID) error`
returns nil if object is controled

#### method Service Owner
Type: `func() engine/modules/uuid.LinkService[core/modules/player.OwnerLink]`

#### method Service Player
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.PlayerComponent]`

### type PlayerComponent
Type: `core/modules/player.PlayerComponent`
marks that player is performing a move

#### property PlayerComponent Name
Type: `string`

### type ActingPlayerComponent
Type: `core/modules/player.ActingPlayerComponent`

### type OwnerLink
Type: `core/modules/player.OwnerLink`

## Variables
### var ErrRequiresOwner
Type: `error`

### var ErrRequiresControl
Type: `error`

### var ErrRequiresToBeEnemy
Type: `error`

## Functions
### func NewPlayer
Type: `func(name string) core/modules/player.PlayerComponent`

### func NewActingPlayer
Type: `func() core/modules/player.ActingPlayerComponent`


## Dependencies
`core/game`:
  - `core/game.Economy`
  - `core/game.GameWorld`

`core/modules/economy`:
  - `core/modules/economy.NewWallet`
  - `core/modules/economy.Wallet`

`core/modules/player`:
  - `core/modules/player.ActingPlayerComponent`
  - `core/modules/player.ErrRequiresControl`
  - `core/modules/player.ErrRequiresOwner`
  - `core/modules/player.OwnerLink`
  - `core/modules/player.PlayerComponent`
  - `core/modules/player.Service`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.Get`
  - `engine/modules/uuid.LinkService`
  - `engine/modules/uuid.New`
  - `engine/modules/uuid.NewUUID`

`engine/modules/uuid/pkg`:
  - `engine/modules/uuid/pkg.LinkPkgT`

### Third Party
- `github.com/ogiusek/ioc/v2`