package uiservice

import (
	"core/modules/ui"
	"engine/modules/collider"
	"engine/modules/ecs"
	"engine/modules/inputs"
	"engine/modules/layout"
	"engine/modules/render"
	"engine/modules/text"
	"engine/modules/transform"

	"github.com/go-gl/mathgl/mgl32"
)

func (s *service) EnsureExists() {
mainLoop:
	for _, camera := range s.uiCameraArray.GetEntities() {
		// objects
		// menu
		for _, child := range s.Hierarchy().Children(camera).GetIndices() {
			if _, ok := s.menuArray.Get(child); ok {
				continue mainLoop
			}
		}
		menu := s.World().NewEntity()
		s.Hierarchy().SetParent(menu, camera)
		s.Transform().ParentPivotPoint().Set(menu, transform.NewParentPivotPoint(1, 1, .5))
		s.Transform().Pos().Set(menu, transform.NewPos(0, 0, 1))
		s.Transform().Size().Set(menu, transform.NewSize(.2, 1, 1))
		s.Transform().PivotPoint().Set(menu, transform.NewPivotPoint(0, 1, .5))

		s.Render().Color().Set(menu, render.NewColor(mgl32.Vec4{1, 1, 1, .5}))
		s.AnimatedBackground().Set(menu, ui.AnimatedBackgroundComponent{})

		s.Groups().InheritGroups(menu)
		s.Collider().Component().Set(menu, collider.NewCollider(s.Definitions().Assets().SquareCollider))
		s.Inputs().KeepSelected().Set(menu, inputs.KeepSelectedComponent{})
		s.menuArray.Set(menu, menuComponent{})

		// quit btn
		quit := s.World().NewEntity()

		s.Hierarchy().SetParent(quit, menu)
		s.Groups().InheritGroups(quit)

		s.Transform().Pos().Set(quit, transform.NewPos(0, 0, 1))
		s.Transform().Parent().Set(quit, transform.NewParent(transform.RelativePos))
		s.Transform().ParentPivotPoint().Set(quit, transform.NewParentPivotPoint(1, 1, .5))
		s.Transform().Size().Set(quit, transform.NewSize(25, 25, 1))
		s.Transform().PivotPoint().Set(quit, transform.NewPivotPoint(1, 1, .5))

		s.Text().Content().Set(quit, text.NewText("X"))
		s.Text().FontSize().Set(quit, text.NewFontSize(25))
		s.Text().Align().Set(quit, text.NewAlign(.5, .5))

		s.Render().Color().Set(quit, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
		s.Render().Mesh().Set(quit, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(quit, render.NewTexture(s.Definitions().Assets().Blank))

		removeEntityEvent := ecs.NewRemoveEntityEvent(s.Interactions().FeatureEntity())
		s.Inputs().LeftClick().Set(quit, inputs.NewLeftClick(removeEntityEvent)) // remove entity
		s.Inputs().KeepSelected().Set(quit, inputs.KeepSelectedComponent{})
		s.Collider().Component().Set(quit, collider.NewCollider(s.Definitions().Assets().SquareCollider))

		// child wrapper
		childWrapper := s.World().NewEntity()
		s.Hierarchy().SetParent(childWrapper, menu)
		s.Groups().InheritGroups(childWrapper)
		s.Transform().Pos().Set(childWrapper, transform.NewPos(0, -30 /* quit height + margin */, 0))
		s.Transform().Parent().Set(childWrapper, transform.NewParent(transform.RelativePos|transform.RelativeSizeXY))

		s.Layout().Order().Set(childWrapper, layout.NewOrder(layout.OrderVectical))
		s.Layout().Align().Set(childWrapper, layout.NewAlign(0, .5))
		s.Layout().Gap().Set(childWrapper, layout.NewGap(10))
		s.childrenWrapperArray.Set(childWrapper, childrenComponent{})
	}
}
