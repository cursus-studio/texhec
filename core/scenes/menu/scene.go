package menuscene

import (
	"core/modules/ui"
	gamescenes "core/scenes"
	"engine/modules/camera"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/layout"
	"engine/modules/scene"
	"engine/modules/text"
	"engine/modules/transform"
	"engine/services/ecs"
	"strings"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) gamescenes.MenuBuilder {
		return func(sceneParent ecs.EntityID) {
			world := ioc.GetServices[gamescenes.World](c)
			cameraEntity := world.NewEntity()
			world.Hierarchy.SetParent(cameraEntity, sceneParent)
			world.Groups.Component().Set(cameraEntity, groups.DefaultGroups())
			world.Camera.Ortho().Set(cameraEntity, camera.NewOrtho(-1000, 1000))
			world.Ui.CursorCamera().Set(cameraEntity, ui.CursorCameraComponent{})

			signature := world.NewEntity()
			world.Hierarchy.SetParent(signature, cameraEntity)
			world.Transform.Pos().Set(signature, transform.NewPos(5, 5, 0))
			world.Transform.Size().Set(signature, transform.NewSize(100, 50, 1))
			world.Transform.PivotPoint().Set(signature, transform.NewPivotPoint(0, .5, .5))
			world.Transform.Parent().Set(signature, transform.NewParent(transform.RelativePos))
			world.Transform.ParentPivotPoint().Set(signature, transform.NewParentPivotPoint(0, 0, .5))

			world.Text.Content().Set(signature, text.TextComponent{Text: "menu"})
			world.Text.FontSize().Set(signature, text.FontSizeComponent{FontSize: 32})
			world.Text.Break().Set(signature, text.BreakComponent{Break: text.BreakNone})

			background := world.NewEntity()
			world.Hierarchy.SetParent(background, cameraEntity)
			world.Transform.Pos().Set(background, transform.NewPos(0, 0, 1))
			world.Transform.PivotPoint().Set(background, transform.NewPivotPoint(.5, .5, 0))
			world.Transform.ParentPivotPoint().Set(background, transform.NewParentPivotPoint(.5, .5, 0))
			world.Ui.AnimatedBackground().Set(background, ui.AnimatedBackgroundComponent{})

			buttonArea := world.NewEntity()
			world.Hierarchy.SetParent(buttonArea, cameraEntity)
			world.Groups.Inherit().Set(buttonArea, groups.InheritGroupsComponent{})
			world.Transform.Pos().Set(buttonArea, transform.NewPos(0, 0, 1))
			world.Transform.Parent().Set(buttonArea, transform.NewParent(transform.RelativePos|transform.RelativeSizeX))

			world.Layout.Order().Set(buttonArea, layout.NewOrder(layout.OrderVectical))
			world.Layout.Align().Set(buttonArea, layout.NewAlign(.5, .5))
			world.Layout.Gap().Set(buttonArea, layout.NewGap(10))

			type Button struct {
				Text    string
				OnClick any
			}
			buttons := []Button{
				{"play", scene.NewChangeSceneEvent(gamescenes.GameID)},
				{"settings", scene.NewChangeSceneEvent(gamescenes.SettingsID)},
				{"credits", scene.NewChangeSceneEvent(gamescenes.CreditsID)},
				{"exit", inputs.QuitEvent{}},
			}

			for _, button := range buttons {
				btn := world.Prototype.Clone(world.Definitions.Hud.Btn)

				world.Hierarchy.SetParent(btn, buttonArea)
				world.Inputs.LeftClick().Set(btn, inputs.NewLeftClick(button.OnClick))
				world.Text.Content().Set(btn, text.TextComponent{Text: strings.ToUpper(button.Text)})
			}
		}
	})
})
