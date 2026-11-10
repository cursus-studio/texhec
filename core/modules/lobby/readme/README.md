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
Go                               7            141             52            751
-------------------------------------------------------------------------------
SUM:                             7            141             52            751
-------------------------------------------------------------------------------
```
## Dependencies
`core/game`:
  - `core/game.Definitions`
  - `core/game.Deploy`
  - `core/game.Economy`
  - `core/game.GameWorld`
  - `core/game.Player`
  - `core/game.Tile`
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
  - `core/modules/tile.GetNthNeighbor`
  - `core/modules/tile.GetPos`
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
  - `engine/modules/grid.ChunkSize`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.NewChunkCoords`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/modules/groups`:
  - `engine/modules/groups.DefaultGroups`
  - `engine/modules/groups.Groups`
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
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.UUID`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`