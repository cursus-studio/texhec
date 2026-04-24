package texturearray

import "github.com/go-gl/gl/v4.5-core/gl"

type textureArray struct {
	texture     uint32
	imagesCount int
}

func (s *textureArray) Texture() uint32  { return s.texture }
func (s *textureArray) ImagesCount() int { return s.imagesCount }
func (s *textureArray) Release()         { gl.DeleteTextures(1, &s.texture) }
func (s *textureArray) Bind()            { gl.BindTexture(gl.TEXTURE_2D_ARRAY, s.texture) }
