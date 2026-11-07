package clienttypes

import (
	"engine/modules/connection"
	"engine/modules/netsync/internal/prediction"
)

// client messages

type ReconcileDTO struct {
	connection.MsgCtx
}

type PredictedEventDTO struct {
	connection.MsgCtx
	ID    prediction.ClientEventID
	Event any
}
