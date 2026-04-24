package transformpkg

import (
	codecpkg "engine/modules/codec/pkg"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/modules/transform"
	"engine/modules/transform/internal/transformservice"
	transitionpkg "engine/modules/transition/pkg"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[transform.PosComponent],
		codecpkg.PkgT[transform.RotationComponent],
		codecpkg.PkgT[transform.SizeComponent],

		codecpkg.PkgT[transform.MaxSizeComponent],
		codecpkg.PkgT[transform.MinSizeComponent],

		codecpkg.PkgT[transform.AspectRatioComponent],
		codecpkg.PkgT[transform.PivotPointComponent],

		codecpkg.PkgT[transform.ParentComponent],
		codecpkg.PkgT[transform.ParentPivotPointComponent],
		//
		prototypepkg.PkgT[transform.PosComponent],
		prototypepkg.PkgT[transform.RotationComponent],
		prototypepkg.PkgT[transform.SizeComponent],

		prototypepkg.PkgT[transform.MaxSizeComponent],
		prototypepkg.PkgT[transform.MinSizeComponent],

		prototypepkg.PkgT[transform.AspectRatioComponent],
		prototypepkg.PkgT[transform.PivotPointComponent],

		prototypepkg.PkgT[transform.ParentComponent],
		prototypepkg.PkgT[transform.ParentPivotPointComponent],
		//
		transitionpkg.PkgT[transform.PosComponent],
		transitionpkg.PkgT[transform.RotationComponent],
		transitionpkg.PkgT[transform.SizeComponent],

		transitionpkg.PkgT[transform.MaxSizeComponent],
		transitionpkg.PkgT[transform.MinSizeComponent],

		transitionpkg.PkgT[transform.AspectRatioComponent],
		transitionpkg.PkgT[transform.PivotPointComponent],

		transitionpkg.PkgT[transform.ParentPivotPointComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) transform.Service {
		return transformservice.NewService(c,
			transform.NewRotation(mgl32.QuatIdent()),
			transform.NewSize(1, 1, 1),
			transform.NewPivotPoint(.5, .5, .5),
			transform.NewParentPivotPoint(.5, .5, .5),
		)
	})
})
