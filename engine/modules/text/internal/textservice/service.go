package textservice

import (
	"engine"
	"engine/modules/text"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`

	breakArray      ecs.ComponentsArray[text.BreakComponent]
	textArray       ecs.ComponentsArray[text.TextComponent]
	textAlignArray  ecs.ComponentsArray[text.TextAlignComponent]
	textColorArray  ecs.ComponentsArray[text.TextColorComponent]
	fontFamilyArray ecs.ComponentsArray[text.FontFamilyComponent]
	fontSizeArray   ecs.ComponentsArray[text.FontSizeComponent]
}

func NewService(c ioc.Dic) text.Service {
	s := ioc.GetServices[*service](c)
	s.breakArray = ecs.GetComponentsArray[text.BreakComponent](s.World())
	s.textArray = ecs.GetComponentsArray[text.TextComponent](s.World())
	s.textAlignArray = ecs.GetComponentsArray[text.TextAlignComponent](s.World())
	s.textColorArray = ecs.GetComponentsArray[text.TextColorComponent](s.World())
	s.fontFamilyArray = ecs.GetComponentsArray[text.FontFamilyComponent](s.World())
	s.fontSizeArray = ecs.GetComponentsArray[text.FontSizeComponent](s.World())
	return s
}

func (t *service) Break() ecs.ComponentsArray[text.BreakComponent]     { return t.breakArray }
func (t *service) Content() ecs.ComponentsArray[text.TextComponent]    { return t.textArray }
func (t *service) Align() ecs.ComponentsArray[text.TextAlignComponent] { return t.textAlignArray }
func (t *service) Color() ecs.ComponentsArray[text.TextColorComponent] { return t.textColorArray }
func (t *service) FontFamily() ecs.ComponentsArray[text.FontFamilyComponent] {
	return t.fontFamilyArray
}
func (t *service) FontSize() ecs.ComponentsArray[text.FontSizeComponent] { return t.fontSizeArray }

func (t *service) AddDirtySet(set ecs.DirtySet) {
	t.breakArray.AddDirtySet(set)
	t.textAlignArray.AddDirtySet(set)
	t.textColorArray.AddDirtySet(set)
	t.fontFamilyArray.AddDirtySet(set)
	t.fontSizeArray.AddDirtySet(set)
}
