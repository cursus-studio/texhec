package prediction

import (
	"engine"
	"engine/modules/netsync/internal/config"

	"github.com/ogiusek/ioc/v2"
)

// client event IDs aren't unique
// they should only be unique in context of maximum predicted client event IDs
type ClientEventID uint8

const IgnoreEventID ClientEventID = 0

type IDService interface {
	GetID() (id ClientEventID, queued bool)
	QueueID(ClientEventID)
}

type eventIDService struct {
	World     engine.EngineWorld      `inject:""`
	Config    ioc.Lazy[config.Config] `inject:""`
	minID     ClientEventID
	id        ClientEventID
	queuedIDs []ClientEventID
}

func NewEventIDService(c ioc.Dic) IDService {
	s := ioc.GetServices[*eventIDService](c)
	maxID := ^ClientEventID(0)
	s.minID = max(1, maxID-ClientEventID(s.Config().MaxPredictions-1))
	s.id = s.minID
	return s
}

func (s *eventIDService) GetID() (id ClientEventID, queued bool) {
	if len(s.queuedIDs) > 0 {
		id = s.queuedIDs[0]
		s.queuedIDs = s.queuedIDs[1:]
		return id, true
	}
	id = s.id
	s.id = max(s.minID, s.id+1)
	return id, false
}
func (s *eventIDService) QueueID(id ClientEventID) {
	s.queuedIDs = append(s.queuedIDs, id)
}
