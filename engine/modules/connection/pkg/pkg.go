package connectionpkg

import (
	"engine"
	"engine/modules/connection"
	"engine/modules/connection/internal"
	"engine/modules/ecs"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) connection.Service {
		return internal.NewService(c)
	})
})

func ParentPkg[Parent any](reaction func(c ioc.Dic) func(parent, conn ecs.EntityID, comp Parent)) ioc.Pkg {
	return ioc.NewPkg(func(b ioc.Builder) {
		ioc.Wrap(b, func(c ioc.Dic, connection connection.Service) {
			reaction := reaction(c)

			s := ioc.Get[engine.EngineWorld](c)
			parent := ecs.GetComponentArray[Parent](s.World())
			connection.Connection().OnUpsert(func(entity ecs.EntityID) {
				parentEntity, ok := s.Hierarchy().GetParent(entity)
				if !ok {
					return
				}
				parentComp, ok := parent.Get(parentEntity)
				if !ok {
					return
				}
				reaction(parentEntity, entity, parentComp)
			})
		})
	})
}
