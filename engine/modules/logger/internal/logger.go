package internal

import (
	"engine/modules/logger"
	"errors"
	"os"
)

func FlattenError(err error) []error {
	toHandle := []error{err}
	errs := []error{}
	for len(toHandle) != 0 {
		err := toHandle[0]
		toHandle[0] = nil
		toHandle = toHandle[1:]
		if errWrap, ok := err.(interface{ Unwrap() []error }); ok {
			toHandle = append(errWrap.Unwrap(), toHandle...)
			continue
		}
		errs = append(errs, err)
	}
	return errs
}

// meta can be used in `errors.Is` to check data about error like severity or audience
type Config interface {
	// stages:
	// 1. formating error
	AddFormatHandler(func(meta, msg error) error)
	AddKeyedFormatHandler(error, func(meta, msg error) error)
	// 2. sending error either in user ui or on slack to developers
	AddDeliverHandler(func(meta, msg error))
	AddKeyedDeliverHandler(error, func(meta, msg error))
}

type Service interface {
	logger.Service
	Config
}

type service struct {
	KeyedFormatHandlers map[error]func(meta, msg error) error
	FormatHandlers      []func(meta, msg error) error

	KeyedDeliverHandlers map[error]func(meta, msg error)
	DeliverHandlers      []func(meta, msg error)
}

func NewService() Service {
	return &service{
		KeyedFormatHandlers: make(map[error]func(meta error, err error) error),
		FormatHandlers:      make([]func(meta error, msg error) error, 0),

		KeyedDeliverHandlers: make(map[error]func(meta error, err error)),
		DeliverHandlers:      make([]func(meta error, msg error), 0),
	}
}

func (s *service) Log(err error) {
	if err == nil {
		return
	}
	errs := FlattenError(err)

	meta := errs[:len(errs)-1]
	metaMerged := errors.Join(meta...)
	errMsg := errs[len(errs)-1]

	for _, handler := range s.FormatHandlers {
		errMsg = handler(metaMerged, errMsg)
	}
	for _, tag := range meta {
		if handler, ok := s.KeyedFormatHandlers[tag]; ok {
			errMsg = handler(metaMerged, errMsg)
		}
	}

	for _, handler := range s.DeliverHandlers {
		handler(metaMerged, errMsg)
	}
	for _, tag := range meta {
		if handler, ok := s.KeyedDeliverHandlers[tag]; ok {
			handler(metaMerged, errMsg)
		}
	}

	if errors.Is(metaMerged, logger.ErrFatal) {
		os.Exit(1)
	}
}

func (s *service) AddFormatHandler(handler func(meta, msg error) error) {
	s.FormatHandlers = append(s.FormatHandlers, handler)
}
func (s *service) AddKeyedFormatHandler(metaKey error, handler func(meta, msg error) error) {
	s.KeyedFormatHandlers[metaKey] = handler
}

func (s *service) AddDeliverHandler(handler func(meta, msg error)) {
	s.DeliverHandlers = append(s.DeliverHandlers, handler)
}
func (s *service) AddKeyedDeliverHandler(metaKey error, handler func(meta, msg error)) {
	s.KeyedDeliverHandlers[metaKey] = handler
}
