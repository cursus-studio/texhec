package client

import (
	"engine"
	"engine/modules/connection"
	"engine/modules/ecs"
	"engine/modules/loop"
	"engine/modules/netsync/internal/clienttypes"
	"engine/modules/netsync/internal/config"
	"engine/modules/netsync/internal/prediction"
	"engine/modules/netsync/internal/servertypes"
	"errors"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

var (
	ErrInvalidPrediction    error = errors.New("client had invalid prediction")
	ErrExceededPredictions  error = errors.New("predictions count exceed maximal amount")
	ErrHasMoreThanOneServer error = errors.New("has more than one server")
)

// can:
// - apply server event
// - apply predicted event (starts and ends prediction)
type Service struct {
	engine.EngineWorld `inject:""`
	Prediction         prediction.Service      `inject:""`
	Config             ioc.Lazy[config.Config] `inject:""`

	dirtySet ecs.DirtySet
}

func NewService(c ioc.Dic) *Service {
	s := ioc.GetServices[*Service](c)

	s.dirtySet = ecs.NewDirtySet()

	s.NetSync().Server().AddDirtySet(s.dirtySet)
	s.Connection().Component().AddDirtySet(s.dirtySet)

	events.Listen(s.EventsBuilder(), s.ListenReconcile)
	events.Listen(s.EventsBuilder(), s.ListenEvent)
	events.Listen(s.EventsBuilder(), s.ListenPredicted)

	// flush messages connections
	events.Listen(s.EventsBuilder(), func(loop.FrameEvent) {
		predictions := s.Prediction.FlushNewPredictions()
		if len(predictions) == 0 {
			return
		}
		for _, entity := range s.NetSync().Server().GetEntities() {
			conn, ok := s.Connection().Component().Get(entity)
			if !ok {
				continue
			}
			for _, prediction := range predictions {
				msg := clienttypes.PredictedEventDTO{
					ID:    prediction.ID,
					Event: prediction.Event,
				}
				s.Logger().Warn(conn.Conn().Send(msg))
			}
		}
	})

	// on connect reconcile
	events.Listen(s.EventsBuilder(), func(loop.FrameEvent) {
		for _, entity := range s.dirtySet.Get() {
			if _, ok := s.NetSync().Server().Get(entity); !ok {
				continue
			}
			conn, ok := s.Connection().Component().Get(entity)
			if !ok {
				continue
			}
			err := conn.Conn().Send(clienttypes.ReconcileDTO{})
			s.Logger().Warn(err)
		}
	})

	return s
}

// reconcile
func (s *Service) ListenReconcile(dto servertypes.ReconcileDTO) {
	conn := s.GetConnection()
	if conn == nil {
		return
	}

	predictions := s.Prediction.UndoPredictions()

	if dto.Error.Error() != nil {
		s.Logger().Log(dto.Error.Error())
		conn.Close()
		return
	}
	s.Record().ApplyUUIDState(s.Config().RecordConfig, dto.State)
	s.Loop().SyncToUnixNano(dto.TickUnixNano)

	s.Prediction.ReApplyPredictions(predictions)
}

func (s *Service) ListenEvent(dto servertypes.EventDTO) {
	s.Prediction.QueueServerEvent(prediction.ServerEvent{Event: dto.Event})
}

func (s *Service) ListenPredicted(dto servertypes.PredictionResponseDTO) {
	err := s.Prediction.ServerResponseToPrediction(dto.ID, dto.Accepted)
	s.Logger().Warn(err)
}

// private methods

func (s *Service) GetConnection() connection.Conn {
	if entities := s.NetSync().Server().GetEntities(); len(entities) == 1 {
		server := entities[0]
		comp, ok := s.Connection().Component().Get(server)
		if ok {
			return comp.Conn()
		}
	}
	// ensure prediction state is cleared
	return nil
}
