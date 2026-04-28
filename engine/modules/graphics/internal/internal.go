package internal

import (
	"engine/modules/graphics"
	"engine/modules/graphics/internal/ebo"
	"engine/modules/graphics/internal/img"
	"engine/modules/graphics/internal/program"
	"engine/modules/graphics/internal/shader"
	"engine/modules/graphics/internal/texture"
	"engine/modules/graphics/internal/texturearray"
	"engine/modules/graphics/internal/vao"
	"image"
)

type service struct {
	textureArrayFactory graphics.TextureArrayFactory
	textureFactory      graphics.TextureFactory
}

func NewService() graphics.Service {
	return &service{
		texturearray.NewFactory(),
		texture.NewFactory(),
	}
}

func (s *service) NewImage(image image.Image) graphics.Image  { return img.NewImage(image) }
func (s *service) TextureArray() graphics.TextureArrayFactory { return s.textureArrayFactory }
func (s *service) Texture() graphics.TextureFactory           { return s.textureFactory }

func (s *service) NewVAO(VBO graphics.VBO, EBO graphics.EBO) graphics.VAO {
	return vao.NewVAO(VBO, EBO)
}
func (s *service) NewEBO() graphics.EBO {
	return ebo.NewEBO()
}

func (s *service) NewProgram(p uint32, parameters []graphics.Parameter) (graphics.Program, error) {
	return program.NewProgram(p, parameters)
}
func (s *service) NewShader(shaderSource string, shaderType uint32) (graphics.Shader, error) {
	return shader.NewShader(shaderSource, shaderType)
}
