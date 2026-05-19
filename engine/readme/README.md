# TEXHEC engine

## Architecture
This module defines everything what every game engine should have but with one caveat.
It places efficient data layout as a first class citizen.

It follows **DOD** (data oriented design) and stores all game objects using **ECS** (entity component system).

## Types
### type EngineWorld
Type: `engine.EngineWorld`

#### property EngineWorld World
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/services/ecs.World]`

#### property EngineWorld EventsBuilder
Type: `github.com/ogiusek/ioc/v2.Lazy[github.com/ogiusek/events.Builder]`

#### property EngineWorld Events
Type: `github.com/ogiusek/ioc/v2.Lazy[github.com/ogiusek/events.Events]`

#### property EngineWorld Assets
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/assets.Service]`

#### property EngineWorld Audio
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/audio.Service]`

#### property EngineWorld Batcher
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/batcher.Service]`

#### property EngineWorld Camera
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/camera.Service]`

#### property EngineWorld Codec
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/codec.Service]`

#### property EngineWorld Collider
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/collider.Service]`

#### property EngineWorld Connection
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/connection.Service]`

#### property EngineWorld Drag
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/drag.Service]`

#### property EngineWorld EntityRegistry
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/entityregistry.Service]`

#### property EngineWorld Graphics
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/graphics.Service]`

#### property EngineWorld Groups
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/groups.Service]`

#### property EngineWorld Hierarchy
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/hierarchy.Service]`

#### property EngineWorld Inputs
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/inputs.Service]`

#### property EngineWorld Layout
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/layout.Service]`

#### property EngineWorld Logger
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/logger.Service]`

#### property EngineWorld Loop
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/loop.Service]`

#### property EngineWorld Metadata
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/metadata.Service]`

#### property EngineWorld NetSync
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/netsync.Service]`

#### property EngineWorld Noise
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/noise.Service]`

#### property EngineWorld Prototype
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/prototype.Service]`

#### property EngineWorld Record
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/record.Service]`

#### property EngineWorld Render
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/render.Service]`

#### property EngineWorld Scene
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/scene.Service]`

#### property EngineWorld Smooth
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/smooth.Service]`

#### property EngineWorld Text
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/text.Service]`

#### property EngineWorld Transform
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/transform.Service]`

#### property EngineWorld Transition
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/transition.Service]`

#### property EngineWorld UUID
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/uuid.Service]`

#### property EngineWorld WarmUp
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/warmup.Service]`

#### property EngineWorld Window
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/modules/window.Service]`

#### property EngineWorld Clock
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/services/clock.Clock]`

#### property EngineWorld Console
Type: `github.com/ogiusek/ioc/v2.Lazy[engine/services/console.Console]`


## Challenges
Biggest challenge was creating framework while building on top of it.
Testing own framework edgecases while delivering is impossible without frequent refactors.

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                             265           2902            766          13685
GLSL                             5             35              4             99
Markdown                         5              4              0             50
-------------------------------------------------------------------------------
SUM:                           275           2941            770          13834
-------------------------------------------------------------------------------

```
## Dependencies
`engine`:
  - `engine.Assets`
  - `engine.Camera`
  - `engine.Clock`
  - `engine.Codec`
  - `engine.Collider`
  - `engine.Connection`
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Graphics`
  - `engine.Groups`
  - `engine.Hierarchy`
  - `engine.Inputs`
  - `engine.Logger`
  - `engine.Loop`
  - `engine.NetSync`
  - `engine.Record`
  - `engine.Render`
  - `engine.Scene`
  - `engine.Text`
  - `engine.Transform`
  - `engine.Transition`
  - `engine.UUID`
  - `engine.WarmUp`
  - `engine.Window`
  - `engine.World`

`engine/modules/assets`:
  - `engine/modules/assets.Asset`
  - `engine/modules/assets.Cache`
  - `engine/modules/assets.CacheComponent`
  - `engine/modules/assets.ErrAssetNotFound`
  - `engine/modules/assets.Extension`
  - `engine/modules/assets.GetAsset`
  - `engine/modules/assets.NewCache`
  - `engine/modules/assets.NewPath`
  - `engine/modules/assets.Path`
  - `engine/modules/assets.PathComponent`
  - `engine/modules/assets.Register`
  - `engine/modules/assets.Release`
  - `engine/modules/assets.Service`

`engine/modules/assets/pkg`:
  - `engine/modules/assets/pkg.Pkg`

`engine/modules/audio`:
  - `engine/modules/audio.Asset`
  - `engine/modules/audio.AudioAsset`
  - `engine/modules/audio.Channel`
  - `engine/modules/audio.Chunk`
  - `engine/modules/audio.NewAudioAsset`
  - `engine/modules/audio.Play`
  - `engine/modules/audio.PlayEvent`
  - `engine/modules/audio.PlayerService`
  - `engine/modules/audio.Queue`
  - `engine/modules/audio.QueueEndless`
  - `engine/modules/audio.QueueEndlessEvent`
  - `engine/modules/audio.QueueEvent`
  - `engine/modules/audio.Service`
  - `engine/modules/audio.SetChannelVolume`
  - `engine/modules/audio.SetChannelVolumeEvent`
  - `engine/modules/audio.SetMasterVolume`
  - `engine/modules/audio.SetMasterVolumeEvent`
  - `engine/modules/audio.Stop`
  - `engine/modules/audio.StopEvent`
  - `engine/modules/audio.Volume`
  - `engine/modules/audio.VolumeService`

`engine/modules/audio/pkg`:
  - `engine/modules/audio/pkg.Pkg`

`engine/modules/batcher`:
  - `engine/modules/batcher.AddOrderedBatch`
  - `engine/modules/batcher.Batch`
  - `engine/modules/batcher.Handler`
  - `engine/modules/batcher.NewBatch`
  - `engine/modules/batcher.Progress`
  - `engine/modules/batcher.Service`
  - `engine/modules/batcher.Step`
  - `engine/modules/batcher.Steps`
  - `engine/modules/batcher.Task`
  - `engine/modules/batcher.TaskFactory`

`engine/modules/batcher/pkg`:
  - `engine/modules/batcher/pkg.Pkg`

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
  - `engine/modules/camera.Mat4`
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
  - `engine/modules/camera.OrderedCameras`
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

`engine/modules/camera/pkg`:
  - `engine/modules/camera/pkg.Pkg`

`engine/modules/codec`:
  - `engine/modules/codec.Decode`
  - `engine/modules/codec.Encode`
  - `engine/modules/codec.ErrCannotDecodeBytes`
  - `engine/modules/codec.ErrCannotEncodeType`
  - `engine/modules/codec.Service`

`engine/modules/codec/pkg`:
  - `engine/modules/codec/pkg.Pkg`
  - `engine/modules/codec/pkg.PkgT`

`engine/modules/collider`:
  - `engine/modules/collider.A`
  - `engine/modules/collider.AABB`
  - `engine/modules/collider.AABBs`
  - `engine/modules/collider.AddRayFallThroughPolicy`
  - `engine/modules/collider.Apply`
  - `engine/modules/collider.B`
  - `engine/modules/collider.Branch`
  - `engine/modules/collider.C`
  - `engine/modules/collider.ColliderAsset`
  - `engine/modules/collider.Component`
  - `engine/modules/collider.Count`
  - `engine/modules/collider.Direction`
  - `engine/modules/collider.Distance`
  - `engine/modules/collider.Entity`
  - `engine/modules/collider.FallThrough`
  - `engine/modules/collider.FallTroughPolicy`
  - `engine/modules/collider.First`
  - `engine/modules/collider.Groups`
  - `engine/modules/collider.Hit`
  - `engine/modules/collider.ID`
  - `engine/modules/collider.Leaf`
  - `engine/modules/collider.Max`
  - `engine/modules/collider.MaxDistance`
  - `engine/modules/collider.Min`
  - `engine/modules/collider.NewAABB`
  - `engine/modules/collider.NewObjectRayCollision`
  - `engine/modules/collider.NewRange`
  - `engine/modules/collider.NewRay`
  - `engine/modules/collider.NewRayHit`
  - `engine/modules/collider.ObjectObjectCollision`
  - `engine/modules/collider.ObjectRayCollision`
  - `engine/modules/collider.Point`
  - `engine/modules/collider.Polygon`
  - `engine/modules/collider.Polygons`
  - `engine/modules/collider.Pos`
  - `engine/modules/collider.Range`
  - `engine/modules/collider.Ranges`
  - `engine/modules/collider.Ray`
  - `engine/modules/collider.RayHit`
  - `engine/modules/collider.RaycastAll`
  - `engine/modules/collider.Service`
  - `engine/modules/collider.Target`

`engine/modules/collider/pkg`:
  - `engine/modules/collider/pkg.Pkg`

`engine/modules/connection`:
  - `engine/modules/connection.Close`
  - `engine/modules/connection.Component`
  - `engine/modules/connection.Conn`
  - `engine/modules/connection.ConnectionComponent`
  - `engine/modules/connection.Listener`
  - `engine/modules/connection.ListenerComponent`
  - `engine/modules/connection.Messages`
  - `engine/modules/connection.NewConnection`
  - `engine/modules/connection.NewListener`
  - `engine/modules/connection.Send`
  - `engine/modules/connection.Service`

`engine/modules/connection/pkg`:
  - `engine/modules/connection/pkg.Pkg`

`engine/modules/drag`:
  - `engine/modules/drag.Drag`
  - `engine/modules/drag.DraggableEvent`
  - `engine/modules/drag.Entity`
  - `engine/modules/drag.Service`

`engine/modules/drag/pkg`:
  - `engine/modules/drag/pkg.Pkg`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.ErrExpectedPointerToAStruct`
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/entityregistry/pkg`:
  - `engine/modules/entityregistry/pkg.Pkg`

`engine/modules/graphics`:
  - `engine/modules/graphics.Bind`
  - `engine/modules/graphics.Buffer`
  - `engine/modules/graphics.Configure`
  - `engine/modules/graphics.EBO`
  - `engine/modules/graphics.ErrInvalidLocation`
  - `engine/modules/graphics.ErrNotALocation`
  - `engine/modules/graphics.ErrProgramHasOtherLocations`
  - `engine/modules/graphics.ErrTexturesHaveToShareSize`
  - `engine/modules/graphics.FlipV`
  - `engine/modules/graphics.Flush`
  - `engine/modules/graphics.FragmentShader`
  - `engine/modules/graphics.GeomShader`
  - `engine/modules/graphics.GetProgramLocations`
  - `engine/modules/graphics.ID`
  - `engine/modules/graphics.Image`
  - `engine/modules/graphics.ImagesCount`
  - `engine/modules/graphics.Index`
  - `engine/modules/graphics.Len`
  - `engine/modules/graphics.Name`
  - `engine/modules/graphics.New`
  - `engine/modules/graphics.NewBuffer`
  - `engine/modules/graphics.NewEBO`
  - `engine/modules/graphics.NewFromSlice`
  - `engine/modules/graphics.NewImage`
  - `engine/modules/graphics.NewProgram`
  - `engine/modules/graphics.NewShader`
  - `engine/modules/graphics.NewVAO`
  - `engine/modules/graphics.NewVBO`
  - `engine/modules/graphics.Parameter`
  - `engine/modules/graphics.Program`
  - `engine/modules/graphics.Release`
  - `engine/modules/graphics.Remove`
  - `engine/modules/graphics.Service`
  - `engine/modules/graphics.Set`
  - `engine/modules/graphics.SetIndices`
  - `engine/modules/graphics.SetVertices`
  - `engine/modules/graphics.Shader`
  - `engine/modules/graphics.Texture`
  - `engine/modules/graphics.TextureArray`
  - `engine/modules/graphics.TextureArrayFactory`
  - `engine/modules/graphics.TextureFactory`
  - `engine/modules/graphics.TrimTransparentBackground`
  - `engine/modules/graphics.VAO`
  - `engine/modules/graphics.VBO`
  - `engine/modules/graphics.VBOFactory`
  - `engine/modules/graphics.VBOSetter`
  - `engine/modules/graphics.Value`
  - `engine/modules/graphics.VertexShader`

`engine/modules/graphics/pkg`:
  - `engine/modules/graphics/pkg.Pkg`

`engine/modules/grid`:
  - `engine/modules/grid.Component`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.GetIndex`
  - `engine/modules/grid.GetTile`
  - `engine/modules/grid.Height`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.Service`
  - `engine/modules/grid.SquareGridComponent`
  - `engine/modules/grid.TileConstraint`
  - `engine/modules/grid.Width`

`engine/modules/groups`:
  - `engine/modules/groups.Component`
  - `engine/modules/groups.DefaultGroups`
  - `engine/modules/groups.GetSharedWith`
  - `engine/modules/groups.GroupsComponent`
  - `engine/modules/groups.InheritGroupsComponent`
  - `engine/modules/groups.Mask`
  - `engine/modules/groups.Service`
  - `engine/modules/groups.SharesAnyGroup`

`engine/modules/groups/pkg`:
  - `engine/modules/groups/pkg.Pkg`

`engine/modules/hierarchy`:
  - `engine/modules/hierarchy.Children`
  - `engine/modules/hierarchy.Component`
  - `engine/modules/hierarchy.GetOrderedParents`
  - `engine/modules/hierarchy.NewParent`
  - `engine/modules/hierarchy.Parent`
  - `engine/modules/hierarchy.Service`
  - `engine/modules/hierarchy.SetParent`

`engine/modules/hierarchy/pkg`:
  - `engine/modules/hierarchy/pkg.Pkg`

`engine/modules/inputs`:
  - `engine/modules/inputs.ApplyDrag`
  - `engine/modules/inputs.ApplyDragEvent`
  - `engine/modules/inputs.ApplyEntity`
  - `engine/modules/inputs.ApplyEntityEvent`
  - `engine/modules/inputs.Camera`
  - `engine/modules/inputs.Capture`
  - `engine/modules/inputs.CaptureKeyboard`
  - `engine/modules/inputs.CaptureKeyboardComponent`
  - `engine/modules/inputs.CursorCol`
  - `engine/modules/inputs.CursorRow`
  - `engine/modules/inputs.DefaultFocusEvent`
  - `engine/modules/inputs.DefaultFocused`
  - `engine/modules/inputs.DefaultFocusedComponent`
  - `engine/modules/inputs.DoubleLeftClick`
  - `engine/modules/inputs.DoubleLeftClickComponent`
  - `engine/modules/inputs.DoubleRightClick`
  - `engine/modules/inputs.DoubleRightClickComponent`
  - `engine/modules/inputs.Drag`
  - `engine/modules/inputs.DragComponent`
  - `engine/modules/inputs.DragEvent`
  - `engine/modules/inputs.DraggedComponent`
  - `engine/modules/inputs.Entity`
  - `engine/modules/inputs.Event`
  - `engine/modules/inputs.EventTargetSetter`
  - `engine/modules/inputs.Fallthrough`
  - `engine/modules/inputs.FocusEvent`
  - `engine/modules/inputs.Focused`
  - `engine/modules/inputs.FocusedComponent`
  - `engine/modules/inputs.From`
  - `engine/modules/inputs.Hover`
  - `engine/modules/inputs.HoverComponent`
  - `engine/modules/inputs.Hovered`
  - `engine/modules/inputs.HoveredComponent`
  - `engine/modules/inputs.IsCaptured`
  - `engine/modules/inputs.KeepSelected`
  - `engine/modules/inputs.KeepSelectedComponent`
  - `engine/modules/inputs.KeyboardEvent`
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.LeftClickComponent`
  - `engine/modules/inputs.Lines`
  - `engine/modules/inputs.MouseEnter`
  - `engine/modules/inputs.MouseEnterComponent`
  - `engine/modules/inputs.MouseLeave`
  - `engine/modules/inputs.MouseLeaveComponent`
  - `engine/modules/inputs.NewDefaultFocusEvent`
  - `engine/modules/inputs.NewDefaultFocused`
  - `engine/modules/inputs.NewFocused`
  - `engine/modules/inputs.NewHoverComponent`
  - `engine/modules/inputs.NewKeyboardEvent`
  - `engine/modules/inputs.NewLeftClick`
  - `engine/modules/inputs.ObjectRayCollision`
  - `engine/modules/inputs.RightClick`
  - `engine/modules/inputs.RightClickComponent`
  - `engine/modules/inputs.Service`
  - `engine/modules/inputs.SetTarget`
  - `engine/modules/inputs.Stack`
  - `engine/modules/inputs.StackComponent`
  - `engine/modules/inputs.Stacked`
  - `engine/modules/inputs.StackedComponent`
  - `engine/modules/inputs.StackedData`
  - `engine/modules/inputs.SynchronizePositionEvent`
  - `engine/modules/inputs.Target`
  - `engine/modules/inputs.TextInput`
  - `engine/modules/inputs.TextInputComponent`
  - `engine/modules/inputs.TextInputEvent`
  - `engine/modules/inputs.To`
  - `engine/modules/inputs.UnfocusEvent`

`engine/modules/inputs/pkg`:
  - `engine/modules/inputs/pkg.Pkg`

`engine/modules/layout`:
  - `engine/modules/layout.AlignComponent`
  - `engine/modules/layout.Gap`
  - `engine/modules/layout.GapComponent`
  - `engine/modules/layout.NewAlign`
  - `engine/modules/layout.NewGap`
  - `engine/modules/layout.Order`
  - `engine/modules/layout.OrderComponent`
  - `engine/modules/layout.Primary`
  - `engine/modules/layout.Secondary`
  - `engine/modules/layout.Service`

`engine/modules/layout/pkg`:
  - `engine/modules/layout/pkg.Pkg`

`engine/modules/logger`:
  - `engine/modules/logger.ErrFatal`
  - `engine/modules/logger.ErrInfo`
  - `engine/modules/logger.Fatal`
  - `engine/modules/logger.Log`
  - `engine/modules/logger.Service`

`engine/modules/logger/pkg`:
  - `engine/modules/logger/pkg.Pkg`

`engine/modules/loop`:
  - `engine/modules/loop.ConfigureEvent`
  - `engine/modules/loop.Delta`
  - `engine/modules/loop.FPS`
  - `engine/modules/loop.FrameBudgetLeft`
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.NewStopEvent`
  - `engine/modules/loop.Service`
  - `engine/modules/loop.Stats`
  - `engine/modules/loop.StopEvent`
  - `engine/modules/loop.TPS`
  - `engine/modules/loop.TickEvent`

`engine/modules/loop/pkg`:
  - `engine/modules/loop/pkg.Pkg`

`engine/modules/metadata`:
  - `engine/modules/metadata.Description`
  - `engine/modules/metadata.DescriptionComponent`
  - `engine/modules/metadata.Link`
  - `engine/modules/metadata.LinkComponent`
  - `engine/modules/metadata.Name`
  - `engine/modules/metadata.NameComponent`
  - `engine/modules/metadata.NewDescription`
  - `engine/modules/metadata.NewLink`
  - `engine/modules/metadata.NewName`
  - `engine/modules/metadata.Service`

`engine/modules/metadata/pkg`:
  - `engine/modules/metadata/pkg.Pkg`

`engine/modules/netsync`:
  - `engine/modules/netsync.AuthorizedEvent`
  - `engine/modules/netsync.Client`
  - `engine/modules/netsync.ClientComponent`
  - `engine/modules/netsync.Server`
  - `engine/modules/netsync.ServerComponent`
  - `engine/modules/netsync.Service`
  - `engine/modules/netsync.SetConnection`

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.Pkg`

`engine/modules/noise`:
  - `engine/modules/noise.CellSize`
  - `engine/modules/noise.Factory`
  - `engine/modules/noise.LayerConfig`
  - `engine/modules/noise.NewLayer`
  - `engine/modules/noise.NewNoise`
  - `engine/modules/noise.Noise`
  - `engine/modules/noise.Read`
  - `engine/modules/noise.Service`
  - `engine/modules/noise.Weight`

`engine/modules/noise/pkg`:
  - `engine/modules/noise/pkg.Pkg`

`engine/modules/prototype`:
  - `engine/modules/prototype.NewCloneEvent`
  - `engine/modules/prototype.Service`

`engine/modules/prototype/pkg`:
  - `engine/modules/prototype/pkg.Pkg`
  - `engine/modules/prototype/pkg.PkgT`

`engine/modules/record`:
  - `engine/modules/record.AddToConfig`
  - `engine/modules/record.Apply`
  - `engine/modules/record.ComponentsOrder`
  - `engine/modules/record.Config`
  - `engine/modules/record.Entities`
  - `engine/modules/record.Entity`
  - `engine/modules/record.EntityKeyedRecorder`
  - `engine/modules/record.GetState`
  - `engine/modules/record.InheritZero`
  - `engine/modules/record.NewConfig`
  - `engine/modules/record.RecordedComponents`
  - `engine/modules/record.Recording`
  - `engine/modules/record.RecordingID`
  - `engine/modules/record.Service`
  - `engine/modules/record.StartBackwardsRecording`
  - `engine/modules/record.StartRecording`
  - `engine/modules/record.Stop`
  - `engine/modules/record.UUID`
  - `engine/modules/record.UUIDKeyedRecorder`
  - `engine/modules/record.UUIDRecording`
  - `engine/modules/record.UUIDRecordingID`

`engine/modules/record/pkg`:
  - `engine/modules/record/pkg.Pkg`

`engine/modules/relation`:
  - `engine/modules/relation.Get`
  - `engine/modules/relation.Service`

`engine/modules/relation/pkg`:
  - `engine/modules/relation/pkg.MapRelationPkg`

`engine/modules/render`:
  - `engine/modules/render.Asset`
  - `engine/modules/render.Camera`
  - `engine/modules/render.Color`
  - `engine/modules/render.ColorComponent`
  - `engine/modules/render.Error`
  - `engine/modules/render.GetFrame`
  - `engine/modules/render.ID`
  - `engine/modules/render.Images`
  - `engine/modules/render.Indices`
  - `engine/modules/render.Mesh`
  - `engine/modules/render.MeshAsset`
  - `engine/modules/render.MeshComponent`
  - `engine/modules/render.NewColor`
  - `engine/modules/render.NewTextureAsset`
  - `engine/modules/render.NewTextureFrame`
  - `engine/modules/render.Pos`
  - `engine/modules/render.RenderEvent`
  - `engine/modules/render.Service`
  - `engine/modules/render.Texture`
  - `engine/modules/render.TextureAsset`
  - `engine/modules/render.TextureComponent`
  - `engine/modules/render.TextureFrame`
  - `engine/modules/render.TextureFrameComponent`
  - `engine/modules/render.TexturePos`
  - `engine/modules/render.Vertex`
  - `engine/modules/render.Vertices`

`engine/modules/render/pkg`:
  - `engine/modules/render/pkg.Pkg`

`engine/modules/scene`:
  - `engine/modules/scene.ChangeSceneEvent`
  - `engine/modules/scene.ID`
  - `engine/modules/scene.Scene`
  - `engine/modules/scene.Service`

`engine/modules/scene/pkg`:
  - `engine/modules/scene/pkg.Pkg`

`engine/modules/seed`:
  - `engine/modules/seed.New`
  - `engine/modules/seed.Seed`
  - `engine/modules/seed.Value`

`engine/modules/smooth`:
  - `engine/modules/smooth.Service`
  - `engine/modules/smooth.SmoothConstraint`

`engine/modules/smooth/pkg`:
  - `engine/modules/smooth/pkg.Pkg`
  - `engine/modules/smooth/pkg.PkgT`

`engine/modules/text`:
  - `engine/modules/text.Align`
  - `engine/modules/text.AlignComponent`
  - `engine/modules/text.Break`
  - `engine/modules/text.BreakAny`
  - `engine/modules/text.BreakComponent`
  - `engine/modules/text.BreakNone`
  - `engine/modules/text.BreakWord`
  - `engine/modules/text.Color`
  - `engine/modules/text.ColorComponent`
  - `engine/modules/text.Content`
  - `engine/modules/text.FontFaceAsset`
  - `engine/modules/text.FontFamily`
  - `engine/modules/text.FontFamilyComponent`
  - `engine/modules/text.FontSize`
  - `engine/modules/text.FontSizeComponent`
  - `engine/modules/text.Glyphs`
  - `engine/modules/text.GlyphsWidth`
  - `engine/modules/text.Horizontal`
  - `engine/modules/text.Images`
  - `engine/modules/text.NewAlign`
  - `engine/modules/text.NewBreak`
  - `engine/modules/text.NewColor`
  - `engine/modules/text.NewFontAsset`
  - `engine/modules/text.NewFontSize`
  - `engine/modules/text.NewText`
  - `engine/modules/text.Service`
  - `engine/modules/text.Text`
  - `engine/modules/text.TextComponent`
  - `engine/modules/text.Vertical`

`engine/modules/text/pkg`:
  - `engine/modules/text/pkg.Pkg`

`engine/modules/transform`:
  - `engine/modules/transform.AbsolutePos`
  - `engine/modules/transform.AbsolutePosComponent`
  - `engine/modules/transform.AbsoluteRotation`
  - `engine/modules/transform.AbsoluteRotationComponent`
  - `engine/modules/transform.AbsoluteSize`
  - `engine/modules/transform.AbsoluteSizeComponent`
  - `engine/modules/transform.AddDirtySet`
  - `engine/modules/transform.AspectRatio`
  - `engine/modules/transform.AspectRatioComponent`
  - `engine/modules/transform.Mat4`
  - `engine/modules/transform.MaxSizeComponent`
  - `engine/modules/transform.MinSizeComponent`
  - `engine/modules/transform.NewAspectRatio`
  - `engine/modules/transform.NewMaxSize`
  - `engine/modules/transform.NewMinSize`
  - `engine/modules/transform.NewParent`
  - `engine/modules/transform.NewParentPivotPoint`
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.NewRotation`
  - `engine/modules/transform.NewSize`
  - `engine/modules/transform.ParentComponent`
  - `engine/modules/transform.ParentPivotPoint`
  - `engine/modules/transform.ParentPivotPointComponent`
  - `engine/modules/transform.PivotPoint`
  - `engine/modules/transform.PivotPointComponent`
  - `engine/modules/transform.Point`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.PosComponent`
  - `engine/modules/transform.PrimaryAxis`
  - `engine/modules/transform.PrimaryAxisX`
  - `engine/modules/transform.PrimaryAxisY`
  - `engine/modules/transform.PrimaryAxisZ`
  - `engine/modules/transform.RelativeMask`
  - `engine/modules/transform.RelativePos`
  - `engine/modules/transform.RelativeRotation`
  - `engine/modules/transform.RelativeSizeX`
  - `engine/modules/transform.RelativeSizeXYZ`
  - `engine/modules/transform.RelativeSizeY`
  - `engine/modules/transform.RelativeSizeZ`
  - `engine/modules/transform.Rotation`
  - `engine/modules/transform.RotationComponent`
  - `engine/modules/transform.Service`
  - `engine/modules/transform.Size`
  - `engine/modules/transform.SizeComponent`

`engine/modules/transform/pkg`:
  - `engine/modules/transform/pkg.Pkg`

`engine/modules/transition`:
  - `engine/modules/transition.Component`
  - `engine/modules/transition.DelayedEvent`
  - `engine/modules/transition.Duration`
  - `engine/modules/transition.EasingComponent`
  - `engine/modules/transition.EasingFunction`
  - `engine/modules/transition.EasingFunctionComponent`
  - `engine/modules/transition.Entity`
  - `engine/modules/transition.Event`
  - `engine/modules/transition.From`
  - `engine/modules/transition.ID`
  - `engine/modules/transition.Lerp`
  - `engine/modules/transition.LerpConstraint`
  - `engine/modules/transition.NewTransition`
  - `engine/modules/transition.Progress`
  - `engine/modules/transition.Service`
  - `engine/modules/transition.To`
  - `engine/modules/transition.TransitionComponent`
  - `engine/modules/transition.TransitionEvent`

`engine/modules/transition/pkg`:
  - `engine/modules/transition/pkg.Pkg`
  - `engine/modules/transition/pkg.PkgT`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.Entity`
  - `engine/modules/uuid.Factory`
  - `engine/modules/uuid.ID`
  - `engine/modules/uuid.New`
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.Service`
  - `engine/modules/uuid.UUID`

`engine/modules/uuid/pkg`:
  - `engine/modules/uuid/pkg.Pkg`

`engine/modules/warmup`:
  - `engine/modules/warmup.Event`
  - `engine/modules/warmup.Service`
  - `engine/modules/warmup.WarmUp`

`engine/modules/warmup/pkg`:
  - `engine/modules/warmup/pkg.Pkg`

`engine/modules/window`:
  - `engine/modules/window.Elem`
  - `engine/modules/window.GetMousePos`
  - `engine/modules/window.MousePos`
  - `engine/modules/window.NewMousePos`
  - `engine/modules/window.Service`
  - `engine/modules/window.Window`
  - `engine/modules/window.X`
  - `engine/modules/window.Y`

`engine/modules/window/pkg`:
  - `engine/modules/window/pkg.Pkg`

`engine/pkg`:
  - `engine/pkg.Pkg`

`engine/services/clock`:
  - `engine/services/clock.Clock`
  - `engine/services/clock.Now`
  - `engine/services/clock.Pkg`

`engine/services/console`:
  - `engine/services/console.Console`
  - `engine/services/console.Pkg`

`engine/services/datastructures`:
  - `engine/services/datastructures.Add`
  - `engine/services/datastructures.Changes`
  - `engine/services/datastructures.ClearChanges`
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.GetIndex`
  - `engine/services/datastructures.GetIndices`
  - `engine/services/datastructures.GetValues`
  - `engine/services/datastructures.NewSet`
  - `engine/services/datastructures.NewSparseArray`
  - `engine/services/datastructures.NewSparseSet`
  - `engine/services/datastructures.NewSparseSetWithPaging`
  - `engine/services/datastructures.NewThreadSafeTrackingArray`
  - `engine/services/datastructures.Remove`
  - `engine/services/datastructures.RemoveElements`
  - `engine/services/datastructures.Set`
  - `engine/services/datastructures.Size`
  - `engine/services/datastructures.SparseArray`
  - `engine/services/datastructures.SparseSet`
  - `engine/services/datastructures.SparseSetReader`
  - `engine/services/datastructures.TrackingArray`

`engine/services/ecs`:
  - `engine/services/ecs.AddDependency`
  - `engine/services/ecs.AddDirtySet`
  - `engine/services/ecs.AnyComponentArray`
  - `engine/services/ecs.BeforeGet`
  - `engine/services/ecs.Clear`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.Dirty`
  - `engine/services/ecs.DirtySet`
  - `engine/services/ecs.EnsureExists`
  - `engine/services/ecs.EntityExists`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetAny`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEmpty`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.NewDirtySet`
  - `engine/services/ecs.NewEntity`
  - `engine/services/ecs.NewSystemRegister`
  - `engine/services/ecs.NewWorld`
  - `engine/services/ecs.OnMod`
  - `engine/services/ecs.OnRemove`
  - `engine/services/ecs.OnUpsert`
  - `engine/services/ecs.Pkg`
  - `engine/services/ecs.Register`
  - `engine/services/ecs.RegisterSystems`
  - `engine/services/ecs.Release`
  - `engine/services/ecs.Remove`
  - `engine/services/ecs.RemoveEntity`
  - `engine/services/ecs.SaveComponent`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SetAny`
  - `engine/services/ecs.SetEmpty`
  - `engine/services/ecs.SystemRegister`
  - `engine/services/ecs.WarmUp`
  - `engine/services/ecs.World`

### Third Party
- `github.com/go-gl/gl/v4.5-core/gl`
- `github.com/go-gl/mathgl/mgl32`
- `github.com/go-gl/mathgl/mgl64`
- `github.com/google/uuid`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/veandco/go-sdl2/mix`
- `github.com/veandco/go-sdl2/sdl`
- `golang.org/x/exp/constraints`
- `golang.org/x/image/draw`
- `golang.org/x/image/font`
- `golang.org/x/image/font/opentype`
- `golang.org/x/image/math/fixed`