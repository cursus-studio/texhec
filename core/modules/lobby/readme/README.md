# lobby
## Architecture
we define how lobby looks.
it should allow inviting multiple players

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               7            138             48            732
-------------------------------------------------------------------------------
SUM:                             7            138             48            732
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/lobby.Service`

#### method Service GameCamera
Type: `func() engine/modules/ecs.ComponentArray[core/modules/lobby.GameCameraComponent]`

#### method Service Lobby
Type: `func() engine/modules/ecs.ComponentArray[core/modules/lobby.LobbyComponent]`

#### method Service LobbyUI
Type: `func() engine/modules/ecs.ComponentArray[core/modules/lobby.LobbyUIComponent]`

#### method Service NewPlayersInitials
Type: `func(wallet core/modules/economy.WalletComponent, units []engine/modules/ecs.EntityID) core/modules/lobby.PlayersInitialsComponent`

#### method Service PlayersInitials
Type: `func() engine/modules/ecs.ComponentArray[core/modules/lobby.PlayersInitialsComponent]`

### type PlayersInitialsComponent
Type: `core/modules/lobby.PlayersInitialsComponent`

#### property PlayersInitialsComponent Wallet
Type: `core/modules/economy.WalletComponent`

#### property PlayersInitialsComponent Units
Type: `[]engine/modules/uuid.UUID`

### type GameCameraComponent
Type: `core/modules/lobby.GameCameraComponent`

#### property GameCameraComponent MaxChunksZoomOut
Type: `core/modules/ui.MaxChunksZoomOut`

### type LobbyComponent
Type: `core/modules/lobby.LobbyComponent`

### type LobbyUIComponent
Type: `core/modules/lobby.LobbyUIComponent`

### type ReRollSeedEvent
Type: `core/modules/lobby.ReRollSeedEvent`

### type PasteSeedEvent
Type: `core/modules/lobby.PasteSeedEvent`

### type SetSeedEvent
Type: `core/modules/lobby.SetSeedEvent`

#### property SetSeedEvent Seed
Type: `engine/modules/seed.Seed`

### type InputAddPlayerEvent
Type: `core/modules/lobby.InputAddPlayerEvent`

### type AddPlayerEvent
Type: `core/modules/lobby.AddPlayerEvent`

#### property AddPlayerEvent NewPlayerUUID
Type: `engine/modules/uuid.UUID`

### type RemovePlayerEvent
Type: `core/modules/lobby.RemovePlayerEvent`

#### property RemovePlayerEvent PlayerUUID
Type: `engine/modules/uuid.UUID`

### type SetPlayerNameEvent
Type: `core/modules/lobby.SetPlayerNameEvent`

#### property SetPlayerNameEvent PlayerUUID
Type: `engine/modules/uuid.UUID`

#### property SetPlayerNameEvent Name
Type: `string`

### type SetPlayerColorEvent
Type: `core/modules/lobby.SetPlayerColorEvent`

#### property SetPlayerColorEvent PlayerUUID
Type: `engine/modules/uuid.UUID`

#### property SetPlayerColorEvent Color
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

### type SetPlayersWalletEvent
Type: `core/modules/lobby.SetPlayersWalletEvent`

#### property SetPlayersWalletEvent Wallet
Type: `core/modules/economy.WalletComponent`

### type AddPlayersUnitEvent
Type: `core/modules/lobby.AddPlayersUnitEvent`

#### property AddPlayersUnitEvent UUID
Type: `engine/modules/uuid.UUID`

### type RemovePlayersUnitEvent
Type: `core/modules/lobby.RemovePlayersUnitEvent`

#### property RemovePlayersUnitEvent UUID
Type: `engine/modules/uuid.UUID`

### type SetConnectionPlayerEvent
Type: `core/modules/lobby.SetConnectionPlayerEvent`

#### property SetConnectionPlayerEvent Player
Type: `engine/modules/uuid.UUID`

#### property SetConnectionPlayerEvent Connection
Type: `engine/modules/uuid.UUID`

### type SetPlayersConnectionEvent
Type: `core/modules/lobby.SetPlayersConnectionEvent`

#### property SetPlayersConnectionEvent AcceptNewConnections
Type: `bool`

### type GenerateEvent
Type: `core/modules/lobby.GenerateEvent`

## Functions
### func NewPlayersInitials
Type: `func(wallet core/modules/economy.WalletComponent, units []engine/modules/uuid.UUID) core/modules/lobby.PlayersInitialsComponent`

### func NewGameCamera
Type: `func(maxChunksZoomOut core/modules/ui.MaxChunksZoomOut) core/modules/lobby.GameCameraComponent`

### func NewLobby
Type: `func() core/modules/lobby.LobbyComponent`

### func NewLobbyUI
Type: `func() core/modules/lobby.LobbyUIComponent`

### func NewReRollSeed
Type: `func() core/modules/lobby.ReRollSeedEvent`

### func NewPasteSeed
Type: `func() core/modules/lobby.PasteSeedEvent`

### func NewSetSeed
Type: `func(seed engine/modules/seed.Seed) core/modules/lobby.SetSeedEvent`

### func NewInputAddPlayerEvent
Type: `func() core/modules/lobby.InputAddPlayerEvent`

### func NewAddPlayerEvent
Type: `func(uuid engine/modules/uuid.UUID) core/modules/lobby.AddPlayerEvent`

### func NewRemovePlayerEvent
Type: `func(uuid engine/modules/uuid.UUID) core/modules/lobby.RemovePlayerEvent`

### func NewSetPlayerNameEvent
Type: `func(uuid engine/modules/uuid.UUID, name string) core/modules/lobby.SetPlayerNameEvent`

### func NewSetPlayerColorEvent
Type: `func(uuid engine/modules/uuid.UUID, color github.com/go-gl/mathgl/mgl32.Vec3) core/modules/lobby.SetPlayerColorEvent`

### func NewSetPlayersWalletEvent
Type: `func(wallet core/modules/economy.WalletComponent) core/modules/lobby.SetPlayersWalletEvent`

### func NewAddPlayersUnitEvent
Type: `func(uuid engine/modules/uuid.UUID) core/modules/lobby.AddPlayersUnitEvent`

### func NewRemovePlayersUnitEvent
Type: `func(uuid engine/modules/uuid.UUID) core/modules/lobby.RemovePlayersUnitEvent`

### func NewSetConnectionPlayerEvent
Type: `func(player engine/modules/uuid.UUID, connection engine/modules/uuid.UUID) core/modules/lobby.SetConnectionPlayerEvent`

### func NewSetPlayersConnectionEvent
Type: `func(acceptNewConnection bool) core/modules/lobby.SetPlayersConnectionEvent`

### func NewGenerateEvent
Type: `func() core/modules/lobby.GenerateEvent`


## Dependencies
`core/game`:
  - `core/game.Definitions`
  - `core/game.Deploy`
  - `core/game.Economy`
  - `core/game.GameWorld`
  - `core/game.Player`
  - `core/game.Ui`

`core/modules/economy`:
  - `core/modules/economy.NewWallet`
  - `core/modules/economy.Wallet`
  - `core/modules/economy.WalletComponent`

`core/modules/lobby`:
  - `core/modules/lobby.AcceptNewConnections`
  - `core/modules/lobby.AddPlayerEvent`
  - `core/modules/lobby.AddPlayersUnitEvent`
  - `core/modules/lobby.Color`
  - `core/modules/lobby.Connection`
  - `core/modules/lobby.GameCameraComponent`
  - `core/modules/lobby.GenerateEvent`
  - `core/modules/lobby.InputAddPlayerEvent`
  - `core/modules/lobby.LobbyComponent`
  - `core/modules/lobby.LobbyUIComponent`
  - `core/modules/lobby.Name`
  - `core/modules/lobby.NewAddPlayerEvent`
  - `core/modules/lobby.NewGameCamera`
  - `core/modules/lobby.NewGenerateEvent`
  - `core/modules/lobby.NewInputAddPlayerEvent`
  - `core/modules/lobby.NewLobbyUI`
  - `core/modules/lobby.NewPasteSeed`
  - `core/modules/lobby.NewPlayerUUID`
  - `core/modules/lobby.NewPlayersInitials`
  - `core/modules/lobby.NewReRollSeed`
  - `core/modules/lobby.NewRemovePlayerEvent`
  - `core/modules/lobby.NewSetPlayerColorEvent`
  - `core/modules/lobby.NewSetPlayerNameEvent`
  - `core/modules/lobby.NewSetPlayersConnectionEvent`
  - `core/modules/lobby.NewSetSeed`
  - `core/modules/lobby.PasteSeedEvent`
  - `core/modules/lobby.Player`
  - `core/modules/lobby.PlayerUUID`
  - `core/modules/lobby.PlayersInitialsComponent`
  - `core/modules/lobby.ReRollSeedEvent`
  - `core/modules/lobby.RemovePlayerEvent`
  - `core/modules/lobby.RemovePlayersUnitEvent`
  - `core/modules/lobby.Seed`
  - `core/modules/lobby.Service`
  - `core/modules/lobby.SetConnectionPlayerEvent`
  - `core/modules/lobby.SetPlayerColorEvent`
  - `core/modules/lobby.SetPlayerNameEvent`
  - `core/modules/lobby.SetPlayersConnectionEvent`
  - `core/modules/lobby.SetPlayersWalletEvent`
  - `core/modules/lobby.SetSeedEvent`
  - `core/modules/lobby.UUID`
  - `core/modules/lobby.Units`
  - `core/modules/lobby.Wallet`

`core/modules/player`:
  - `core/modules/player.AcceptNewConnections`
  - `core/modules/player.ActingConnection`
  - `core/modules/player.ActingPlayer`
  - `core/modules/player.Color`
  - `core/modules/player.ErrNameCollision`
  - `core/modules/player.Name`
  - `core/modules/player.NameRelation`
  - `core/modules/player.NewActingPlayer`
  - `core/modules/player.NewActiongConnection`
  - `core/modules/player.NewColor`
  - `core/modules/player.NewName`
  - `core/modules/player.NewPlayersConnection`
  - `core/modules/player.PlayersConnection`

`core/modules/tile`:
  - `core/modules/tile.NewMissingChunkEvent`

`core/modules/ui`:
  - `core/modules/ui.AddBgCamera`
  - `core/modules/ui.AddUICamera`
  - `core/modules/ui.CursorCamera`
  - `core/modules/ui.CursorCameraComponent`
  - `core/modules/ui.GameCamera`
  - `core/modules/ui.GameCameraComponent`
  - `core/modules/ui.MaxChunksZoomOut`
  - `core/modules/ui.NewGameCamera`

`engine/modules/batcher`:
  - `engine/modules/batcher.AddOrderedBatch`
  - `engine/modules/batcher.Build`
  - `engine/modules/batcher.NewBatch`
  - `engine/modules/batcher.NewTask`
  - `engine/modules/batcher.Perform`
  - `engine/modules/batcher.Queue`
  - `engine/modules/batcher.Tasks`

`engine/modules/camera`:
  - `engine/modules/camera.NewOrtho`
  - `engine/modules/camera.Ortho`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/grid`:
  - `engine/modules/grid.AbsoluteCoords`
  - `engine/modules/grid.ChunkSize`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.NewChunkCoords`
  - `engine/modules/grid.NewCoords`

`engine/modules/groups`:
  - `engine/modules/groups.Component`
  - `engine/modules/groups.DefaultGroups`
  - `engine/modules/groups.InheritGroups`

`engine/modules/inputs`:
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewLeftClick`
  - `engine/modules/inputs.Text`
  - `engine/modules/inputs.TextInput`
  - `engine/modules/inputs.TextInputComponent`

`engine/modules/layout`:
  - `engine/modules/layout.Align`
  - `engine/modules/layout.Gap`
  - `engine/modules/layout.NewAlign`
  - `engine/modules/layout.NewGap`
  - `engine/modules/layout.NewOrder`
  - `engine/modules/layout.Order`
  - `engine/modules/layout.OrderHorizontal`
  - `engine/modules/layout.OrderVectical`

`engine/modules/netsync`:
  - `engine/modules/netsync.Clients`
  - `engine/modules/netsync.ClientsComponent`
  - `engine/modules/netsync.Server`

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.ComponentPkg`
  - `engine/modules/netsync/pkg.EventPkg`

`engine/modules/render`:
  - `engine/modules/render.Color`
  - `engine/modules/render.Mesh`
  - `engine/modules/render.NewColor`
  - `engine/modules/render.NewMesh`
  - `engine/modules/render.NewTexture`
  - `engine/modules/render.Texture`

`engine/modules/seed`:
  - `engine/modules/seed.ErrWorldShouldHaveOneSeed`
  - `engine/modules/seed.GenerateSeed`
  - `engine/modules/seed.New`
  - `engine/modules/seed.NewSeed`
  - `engine/modules/seed.Seed`

`engine/modules/text`:
  - `engine/modules/text.Align`
  - `engine/modules/text.Color`
  - `engine/modules/text.Content`
  - `engine/modules/text.FontSize`
  - `engine/modules/text.NewAlign`
  - `engine/modules/text.NewColor`
  - `engine/modules/text.NewFontSize`
  - `engine/modules/text.NewText`

`engine/modules/transform`:
  - `engine/modules/transform.AspectRatio`
  - `engine/modules/transform.Inherit`
  - `engine/modules/transform.MaxSize`
  - `engine/modules/transform.MinSize`
  - `engine/modules/transform.NewAspectRatio`
  - `engine/modules/transform.NewInherit`
  - `engine/modules/transform.NewMaxSize`
  - `engine/modules/transform.NewMinSize`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.NewSize`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.RelativePos`
  - `engine/modules/transform.RelativeSizeX`
  - `engine/modules/transform.Size`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.New`
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`