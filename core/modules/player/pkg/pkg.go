package playerpkg

import (
	"core/game"
	"core/modules/actions"
	"core/modules/player"
	"core/modules/player/internal"
	"engine/modules/ecs"
	featurepkg "engine/modules/feature/pkg"
	netsyncpkg "engine/modules/netsync/pkg"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/modules/uuid"
	uuidpkg "engine/modules/uuid/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		uuidpkg.LinkPkgT[player.OwnerLink],
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
			context, ok := event.(actions.PlayerContextGetter)
			if !ok {
				return nil
			}
			ctx := context.Context()

			// require existing player
			playerEntity, ok := world.UUID().Entity(uuid.UUID(ctx))
			if !ok {
				return world.Actions().Own().Error()
			}
			// require player to be controled by client
			conn, ok := world.Player().ActingConnection().Get(playerEntity)
			if !ok || conn.Connection != client {
				return world.Actions().Own().Error()
			}
			return nil
		})
	})
	ioc.Wrap(b, func(c ioc.Dic, config featurepkg.Config) {
		featurepkg.RegisterFieldProvider(config, ioc.Get[game.GameWorld](c).Player().GetContext)
	})
})
