package servertypes

import (
	"engine/modules/connection"
	"engine/modules/record"
	"engine/modules/uuid"
	"errors"
)

type Error struct {
	Msg string
}

func NewError(err error) Error {
	if err == nil {
		return Error{}
	}
	return Error{err.Error()}
}

func (err Error) Error() error {
	if err.Msg == "" {
		return nil
	}
	return errors.New(err.Msg)
}

// server messages

type SendStateDTO struct {
	connection.MsgCtx
	TickUnixNano int64
	State        record.UUIDRecording
	Error        Error
}

type SendChangeDTO struct {
	connection.MsgCtx
	EventID uuid.UUID
	Changes record.UUIDRecording
	Error   Error
}

type TransparentEventDTO struct {
	connection.MsgCtx
	Event any
	Error Error
}

type VerifyEventHappenDTO struct {
	connection.MsgCtx
	Event any
}

// Add here tick dto to notify about events order
// - use VerityEventHappenDTO and create a queue for it
// - extract prediction machine
