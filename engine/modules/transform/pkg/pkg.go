package transformpkg

import (
	prototypepkg "engine/modules/prototype/pkg"
	"engine/modules/transform"
	"engine/modules/transform/internal/transformservice"
	transitionpkg "engine/modules/transition/pkg"
	"engine/services/codec"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PkgT[transform.PosComponent](),
		prototypepkg.PkgT[transform.RotationComponent](),
		prototypepkg.PkgT[transform.SizeComponent](),

		prototypepkg.PkgT[transform.MaxSizeComponent](),
		prototypepkg.PkgT[transform.MinSizeComponent](),

		prototypepkg.PkgT[transform.AspectRatioComponent](),
		prototypepkg.PkgT[transform.PivotPointComponent](),

		prototypepkg.PkgT[transform.ParentComponent](),
		prototypepkg.PkgT[transform.ParentPivotPointComponent](),
		//
		transitionpkg.PkgT[transform.PosComponent](),
		transitionpkg.PkgT[transform.RotationComponent](),
		transitionpkg.PkgT[transform.SizeComponent](),

		transitionpkg.PkgT[transform.MaxSizeComponent](),
		transitionpkg.PkgT[transform.MinSizeComponent](),

		transitionpkg.PkgT[transform.AspectRatioComponent](),
		transitionpkg.PkgT[transform.PivotPointComponent](),

		transitionpkg.PkgT[transform.ParentPivotPointComponent](),
	} {
		pkg(b)
	}
	ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// components
			Register(transform.PosComponent{}).
			Register(transform.RotationComponent{}).
			Register(transform.SizeComponent{}).
			//
			Register(transform.MaxSizeComponent{}).
			Register(transform.MinSizeComponent{}).
			//
			Register(transform.AspectRatioComponent{}).
			Register(transform.PivotPointComponent{}).
			//
			Register(transform.ParentComponent{}).
			Register(transform.ParentPivotPointComponent{})
	})

	ioc.Register(b, func(c ioc.Dic) transform.Service {
		return transformservice.NewService(c,
			transform.NewRotation(mgl32.QuatIdent()),
			transform.NewSize(1, 1, 1),
			transform.NewPivotPoint(.5, .5, .5),
			transform.NewParentPivotPoint(.5, .5, .5),
		)
	})
})
