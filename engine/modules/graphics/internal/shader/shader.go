package shader

import (
	"engine/modules/graphics"
	"fmt"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type shader struct {
	id uint32
}

func NewShader(shaderSource string, shaderType uint32) (graphics.Shader, error) {
	s, err := compileShader(shaderSource+"\x00", shaderType)
	if err != nil {
		return nil, fmt.Errorf("failed to compile vertex shader: %v", err)
	}

	return &shader{id: s}, nil
}

func (shader *shader) ID() uint32 { return shader.id }

func (shader *shader) Release() {
	gl.DeleteShader(shader.id)
}
