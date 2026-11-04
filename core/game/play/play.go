package playscene

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/ui"
	"engine/modules/camera"
	"engine/modules/ecs"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/layout"
	"engine/modules/scene"
	"engine/modules/text"
	"engine/modules/transform"
	"strings"

	"github.com/ogiusek/ioc/v2"
)

var playPkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) game.PlayBuilder {
		world := ioc.Get[game.GameWorld](c)
		return func(sceneParent ecs.EntityID) {
			world.Ui().AddBgCamera(sceneParent)

			cameraEntity := world.World().NewEntity()
			world.Hierarchy().SetParent(cameraEntity, sceneParent)
			world.Groups().Groups().Set(cameraEntity, groups.DefaultGroups())
			world.Camera().Ortho().Set(cameraEntity, camera.NewOrtho(-1000, 1000))
			world.Ui().CursorCamera().Set(cameraEntity, ui.CursorCameraComponent{})

			signature := world.World().NewEntity()
			world.Hierarchy().SetParent(signature, cameraEntity)
			world.Transform().Pos().Set(signature, transform.NewPos(5, 5, 0))
			world.Transform().Size().Set(signature, transform.NewSize(100, 50, 1))
			world.Transform().PivotPoint().Set(signature, transform.NewPivotPoint(0, .5, .5))
			world.Transform().Inherit().Set(signature, transform.NewInherit(transform.RelativePos))
			world.Transform().ParentPivotPoint().Set(signature, transform.NewParentPivotPoint(0, 0, .5))

			world.Text().Content().Set(signature, text.NewText("play"))
			world.Text().FontSize().Set(signature, text.NewFontSize(32))
			world.Text().Break().Set(signature, text.NewBreak(text.BreakNone))

			buttonArea := world.World().NewEntity()
			world.Hierarchy().SetParent(buttonArea, cameraEntity)
			world.Groups().InheritGroups(buttonArea)
			world.Transform().Pos().Set(buttonArea, transform.NewPos(0, 0, 1))
			world.Transform().Inherit().Set(buttonArea, transform.NewInherit(transform.RelativePos|transform.RelativeSizeX))

			world.Layout().Order().Set(buttonArea, layout.NewOrder(layout.OrderVectical))
			world.Layout().Align().Set(buttonArea, layout.NewAlign(.5, .5))
			world.Layout().Gap().Set(buttonArea, layout.NewGap(10))

			type Button struct {
				Text    string
				OnClick any
			}
			buttons := []Button{
				{"connect", scene.NewChangeSceneEvent(definitions.PlayConnectID)},
				{"generate", scene.NewChangeSceneEvent(definitions.PlayGenerateID)},
				{"main menu", scene.NewChangeSceneEvent(definitions.MainMenuID)},
			}

			for _, button := range buttons {
				btn := world.Prototype().Clone(world.Definitions().Hud().Btn)

				world.Hierarchy().SetParent(btn, buttonArea)
				world.Inputs().LeftClick().Set(btn, inputs.NewLeftClick(button.OnClick))
				world.Text().Content().Set(btn, text.NewText(strings.ToUpper(button.Text)))
			}
		}
	})
})
