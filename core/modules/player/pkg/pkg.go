package playerpkg

import (
	"core/game"
	"core/modules/player"
	"core/modules/player/internal"
	"engine/modules/ecs"
	netsyncpkg "engine/modules/netsync/pkg"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	uuidpkg "engine/modules/uuid/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		uuidpkg.LinkPkgT[player.OwnerLink],
		typeregistrypkg.PkgT[player.PlayerContext],
		typeregistrypkg.PkgT[player.PlayerComponent],
		typeregistrypkg.PkgT[player.PlayerUUIDComponent],
		typeregistrypkg.PkgT[player.ActingPlayerComponent],
		typeregistrypkg.PkgT[player.AssignActingPlayerDTO],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) player.Service {
		return internal.NewService(c)
	})
	ioc.Wrap(b, func(c ioc.Dic, config netsyncpkg.Config) {
		world := ioc.Get[game.GameWorld](c)
		netsyncpkg.AddGenericEventValidation(config, func(client ecs.EntityID, event any) error {
			context, ok := event.(player.PlayerContextGetter)
			if !ok {
				return nil
			}
			ctx := context.Context()

			// require existing player
			playerEntity, ok := world.UUID().Entity(ctx.PlayerUUID)
			if !ok {
				return player.ErrRequiresControl
			}
			// require player to be controled by client
			conn, ok := world.Player().ActingConnection().Get(playerEntity)
			if !ok || conn.Connection != client {
				return player.ErrRequiresControl
			}
			return nil
		})
	})
})
