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
