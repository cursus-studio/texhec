package playscene

import (
	"core/game"
	"core/modules/lobby"
	"core/modules/ui"
	"engine/modules/camera"
	"engine/modules/ecs"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/layout"
	"engine/modules/netsync"
	"engine/modules/text"
	"engine/modules/transform"
	"strings"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

const defaultConnectAddr = "localhost:8080"

// SetIPEvent is emitted when the connect button is pressed.
type SetIPEvent struct {
	Entity ecs.EntityID
}

var playConnectPkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) game.PlayConnectBuilder {
		s := ioc.Get[game.GameWorld](c)

		events.Listen(s.EventsBuilder(), func(event SetIPEvent) {
			input, ok := s.Inputs().TextInput().Get(event.Entity)
			if !ok {
				return
			}

			addr := strings.TrimSpace(input.Text())
			if addr == "" {
				return
			}

			connectionEntity := s.World().NewEntity()
			s.Hierarchy().SetParent(connectionEntity, s.Scene().Scene())
			s.NetSync().Server().Set(connectionEntity, netsync.ServerComponent{})
			if err := s.Connection().Connect(connectionEntity, addr); err != nil {
				s.Logger().Warn(err)
				s.World().RemoveEntity(connectionEntity)
			}
		})

		return func(sceneParent ecs.EntityID) {
			lobbyUI := s.World().NewEntity()
			s.Hierarchy().SetParent(lobbyUI, sceneParent)
			s.Lobby().LobbyUI().Set(lobbyUI, lobby.NewLobbyUI())

			s.Ui().AddBgCamera(lobbyUI)

			cameraEntity := s.World().NewEntity()
			s.Hierarchy().SetParent(cameraEntity, lobbyUI)
			s.Groups().Groups().Set(cameraEntity, groups.DefaultGroups())
			s.Camera().Ortho().Set(cameraEntity, camera.NewOrtho(-1000, 1000))
			s.Ui().CursorCamera().Set(cameraEntity, ui.CursorCameraComponent{})

			buttonArea := s.World().NewEntity()
			s.Hierarchy().SetParent(buttonArea, cameraEntity)
			s.Groups().InheritGroups(buttonArea)
			s.Transform().Pos().Set(buttonArea, transform.NewPos(0, 0, 1))
			s.Transform().Inherit().Set(buttonArea, transform.NewInherit(transform.RelativePos|transform.RelativeSizeX))

			s.Layout().Order().Set(buttonArea, layout.NewOrder(layout.OrderVectical))
			s.Layout().Align().Set(buttonArea, layout.NewAlign(.5, .5))
			s.Layout().Gap().Set(buttonArea, layout.NewGap(10))

			ipInput := s.Prototype().Clone(s.Definitions().Hud().Input)
			s.Hierarchy().SetParent(ipInput, buttonArea)
			s.Inputs().TextInput().Set(ipInput, inputs.TextInputComponent{
				Lines:     []string{defaultConnectAddr},
				CursorRow: 0,
				CursorCol: len(defaultConnectAddr),
			})

			connectBtn := s.Prototype().Clone(s.Definitions().Hud().Btn)
			s.Hierarchy().SetParent(connectBtn, buttonArea)
			s.Inputs().LeftClick().Set(connectBtn, inputs.NewLeftClick(SetIPEvent{Entity: ipInput}))
			s.Text().Content().Set(connectBtn, text.NewText("CONNECT"))
		}
	})
})
