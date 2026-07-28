package internal

import (
	"engine"
	"engine/modules/codec"
	"reflect"

	"github.com/ogiusek/ioc/v2"
)

type Builder interface {
	Register(any)
	Build() codec.Service
}

type builder struct {
	engine.EngineWorld `inject:""`
	types              map[reflect.Type]struct{}
}

func NewBuilder(c ioc.Dic) Builder {
	b := ioc.GetServices[*builder](c)
	b.types = make(map[reflect.Type]struct{})
	return b
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
	return newCodec(b.EngineWorld, types)
}
