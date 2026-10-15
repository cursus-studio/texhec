package internal

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/feature"
	"reflect"

	"github.com/ogiusek/ioc/v2"
)

type fieldService[FieldType any] struct {
	World   engine.EngineWorld `inject:""`
	Service ioc.Lazy[Service]  `inject:""`
	Config  Config             `inject:""`

	renderMark ecs.ComponentArray[feature.RenderComponent]
	render     ecs.ComponentArray[feature.FieldRenderComponent[FieldType]]
}

func NewFieldService[FieldType any](c ioc.Dic) feature.FieldService[FieldType] {
	s := ioc.GetServices[*fieldService[FieldType]](c)
	s.renderMark = ecs.GetComponentArray[feature.RenderComponent](s.World.World())
	s.render = ecs.GetComponentArray[feature.FieldRenderComponent[FieldType]](s.World.World())

	s.render.OnUpsert(func(entity ecs.EntityID) {
		s.renderMark.Set(entity, feature.NewRender())
	})

	s.Service().RegisterFieldService(reflect.TypeFor[FieldType](), s)
	return s
}

func (s *fieldService[FieldType]) RenderAny() ecs.AnyComponentArray {
	return s.render
}
func (s *fieldService[FieldType]) Render() ecs.ComponentArray[feature.FieldRenderComponent[FieldType]] {
	return s.render
}

func (s *fieldService[FieldType]) NewAnyRenderComponent(value any, fieldRender feature.FieldRender) any {
	val, ok := value.(FieldType)
	if !ok {
		return feature.NewFieldRender(val, feature.FieldRenderMissing)
	}
	return feature.NewFieldRender(val, fieldRender)
}
