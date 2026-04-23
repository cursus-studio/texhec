package window

import (
	"engine/services/logger"
	"sync"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type MousePos struct{ X, Y int32 }

func NewMousePos(x, y int32) MousePos  { return MousePos{x, y} }
func (p *MousePos) Elem() (x, y int32) { return p.X, p.Y }

type Api interface {
	NormalizeMousePos(MousePos) mgl32.Vec2
	GetMousePos() MousePos
	Window() *sdl.Window
	Ctx() sdl.GLContext
}

type api struct {
	Logger  logger.Logger `inject:""`
	window  *sdl.Window
	context sdl.GLContext
	once    sync.Once
}

func newApi(c ioc.Dic) Api {
	s := ioc.GetServices[*api](c)
	return s
}

func (api *api) init() {
	api.once.Do(func() {
		var err error
		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
			api.Logger.Fatal(err)
			return
		}

		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4 /* 3 */)
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1 /* 3 */)
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
		_ = sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1) // Essential for GLSwap
		_ = sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 24)  // Good practice for depth testing

		// audio
		if err := mix.OpenAudio(48000, sdl.AUDIO_F32SYS, 2, 1024); err != nil {
			api.Logger.Fatal(err)
			return
		}

		// window and opengl
		api.window, err = sdl.CreateWindow(
			"",
			sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			800, 600,
			sdl.WINDOW_SHOWN|sdl.WINDOW_OPENGL,
		)
		if err != nil {
			api.Logger.Fatal(err)
			return
		}

		api.context, err = api.window.GLCreateContext()
		if err != nil {
			api.Logger.Fatal(err)
			return
		}
		if err := gl.Init(); err != nil {
			api.Logger.Fatal(err)
			return
		}
		if err := api.window.GLMakeCurrent(api.context); err != nil {
			api.Logger.Fatal(err)
			return
		}
		_ = sdl.GLSetSwapInterval(0)

		// render settings
		gl.Enable(gl.CULL_FACE)
		gl.CullFace(gl.FRONT)
		gl.FrontFace(gl.CCW)

		gl.Enable(gl.DEPTH_TEST)
		gl.DepthFunc(gl.LEQUAL) // less or equal

		gl.Enable(gl.BLEND)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	})
}

func (api *api) NormalizeMousePos(mousePos MousePos) mgl32.Vec2 {
	api.init()
	x, y := mousePos.Elem()
	w, h := api.Window().GetSize()
	return mgl32.Vec2{
		(2*float32(x)/float32(w) - 1),
		-(2*float32(y)/float32(h) - 1),
	}
}
func (api *api) GetMousePos() MousePos {
	x, y, _ := sdl.GetMouseState()
	return NewMousePos(x, y)
}
func (api *api) Window() *sdl.Window {
	api.init()
	return api.window
}
func (api *api) Ctx() sdl.GLContext {
	api.init()
	return api.context
}
