package inputspkg

import (
	"engine/modules/inputs"
	"engine/modules/inputs/internal/mouse"
	"engine/modules/inputs/internal/service"
	"engine/modules/inputs/internal/systems"
	"engine/modules/loop"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[inputs.HoveredComponent],
		typeregistrypkg.PkgT[inputs.DraggedComponent],
		typeregistrypkg.PkgT[inputs.StackComponent],
		typeregistrypkg.PkgT[inputs.StackedComponent],
		typeregistrypkg.PkgT[inputs.KeepSelectedComponent],
		typeregistrypkg.PkgT[inputs.LeftClickComponent],
		typeregistrypkg.PkgT[inputs.DoubleLeftClickComponent],
		typeregistrypkg.PkgT[inputs.RightClickComponent],
		typeregistrypkg.PkgT[inputs.DoubleRightClickComponent],
		typeregistrypkg.PkgT[inputs.MouseEnterComponent],
		typeregistrypkg.PkgT[inputs.MouseLeaveComponent],
		typeregistrypkg.PkgT[inputs.HoverComponent],
		typeregistrypkg.PkgT[inputs.DragComponent],

		// events
		typeregistrypkg.PkgT[inputs.DragEvent],
		typeregistrypkg.PkgT[inputs.SynchronizePositionEvent],
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
