package text

import (
	"engine/services/ecs"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	BreakNone uint8 = iota
	BreakWord
	BreakAny
)

// this is required to render text
// every other component is optional and has default value
type TextComponent struct {
	Text string
}
type AlignComponent struct {
	// value between 0 and 1 where 0 means aligned to left and 1 aligned to right
	Vertical, Horizontal float32 // default is 0
}
type ColorComponent struct {
	Color mgl32.Vec4
}
type FontFamilyComponent struct {
	FontFamily ecs.EntityID
}
type FontSizeComponent struct {
	FontSize uint
}
type BreakComponent struct {
	Break uint8
}

func NewText(text string) TextComponent { return TextComponent{text} }
func NewAlign(vertical, horizontal float32) AlignComponent {
	return AlignComponent{vertical, horizontal}
}
func NewColor(color mgl32.Vec4) ColorComponent { return ColorComponent{color} }
func NewFontFamily(fontFamily ecs.EntityID) FontFamilyComponent {
	return FontFamilyComponent{fontFamily}
}
func NewFontSize(fontSize uint) FontSizeComponent { return FontSizeComponent{fontSize} }
func NewBreak(b uint8) BreakComponent             { return BreakComponent{b} }

//

type SystemRenderer ecs.SystemRegister

type Service interface {
	Break() ecs.ComponentsArray[BreakComponent]
	Content() ecs.ComponentsArray[TextComponent]
	Align() ecs.ComponentsArray[AlignComponent]
	Color() ecs.ComponentsArray[ColorComponent]
	FontFamily() ecs.ComponentsArray[FontFamilyComponent]
	FontSize() ecs.ComponentsArray[FontSizeComponent]

	AddDirtySet(ecs.DirtySet)
}
