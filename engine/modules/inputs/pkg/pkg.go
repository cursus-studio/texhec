package inputspkg

import (
	codecpkg "engine/modules/codec/pkg"
	"engine/modules/inputs"
	"engine/modules/inputs/internal/mouse"
	"engine/modules/inputs/internal/service"
	"engine/modules/inputs/internal/systems"
	"engine/modules/loop"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[inputs.HoveredComponent],
		codecpkg.PkgT[inputs.DraggedComponent],
		codecpkg.PkgT[inputs.StackComponent],
		codecpkg.PkgT[inputs.StackedComponent],
		codecpkg.PkgT[inputs.KeepSelectedComponent],
		codecpkg.PkgT[inputs.LeftClickComponent],
		codecpkg.PkgT[inputs.DoubleLeftClickComponent],
		codecpkg.PkgT[inputs.RightClickComponent],
		codecpkg.PkgT[inputs.DoubleRightClickComponent],
		codecpkg.PkgT[inputs.MouseEnterComponent],
		codecpkg.PkgT[inputs.MouseLeaveComponent],
		codecpkg.PkgT[inputs.HoverComponent],
		codecpkg.PkgT[inputs.DragComponent],

		// events
		codecpkg.PkgT[inputs.DragEvent],
		codecpkg.PkgT[inputs.SynchronizePositionEvent],

		prototypepkg.PkgT[inputs.HoveredComponent],
		prototypepkg.PkgT[inputs.DraggedComponent],
		prototypepkg.PkgT[inputs.StackComponent],
		prototypepkg.PkgT[inputs.StackedComponent],
		prototypepkg.PkgT[inputs.KeepSelectedComponent],
		prototypepkg.PkgT[inputs.LeftClickComponent],
		prototypepkg.PkgT[inputs.DoubleLeftClickComponent],
		prototypepkg.PkgT[inputs.RightClickComponent],
		prototypepkg.PkgT[inputs.DoubleRightClickComponent],
		prototypepkg.PkgT[inputs.MouseEnterComponent],
		prototypepkg.PkgT[inputs.MouseLeaveComponent],
		prototypepkg.PkgT[inputs.HoverComponent],
		prototypepkg.PkgT[inputs.DragComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) inputs.Service {
		return service.NewService(c)
	})

	ioc.Register(b, func(c ioc.Dic) inputs.System {
		return ecs.NewSystemRegister(func() error {
			eventsBuilder := ioc.Get[events.Builder](c)
			events.Listen(eventsBuilder, func(loop.FrameEvent) {
				events.Emit(eventsBuilder.Events(), mouse.NewShootRayEvent())
			})
			events.Listen(eventsBuilder, func(sdl.QuitEvent) {
				events.Emit(eventsBuilder.Events(), loop.NewStopEvent())
			})
			ecs.RegisterSystems(
				systems.NewInputsSystem(c),
				mouse.NewCameraRaySystem(c),
				mouse.NewHoverSystem(c),
				mouse.NewHoverEventsSystem(c),
				mouse.NewClickSystem(c),
			)
			return nil
		})
	})
})
