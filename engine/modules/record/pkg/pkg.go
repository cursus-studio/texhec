package recordpkg

import (
	codecpkg "engine/modules/codec/pkg"
	"engine/modules/record"
	"engine/modules/record/internal/recordimpl"
	"engine/modules/uuid"

	"github.com/ogiusek/ioc/v2"
)

// this is parent configuration.
// it should have all used recorded components in any configuration.
// note: each new recorded component in configuration adds new BeforeGet to this type
// so do not add it automatically to everyhing because it can result in performance loss
var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[record.Recording],
		// TODO without this package recording.Recording cannot be sent using netsync
		// codecpkg.PkgT[datastructures.SparseArray[ecs.EntityID, []any]],

		codecpkg.PkgT[record.UUIDRecording],
		codecpkg.PkgT[map[uuid.UUID][]any],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) record.Service {
		return recordimpl.NewService(c)
	})
})
