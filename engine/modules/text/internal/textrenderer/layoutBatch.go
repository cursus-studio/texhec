package textrenderer

import (
	"engine/modules/graphics"
	"fmt"
	"math"
)

type layoutBatch struct {
	vao           graphics.VAO
	vertices      graphics.VBOSetter[Glyph]
	verticesCount int32

	Layout Layout
}

func NewLayoutBatch(
	s graphics.Service,
	v graphics.VBOFactory[Glyph],
	layout Layout,
) (layoutBatch, error) {
	VBO := v()
	if err := VBO.SetVertices(layout.Glyphs); err != nil {
		return layoutBatch{}, err
	}
	glyphsLen := len(layout.Glyphs)
	if glyphsLen < 0 || glyphsLen > math.MaxInt32 {
		return layoutBatch{}, fmt.Errorf("there cannot be more then %v glyphs", math.MaxInt32)
	}
	VAO := s.NewVAO(VBO, nil)
	return layoutBatch{
		vao:           VAO,
		vertices:      VBO,
		verticesCount: int32(glyphsLen),

		Layout: layout,
	}, nil
}

func (b layoutBatch) Release() {
	b.vao.Release()
}
