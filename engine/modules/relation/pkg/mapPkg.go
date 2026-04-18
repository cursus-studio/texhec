package relationpkg

import (
	"engine/modules/relation"
	"engine/modules/relation/internal/onetokey"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

func MapRelationPkg[IndexType comparable](
	queryFactory func(ecs.World) ecs.DirtySet,
	componentIndex func(ecs.World) func(entity ecs.EntityID) (indexType IndexType, ok bool),
) ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		ioc.Register(b, func(c ioc.Dic) relation.Service[IndexType] {
			w := ioc.Get[ecs.World](c)
			return onetokey.NewMapIndex(
				w,
				queryFactory,
				componentIndex(w),
			)
		})
	})
}
