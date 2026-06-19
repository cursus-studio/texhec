package textservice

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/text"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	renderer           ecs.SystemRegister

	breakArray      ecs.ComponentArray[text.BreakComponent]
	textArray       ecs.ComponentArray[text.TextComponent]
	alignArray      ecs.ComponentArray[text.AlignComponent]
	colorArray      ecs.ComponentArray[text.ColorComponent]
	fontFamilyArray ecs.ComponentArray[text.FontFamilyComponent]
	fontSizeArray   ecs.ComponentArray[text.FontSizeComponent]
}

func NewService(c ioc.Dic, register ecs.SystemRegister) text.Service {
	s := ioc.GetServices[*service](c)
	s.renderer = register
	s.breakArray = ecs.GetComponentArray[text.BreakComponent](s.World())
	s.textArray = ecs.GetComponentArray[text.TextComponent](s.World())
	s.alignArray = ecs.GetComponentArray[text.AlignComponent](s.World())
	s.colorArray = ecs.GetComponentArray[text.ColorComponent](s.World())
	s.fontFamilyArray = ecs.GetComponentArray[text.FontFamilyComponent](s.World())
	s.fontSizeArray = ecs.GetComponentArray[text.FontSizeComponent](s.World())

	s.breakArray.SetEmpty(text.NewBreak(text.BreakWord))
	s.alignArray.SetEmpty(text.NewAlign(0, 0))
	s.colorArray.SetEmpty(text.NewColor(mgl32.Vec4{1, 1, 1, 1}))
	s.fontSizeArray.SetEmpty(text.NewFontSize(16))
	return s
}

func (s *service) Renderer() ecs.SystemRegister { return s.renderer }

func (s *service) Break() ecs.ComponentArray[text.BreakComponent]  { return s.breakArray }
func (s *service) Content() ecs.ComponentArray[text.TextComponent] { return s.textArray }
func (s *service) Align() ecs.ComponentArray[text.AlignComponent]  { return s.alignArray }
func (s *service) Color() ecs.ComponentArray[text.ColorComponent]  { return s.colorArray }
func (s *service) FontFamily() ecs.ComponentArray[text.FontFamilyComponent] {
	return s.fontFamilyArray
}
func (s *service) FontSize() ecs.ComponentArray[text.FontSizeComponent] { return s.fontSizeArray }

func (s *service) AddDirtySet(set ecs.DirtySet) {
	s.breakArray.AddDirtySet(set)
	s.alignArray.AddDirtySet(set)
	s.colorArray.AddDirtySet(set)
	s.fontFamilyArray.AddDirtySet(set)
	s.fontSizeArray.AddDirtySet(set)
}
