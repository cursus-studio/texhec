# camera
## Architecture
this module is responsible for cameras. Responsibilities:
- projections
- sets [size](/engine/modules/transform/readme/README.md#type-SizeComponent) for objects with projections
- shots rays
- parses cameras to mgl32.Mat4

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              11            171             32            833
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                            12            171             32            834
-------------------------------------------------------------------------------
```
## TODO
clean up pkg and move logic to internal

## Types
### type Service
Type: `engine/modules/camera.Service`

#### method Service Component
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/camera.Component]`

#### method Service DynamicPerspective
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/camera.DynamicPerspectiveComponent]`

#### method Service GetViewport
Type: `func(camera engine/modules/ecs.EntityID) (x int32, y int32, w int32, h int32)`

#### method Service Limits
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/camera.CameraLimitsComponent]`

#### method Service Mat4
Type: `func(caemra engine/modules/ecs.EntityID) github.com/go-gl/mathgl/mgl32.Mat4`

#### method Service Mobile
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/camera.MobileCameraComponent]`

#### method Service NormalizedViewport
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/camera.NormalizedViewportComponent]`

#### method Service OrderedCameras
Type: `func() []engine/modules/ecs.EntityID`
returns cameras from smallest to biggest

#### method Service Ortho
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/camera.OrthoComponent]`

#### method Service OrthoResolution
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/camera.OrthoResolutionComponent]`

#### method Service Perspective
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/camera.PerspectiveComponent]`

#### method Service Priority
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/camera.PriorityComponent]`

#### method Service Register
Type: `func() error`

#### method Service ShootRay
Type: `func(camera engine/modules/ecs.EntityID, mousePos engine/modules/window.MousePos) engine/modules/collider.Ray`

#### method Service Viewport
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/camera.ViewportComponent]`

### type Component
Type: `engine/modules/camera.Component`

#### property Component Projection
Type: `reflect.Type`

### type PriorityComponent
Type: `engine/modules/camera.PriorityComponent`

#### property PriorityComponent Priority
Type: `int`
biggest camera is uppermost

### type MobileCameraComponent
Type: `engine/modules/camera.MobileCameraComponent`
component specifying that camera can be freely moved on map

### type CameraLimitsComponent
Type: `engine/modules/camera.CameraLimitsComponent`

#### property CameraLimitsComponent MinZoom
Type: `float32`

#### property CameraLimitsComponent MaxZoom
Type: `float32`

#### property CameraLimitsComponent Min
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### property CameraLimitsComponent Max
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

### type ViewportComponent
Type: `engine/modules/camera.ViewportComponent`

#### property ViewportComponent X
Type: `int32`

#### property ViewportComponent Y
Type: `int32`

#### property ViewportComponent W
Type: `int32`

#### property ViewportComponent H
Type: `int32`

#### method ViewportComponent Viewport
Type: `func() (x int32, y int32, w int32, h int32)`

### type NormalizedViewportComponent
Type: `engine/modules/camera.NormalizedViewportComponent`

#### property NormalizedViewportComponent X
Type: `float32`

#### property NormalizedViewportComponent Y
Type: `float32`

#### property NormalizedViewportComponent W
Type: `float32`

#### property NormalizedViewportComponent H
Type: `float32`

#### method NormalizedViewportComponent Viewport
Type: `func(fullW int32, fullH int32) (rx int32, ry int32, rw int32, rh int32)`

### type OrthoComponent
Type: `engine/modules/camera.OrthoComponent`

#### property OrthoComponent Near
Type: `float32`

#### property OrthoComponent Far
Type: `float32`

#### property OrthoComponent Zoom
Type: `float32`

#### method OrthoComponent GetMatrix
Type: `func(w int32, h int32) github.com/go-gl/mathgl/mgl32.Mat4`

### type OrthoResolutionComponent
Type: `engine/modules/camera.OrthoResolutionComponent`

#### property OrthoResolutionComponent W
Type: `int32`

#### property OrthoResolutionComponent H
Type: `int32`

#### method OrthoResolutionComponent Elem
Type: `func() (w int32, h int32)`

### type CameraUp
Type: `engine/modules/camera.CameraUp`

### type CameraForward
Type: `engine/modules/camera.CameraForward`

### type PerspectiveComponent
Type: `engine/modules/camera.PerspectiveComponent`

#### property PerspectiveComponent FovY
Type: `float32`

#### property PerspectiveComponent AspectRatio
Type: `float32`

#### property PerspectiveComponent Near
Type: `float32`

#### property PerspectiveComponent Far
Type: `float32`

### type DynamicPerspectiveComponent
Type: `engine/modules/camera.DynamicPerspectiveComponent`

#### property DynamicPerspectiveComponent FovY
Type: `float32`

#### property DynamicPerspectiveComponent Near
Type: `float32`

#### property DynamicPerspectiveComponent Far
Type: `float32`

### type ChangedResolutionEvent
Type: `engine/modules/camera.ChangedResolutionEvent`
updates dynamic projections

## Variables
### var ErrNotCamera
Type: `error`

## Functions
### func NewCamera
Type: `func[Projection any]() engine/modules/camera.Component`

### func NewPriority
Type: `func(priority int) engine/modules/camera.PriorityComponent`

### func NewMobileCamera
Type: `func() engine/modules/camera.MobileCameraComponent`

### func NewCameraLimits
Type: `func(minZ float32, maxZ float32, min github.com/go-gl/mathgl/mgl32.Vec3, max github.com/go-gl/mathgl/mgl32.Vec3) engine/modules/camera.CameraLimitsComponent`

### func NewViewport
Type: `func(x int32, y int32, w int32, h int32) engine/modules/camera.ViewportComponent`

### func NewNormalizedViewport
Type: `func(x float32, y float32, w float32, h float32) engine/modules/camera.NormalizedViewportComponent`

### func NewOrtho
Type: `func(near float32, far float32) engine/modules/camera.OrthoComponent`

### func NewOrthoResolution
Type: `func(w int32, h int32) engine/modules/camera.OrthoResolutionComponent`

### func GetViewportOrthoResolution
Type: `func(x int32, y int32, w int32, h int32) engine/modules/camera.OrthoResolutionComponent`

### func NewPerspective
Type: `func(fovY float32, aspectRatio float32, near float32, far float32) engine/modules/camera.PerspectiveComponent`

### func NewDynamicPerspective
Type: `func(fovY float32, near float32, far float32) engine/modules/camera.DynamicPerspectiveComponent`

### func NewUpdateProjectionsEvent
Type: `func() engine/modules/camera.ChangedResolutionEvent`


## Dependencies
`engine`:
  - `engine.Camera`
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Focus`
  - `engine.Groups`
  - `engine.Logger`
  - `engine.Transform`
  - `engine.Window`
  - `engine.World`

`engine/modules/camera`:
  - `engine/modules/camera.AspectRatio`
  - `engine/modules/camera.CameraForward`
  - `engine/modules/camera.CameraLimitsComponent`
  - `engine/modules/camera.CameraUp`
  - `engine/modules/camera.ChangedResolutionEvent`
  - `engine/modules/camera.Component`
  - `engine/modules/camera.DynamicPerspective`
  - `engine/modules/camera.DynamicPerspectiveComponent`
  - `engine/modules/camera.Elem`
  - `engine/modules/camera.Far`
  - `engine/modules/camera.FovY`
  - `engine/modules/camera.GetMatrix`
  - `engine/modules/camera.GetViewport`
  - `engine/modules/camera.GetViewportOrthoResolution`
  - `engine/modules/camera.Limits`
  - `engine/modules/camera.Max`
  - `engine/modules/camera.MaxZoom`
  - `engine/modules/camera.Min`
  - `engine/modules/camera.MinZoom`
  - `engine/modules/camera.Mobile`
  - `engine/modules/camera.MobileCameraComponent`
  - `engine/modules/camera.Near`
  - `engine/modules/camera.NewCamera`
  - `engine/modules/camera.NewPerspective`
  - `engine/modules/camera.NewUpdateProjectionsEvent`
  - `engine/modules/camera.NormalizedViewport`
  - `engine/modules/camera.NormalizedViewportComponent`
  - `engine/modules/camera.Ortho`
  - `engine/modules/camera.OrthoComponent`
  - `engine/modules/camera.OrthoResolution`
  - `engine/modules/camera.OrthoResolutionComponent`
  - `engine/modules/camera.Perspective`
  - `engine/modules/camera.PerspectiveComponent`
  - `engine/modules/camera.Priority`
  - `engine/modules/camera.PriorityComponent`
  - `engine/modules/camera.Projection`
  - `engine/modules/camera.Service`
  - `engine/modules/camera.ShootRay`
  - `engine/modules/camera.Viewport`
  - `engine/modules/camera.ViewportComponent`
  - `engine/modules/camera.Zoom`

`engine/modules/collider`:
  - `engine/modules/collider.Direction`
  - `engine/modules/collider.Groups`
  - `engine/modules/collider.NewRay`
  - `engine/modules/collider.Pos`
  - `engine/modules/collider.Ray`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.SparseArray`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.NewSystemRegister`
  - `engine/modules/ecs.RegisterSystems`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/groups`:
  - `engine/modules/groups.Component`
  - `engine/modules/groups.DefaultGroups`

`engine/modules/inputs`:
  - `engine/modules/inputs.DragEvent`
  - `engine/modules/inputs.From`
  - `engine/modules/inputs.KeyboardEvent`
  - `engine/modules/inputs.To`

`engine/modules/loop`:
  - `engine/modules/loop.Delta`
  - `engine/modules/loop.FrameEvent`

`engine/modules/transform`:
  - `engine/modules/transform.AbsolutePos`
  - `engine/modules/transform.AbsolutePosComponent`
  - `engine/modules/transform.AbsoluteRotation`
  - `engine/modules/transform.AbsoluteSize`
  - `engine/modules/transform.AbsoluteSizeComponent`
  - `engine/modules/transform.AddDirtySet`
  - `engine/modules/transform.NewSize`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.Rotation`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/window`:
  - `engine/modules/window.Elem`
  - `engine/modules/window.GetMousePos`
  - `engine/modules/window.MousePos`
  - `engine/modules/window.Window`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/veandco/go-sdl2/sdl`