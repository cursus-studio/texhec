package vao

import (
	"engine/modules/graphics"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type vao struct {
	id  uint32
	vbo graphics.VBO
	ebo graphics.EBO
}

func NewVAO(VBO graphics.VBO, EBO graphics.EBO) graphics.VAO {
	var id uint32
	gl.GenVertexArrays(1, &id)

	gl.BindVertexArray(id)
	if VBO != nil {
		VBO.Configure()
	}
	if EBO != nil {
		EBO.Configure()
	}
	gl.BindVertexArray(0)

	return &vao{
		id:  id,
		vbo: VBO,
		ebo: EBO,
	}
}

func (vao *vao) ID() uint32 { return vao.id }

func (vao *vao) VBO() graphics.VBO { return vao.vbo }
func (vao *vao) EBO() graphics.EBO { return vao.ebo }

func (vao *vao) ReleaseVAO() {
	gl.DeleteVertexArrays(1, &vao.id)
}

func (vao *vao) Release() {
	vao.vbo.Release()
	if vao.ebo != nil {
		vao.ebo.Release()
	}
	vao.ReleaseVAO()
}

func (vao *vao) Bind() {
	gl.BindVertexArray(vao.id)
}
