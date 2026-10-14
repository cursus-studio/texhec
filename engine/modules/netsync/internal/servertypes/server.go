package servertypes

import (
	"engine/modules/connection"
	"engine/modules/netsync/internal/prediction"
	"engine/modules/record"
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

type ReconcileDTO struct {
	connection.MsgCtx
	TickUnixNano int64
	State        record.UUIDRecording
	Error        Error
}

type EventDTO struct {
	connection.MsgCtx
	Event any
}

type PredictionResponseDTO struct {
	connection.MsgCtx
	ID       prediction.ClientEventID
	Accepted bool
}
