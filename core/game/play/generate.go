package playscene

import (
	"core/game"
	"core/modules/lobby"
	"engine/modules/ecs"

	"github.com/ogiusek/ioc/v2"
)

var playGeneratePkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) game.PlayGenerateBuilder {
		s := ioc.Get[game.GameWorld](c)
		return func(sceneParent ecs.EntityID) {
			entity := s.World().NewEntity()
			s.Hierarchy().SetParent(entity, sceneParent)
			s.Lobby().Lobby().Set(entity, lobby.NewLobby())
		}
	})
})
