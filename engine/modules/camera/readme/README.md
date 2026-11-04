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
Go                              11            172             31            845
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                            12            172             31            846
-------------------------------------------------------------------------------
```
## TODO
clean up pkg and move logic to internal

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
  - `engine/modules/camera.CameraComponent`
  - `engine/modules/camera.CameraForward`
  - `engine/modules/camera.CameraLimitsComponent`
  - `engine/modules/camera.CameraUp`
  - `engine/modules/camera.ChangedResolutionEvent`
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
  - `engine/modules/groups.DefaultGroups`
  - `engine/modules/groups.Groups`

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