package transformpkg

import (
	"engine/modules/transform"
	"engine/modules/transform/internal/transformservice"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[transform.PosComponent],
		typeregistrypkg.PkgT[transform.RotationComponent],
		typeregistrypkg.PkgT[transform.SizeComponent],

		typeregistrypkg.PkgT[transform.MaxSizeComponent],
		typeregistrypkg.PkgT[transform.MinSizeComponent],

		typeregistrypkg.PkgT[transform.AspectRatioComponent],
		typeregistrypkg.PkgT[transform.PivotPointComponent],

		typeregistrypkg.PkgT[transform.ParentComponent],
		typeregistrypkg.PkgT[transform.ParentPivotPointComponent],
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
