package render

import (
	"engine/services/ecs"
)

type RenderEvent struct {
	Camera ecs.EntityID
}

type Service interface {
	ecs.SystemRegister
	Renderer() ecs.SystemRegister

	Color() ecs.ComponentsArray[ColorComponent]
	Mesh() ecs.ComponentsArray[MeshComponent]
	Texture() ecs.ComponentsArray[TextureComponent]
	TextureFrame() ecs.ComponentsArray[TextureFrameComponent]

	Error() error
}
