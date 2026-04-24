package internal

import (
	"engine/modules/codec"
	"engine/services/logger"
	"reflect"
)

type Builder interface {
	Register(any)
	Build() codec.Service
}

type builder struct {
	logger logger.Logger
	types  map[reflect.Type]struct{}
}

func NewBuilder(logger logger.Logger) Builder {
	return &builder{
		logger: logger,
		types:  make(map[reflect.Type]struct{}),
	}
}

type GobTypesHook interface { // types to register
	GobTypes() []any
}

func (b *builder) Register(codecExample any) {
	codecType := reflect.TypeOf(codecExample)
	if codecType == nil {
		panic("WTF?")
	}
	if _, ok := b.types[codecType]; ok {
		return
	}
	b.types[codecType] = struct{}{}

	if h, ok := codecExample.(GobTypesHook); ok {
		for _, t := range h.GobTypes() {
			b.Register(t)
		}
	}
}

func (b *builder) Build() codec.Service {
	types := make([]reflect.Type, 0, len(b.types))
	for codecType := range b.types {
		types = append(types, codecType)
	}
	return newCodec(b.logger, types)
}
