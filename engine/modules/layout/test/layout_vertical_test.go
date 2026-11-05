package test

import (
	"engine/modules/layout"
	"engine/modules/transform"
	"testing"
)

func TestLayoutVertical(t *testing.T) {
	setup := NewSetup(t)

	parent := setup.World().NewEntity()

	setup.Layout().Order().Set(parent, layout.NewOrder(layout.OrderVectical))
	setup.Layout().Align().Set(parent, layout.NewAlign(.5, .5))
	setup.Layout().Gap().Set(parent, layout.NewGap(10))

	// test one button
	btn1 := setup.World().NewEntity()
	setup.Hierarchy().SetParent(btn1, parent)
	setup.Transform().Size().Set(btn1, transform.NewSize(10, 10, 10))
	setup.Transform().Inherit().Set(btn1, transform.NewInherit(transform.RelativePos))

	setup.Expect(btn1, 0, 0)

	setup.Layout().Align().Set(parent, layout.NewAlign(1, 1))
	setup.Expect(btn1, -4.5, 4.5)

	setup.Layout().Align().Set(parent, layout.NewAlign(0, 0))
	setup.Expect(btn1, 4.5, -4.5)

	setup.Transform().Size().Set(btn1, transform.NewSize(0, 0, 0))
	setup.Expect(btn1, -0.5, 0.5)

	setup.Layout().Align().Set(parent, layout.NewAlign(1, 1))
	setup.Expect(btn1, 0.5, -0.5)

	// reset to the beginning state
	setup.Layout().Order().Set(parent, layout.NewOrder(layout.OrderVectical))
	setup.Layout().Align().Set(parent, layout.NewAlign(.5, .5))

	setup.Transform().Size().Set(btn1, transform.NewSize(0, 0, 0))

	// test two buttons
	btn2 := setup.World().NewEntity()
	setup.Hierarchy().SetParent(btn2, parent)
	setup.Transform().Size().Set(btn2, transform.NewSize(0, 0, 0))
	setup.Transform().Inherit().Set(btn2, transform.NewInherit(transform.RelativePos))

	setup.Expect(btn1, 0, 5)
	setup.Expect(btn2, 0, -5)

	setup.World().RemoveEntity(parent)
}
