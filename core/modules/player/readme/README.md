# player
## Architecture
allowes objects to be owned

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             41             12            305
Markdown                         2              0              0              3
-------------------------------------------------------------------------------
SUM:                             5             41             12            308
-------------------------------------------------------------------------------
```
## TODO
Restrict actions to allow only user to perform his actions.
Perhaps attach `PlayerComponent` to camera.

## Dependencies
`core/game`:
  - `core/game.Actions`
  - `core/game.Economy`
  - `core/game.GameWorld`
  - `core/game.Player`

`core/modules/actions`:
  - `core/modules/actions.Context`
  - `core/modules/actions.Entity`
  - `core/modules/actions.IsEnemy`
  - `core/modules/actions.Own`
  - `core/modules/actions.PlayerContext`
  - `core/modules/actions.PlayerContextGetter`

`core/modules/economy`:
  - `core/modules/economy.NewWallet`
  - `core/modules/economy.Wallet`

`core/modules/lobby`:
  - `core/modules/lobby.NewInputAddPlayerEvent`
  - `core/modules/lobby.NewReloadLobbyEvent`

`core/modules/player`:
  - `core/modules/player.AcceptNewConnections`
  - `core/modules/player.ActingConnection`
  - `core/modules/player.ActingConnectionComponent`
  - `core/modules/player.ActingPlayer`
  - `core/modules/player.ActingPlayerComponent`
  - `core/modules/player.AssignActingPlayerDTO`
  - `core/modules/player.ColorComponent`
  - `core/modules/player.Connection`
  - `core/modules/player.GetContext`
  - `core/modules/player.Name`
  - `core/modules/player.NameComponent`
  - `core/modules/player.NewActingPlayer`
  - `core/modules/player.NewActiongConnection`
  - `core/modules/player.NewAssignActingPlayerDTO`
  - `core/modules/player.NewPlayerConnection`
  - `core/modules/player.NewPlayerUUID`
  - `core/modules/player.OwnerLink`
  - `core/modules/player.PlayerConnection`
  - `core/modules/player.PlayerConnectionComponent`
  - `core/modules/player.PlayerUUID`
  - `core/modules/player.PlayerUUIDComponent`
  - `core/modules/player.PlayersConnectionComponent`
  - `core/modules/player.Service`
  - `core/modules/player.UUID`

`engine/modules/connection`:
  - `engine/modules/connection.Conn`
  - `engine/modules/connection.ConnEntity`
  - `engine/modules/connection.Connection`
  - `engine/modules/connection.MsgCtx`
  - `engine/modules/connection.Send`

`engine/modules/connection/pkg`:
  - `engine/modules/connection/pkg.ParentPkg`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.World`

`engine/modules/feature/pkg`:
  - `engine/modules/feature/pkg.Config`
  - `engine/modules/feature/pkg.RegisterFieldProvider`

`engine/modules/loop`:
  - `engine/modules/loop.NewEmitOnFrameEvent`

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.AddGenericEventValidation`
  - `engine/modules/netsync/pkg.ComponentPkg`
  - `engine/modules/netsync/pkg.Config`

`engine/modules/relation`:
  - `engine/modules/relation.Service`

`engine/modules/relation/pkg`:
  - `engine/modules/relation/pkg.MapRelationPkg`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.Get`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.LinkService`
  - `engine/modules/uuid.LinkUUIDComponent`
  - `engine/modules/uuid.NewDeterministicUUID`
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.UUID`

`engine/modules/uuid/pkg`:
  - `engine/modules/uuid/pkg.LinkPkgT`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`