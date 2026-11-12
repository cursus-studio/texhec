package lobbypkg

import (
	"core/modules/lobby"
	"core/modules/lobby/internal"
	netsyncpkg "engine/modules/netsync/pkg"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[lobby.PlayersInitialsComponent],
		typeregistrypkg.PkgT[lobby.GameCameraComponent],
		typeregistrypkg.PkgT[lobby.LobbyComponent],
		typeregistrypkg.PkgT[lobby.LobbyUIComponent],

		typeregistrypkg.PkgT[lobby.SetSeedEvent],
		typeregistrypkg.PkgT[lobby.AddPlayerEvent],
		typeregistrypkg.PkgT[lobby.RemovePlayerEvent],
		typeregistrypkg.PkgT[lobby.SetPlayerNameEvent],
		typeregistrypkg.PkgT[lobby.SetPlayerColorEvent],

		typeregistrypkg.PkgT[lobby.SetPlayersWalletEvent],
		typeregistrypkg.PkgT[lobby.AddPlayersUnitEvent],
		typeregistrypkg.PkgT[lobby.RemovePlayersUnitEvent],
		typeregistrypkg.PkgT[lobby.SetConnectionPlayerEvent],
		typeregistrypkg.PkgT[lobby.SetPlayersConnectionEvent],

		typeregistrypkg.PkgT[lobby.GenerateEvent],

		//

		netsyncpkg.ComponentPkg[lobby.PlayersInitialsComponent],
		netsyncpkg.ComponentPkg[lobby.GameCameraComponent],
		netsyncpkg.ComponentPkg[lobby.LobbyComponent],

		netsyncpkg.EventPkg[lobby.SetSeedEvent],
		netsyncpkg.EventPkg[lobby.AddPlayerEvent],
		netsyncpkg.EventPkg[lobby.RemovePlayerEvent],
		netsyncpkg.EventPkg[lobby.SetPlayerNameEvent],
		netsyncpkg.EventPkg[lobby.SetPlayerColorEvent],

		netsyncpkg.EventPkg[lobby.SetPlayersWalletEvent],
		netsyncpkg.EventPkg[lobby.AddPlayersUnitEvent],
		netsyncpkg.EventPkg[lobby.RemovePlayersUnitEvent],
		netsyncpkg.EventPkg[lobby.SetConnectionPlayerEvent],
		netsyncpkg.EventPkg[lobby.SetPlayersConnectionEvent],

		netsyncpkg.EventPkg[lobby.GenerateEvent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, internal.NewService)
})
