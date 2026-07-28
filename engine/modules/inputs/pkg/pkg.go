package inputspkg

import (
	focuspkg "engine/modules/focus/pkg"
	"engine/modules/inputs"
	"engine/modules/inputs/internal/service"
	typeregistrypkg "engine/modules/typeregistry/pkg"

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
		typeregistrypkg.PkgT[inputs.TextInputComponent],

		// events
		typeregistrypkg.PkgT[inputs.DragEvent],
		typeregistrypkg.PkgT[inputs.SynchronizePositionEvent],
		typeregistrypkg.PkgT[inputs.KeyboardEvent],
		typeregistrypkg.PkgT[inputs.TextInputEvent],

		focuspkg.BubblePkgT(inputs.NewKeyboardEvent),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) inputs.Service {
		return service.NewService(c)
	})
})
