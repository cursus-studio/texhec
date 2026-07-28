package ebo

import (
	"engine/modules/graphics"
	"fmt"
	"math"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type ebo struct {
	id  uint32
	len int32
}

func NewEBO() graphics.EBO {
	var id uint32
	gl.GenBuffers(1, &id)
	return &ebo{
		id:  id,
		len: 0,
	}
}

func (ebo *ebo) ID() uint32 { return ebo.id }
func (ebo *ebo) Len() int32 { return ebo.len }

func (ebo *ebo) Configure() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo.id)
}

func (ebo *ebo) Release() {
	gl.DeleteBuffers(1, &ebo.id)
}

func (ebo *ebo) SetIndices(indices []graphics.Index) error {
	indicesLen := len(indices)
	if indicesLen < 0 || indicesLen > math.MaxInt32 {
		return fmt.Errorf("indices length %d exceeds maximum allowed size (%d)", indicesLen, math.MaxInt32)
	}
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo.id)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, indicesLen*4, gl.Ptr(indices), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	ebo.len = int32(indicesLen)
	return nil
}
