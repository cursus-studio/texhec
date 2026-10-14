package prediction_test

import (
	"engine/modules/netsync/internal/prediction"
	"testing"

	"github.com/ogiusek/events"
)

func TestPredictions(t *testing.T) {
	w := NewWorld(t)

	lastCompValue := Component(0)
	w.AssertState(NewValue(int(lastCompValue)))

	w.PredictionID.QueueID(prediction.IgnoreEventID)
	events.Emit(w.Events(), IncreaseToNextOddEvent{})
	w.AssertWorldState(lastCompValue.NextOdd(), Queued(0), Flushed{})

	lastCompValue = lastCompValue.NextOdd().NextEven()

	w.PredictionID.QueueID(prediction.IgnoreEventID)
	events.Emit(w.Events(), IncreaseToNextEvenEvent{})
	w.AssertWorldState(lastCompValue, Queued(0), Flushed{})

	// client predicts ticks and server sends them and [IncreaseByTenEvent] should be applied only once per its emission
	{
		events.Emit(w.Events(), IncreaseByTenEvent{})
		w.AssertWorldState(lastCompValue.IncrementByTen(), Queued(1), Flushed{})

		events.Emit(w.Events(), IncreaseByTenEvent{})
		w.AssertWorldState(lastCompValue.IncrementByTen().IncrementByTen(), Queued(2), Flushed{})

		lastCompValue = lastCompValue.IncrementByTen().IncrementByTen()

		w.Prediction.QueueServerEvent(prediction.ServerEvent{Event: IncreaseByTenEvent{}})
		w.AssertWorldState(lastCompValue, Queued(1), Flushed{})

		w.Prediction.QueueServerEvent(prediction.ServerEvent{Event: IncreaseByTenEvent{}})
		w.AssertWorldState(lastCompValue, Queued(0), Flushed{})
	}

	// client sends [prediction.ClientEvent] while server executed [prediction.ServerEvent] :
	// - client emits [prediction.ClientEvent] and sends it to [prediction.Service.QueuePrediction]
	//   creates [prediction.Prediction]
	// - server sends [prediction.ServerEvent] to [prediction.Service.QueueServerEvent]
	//   removes [prediction.Prediction]
	//   creates it again [prediction.Prediction]
	// - server mirrors [prediction.ClientEvent] and send it to [prediction.Service.QueueServerEvent]
	//   removes [prediction.Prediction]
	{
		events.Emit(w.Events(), IncreaseToNextOddEvent{})
		predictionID := w.AssertWorldState(
			lastCompValue.NextOdd(), Queued(1), Flushed{IncreaseToNextOddEvent{}})[0]

		w.Prediction.QueueServerEvent(prediction.ServerEvent{Event: IncreaseToNextEvenEvent{}})
		w.AssertWorldState(
			lastCompValue.NextEven().NextOdd(), Queued(1), Flushed{})

		lastCompValue = lastCompValue.NextEven().NextOdd()

		err := w.Prediction.ServerResponseToPrediction(predictionID, true)
		w.AssertSuccess(err)
		w.AssertWorldState(lastCompValue, Queued(0), Flushed{})
	}

	// client prediction in between ticks
	{
		events.Emit(w.Events(), IncreaseByTenEvent{})
		w.AssertWorldState(lastCompValue.IncrementByTen(), Queued(1), Flushed{})

		events.Emit(w.Events(), IncreaseToNextOddEvent{})
		predictionID := w.AssertWorldState(
			lastCompValue.IncrementByTen().NextOdd(), Queued(2), Flushed{IncreaseToNextOddEvent{}})[0]

		events.Emit(w.Events(), IncreaseByTenEvent{})
		w.AssertWorldState(lastCompValue.IncrementByTen().NextOdd().IncrementByTen(), Queued(3), Flushed{})

		w.Prediction.QueueServerEvent(prediction.ServerEvent{Event: IncreaseByTenEvent{}})
		w.AssertWorldState(lastCompValue.IncrementByTen().NextOdd().IncrementByTen(), Queued(2), Flushed{})

		err := w.Prediction.ServerResponseToPrediction(predictionID, true)
		w.AssertSuccess(err)
		w.AssertWorldState(lastCompValue.IncrementByTen().NextOdd().IncrementByTen(), Queued(1), Flushed{})

		lastCompValue = lastCompValue.IncrementByTen().NextOdd().IncrementByTen()

		w.Prediction.QueueServerEvent(prediction.ServerEvent{Event: IncreaseByTenEvent{}})
		w.AssertWorldState(lastCompValue, Queued(0), Flushed{})
	}

	// client sends [prediction.ClientEvent] and server rejects it :
	// - client emits [prediction.ClientEvent] and sends it to [prediction.Service.QueuePrediction]
	//   creates [prediction.Prediction]
	// - server sends [prediction.Service.PredictionResponse] with rejected
	//   removes [prediction.Prediction]
	//   creates again other [prediction.Prediction]
	{
		events.Emit(w.Events(), IncreaseToNextEvenEvent{})
		serverResponse := w.AssertWorldState(
			lastCompValue.NextEven(), Queued(1), Flushed{IncreaseToNextEvenEvent{}})[0]

		w.Prediction.QueueServerEvent(prediction.ServerEvent{Event: IncreaseToNextEvenEvent{}})
		w.AssertWorldState(
			lastCompValue.NextEven().NextEven(), Queued(1), Flushed{})

		lastCompValue = lastCompValue.NextEven()

		err := w.Prediction.ServerResponseToPrediction(serverResponse, false)
		w.AssertSuccess(err)
		w.AssertWorldState(lastCompValue, Queued(0), Flushed{})
	}

	// server sends [prediction.ServerEvent] :
	// - server sends [prediction.ServerEvent] and send it to [prediction.Service.QueueServerEvent]
	{
		lastCompValue = lastCompValue.NextEven()

		w.Prediction.QueueServerEvent(prediction.ServerEvent{Event: IncreaseToNextEvenEvent{}})
		w.AssertWorldState(
			lastCompValue, Queued(0), Flushed{})
	}
}
