package textrenderer

import (
	"engine/modules/graphics"
	"engine/modules/text"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type fontBatch struct {
	glyphsWidth graphics.Buffer[float32]
	textures    graphics.TextureArray

	font text.Glyphs
}

func NewFontBatch(
	s graphics.Service,
	font text.Glyphs,
) (fontBatch, error) {
	textureArray, err := s.TextureArray().New(font.Images)
	if err != nil {
		return fontBatch{}, err
	}

	glyphsWidth := graphics.NewBuffer[float32](gl.SHADER_STORAGE_BUFFER, gl.DYNAMIC_DRAW, 0)

	for _, index := range font.GlyphsWidth.GetIndices() {
		width, _ := font.GlyphsWidth.Get(index)
		glyphsWidth.Set(int(index), width)
	}
	glyphsWidth.Flush()

	return fontBatch{
		glyphsWidth: glyphsWidth,
		textures:    textureArray,
		font:        font,
	}, nil
}

func (b *fontBatch) Release() {
	b.glyphsWidth.Release()
	b.textures.Release()
}
