package servertypes

import (
	"engine/modules/connection"
	"engine/modules/record"
	"engine/modules/uuid"
)

// server messages

type SendStateDTO struct {
	connection.MsgCtx
	TickUnixNano int64
	State        record.UUIDRecording
	Error        error
}

type SendChangeDTO struct {
	connection.MsgCtx
	EventID uuid.UUID
	Changes record.UUIDRecording
	Error   error
}

type TransparentEventDTO struct {
	connection.MsgCtx
	Event any
	Error error
}

type VerifyEventHappenDTO struct {
	connection.MsgCtx
	Event any
}

// Add here tick dto to notify about events order
// - use VerityEventHappenDTO and create a queue for it
// - extract prediction machine
