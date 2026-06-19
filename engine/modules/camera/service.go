// this module is responsible for cameras. Responsibilities:
// - projections
// - sets [size](/engine/modules/transform/readme/README.md#type-SizeComponent) for objects with projections
// - shots rays
// - parses cameras to mgl32.Mat4
package camera

import (
	"engine/modules/collider"
	"engine/modules/ecs"
	"engine/modules/window"
	"errors"

	"github.com/go-gl/mathgl/mgl32"
)

// updates dynamic projections
type ChangedResolutionEvent struct{}

func NewUpdateProjectionsEvent() ChangedResolutionEvent {
	return ChangedResolutionEvent{}
}

var (
	ErrNotCamera error = errors.New("this isn't a camera")
)

type Service interface {
	ecs.SystemRegister
	Component() ecs.ComponentArray[Component]
	Priority() ecs.ComponentArray[PriorityComponent]

	Mobile() ecs.ComponentArray[MobileCameraComponent]
	Limits() ecs.ComponentArray[CameraLimitsComponent]
	Viewport() ecs.ComponentArray[ViewportComponent]
	NormalizedViewport() ecs.ComponentArray[NormalizedViewportComponent]

	Ortho() ecs.ComponentArray[OrthoComponent]
	OrthoResolution() ecs.ComponentArray[OrthoResolutionComponent]
	Perspective() ecs.ComponentArray[PerspectiveComponent]
	DynamicPerspective() ecs.ComponentArray[DynamicPerspectiveComponent]

	// returns cameras from smallest to biggest
	OrderedCameras() []ecs.EntityID

	GetViewport(camera ecs.EntityID) (x, y, w, h int32)
	Mat4(caemra ecs.EntityID) mgl32.Mat4
	ShootRay(camera ecs.EntityID, mousePos window.MousePos) collider.Ray
}
