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
	PredictionID       prediction.IDService    `inject:""`
	Config             ioc.Lazy[config.Config] `inject:""`
}

func NewService(c ioc.Dic) *Service {
	s := ioc.GetServices[*Service](c)

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
			conn, ok := s.Connection().Connection().Get(entity)
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
	onConnAdd := func(entity ecs.EntityID) {
		if _, ok := s.NetSync().Server().Get(entity); !ok {
			return
		}
		conn, ok := s.Connection().Connection().Get(entity)
		if !ok {
			return
		}
		err := conn.Conn().Send(clienttypes.ReconcileDTO{})
		s.Logger().Warn(err)
	}
	s.NetSync().Server().OnUpsert(onConnAdd)
	s.Connection().Connection().OnUpsert(onConnAdd)

	return s
}

// reconcile
func (s *Service) ListenReconcile(dto servertypes.ReconcileDTO) {
	conn := s.GetConnection()
	if conn == nil {
		return
	}

	_ = s.Prediction.UndoPredictions()

	if dto.Error.Error() != nil {
		s.Logger().Log(dto.Error.Error())
		conn.Close()
		return
	}
	s.Loop().SyncToUnixNano(dto.TickUnixNano)
	s.Record().ApplyUUIDState(s.Config().RecordConfig, dto.TickState)
	for _, event := range dto.PostTickEvents {
		s.PredictionID.QueueID(prediction.IgnoreEventID)
		events.EmitAny(s.Events(), event)
	}
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
		comp, ok := s.Connection().Connection().Get(server)
		if ok {
			return comp.Conn()
		}
	}
	// ensure prediction state is cleared
	return nil
}
