package internal

import (
	"core/game"
	"core/modules/economy"
	"core/modules/lobby"
	"core/modules/tile"
	"core/modules/ui"
	"engine/modules/batcher"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/uuid"
	"fmt"
	"slices"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld  `inject:""`
	playersInitials ecs.ComponentArray[lobby.PlayersInitialsComponent]
	gameCamera      ecs.ComponentArray[lobby.GameCameraComponent]
	lobby           ecs.ComponentArray[lobby.LobbyComponent]
	lobbyUI         ecs.ComponentArray[lobby.LobbyUIComponent]

	inputChanges map[ecs.EntityID]func(string)
}

func wrapCall[Wrapped any](call func(), wrap func(Wrapped)) func(Wrapped) {
	return func(w Wrapped) {
		wrap(w)
		call()
	}
}

func NewService(c ioc.Dic) lobby.Service {
	s := ioc.GetServices[*service](c)
	s.playersInitials = ecs.GetComponentArray[lobby.PlayersInitialsComponent](s.World())
	s.gameCamera = ecs.GetComponentArray[lobby.GameCameraComponent](s.World())
	s.lobby = ecs.GetComponentArray[lobby.LobbyComponent](s.World())
	s.lobbyUI = ecs.GetComponentArray[lobby.LobbyUIComponent](s.World())
	s.inputChanges = make(map[ecs.EntityID]func(string))

	s.Inputs().TextInput().OnMod(s.onTextInputMod)

	s.Lobby().OnUpsert(func(ecs.EntityID) { s.onLobbyUpsert() })

	events.Listen(s.EventsBuilder(), s.OnReRollSeed)
	events.Listen(s.EventsBuilder(), s.OnPasteSeed)
	events.Listen(s.EventsBuilder(), wrapCall(s.onLobbyUpsert, s.OnSetSeed))

	events.Listen(s.EventsBuilder(), s.OnInputAddPlayer)
	events.Listen(s.EventsBuilder(), wrapCall(s.onLobbyUpsert, s.OnAddPlayer))
	events.Listen(s.EventsBuilder(), wrapCall(s.onLobbyUpsert, s.OnRemovePlayer))

	events.Listen(s.EventsBuilder(), wrapCall(s.onLobbyUpsert, s.OnSetPlayerName))
	events.Listen(s.EventsBuilder(), wrapCall(s.onLobbyUpsert, s.OnSetPlayerColor))
	events.Listen(s.EventsBuilder(), wrapCall(s.onLobbyUpsert, s.OnSetPlayersWallet))
	events.Listen(s.EventsBuilder(), wrapCall(s.onLobbyUpsert, s.OnAddPlayersUnitEvent))
	events.Listen(s.EventsBuilder(), wrapCall(s.onLobbyUpsert, s.OnRemovePlayersUnitEvent))
	events.Listen(s.EventsBuilder(), wrapCall(s.onLobbyUpsert, s.OnSetConnectionPlayer))
	events.Listen(s.EventsBuilder(), wrapCall(s.onLobbyUpsert, s.OnSetPlayersConnection))

	events.Listen(s.EventsBuilder(), s.onGenerate)

	return s
}

func (s *service) PlayersInitials() ecs.ComponentArray[lobby.PlayersInitialsComponent] {
	return s.playersInitials
}
func (s *service) GameCamera() ecs.ComponentArray[lobby.GameCameraComponent] {
	return s.gameCamera
}
func (s *service) Lobby() ecs.ComponentArray[lobby.LobbyComponent] {
	return s.lobby
}
func (s *service) LobbyUI() ecs.ComponentArray[lobby.LobbyUIComponent] {
	return s.lobbyUI
}

//

func (s *service) onLobbyUpsert() {
	if len(s.Lobby().GetEntities()) == 0 {
		return
	}
	clear(s.inputChanges)
	isServer := len(s.NetSync().Server().GetEntities()) == 0

	if isServer {
		s.setupServerLobbyVariables()
	} else {
		s.setupClientLobbyVariables()
	}

	for _, entity := range s.LobbyUI().GetEntities() {
		s.World().RemoveEntity(entity)
	}

	if isServer {
		s.renderServerLobby()
	} else {
		s.renderClientLobby()
	}
}
func (s *service) onGenerate(lobby.GenerateEvent) {
	for _, lobby := range s.Lobby().GetEntities() {
		s.Lobby().Remove(lobby)
	}

	scene := s.Scene().Scene()
	gameCameraEntity := s.GameCamera().GetEntities()[0]
	gameCameraComp, _ := s.GameCamera().Get(gameCameraEntity)
	playersInitials, _ := s.PlayersInitials().Get(gameCameraEntity)

	clear(s.inputChanges)
	for _, entity := range s.LobbyUI().GetEntities() {
		s.World().RemoveEntity(entity)
	}

	s.Ui().AddUICamera(scene)
	s.Ui().AddBgCamera(scene)
	s.Ui().GameCamera().Set(gameCameraEntity, ui.NewGameCamera(gameCameraComp.MaxChunksZoomOut))

	for x := range grid.Coord(gameCameraComp.MaxChunksZoomOut) {
		for y := range grid.Coord(gameCameraComp.MaxChunksZoomOut) {
			events.Emit(s.Events(), tile.NewMissingChunkEvent(grid.NewChunkCoords(x, y)))
		}
	}

	//

	units := make([]ecs.EntityID, 0, len(playersInitials.Units))
	for _, unit := range playersInitials.Units {
		entity, ok := s.UUID().Entity(unit)
		if ok {
			units = append(units, entity)
		}
	}

	unitsBatch := batcher.NewBatch(1, func(i int) {
		for _, playerEntity := range s.Player().Name().GetEntities() {
			s.Economy().Wallet().Set(playerEntity, playersInitials.Wallet)

			units := slices.Clone(units)
		loop:
			for y := range s.Grid().ChunkSize() {
				for x := range s.Grid().ChunkSize() {
					coords := grid.NewCoords(x, y)
					coords = s.Grid().AbsoluteCoords(grid.NewChunkCoords(0, 0), coords)
					deployed := units[0]
					if _, err := s.Deploy().Deploy(deployed, playerEntity, coords); err == nil {
						units = units[1:]
					}
					if len(units) == 0 {
						break loop
					}
				}
			}
		}
	})

	// task
	task := s.Batcher().NewTask()
	task.AddOrderedBatch(unitsBatch)
	s.Batcher().Queue(task.Build())

	for _, task := range s.Batcher().Tasks() {
		task.Perform()
	}

}

func (s *service) NewPlayersInitials(wallet economy.WalletComponent, units []ecs.EntityID) lobby.PlayersInitialsComponent {
	uuids := make([]uuid.UUID, 0, len(units))
	for _, entity := range units {
		uuid, ok := s.UUID().Component().Get(entity)
		if !ok {
			s.Logger().Warn(fmt.Errorf("entity is missing uuid"))
			continue
		}
		uuids = append(uuids, uuid.ID)
	}
	return lobby.NewPlayersInitials(wallet, uuids)
}

// add getters for:
// - seed
// - game camera
// - netsync

// render:
// - run btn
// - re-roll seed btn
// - change max zoom btn
// - max entities zoom out

// render menu where:
// - I can set seed
// - Map size
// - Players
// - Who controls players
// - Players starting tools
// create a form to generate this along with lobby
