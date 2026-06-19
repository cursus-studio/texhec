// this module is respinsible for defining unified events and components for rendering and
// for providing basic instancing renderer
package render

import (
	"engine/modules/ecs"
)

type RenderEvent struct {
	Camera ecs.EntityID
}

type Service interface {
	ecs.SystemRegister
	Renderer() ecs.SystemRegister

	Color() ecs.ComponentArray[ColorComponent]
	Mesh() ecs.ComponentArray[MeshComponent]
	Texture() ecs.ComponentArray[TextureComponent]
	TextureFrame() ecs.ComponentArray[TextureFrameComponent]

	Error() error
}
