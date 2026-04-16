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

type pkg struct {
	defaultPos         transform.PosComponent
	defaultRot         transform.RotationComponent
	defaultSize        transform.SizeComponent
	defaultPivot       transform.PivotPointComponent
	defaultParentPivot transform.ParentPivotPointComponent
}

func Package() ioc.Pkg {
	return pkg{
		transform.NewPos(0, 0, 0),
		transform.NewRotation(mgl32.QuatIdent()),
		transform.NewSize(1, 1, 1),
		transform.NewPivotPoint(.5, .5, .5),
		transform.NewParentPivotPoint(.5, .5, .5),
	}
}

func (pkg pkg) Register(b ioc.Builder) {
	for _, pkg := range []ioc.Pkg{
		prototypepkg.PackageT[transform.PosComponent](),
		prototypepkg.PackageT[transform.RotationComponent](),
		prototypepkg.PackageT[transform.SizeComponent](),

		prototypepkg.PackageT[transform.MaxSizeComponent](),
		prototypepkg.PackageT[transform.MinSizeComponent](),

		prototypepkg.PackageT[transform.AspectRatioComponent](),
		prototypepkg.PackageT[transform.PivotPointComponent](),

		prototypepkg.PackageT[transform.ParentComponent](),
		prototypepkg.PackageT[transform.ParentPivotPointComponent](),
		//
		transitionpkg.PackageT[transform.PosComponent](),
		transitionpkg.PackageT[transform.RotationComponent](),
		transitionpkg.PackageT[transform.SizeComponent](),

		transitionpkg.PackageT[transform.MaxSizeComponent](),
		transitionpkg.PackageT[transform.MinSizeComponent](),

		transitionpkg.PackageT[transform.AspectRatioComponent](),
		transitionpkg.PackageT[transform.PivotPointComponent](),

		transitionpkg.PackageT[transform.ParentPivotPointComponent](),
	} {
		pkg.Register(b)
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
			pkg.defaultRot,
			pkg.defaultSize,
			pkg.defaultPivot,
			pkg.defaultParentPivot,
		)
	})

}
