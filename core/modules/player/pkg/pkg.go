package playerpkg

import (
	"core/game"
	"core/modules/actions"
	"core/modules/lobby"
	"core/modules/player"
	"core/modules/player/internal"
	connectionpkg "engine/modules/connection/pkg"
	"engine/modules/ecs"
	featurepkg "engine/modules/feature/pkg"
	netsyncpkg "engine/modules/netsync/pkg"
	relationpkg "engine/modules/relation/pkg"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/modules/uuid"
	uuidpkg "engine/modules/uuid/pkg"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		relationpkg.MapRelationPkg(func(w ecs.World) ecs.DirtySet {
			set := ecs.NewDirtySet()
			ecs.GetComponentArray[player.NameComponent](w).AddDirtySet(set)
			return set
		}, func(w ecs.World) func(entity ecs.EntityID) (indexType player.NameComponent, ok bool) {
			arr := ecs.GetComponentArray[player.NameComponent](w)
			return func(entity ecs.EntityID) (indexType player.NameComponent, ok bool) {
				return arr.Get(entity)
			}
		}),
		uuidpkg.LinkPkgT[player.OwnerLink],
		typeregistrypkg.PkgT[player.NameComponent],
		typeregistrypkg.PkgT[player.ColorComponent],
		typeregistrypkg.PkgT[player.PlayerUUIDComponent],
		typeregistrypkg.PkgT[player.ActingPlayerComponent],
		typeregistrypkg.PkgT[player.AssignActingPlayerDTO],

		netsyncpkg.ComponentPkg[player.NameComponent],
		netsyncpkg.ComponentPkg[player.ColorComponent],
		netsyncpkg.ComponentPkg[player.PlayerUUIDComponent],
		netsyncpkg.ComponentPkg[uuid.LinkUUIDComponent[player.OwnerLink]],

		connectionpkg.ParentPkg(func(c ioc.Dic) func(parent ecs.EntityID, conn ecs.EntityID, comp player.PlayersConnectionComponent) {
			s := ioc.Get[game.GameWorld](c)
			getFreePlayer := func() (ecs.EntityID, uuid.UUID) {
				for _, playerEntity := range s.Player().Name().GetEntities() {
					if _, ok := s.Player().ActingPlayer().Get(playerEntity); ok {
						continue
					}
					if _, ok := s.Player().ActingConnection().Get(playerEntity); ok {
						continue
					}
					uuid, ok := s.UUID().UUID().Get(playerEntity)
					if !ok {
						continue
					}
					return playerEntity, uuid.ID
				}
				var uuid uuid.UUID
				return 0, uuid
			}
			return func(parent, connEntity ecs.EntityID, comp player.PlayersConnectionComponent) {
				conn, ok := s.Connection().Connection().Get(connEntity)
				if !ok {
					return
				}

				playerEntity, playerUUID := getFreePlayer()
				if playerEntity == 0 {
					if !comp.AcceptNewConnections {
						s.World().RemoveEntity(connEntity)
						return
					}
					events.Emit(s.Events(), lobby.NewInputAddPlayerEvent())
					playerEntity, playerUUID = getFreePlayer()
				}
				s.Player().PlayerConnection().Set(connEntity, player.NewPlayerConnection(playerUUID))
				s.Player().ActingConnection().Set(playerEntity, player.NewActiongConnection(connEntity))
				if err := conn.Conn().Send(player.NewAssignActingPlayerDTO(playerUUID)); err != nil {
					s.Logger().Warn(err)
				}
			}
		}),
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
