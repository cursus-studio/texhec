package focuspkg

import (
	"engine"
	"engine/modules/focus"
	"engine/modules/focus/internal"
	"engine/modules/scene"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[focus.BubblingComponent],
		typeregistrypkg.PkgT[focus.BubbleEvent],

		typeregistrypkg.PkgT[focus.UnfocusEvent],

		typeregistrypkg.PkgT[focus.FocusEvent],
		typeregistrypkg.PkgT[focus.DefaultFocusEvent],

		typeregistrypkg.PkgT[focus.FocusedComponent],
		typeregistrypkg.PkgT[focus.DefaultFocusedComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, internal.NewService)

	ioc.Wrap(b, func(c ioc.Dic, b events.Builder) {
		world := ioc.Get[engine.EngineWorld](c)
		world.Scene() // ensure scene events are loaded before current event
		events.Listen(b, func(e scene.ChangeSceneEvent) {
			events.Emit(world.Events(), focus.NewDefaultFocusEvent(world.Scene().Scene()))
		})
	})
})

func BubblePkgT[SrcEvent any, TgtEvent any](ctor func(SrcEvent) TgtEvent) ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		ioc.Wrap(b, func(c ioc.Dic, b events.Builder) {
			world := ioc.Get[engine.EngineWorld](c)
			events.Listen(world.EventsBuilder(), func(event SrcEvent) {
				bubbleEvent, ok := world.Focus().NewFocusedBubbleEvent(ctor(event))
				if !ok {
					return
				}
				world.Focus().Emit(bubbleEvent)
			})
		})
	})
}
