package creditsscene

import (
	"core/modules/ui"
	gamescenes "core/scenes"
	"engine/modules/camera"
	"engine/modules/collider"
	"engine/modules/drag"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/render"
	"engine/modules/scene"
	"engine/modules/text"
	"engine/modules/transform"
	"engine/services/ecs"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) gamescenes.CreditsBuilder {
		return func(sceneParent ecs.EntityID) {
			world := ioc.GetServices[gamescenes.GameWorld](c)
			cameraEntity := world.World().NewEntity()
			world.Hierarchy().SetParent(cameraEntity, sceneParent)
			world.Groups().Component().Set(cameraEntity, groups.DefaultGroups())
			world.Camera().Ortho().Set(cameraEntity, camera.NewOrtho(-1000, +1000))
			world.Ui().CursorCamera().Set(cameraEntity, ui.CursorCameraComponent{})

			signature := world.World().NewEntity()
			world.Hierarchy().SetParent(signature, cameraEntity)
			world.Transform().Pos().Set(signature, transform.NewPos(5, 5, 0))
			world.Transform().Size().Set(signature, transform.NewSize(100, 50, 1))
			world.Transform().PivotPoint().Set(signature, transform.NewPivotPoint(0, .5, .5))
			world.Transform().Parent().Set(signature, transform.NewParent(transform.RelativePos))
			world.Transform().ParentPivotPoint().Set(signature, transform.NewParentPivotPoint(0, 0, .5))

			world.Text().Content().Set(signature, text.NewText("credits"))
			world.Text().FontSize().Set(signature, text.NewFontSize(32))
			world.Text().Break().Set(signature, text.NewBreak(text.BreakNone))

			background := world.World().NewEntity()
			world.Hierarchy().SetParent(background, cameraEntity)
			world.Ui().AnimatedBackground().Set(background, ui.AnimatedBackgroundComponent{})

			buttonArea := world.World().NewEntity()
			world.Hierarchy().SetParent(buttonArea, cameraEntity)
			world.Groups().Inherit().Set(buttonArea, groups.InheritGroupsComponent{})
			world.Transform().Pos().Set(buttonArea, transform.NewPos(0, 0, 1))
			world.Transform().Size().Set(buttonArea, transform.NewSize(500, 200, 1))
			world.Transform().Parent().Set(buttonArea, transform.NewParent(transform.RelativePos))

			draggable := world.World().NewEntity()
			world.Hierarchy().SetParent(draggable, cameraEntity)
			world.Transform().Pos().Set(draggable, transform.NewPos(0, 0, 2))
			world.Transform().Size().Set(draggable, transform.NewSize(50, 50, 1))
			world.Render().Color().Set(draggable, render.NewColor(mgl32.Vec4{0, 1, 0, 1}))
			world.Render().Mesh().Set(draggable, render.NewMesh(world.Definitions().SquareMesh))
			world.Render().Texture().Set(draggable, render.NewTexture(world.Definitions().Hud().Cursor))

			world.Collider().Component().Set(draggable, collider.NewCollider(world.Definitions().SquareCollider))
			world.Inputs().Drag().Set(draggable, inputs.NewDragComponent(drag.NewDraggable(draggable)))

			world.Text().Content().Set(draggable, text.NewText(strings.ToUpper("drag me")))
			world.Text().Align().Set(draggable, text.NewAlign(.5, .5))
			world.Text().FontSize().Set(draggable, text.NewFontSize(15))
			world.Text().Color().Set(draggable, text.NewColor(mgl32.Vec4{1, 0, 0, 1}))

			btn := world.Prototype().Clone(world.Definitions().Hud().Btn)
			world.Hierarchy().SetParent(btn, buttonArea)
			world.Transform().Size().Set(btn, transform.NewSize(500, 100, 1))
			world.Transform().ParentPivotPoint().Set(btn, transform.NewParentPivotPoint(.5, 0, .5))

			world.Inputs().LeftClick().Set(btn, inputs.NewLeftClick(scene.NewChangeSceneEvent(gamescenes.MenuID)))
			world.Text().Content().Set(btn, text.NewText(strings.ToUpper("return to menu")))
		}
	})
})
