package uuidpkg

import (
	codecpkg "engine/modules/codec/pkg"
	prototypepkg "engine/modules/prototype/pkg"
	relationpkg "engine/modules/relation/pkg"
	uuid "engine/modules/uuid"
	"engine/modules/uuid/internal"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[uuid.UUID],
		codecpkg.PkgT[uuid.Component],

		prototypepkg.PkgT[uuid.Component],
		relationpkg.MapRelationPkg(
			func(w ecs.World) ecs.DirtySet {
				set := ecs.NewDirtySet()
				ecs.GetComponentsArray[uuid.Component](w).AddDirtySet(set)
				return set
			},
			func(w ecs.World) func(entity ecs.EntityID) (indexType uuid.UUID, ok bool) {
				uniqueArray := ecs.GetComponentsArray[uuid.Component](w)
				return func(entity ecs.EntityID) (indexType uuid.UUID, ok bool) {
					component, ok := uniqueArray.Get(entity)
					if !ok {
						return uuid.UUID{}, false
					}
					return component.ID, true
				}
			},
		),
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) uuid.Factory { return internal.NewFactory() })
	ioc.Register(b, func(c ioc.Dic) uuid.Service {
		return internal.NewService(c)
	})
})
