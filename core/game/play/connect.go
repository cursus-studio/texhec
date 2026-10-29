package playscene

import (
	"core/game"
	"engine/modules/ecs"
	"engine/modules/netsync"

	"github.com/ogiusek/ioc/v2"
)

var playConnectPkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) game.PlayConnectBuilder {
		s := ioc.Get[game.GameWorld](c)
		return func(sceneParent ecs.EntityID) {
			connectionEntity := s.World().NewEntity()
			s.Hierarchy().SetParent(connectionEntity, sceneParent)
			s.NetSync().Server().Set(connectionEntity, netsync.ServerComponent{})
			if err := s.Connection().Connect(connectionEntity, ":8080"); err != nil {
				s.Logger().Warn(err)
				return
			}
		}
	})
})
