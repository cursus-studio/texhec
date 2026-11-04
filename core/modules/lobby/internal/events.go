package internal

import (
	"core/modules/lobby"
	"core/modules/player"
	"engine/modules/seed"
	"engine/modules/uuid"
	"fmt"
	"math/rand"
	"slices"
	"strconv"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
)

func (s *service) OnReRollSeed(event lobby.ReRollSeedEvent) {
	events.Emit(s.Events(), lobby.NewSetSeed(s.Seed().GenerateSeed()))
}
func (s *service) OnPasteSeed(event lobby.PasteSeedEvent) {
	value, _ := s.Window().GetClipboard()
	seedValue, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		s.Logger().Warn(err)
	}
	events.Emit(s.Events(), lobby.NewSetSeed(seed.Seed(seedValue)))
}
func (s *service) OnSetSeed(event lobby.SetSeedEvent) {
	seedEntity, _, ok := s.Seed().Seed().GetSingleton()
	if !ok {
		return
	}
	s.Seed().Seed().Set(seedEntity, seed.NewSeed(event.Seed))
}
func (s *service) OnInputAddPlayer(event lobby.InputAddPlayerEvent) {
	events.Emit(s.Events(), lobby.NewAddPlayerEvent(s.UUID().NewUUID()))
}
func (s *service) OnAddPlayer(event lobby.AddPlayerEvent) {
	name := player.NewName("player")
	color := player.NewColor(mgl32.Vec3{0, 1, 0})
	for i := 0; true; i++ {
		newName := player.NewName(fmt.Sprintf("player %v", i))
		if _, ok := s.Player().NameRelation().Get(newName); !ok {
			name = newName
			// #nosec G404
			rng := rand.New(rand.NewSource(int64(i)))
			color.Color = mgl32.Vec3{rng.Float32(), rng.Float32(), rng.Float32()}
			break
		}
	}
	entity := s.World().NewEntity()
	s.UUID().UUID().Set(entity, uuid.NewUUID(event.NewPlayerUUID))
	s.Player().Name().Set(entity, name)
	s.Player().Color().Set(entity, color)
}
func (s *service) OnRemovePlayer(event lobby.RemovePlayerEvent) {
	if entity, ok := s.UUID().Entity(event.PlayerUUID); ok {
		s.World().RemoveEntity(entity)
	}
}
func (s *service) OnSetPlayerName(event lobby.SetPlayerNameEvent) {
	name := player.NewName(event.Name)
	if _, ok := s.Player().NameRelation().Get(name); ok {
		s.Logger().Warn(player.ErrNameCollision)
		return
	}
	if entity, ok := s.UUID().Entity(event.PlayerUUID); ok {
		s.Player().Name().Set(entity, name)
	}
}
func (s *service) OnSetPlayerColor(event lobby.SetPlayerColorEvent) {
	color := player.NewColor(event.Color)
	if entity, ok := s.UUID().Entity(event.PlayerUUID); ok {
		s.Player().Color().Set(entity, color)
	}
}

func (s *service) OnSetPlayersWallet(event lobby.SetPlayersWalletEvent) {
	seedEntity, _, ok := s.Seed().Seed().GetSingleton()
	if !ok {
		s.Logger().Warn(seed.ErrWorldShouldHaveOneSeed)
		return
	}
	initials, _ := s.PlayersInitials().Get(seedEntity)
	initials.Wallet = event.Wallet
	s.PlayersInitials().Set(seedEntity, initials)
}
func (s *service) OnAddPlayersUnitEvent(event lobby.AddPlayersUnitEvent) {
	seedEntity, _, ok := s.Seed().Seed().GetSingleton()
	if !ok {
		s.Logger().Warn(seed.ErrWorldShouldHaveOneSeed)
		return
	}
	initials, _ := s.PlayersInitials().Get(seedEntity)
	initials.Units = append(initials.Units, event.UUID)
	s.PlayersInitials().Set(seedEntity, initials)
}
func (s *service) OnRemovePlayersUnitEvent(event lobby.RemovePlayersUnitEvent) {
	seedEntity, _, ok := s.Seed().Seed().GetSingleton()
	if !ok {
		s.Logger().Warn(seed.ErrWorldShouldHaveOneSeed)
		return
	}
	initials, _ := s.PlayersInitials().Get(seedEntity)
	if i := slices.Index(initials.Units, event.UUID); i >= 0 {
		initials.Units = append(initials.Units[:i], initials.Units[i+1:]...)
	}
	s.PlayersInitials().Set(seedEntity, initials)
}

func (s *service) OnSetConnectionPlayer(event lobby.SetConnectionPlayerEvent) {
	playerEntity, ok := s.UUID().Entity(event.Player)
	if !ok {
		return
	}
	connectionEntity, ok := s.UUID().Entity(event.Connection)
	if !ok {
		return
	}

	if _, isServer := s.Connection().Connection().Get(connectionEntity); isServer {
		s.Player().ActingConnection().Set(playerEntity, player.NewActiongConnection(connectionEntity))
		return
	}

	s.Player().ActingPlayer().Set(playerEntity, player.NewActingPlayer())
}

func (s *service) OnSetPlayersConnection(event lobby.SetPlayersConnectionEvent) {
	entity, comp, ok := s.Player().PlayersConnection().GetSingleton()
	if !ok {
		s.Logger().Warn(fmt.Errorf("no players connection to set"))
		return
	}
	comp.AcceptNewConnections = event.AcceptNewConnections
	s.Player().PlayersConnection().Set(entity, comp)
}
