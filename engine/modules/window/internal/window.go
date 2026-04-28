package internal

import (
	"engine"
	"engine/modules/logger"
	"engine/modules/window"
	"errors"
	"sync"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type service struct {
	engine.EngineWorld `inject:""`
	window             *sdl.Window
	context            sdl.GLContext
	once               sync.Once
}

func NewService(c ioc.Dic) window.Service {
	s := ioc.GetServices[*service](c)
	return s
}

func (s *service) init() {
	s.once.Do(func() {
		var err error
		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
			s.Logger().Log(errors.Join(logger.ErrFatal, err))
			return
		}

		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4 /* 3 */)
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1 /* 3 */)
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
		_ = sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1) // Essential for GLSwap
		_ = sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 24)  // Good practice for depth testing

		// audio
		if err := mix.OpenAudio(48000, sdl.AUDIO_F32SYS, 2, 1024); err != nil {
			s.Logger().Log(errors.Join(logger.ErrFatal, err))
			return
		}

		// window and opengl
		s.window, err = sdl.CreateWindow(
			"ENGINE",
			sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			800, 600,
			sdl.WINDOW_SHOWN|sdl.WINDOW_OPENGL,
		)
		if err != nil {
			s.Logger().Log(errors.Join(logger.ErrFatal, err))
			return
		}

		s.context, err = s.window.GLCreateContext()
		if err != nil {
			s.Logger().Log(errors.Join(logger.ErrFatal, err))
			return
		}
		if err := gl.Init(); err != nil {
			s.Logger().Log(errors.Join(logger.ErrFatal, err))
			return
		}
		if err := s.window.GLMakeCurrent(s.context); err != nil {
			s.Logger().Log(errors.Join(logger.ErrFatal, err))
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

func (s *service) NormalizeMousePos(mousePos window.MousePos) mgl32.Vec2 {
	s.init()
	x, y := mousePos.Elem()
	w, h := s.Window().GetSize()
	return mgl32.Vec2{
		(2*float32(x)/float32(w) - 1),
		-(2*float32(y)/float32(h) - 1),
	}
}
func (s *service) GetMousePos() window.MousePos {
	s.init()
	x, y, _ := sdl.GetMouseState()
	return window.NewMousePos(x, y)
}
func (s *service) Window() *sdl.Window {
	s.init()
	return s.window
}
func (s *service) Ctx() sdl.GLContext {
	s.init()
	return s.context
}
