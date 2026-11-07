package internal

import (
	"core/modules/lobby"
	"core/modules/ui"
	"engine/modules/camera"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/layout"
	"engine/modules/text"
	"engine/modules/transform"
	"strings"
)

func (s *service) setupClientLobbyVariables() {
}

func (s *service) renderClientLobby() {
	lobbyUi := s.World().NewEntity()
	s.Hierarchy().SetParent(lobbyUi, s.Scene().Scene())
	s.LobbyUI().Set(lobbyUi, lobby.NewLobbyUI())

	cameraEntity := s.World().NewEntity()
	s.Hierarchy().SetParent(cameraEntity, lobbyUi)
	s.Groups().Groups().Set(cameraEntity, groups.DefaultGroups())
	s.Camera().Ortho().Set(cameraEntity, camera.NewOrtho(-1000, 1000))
	s.Ui().CursorCamera().Set(cameraEntity, ui.CursorCameraComponent{})

	s.Ui().AddBgCamera(lobbyUi)

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

	buttons := []Button{
		{"awaiting for the host...", nil},
	}

	for _, button := range buttons {
		btn := s.Prototype().Clone(s.Definitions().Hud().Btn)

		s.Hierarchy().SetParent(btn, buttonArea)
		s.Inputs().LeftClick().Set(btn, inputs.NewLeftClick(button.OnClick))
		s.Text().Content().Set(btn, text.NewText(strings.ToUpper(button.Text)))
	}
}
