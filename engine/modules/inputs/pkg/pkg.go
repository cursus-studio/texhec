package inputspkg

import (
	"engine"
	"engine/modules/inputs"
	"engine/modules/inputs/internal/service"
	"engine/modules/scene"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
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

	ioc.Wrap(b, func(c ioc.Dic, b events.Builder) {
		world := ioc.GetServices[engine.EngineWorld](c)
		world.Scene() // ensure scene events are loaded before current event
		events.Listen(b, func(e scene.ChangeSceneEvent) {
			events.Emit(world.Events(), inputs.NewDefaultFocusEvent(world.Scene().Scene()))
		})
	})

	ioc.Register(b, func(c ioc.Dic) inputs.Service {
		return service.NewService(c)
	})
})
