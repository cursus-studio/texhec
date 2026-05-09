package textrenderer

import "engine/modules/graphics"

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
	VAO := s.NewVAO(VBO, nil)
	return layoutBatch{
		vao:           VAO,
		vertices:      VBO,
		verticesCount: int32(len(layout.Glyphs)),

		Layout: layout,
	}, nil
}

func (b layoutBatch) Release() {
	b.vao.Release()
}
