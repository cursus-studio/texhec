package internal

import (
	"bytes"
	"encoding/gob"
	"engine/modules/codec"
	"engine/services/logger"
	"errors"
	"reflect"
)

type service struct {
	logger logger.Logger
}

func newCodec(
	logger logger.Logger,
	types []reflect.Type,
) codec.Service {
	for _, codecType := range types {
		name := codecType.String()
		value := reflect.New(codecType).Elem().Interface()
		gob.RegisterName(name, value)
	}
	return &service{logger}
}

func (s *service) Encode(model any) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(&model); err != nil {
		return nil, errors.Join(codec.ErrCannotEncodeType, err)
	}
	return buffer.Bytes(), nil
}

func (s *service) Decode(bytesToDecode []byte) (any, error) {
	var value any
	if err := gob.
		NewDecoder(bytes.NewReader(bytesToDecode)).
		Decode(&value); err != nil {
		return nil, errors.Join(codec.ErrCannotDecodeBytes, err)
	}
	if value == nil {
		return nil, errors.Join(codec.ErrCannotDecodeBytes)
	}
	return value, nil
}
