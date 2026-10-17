package prediction

import (
	"engine"
	"engine/modules/datastructures"
	"engine/modules/netsync/internal/config"
	"engine/modules/record"
	"errors"
	"reflect"
	"slices"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/ogiusek/null"
)

var (
	// on this error we need to reconcile
	ErrPredictionMachineFailed error = errors.New("prediction machine failed, need to reconcile")

	ErrAlreadyRecordingPrediction error = errors.New("already started prediction")
)

type EventType reflect.Type
type Event = any

type ServerEvent struct {
	Event Event
}
type ClientEvent struct {
	ID    ClientEventID
	Event Event
}
type Prediction struct {
	ID        ClientEventID
	Recording record.Recording
	Event     Event
}

type RecordedPrediction struct {
	ID        ClientEventID
	Recording record.RecordingID
	Event     Event
}
type ResponseToPrediction struct {
	ID       ClientEventID
	Accepted bool
}

func NewResponseToPrediction(id ClientEventID, accepted bool) ResponseToPrediction {
	return ResponseToPrediction{id, accepted}
}

func (r *ResponseToPrediction) Params() (ClientEventID, bool) {
	return r.ID, r.Accepted
}

// may return [ErrPredictionMachineFailed]
type Service interface {
	QueuedPredictions() int

	StartPrediction(Event) error
	StopPrediction() error

	ReApplyPredictions([]Prediction)
	UndoPredictions() []Prediction

	// removes [Prediction]
	// recreates queued [][Prediction] again with removed one
	QueueServerEvent(ServerEvent)

	// these should be flushed to server
	FlushNewPredictions() []ClientEvent

	// removes [Prediction]
	// recreates queued [][Prediction] again without removed one
	ServerResponseToPrediction(id ClientEventID, accepted bool) error
}

//

type service struct {
	World   engine.EngineWorld `inject:""`
	EventID IDService          `inject:""`
	Config  config.Config      `inject:""`

	newPredictions    []ClientEvent
	currentPrediction null.Nullable[RecordedPrediction]

	revertPrediction null.Nullable[record.RecordingID]

	// we use stable sparse array. not stable sparse array would cause swap and pop
	predictions datastructures.SparseArray[ClientEventID, Prediction]
}

func NewService(c ioc.Dic) Service {
	s := ioc.GetServices[*service](c)

	s.predictions = datastructures.NewStableSparseArray[ClientEventID, Prediction]()
	return s
}

func (s *service) ReApplyPredictions(predictions []Prediction) {
	for _, prediction := range predictions {
		s.EventID.QueueID(prediction.ID)
		events.EmitAny(s.World.Events(), prediction.Event)
	}
}

// returns predictions in order they should be applied to re-apply them
func (s *service) UndoPredictions() []Prediction {
	predictionIDs := slices.Clone(s.predictions.GetIndices())
	slices.Reverse(predictionIDs)
	predictions := make([]Prediction, 0, len(predictionIDs))
	for _, id := range predictionIDs {
		prediction, _ := s.predictions.Get(id)
		predictions = append(predictions, prediction)
		s.predictions.Remove(id)
	}
	recordings := make([]record.Recording, 0, len(predictions))
	for _, prediction := range predictions {
		recordings = append(recordings, prediction.Recording)
	}
	s.World.Record().Apply(s.Config.RecordConfig, recordings...)
	return predictions
}

//

func (s *service) QueuedPredictions() int { return s.predictions.Size() }

func (s *service) StartPrediction(event Event) error {
	if _, ok := s.currentPrediction.Ok(); ok {
		return ErrAlreadyRecordingPrediction
	}
	eventID, queued := s.EventID.GetID()
	if eventID == IgnoreEventID {
		return nil
	}
	// prediction overflow
	if _, ok := s.predictions.Get(eventID); ok {
		s.UndoPredictions()
		s.newPredictions = nil
		s.revertPrediction = null.New(s.World.Record().StartBackwardsRecording(s.Config.RecordConfig))
		return ErrPredictionMachineFailed
	}
	currentPrediction := RecordedPrediction{
		ID:        eventID,
		Recording: s.World.Record().StartBackwardsRecording(s.Config.RecordConfig),
		Event:     event,
	}
	s.currentPrediction = null.New(currentPrediction)
	if !queued && !s.Config.IsTickEvent(event) {
		s.newPredictions = append(s.newPredictions, ClientEvent{eventID, event})
	}

	return nil
}
func (s *service) StopPrediction() error {
	if id, ok := s.revertPrediction.Ok(); ok {
		recording, _ := s.World.Record().Stop(id)
		s.World.Record().Apply(s.Config.RecordConfig, recording)
		s.revertPrediction = null.Null[record.RecordingID]()
		return nil
	}
	prediction, ok := s.currentPrediction.Ok()
	if !ok {
		// this occurs when we used [IgnoreEventID]
		return nil
	}
	s.currentPrediction = null.Null[RecordedPrediction]()
	recording, ok := s.World.Record().Stop(prediction.Recording)
	if !ok {
		return ErrPredictionMachineFailed
	}
	s.predictions.Set(prediction.ID, Prediction{
		prediction.ID,
		recording,
		prediction.Event,
	})
	return nil
}

func (s *service) QueueServerEvent(event ServerEvent) {
	if !s.Config.IsTickEvent(event.Event) || len(s.predictions.GetIndices()) == 0 {
		predictions := s.UndoPredictions()
		s.EventID.QueueID(IgnoreEventID)
		events.EmitAny(s.World.Events(), event.Event)
		s.ReApplyPredictions(predictions)
		return
	}

	// if predicted in correct order
	firstPrediction := s.predictions.GetIndices()[0]
	if prediction, _ := s.predictions.Get(firstPrediction); prediction.Event == event.Event {
		s.predictions.Remove(prediction.ID)
		return
	}

	predictions := s.UndoPredictions()
	predictionIndex := slices.IndexFunc(predictions, func(prediction Prediction) bool {
		return reflect.DeepEqual(prediction.Event, event)
		// return prediction.Event == event
	})
	if predictionIndex != -1 {
		predictions = append(predictions[:predictionIndex], predictions[predictionIndex+1:]...)
	}
	s.EventID.QueueID(IgnoreEventID)
	events.EmitAny(s.World.Events(), event.Event)
	s.ReApplyPredictions(predictions)
}

func (s *service) FlushNewPredictions() []ClientEvent {
	newPredictions := s.newPredictions
	s.newPredictions = nil
	return newPredictions
}

// if received predicted event isn't in queue it means that we ran prediction reset
func (s *service) ServerResponseToPrediction(id ClientEventID, accepted bool) error {
	// this means that buffer was flushed
	if _, containsPrediction := s.predictions.Get(id); !containsPrediction {
		return ErrPredictionMachineFailed
	}

	predictionIndex := -1
	for i, iId := range s.predictions.GetIndices() {
		if iId == id {
			predictionIndex = i
			break
		}
	}
	if predictionIndex == -1 {
		return ErrPredictionMachineFailed
	}

	// if accepted response we just remove it
	if accepted && predictionIndex == 0 {
		s.predictions.Remove(id)
		return nil
	}

	// if rejected response we remove it
	undone := s.UndoPredictions()
	s.predictions.Remove(id)
	s.ReApplyPredictions(append(undone[:predictionIndex], undone[predictionIndex+1:]...))

	if accepted {
		return nil
	}

	// consider logging reason of rejection

	return nil
}
