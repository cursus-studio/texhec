# player
## Architecture
allowes objects to be owned

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               4             45              7            257
Markdown                         2              0              0              3
-------------------------------------------------------------------------------
SUM:                             6             45              7            260
-------------------------------------------------------------------------------
```
## TODO
Restrict actions to allow only user to perform his actions.
Perhaps attach `PlayerComponent` to camera.

## Types
### type Service
Type: `core/modules/player.Service`

#### method Service ActingConnection
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.ActingConnectionComponent]`

#### method Service ActingPlayer
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.ActingPlayerComponent]`

#### method Service ControlsEntity
Type: `func(engine/modules/ecs.EntityID) error`
returns nil if object is controled

#### method Service ControlsUUID
Type: `func(engine/modules/uuid.UUID) error`

#### method Service Owner
Type: `func() engine/modules/uuid.LinkService[core/modules/player.OwnerLink]`

#### method Service Player
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.PlayerComponent]`

#### method Service PlayerConnection
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.PlayerConnectionComponent]`

#### method Service PlayerUUID
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.PlayerUUIDComponent]`

#### method Service PlayersConnection
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.PlayersConnectionComponent]`

### type PlayerContext
Type: `core/modules/player.PlayerContext`
event context

#### property PlayerContext PlayerUUID
Type: `engine/modules/uuid.UUID`

### type PlayerComponent
Type: `core/modules/player.PlayerComponent`
marks that player is performing a move

#### property PlayerComponent Name
Type: `string`

### type PlayerUUIDComponent
Type: `core/modules/player.PlayerUUIDComponent`

#### property PlayerUUIDComponent UUID
Type: `engine/modules/uuid.UUID`

### type ActingPlayerComponent
Type: `core/modules/player.ActingPlayerComponent`

### type ActingConnectionComponent
Type: `core/modules/player.ActingConnectionComponent`

#### property ActingConnectionComponent Connection
Type: `engine/modules/uuid.UUID`

### type OwnerLink
Type: `core/modules/player.OwnerLink`

### type AssignActingPlayerDTO
Type: `core/modules/player.AssignActingPlayerDTO`

#### property AssignActingPlayerDTO MsgCtx
Type: `engine/modules/connection.MsgCtx`

#### property AssignActingPlayerDTO PlayerUUID
Type: `engine/modules/uuid.UUID`

### type PlayersConnectionComponent
Type: `core/modules/player.PlayersConnectionComponent`

### type PlayerConnectionComponent
Type: `core/modules/player.PlayerConnectionComponent`

#### property PlayerConnectionComponent Player
Type: `engine/modules/uuid.UUID`

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

### func NewPlayerUUID
Type: `func(uuid engine/modules/uuid.UUID) core/modules/player.PlayerUUIDComponent`

### func NewActingPlayer
Type: `func() core/modules/player.ActingPlayerComponent`

### func NewActiongConnection
Type: `func(uuid engine/modules/uuid.UUID) core/modules/player.ActingConnectionComponent`

### func NewAssignActingPlayerDTO
Type: `func(playerUUID engine/modules/uuid.UUID) core/modules/player.AssignActingPlayerDTO`

### func NewPlayersConnection
Type: `func() core/modules/player.PlayersConnectionComponent`

### func NewPlayerConnection
Type: `func(player engine/modules/uuid.UUID) core/modules/player.PlayerConnectionComponent`


## Dependencies
`core/game`:
  - `core/game.Economy`
  - `core/game.GameWorld`
  - `core/game.Player`

`core/modules/economy`:
  - `core/modules/economy.NewWallet`
  - `core/modules/economy.Wallet`

`core/modules/player`:
  - `core/modules/player.ActingConnectionComponent`
  - `core/modules/player.ActingPlayer`
  - `core/modules/player.ActingPlayerComponent`
  - `core/modules/player.AssignActingPlayerDTO`
  - `core/modules/player.ErrRequiresControl`
  - `core/modules/player.ErrRequiresOwner`
  - `core/modules/player.NewActingPlayer`
  - `core/modules/player.NewAssignActingPlayerDTO`
  - `core/modules/player.NewPlayerConnection`
  - `core/modules/player.NewPlayerUUID`
  - `core/modules/player.OwnerLink`
  - `core/modules/player.PlayerComponent`
  - `core/modules/player.PlayerConnectionComponent`
  - `core/modules/player.PlayerContext`
  - `core/modules/player.PlayerUUID`
  - `core/modules/player.PlayerUUIDComponent`
  - `core/modules/player.PlayersConnectionComponent`
  - `core/modules/player.Service`
  - `core/modules/player.UUID`

`engine/modules/connection`:
  - `engine/modules/connection.Component`
  - `engine/modules/connection.Conn`
  - `engine/modules/connection.ConnEntity`
  - `engine/modules/connection.MsgCtx`
  - `engine/modules/connection.Send`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/interactions`:
  - `engine/modules/interactions.ContextSetter`
  - `engine/modules/interactions.Feature`

`engine/modules/loop`:
  - `engine/modules/loop.NewEmitOnFrameEvent`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.ErrMissingUUID`
  - `engine/modules/uuid.Get`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.LinkService`
  - `engine/modules/uuid.New`
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.UUID`

`engine/modules/uuid/pkg`:
  - `engine/modules/uuid/pkg.LinkPkgT`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`