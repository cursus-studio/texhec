package service

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/render"
	"engine/modules/render/internal/instancing"
	"engine/modules/render/internal/systems"
	"fmt"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`
	ecs.SystemRegister
	renderer ecs.SystemRegister

	colorArray        ecs.ComponentArray[render.ColorComponent]
	meshArray         ecs.ComponentArray[render.MeshComponent]
	textureArray      ecs.ComponentArray[render.TextureComponent]
	textureFrameArray ecs.ComponentArray[render.TextureFrameComponent]
}

func NewService(c ioc.Dic) render.Service {
	s := ioc.GetServices[*service](c)

	s.SystemRegister = ecs.NewSystemRegister(func() error {
		errs := ecs.RegisterSystems(
			systems.NewErrorLogger(c),
			systems.NewRenderSystem(c),
		)
		if len(errs) != 0 {
			return errs[0]
		}
		return nil
	})

	s.renderer = instancing.NewSystem(c)

	s.colorArray = ecs.GetComponentArray[render.ColorComponent](s.World())
	s.meshArray = ecs.GetComponentArray[render.MeshComponent](s.World())
	s.textureArray = ecs.GetComponentArray[render.TextureComponent](s.World())
	s.textureFrameArray = ecs.GetComponentArray[render.TextureFrameComponent](s.World())

	// defaults
	s.colorArray.SetEmpty(render.NewColor(mgl32.Vec4{1, 1, 1, 1}))
	// no default mesh
	// no default texture
	s.textureFrameArray.SetEmpty(render.NewTextureFrame(0))

	return s
}

//

var GlErrorStrings = map[uint32]string{
	gl.NO_ERROR:                      "GL_NO_ERROR",
	gl.INVALID_ENUM:                  "GL_INVALID_ENUM",
	gl.INVALID_VALUE:                 "GL_INVALID_VALUE",
	gl.INVALID_OPERATION:             "GL_INVALID_OPERATION",
	gl.STACK_OVERFLOW:                "GL_STACK_OVERFLOW",
	gl.STACK_UNDERFLOW:               "GL_STACK_UNDERFLOW",
	gl.OUT_OF_MEMORY:                 "GL_OUT_OF_MEMORY",
	gl.INVALID_FRAMEBUFFER_OPERATION: "GL_INVALID_FRAMEBUFFER_OPERATION",
	gl.CONTEXT_LOST:                  "GL_CONTEXT_LOST",
	// gl.TABLE_TOO_LARGE:               "GL_TABLE_TOO_LARGE", // Less common in modern GL
}

func (s *service) Renderer() ecs.SystemRegister { return s.renderer }

func (s *service) Color() ecs.ComponentArray[render.ColorComponent] {
	return s.colorArray
}
func (s *service) Mesh() ecs.ComponentArray[render.MeshComponent] {
	return s.meshArray
}
func (s *service) Texture() ecs.ComponentArray[render.TextureComponent] {
	return s.textureArray
}
func (s *service) TextureFrame() ecs.ComponentArray[render.TextureFrameComponent] {
	return s.textureFrameArray
}

func (*service) Error() error {
	if glErr := gl.GetError(); glErr != gl.NO_ERROR {
		return fmt.Errorf("opengl error: %x %s", glErr, GlErrorStrings[glErr])
	}
	return nil
}
