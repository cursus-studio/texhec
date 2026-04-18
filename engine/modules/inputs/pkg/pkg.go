package inputspkg

import (
	"engine/modules/inputs"
	"engine/modules/inputs/internal/mouse"
	"engine/modules/inputs/internal/service"
	"engine/modules/inputs/internal/systems"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/services/codec"
	"engine/services/ecs"
	"engine/services/frames"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PkgT[inputs.HoveredComponent](),
		prototypepkg.PkgT[inputs.DraggedComponent](),
		prototypepkg.PkgT[inputs.StackComponent](),
		prototypepkg.PkgT[inputs.StackedComponent](),
		prototypepkg.PkgT[inputs.KeepSelectedComponent](),
		prototypepkg.PkgT[inputs.LeftClickComponent](),
		prototypepkg.PkgT[inputs.DoubleLeftClickComponent](),
		prototypepkg.PkgT[inputs.RightClickComponent](),
		prototypepkg.PkgT[inputs.DoubleRightClickComponent](),
		prototypepkg.PkgT[inputs.MouseEnterComponent](),
		prototypepkg.PkgT[inputs.MouseLeaveComponent](),
		prototypepkg.PkgT[inputs.HoverComponent](),
		prototypepkg.PkgT[inputs.DragComponent](),
	} {
		pkg(b)
	}
	ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// components
			Register(inputs.HoveredComponent{}).
			Register(inputs.DraggedComponent{}).
			Register(inputs.StackComponent{}).
			Register(inputs.StackedComponent{}).
			Register(inputs.KeepSelectedComponent{}).
			Register(inputs.LeftClickComponent{}).
			Register(inputs.DoubleLeftClickComponent{}).
			Register(inputs.RightClickComponent{}).
			Register(inputs.DoubleRightClickComponent{}).
			Register(inputs.MouseEnterComponent{}).
			Register(inputs.MouseLeaveComponent{}).
			Register(inputs.HoverComponent{}).
			Register(inputs.DragComponent{}).

			// events
			Register(inputs.QuitEvent{}).
			Register(inputs.DragEvent{}).
			Register(inputs.SynchronizePositionEvent{})
	})

	ioc.Register(b, func(c ioc.Dic) inputs.Service {
		return service.NewService(c)
	})

	ioc.Register(b, func(c ioc.Dic) inputs.System {
		return ecs.NewSystemRegister(func() error {
			ecs.RegisterSystems(
				systems.NewInputsSystem(c),

				ecs.NewSystemRegister(func() error {
					eventsBuilder := ioc.Get[events.Builder](c)
					events.Listen(eventsBuilder, func(sdl.QuitEvent) {
						events.Emit(eventsBuilder.Events(), inputs.NewQuitEvent())
					})
					return nil
				}),

				mouse.NewCameraRaySystem(c),
				mouse.NewHoverSystem(c),
				mouse.NewHoverEventsSystem(c),
				mouse.NewClickSystem(c),
				ecs.NewSystemRegister(func() error {
					eventsBuilder := ioc.Get[events.Builder](c)
					events.Listen(eventsBuilder, func(frames.FrameEvent) {
						events.Emit(eventsBuilder.Events(), mouse.NewShootRayEvent())
					})
					return nil
				}),
			)
			return nil
		})
	})
	ioc.Register(b, func(c ioc.Dic) inputs.ShutdownSystem {
		return systems.NewQuitSystem(c)
	})
})
