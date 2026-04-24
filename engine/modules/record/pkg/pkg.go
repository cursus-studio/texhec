package recordpkg

import (
	"engine/modules/record"
	"engine/modules/record/internal/recordimpl"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/modules/uuid"

	"github.com/ogiusek/ioc/v2"
)

// this is parent configuration.
// it should have all used recorded components in any configuration.
// note: each new recorded component in configuration adds new BeforeGet to this type
// so do not add it automatically to everyhing because it can result in performance loss
var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[record.Recording],

		typeregistrypkg.PkgT[record.UUIDRecording],
		typeregistrypkg.PkgT[map[uuid.UUID][]any],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) record.Service {
		return recordimpl.NewService(c)
	})
})
