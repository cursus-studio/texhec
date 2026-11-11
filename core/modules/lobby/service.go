// we define how lobby looks.
// it should allow inviting multiple players
package lobby

import (
	"core/modules/economy"
	"core/modules/ui"
	"engine/modules/ecs"
	"engine/modules/seed"
	"engine/modules/uuid"

	"github.com/go-gl/mathgl/mgl32"
)

type PlayersInitialsComponent struct {
	Wallet economy.WalletComponent
	Units  []uuid.UUID
}
type GameCameraComponent ui.GameCameraComponent
type LobbyComponent struct{}
type LobbyUIComponent struct{}

func NewPlayersInitials(wallet economy.WalletComponent, units []uuid.UUID) PlayersInitialsComponent {
	return PlayersInitialsComponent{wallet, units}
}
func NewGameCamera(maxChunksZoomOut ui.MaxChunksZoomOut) GameCameraComponent {
	return GameCameraComponent(ui.NewGameCamera(maxChunksZoomOut))
}
func NewLobby() LobbyComponent {
	return LobbyComponent{}
}
func NewLobbyUI() LobbyUIComponent {
	return LobbyUIComponent{}
}

//

type ReRollSeedEvent struct{}              // isn't deterministic
type PasteSeedEvent struct{}               // isn't deterministic
type SetSeedEvent struct{ Seed seed.Seed } // is [InputReRollSeedEvent] but deterministic

func NewReRollSeed() ReRollSeedEvent         { return ReRollSeedEvent{} }
func NewPasteSeed() PasteSeedEvent           { return PasteSeedEvent{} }
func NewSetSeed(seed seed.Seed) SetSeedEvent { return SetSeedEvent{seed} }

type ReloadLobbyEvent struct{}
type InputAddPlayerEvent struct{}                     // isn't deterministic
type AddPlayerEvent struct{ NewPlayerUUID uuid.UUID } // is [InputAddPlayerEvent] but deterministic
type RemovePlayerEvent struct{ PlayerUUID uuid.UUID }
type SetPlayerNameEvent struct {
	PlayerUUID uuid.UUID
	Name       string
}
type SetPlayerColorEvent struct {
	PlayerUUID uuid.UUID
	Color      mgl32.Vec3
}

func NewReloadLobbyEvent() ReloadLobbyEvent                 { return ReloadLobbyEvent{} }
func NewInputAddPlayerEvent() InputAddPlayerEvent           { return InputAddPlayerEvent{} }
func NewAddPlayerEvent(uuid uuid.UUID) AddPlayerEvent       { return AddPlayerEvent{uuid} }
func NewRemovePlayerEvent(uuid uuid.UUID) RemovePlayerEvent { return RemovePlayerEvent{uuid} }
func NewSetPlayerNameEvent(uuid uuid.UUID, name string) SetPlayerNameEvent {
	return SetPlayerNameEvent{uuid, name}
}
func NewSetPlayerColorEvent(uuid uuid.UUID, color mgl32.Vec3) SetPlayerColorEvent {
	return SetPlayerColorEvent{uuid, color}
}

type SetPlayersWalletEvent struct{ Wallet economy.WalletComponent }
type AddPlayersUnitEvent struct{ UUID uuid.UUID }
type RemovePlayersUnitEvent struct{ UUID uuid.UUID }

func NewSetPlayersWalletEvent(wallet economy.WalletComponent) SetPlayersWalletEvent {
	return SetPlayersWalletEvent{wallet}
}
func NewAddPlayersUnitEvent(uuid uuid.UUID) AddPlayersUnitEvent { return AddPlayersUnitEvent{uuid} }
func NewRemovePlayersUnitEvent(uuid uuid.UUID) RemovePlayersUnitEvent {
	return RemovePlayersUnitEvent{uuid}
}

type SetConnectionPlayerEvent struct{ Player, Connection uuid.UUID }

func NewSetConnectionPlayerEvent(player, connection uuid.UUID) SetConnectionPlayerEvent {
	return SetConnectionPlayerEvent{player, connection}
}

type SetPlayersConnectionEvent struct{ AcceptNewConnections bool }

func NewSetPlayersConnectionEvent(acceptNewConnection bool) SetPlayersConnectionEvent {
	return SetPlayersConnectionEvent{acceptNewConnection}
}

//

type GenerateEvent struct{}

func NewGenerateEvent() GenerateEvent { return GenerateEvent{} }

//

type Service interface {
	PlayersInitials() ecs.ComponentArray[PlayersInitialsComponent]
	GameCamera() ecs.ComponentArray[GameCameraComponent]
	Lobby() ecs.ComponentArray[LobbyComponent]
	LobbyUI() ecs.ComponentArray[LobbyUIComponent]

	NewPlayersInitials(wallet economy.WalletComponent, units []ecs.EntityID) PlayersInitialsComponent
}
