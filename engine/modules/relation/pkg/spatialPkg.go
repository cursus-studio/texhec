package relationpkg

import (
	"engine/modules/relation"
	"engine/modules/relation/internal/onetokey"
	"engine/modules/warmup"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

func SpatialRelationPkg[IndexType any](
	queryFactory func(ecs.World) ecs.DirtySet,
	componentIndex func(ecs.World) func(entity ecs.EntityID) (indexType IndexType, ok bool),
	indexNumber func(index IndexType) uint32,
) ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		ioc.Register(b, func(c ioc.Dic) relation.Service[IndexType] {
			w := ioc.Get[ecs.World](c)
			return onetokey.NewSpatialIndex(
				w,
				queryFactory,
				componentIndex(w),
				indexNumber,
			)
		})

		ioc.Wrap(b, func(c ioc.Dic, b events.Builder) {
			r := ioc.Get[relation.Service[IndexType]](c)
			events.Listen(b, func(warmup.Event) {
				var i IndexType
				r.Get(i)
			})
		})
	})
}
