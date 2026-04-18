package recordpkg

import (
	"engine/modules/record"
	"engine/modules/record/internal/recordimpl"
	"engine/modules/uuid"
	"engine/services/codec"
	"engine/services/datastructures"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

// this is parent configuration.
// it should have all used recorded components in any configuration.
// note: each new recorded component in configuration adds new BeforeGet to this type
// so do not add it automatically to everyhing because it can result in performance loss
var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Wrap(b, func(c ioc.Dic, b codec.Builder) {
		b.
			// recording
			Register(record.Recording{}).
			Register(datastructures.NewSparseArray[ecs.EntityID, []any]()).

			// uuid recording
			Register(record.UUIDRecording{}).
			Register(map[uuid.UUID][]any{})
	})
	ioc.Register(b, func(c ioc.Dic) record.Service {
		return recordimpl.NewService(c)
	})
})
