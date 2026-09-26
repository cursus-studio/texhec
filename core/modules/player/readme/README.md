# player
## Architecture
allowes objects to be owned

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             43             14            279
Markdown                         2              0              0              3
-------------------------------------------------------------------------------
SUM:                             5             43             14            282
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

#### method Service GetContext
Type: `func() core/modules/player.PlayerContext`

#### method Service Owner
Type: `func() engine/modules/uuid.LinkService[core/modules/player.OwnerLink]`
points to player entity

#### method Service Player
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.PlayerComponent]`
attached to player

#### method Service PlayerConnection
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.PlayerConnectionComponent]`

#### method Service PlayerControlsEntity
Type: `func(player engine/modules/uuid.UUID, object engine/modules/ecs.EntityID) error`

#### method Service PlayerControlsUUID
Type: `func(player engine/modules/uuid.UUID, object engine/modules/uuid.UUID) error`

#### method Service PlayerUUID
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.PlayerUUIDComponent]`

#### method Service PlayersConnection
Type: `func() engine/modules/ecs.ComponentArray[core/modules/player.PlayersConnectionComponent]`
attached to host and connection

### type PlayerContextGetter
Type: `core/modules/player.PlayerContextGetter`
event context

#### method PlayerContextGetter Context
Type: `func() core/modules/player.PlayerContext`

### type PlayerContext
Type: `core/modules/player.PlayerContext`

#### property PlayerContext PlayerUUID
Type: `engine/modules/uuid.UUID`

#### method PlayerContext Context
Type: `func() core/modules/player.PlayerContext`

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
Type: `engine/modules/ecs.EntityID`

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
Type: `func(entity engine/modules/ecs.EntityID) core/modules/player.ActingConnectionComponent`

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
  - `core/modules/player.ActingConnection`
  - `core/modules/player.ActingConnectionComponent`
  - `core/modules/player.ActingPlayerComponent`
  - `core/modules/player.AssignActingPlayerDTO`
  - `core/modules/player.Connection`
  - `core/modules/player.Context`
  - `core/modules/player.ErrRequiresControl`
  - `core/modules/player.ErrRequiresOwner`
  - `core/modules/player.NewActingPlayer`
  - `core/modules/player.NewActiongConnection`
  - `core/modules/player.NewAssignActingPlayerDTO`
  - `core/modules/player.NewPlayerConnection`
  - `core/modules/player.NewPlayerUUID`
  - `core/modules/player.OwnerLink`
  - `core/modules/player.PlayerComponent`
  - `core/modules/player.PlayerConnectionComponent`
  - `core/modules/player.PlayerContext`
  - `core/modules/player.PlayerContextGetter`
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

`engine/modules/loop`:
  - `engine/modules/loop.NewEmitOnFrameEvent`

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.AddGenericEventAuthorization`
  - `engine/modules/netsync/pkg.Config`

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