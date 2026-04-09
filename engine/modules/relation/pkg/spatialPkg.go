package relationpkg

import (
	"engine/modules/relation"
	"engine/modules/relation/internal/onetokey"
	"engine/modules/warmup"
	"engine/services/ecs"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type spatialRelationPkg[IndexType any] struct {
	queryFactory   func(ecs.World) ecs.DirtySet
	componentIndex func(ecs.World) func(ecs.EntityID) (IndexType, bool)
	indexNumber    func(IndexType) uint32
}

func SpatialRelationPackage[IndexType any](
	queryFactory func(ecs.World) ecs.DirtySet,
	componentIndex func(ecs.World) func(entity ecs.EntityID) (indexType IndexType, ok bool),
	indexNumber func(index IndexType) uint32,
) ioc.Pkg {
	return spatialRelationPkg[IndexType]{
		queryFactory:   queryFactory,
		componentIndex: componentIndex,
		indexNumber:    indexNumber,
	}
}

func (pkg spatialRelationPkg[IndexType]) Register(b ioc.Builder) {
	ioc.RegisterSingleton(b, func(c ioc.Dic) relation.Service[IndexType] {
		w := ioc.Get[ecs.World](c)
		return onetokey.NewSpatialIndex(
			w,
			pkg.queryFactory,
			pkg.componentIndex(w),
			pkg.indexNumber,
		)
	})

	ioc.WrapService(b, func(c ioc.Dic, b events.Builder) {
		r := ioc.Get[relation.Service[IndexType]](c)
		events.Listen(b, func(warmup.Event) {
			var i IndexType
			r.Get(i)
		})
	})
}
