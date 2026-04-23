package window

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/veandco/go-sdl2/sdl"
)

type MousePos struct{ X, Y int32 }

func NewMousePos(x, y int32) MousePos  { return MousePos{x, y} }
func (p *MousePos) Elem() (x, y int32) { return p.X, p.Y }

//

type Service interface {
	NormalizeMousePos(MousePos) mgl32.Vec2
	GetMousePos() MousePos
	Window() *sdl.Window
	Ctx() sdl.GLContext
}
