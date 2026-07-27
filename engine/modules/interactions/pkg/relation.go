package interactionspkg

import (
	"engine/modules/ecs"
	"engine/modules/interactions/internal"

	"github.com/ogiusek/ioc/v2"
)

type Relation = internal.RawRelation

func NewRelation(
	src, tgt uintptr,
	set func(c ioc.Dic) func(srcEntity ecs.EntityID, tgtEntity ecs.EntityID),
) Relation {
	return internal.NewRawRelation(src, tgt, set)
}
func NewCopyRelation[Component any](src, tgt uintptr) Relation {
	return internal.NewRawRelation(src, tgt, func(c ioc.Dic) func(srcEntity ecs.EntityID, tgtEntity ecs.EntityID) {
		arr := ecs.GetComponentArray[Component](ioc.Get[ecs.World](c))
		return func(srcEntity, tgtEntity ecs.EntityID) {
			comp, _ := arr.Get(srcEntity)
			arr.Set(tgtEntity, comp)
		}
	})
}
