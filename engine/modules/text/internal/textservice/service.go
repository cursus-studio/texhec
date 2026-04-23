package textservice

import (
	"engine"
	"engine/modules/text"
	"engine/services/ecs"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`

	breakArray      ecs.ComponentsArray[text.BreakComponent]
	textArray       ecs.ComponentsArray[text.TextComponent]
	alignArray      ecs.ComponentsArray[text.AlignComponent]
	colorArray      ecs.ComponentsArray[text.ColorComponent]
	fontFamilyArray ecs.ComponentsArray[text.FontFamilyComponent]
	fontSizeArray   ecs.ComponentsArray[text.FontSizeComponent]
}

func NewService(c ioc.Dic) text.Service {
	s := ioc.GetServices[*service](c)
	s.breakArray = ecs.GetComponentsArray[text.BreakComponent](s.World())
	s.textArray = ecs.GetComponentsArray[text.TextComponent](s.World())
	s.alignArray = ecs.GetComponentsArray[text.AlignComponent](s.World())
	s.colorArray = ecs.GetComponentsArray[text.ColorComponent](s.World())
	s.fontFamilyArray = ecs.GetComponentsArray[text.FontFamilyComponent](s.World())
	s.fontSizeArray = ecs.GetComponentsArray[text.FontSizeComponent](s.World())

	s.breakArray.SetEmpty(text.NewBreak(text.BreakWord))
	s.alignArray.SetEmpty(text.NewAlign(0, 0))
	s.colorArray.SetEmpty(text.NewColor(mgl32.Vec4{1, 1, 1, 1}))
	s.fontSizeArray.SetEmpty(text.NewFontSize(16))
	return s
}

func (t *service) Break() ecs.ComponentsArray[text.BreakComponent]  { return t.breakArray }
func (t *service) Content() ecs.ComponentsArray[text.TextComponent] { return t.textArray }
func (t *service) Align() ecs.ComponentsArray[text.AlignComponent]  { return t.alignArray }
func (t *service) Color() ecs.ComponentsArray[text.ColorComponent]  { return t.colorArray }
func (t *service) FontFamily() ecs.ComponentsArray[text.FontFamilyComponent] {
	return t.fontFamilyArray
}
func (t *service) FontSize() ecs.ComponentsArray[text.FontSizeComponent] { return t.fontSizeArray }

func (t *service) AddDirtySet(set ecs.DirtySet) {
	t.breakArray.AddDirtySet(set)
	t.alignArray.AddDirtySet(set)
	t.colorArray.AddDirtySet(set)
	t.fontFamilyArray.AddDirtySet(set)
	t.fontSizeArray.AddDirtySet(set)
}
