package clienttypes

import (
	"engine/modules/connection"
	"engine/modules/uuid"
)

// types

type PredictedEvent struct {
	connection.MsgCtx
	ID    uuid.UUID
	Event any
}

// client messages

type FetchStateDTO struct {
	connection.MsgCtx
}

type EmitEventDTO PredictedEvent

type TransparentEventDTO struct {
	connection.MsgCtx
	Event any
}
