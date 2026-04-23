package test

import (
	"engine/modules/relation"
	"engine/modules/relation/pkg"
	enginepkg "engine/pkg"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type Component struct {
	Index uint32
}

type Setup struct {
	W       ecs.World
	Array   ecs.ComponentsArray[Component]
	Service relation.Service[uint32]
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		enginepkg.Pkg,
		relationpkg.SpatialRelationPkg(
			func(w ecs.World) ecs.DirtySet {
				dirtySet := ecs.NewDirtySet()
				ecs.GetComponentsArray[Component](w).AddDirtySet(dirtySet)
				return dirtySet
			},
			func(w ecs.World) func(entity ecs.EntityID) (uint32, bool) {
				componentArray := ecs.GetComponentsArray[Component](w)
				return func(entity ecs.EntityID) (uint32, bool) {
					comp, ok := componentArray.Get(entity)
					return comp.Index, ok
				}
			},
			func(index uint32) uint32 { return index },
		),
	)

	w := ioc.Get[ecs.World](c)
	return Setup{
		W:       w,
		Array:   ecs.GetComponentsArray[Component](w),
		Service: ioc.Get[relation.Service[uint32]](c),
	}
}
