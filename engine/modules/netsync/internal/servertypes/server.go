package servertypes

import (
	"engine/modules/record"
	"engine/modules/uuid"
)

// server messages

type SendStateDTO struct {
	TickUnixNano int64
	State        record.UUIDRecording
	Error        error
}

type SendChangeDTO struct {
	EventID uuid.UUID
	Changes record.UUIDRecording
	Error   error
}

type TransparentEventDTO struct {
	Event any
	Error error
}
