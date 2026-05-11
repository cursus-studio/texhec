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
	renderer           ecs.SystemRegister

	breakArray      ecs.ComponentsArray[text.BreakComponent]
	textArray       ecs.ComponentsArray[text.TextComponent]
	alignArray      ecs.ComponentsArray[text.AlignComponent]
	colorArray      ecs.ComponentsArray[text.ColorComponent]
	fontFamilyArray ecs.ComponentsArray[text.FontFamilyComponent]
	fontSizeArray   ecs.ComponentsArray[text.FontSizeComponent]
}

func NewService(c ioc.Dic, register ecs.SystemRegister) text.Service {
	s := ioc.GetServices[*service](c)
	s.renderer = register
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

func (s *service) Renderer() ecs.SystemRegister { return s.renderer }

func (s *service) Break() ecs.ComponentsArray[text.BreakComponent]  { return s.breakArray }
func (s *service) Content() ecs.ComponentsArray[text.TextComponent] { return s.textArray }
func (s *service) Align() ecs.ComponentsArray[text.AlignComponent]  { return s.alignArray }
func (s *service) Color() ecs.ComponentsArray[text.ColorComponent]  { return s.colorArray }
func (s *service) FontFamily() ecs.ComponentsArray[text.FontFamilyComponent] {
	return s.fontFamilyArray
}
func (s *service) FontSize() ecs.ComponentsArray[text.FontSizeComponent] { return s.fontSizeArray }

func (s *service) AddDirtySet(set ecs.DirtySet) {
	s.breakArray.AddDirtySet(set)
	s.alignArray.AddDirtySet(set)
	s.colorArray.AddDirtySet(set)
	s.fontFamilyArray.AddDirtySet(set)
	s.fontSizeArray.AddDirtySet(set)
}
