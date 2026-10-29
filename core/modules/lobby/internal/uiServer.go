package internal

import (
	"core/modules/economy"
	"core/modules/lobby"
	"core/modules/player"
	"core/modules/ui"
	"engine/modules/camera"
	"engine/modules/ecs"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/layout"
	"engine/modules/netsync"
	"engine/modules/seed"
	"engine/modules/text"
	"engine/modules/transform"
	"strings"

	"github.com/ogiusek/events"
)

func (s *service) setupServerLobbyVariables() {
	if len(s.NetSync().Clients().GetEntities()) != 0 {
		return
	}

	hostEntity := s.World().NewEntity()
	s.Hierarchy().SetParent(hostEntity, s.Scene().Scene())
	s.Player().PlayersConnection().Set(hostEntity, player.NewPlayersConnection(true))
	s.NetSync().Clients().Set(hostEntity, netsync.ClientsComponent{})
	if err := s.Connection().Host(hostEntity, ":8080"); err != nil {
		s.Logger().Warn(err)
		return
	}
	s.Seed().Seed().Set(hostEntity, seed.NewSeed(seed.New(21377137))) // seed.NewSeed(seed.New(s.Clock().Now().UnixNano()))
	s.PlayersInitials().Set(hostEntity, s.NewPlayersInitials(economy.NewWallet(10), []ecs.EntityID{
		s.Definitions().Objects().Farm,
		s.Definitions().Objects().Tank,
	}))
	s.GameCamera().Set(hostEntity, lobby.NewGameCamera(2))

	events.Emit(s.Events(), lobby.NewInputAddPlayerEvent())
	s.Player().ActingPlayer().Set(s.Player().Name().GetEntities()[0], player.NewActingPlayer())
}

func (s *service) renderServerLobby() {
	cameraEntity := s.World().NewEntity()
	s.Hierarchy().SetParent(cameraEntity, s.Scene().Scene())
	s.LobbyUI().Set(cameraEntity, lobby.NewLobbyUI())
	s.Groups().Component().Set(cameraEntity, groups.DefaultGroups())
	s.Camera().Ortho().Set(cameraEntity, camera.NewOrtho(-1000, 1000))
	s.Ui().CursorCamera().Set(cameraEntity, ui.CursorCameraComponent{})

	s.Ui().AddBgCamera(cameraEntity)

	buttonArea := s.World().NewEntity()
	s.Hierarchy().SetParent(buttonArea, cameraEntity)
	s.Groups().InheritGroups(buttonArea)
	s.Transform().Pos().Set(buttonArea, transform.NewPos(0, 0, 1))
	s.Transform().Inherit().Set(buttonArea, transform.NewInherit(transform.RelativePos|transform.RelativeSizeX))

	s.Layout().Order().Set(buttonArea, layout.NewOrder(layout.OrderVectical))
	s.Layout().Align().Set(buttonArea, layout.NewAlign(.5, .5))
	s.Layout().Gap().Set(buttonArea, layout.NewGap(10))

	type Button struct {
		Text    string
		OnClick any
	}

	// TODO
	// render:
	// - world config(seed, initial wallet and units, max camera zoom out)
	// - player(name, color, controlled by?)
	// - connections(connection->player)

	// type RemovePlayerEvent struct{ PlayerUUID uuid.UUID }
	// type SetPlayerNameEvent struct { PlayerUUID uuid.UUID, Name       string }
	// type SetPlayerColorEvent struct { PlayerUUID uuid.UUID, Color      mgl32.Vec3 }
	//
	// type SetPlayersWalletEvent struct{ Wallet economy.WalletComponent }
	// type AddPlayersUnitEvent struct{ UUID uuid.UUID }
	// type RemovePlayersUnitEvent struct{ UUID uuid.UUID }
	//
	// type SetConnectionPlayerEvent struct{ Player, Connection uuid.UUID }

	buttons := []Button{}
	if _, comp, _ := s.Player().PlayersConnection().GetSingleton(); comp.AcceptNewConnections {
		buttons = append(buttons, Button{"open", lobby.NewSetPlayersConnectionEvent(false)})
	} else {
		buttons = append(buttons, Button{"closed", lobby.NewSetPlayersConnectionEvent(true)})
	}
	buttons = append(buttons,
		// add here add AI player btn
		Button{"re-roll seed", lobby.NewReRollSeed()},
		Button{"paste seed", lobby.NewPasteSeed()},
		Button{"generate", lobby.NewGenerateEvent()})

	for _, button := range buttons {
		btn := s.Prototype().Clone(s.Definitions().Hud().Btn)

		s.Hierarchy().SetParent(btn, buttonArea)
		s.Inputs().LeftClick().Set(btn, inputs.NewLeftClick(button.OnClick))
		s.Text().Content().Set(btn, text.NewText(strings.ToUpper(button.Text)))
	}
	s.renderPlayerCards(buttonArea)
}
