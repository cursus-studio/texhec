# TEXHEC CORE

## Architecture
Core module is responsible for implementing all `TEXHEC` game modules.

Implemented scenes:
- `menu`
- `game`
- `credits`
- `settings`

Map scale preview:
![Map scroll](/readme/map_scroll.gif)

## Challenges
Challenge of this module is to create something on incomplete foundation (`engine`)
without ending up in spagetti relations or broken code.
It needs to balance stability and features.

`grid` module which were made to deliver pre-mature features.
It stores whole map in a single slice where it should chunk the map.

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              61            745            341           4045
GLSL                             3             29              2            114
Markdown                         6             15              0             65
-------------------------------------------------------------------------------
SUM:                            70            789            343           4224
-------------------------------------------------------------------------------

```
## Dependencies
`core/game`:
  - `core/game.CreditsBuilder`
  - `core/game.Definitions`
  - `core/game.Deploy`
  - `core/game.FpsLogger`
  - `core/game.GameBuilder`
  - `core/game.GameWorld`
  - `core/game.Generation`
  - `core/game.Loading`
  - `core/game.MenuBuilder`
  - `core/game.Obstruction`
  - `core/game.Pathfind`
  - `core/game.Pkg`
  - `core/game.Player`
  - `core/game.Settings`
  - `core/game.SettingsBuilder`
  - `core/game.Tile`
  - `core/game.Ui`

`core/game/credits`:
  - `core/game/credits.Pkg`

`core/game/game`:
  - `core/game/game.Pkg`

`core/game/menu`:
  - `core/game/menu.Pkg`

`core/game/settings`:
  - `core/game/settings.Pkg`

`core/modules/definitions`:
  - `core/modules/definitions.AirspaceObstruction`
  - `core/modules/definitions.Assets`
  - `core/modules/definitions.Background1`
  - `core/modules/definitions.Background2`
  - `core/modules/definitions.BgGroup`
  - `core/modules/definitions.Blank`
  - `core/modules/definitions.Btn`
  - `core/modules/definitions.Can`
  - `core/modules/definitions.Cannot`
  - `core/modules/definitions.ConstructLayer`
  - `core/modules/definitions.CreditsID`
  - `core/modules/definitions.Cursor`
  - `core/modules/definitions.EaseOutElastic`
  - `core/modules/definitions.EffectChannel`
  - `core/modules/definitions.ExampleAudio`
  - `core/modules/definitions.Farm`
  - `core/modules/definitions.FontAsset`
  - `core/modules/definitions.GameGroup`
  - `core/modules/definitions.GameID`
  - `core/modules/definitions.HouseT1`
  - `core/modules/definitions.HouseT2`
  - `core/modules/definitions.HouseT3`
  - `core/modules/definitions.HouseT4`
  - `core/modules/definitions.Hud`
  - `core/modules/definitions.Input`
  - `core/modules/definitions.Linear`
  - `core/modules/definitions.Load`
  - `core/modules/definitions.LowlandObstruction`
  - `core/modules/definitions.MenuID`
  - `core/modules/definitions.MyEasing`
  - `core/modules/definitions.ObjectPlaceholderLayer`
  - `core/modules/definitions.Objects`
  - `core/modules/definitions.PathLayer`
  - `core/modules/definitions.Selected`
  - `core/modules/definitions.Service`
  - `core/modules/definitions.Settings`
  - `core/modules/definitions.SettingsID`
  - `core/modules/definitions.SquareCollider`
  - `core/modules/definitions.SquareMesh`
  - `core/modules/definitions.Tank`
  - `core/modules/definitions.Target`
  - `core/modules/definitions.Text`
  - `core/modules/definitions.TileLayer`
  - `core/modules/definitions.TilePlaceholderLayer`
  - `core/modules/definitions.Tiles`
  - `core/modules/definitions.Transitions`
  - `core/modules/definitions.UiGroup`
  - `core/modules/definitions.UnitLayer`
  - `core/modules/definitions.WaterObstruction`

`core/modules/definitions/pkg`:
  - `core/modules/definitions/pkg.Pkg`

`core/modules/deploy`:
  - `core/modules/deploy.ApplyCoords`
  - `core/modules/deploy.Blueprint`
  - `core/modules/deploy.By`
  - `core/modules/deploy.Component`
  - `core/modules/deploy.Coords`
  - `core/modules/deploy.Deploy`
  - `core/modules/deploy.Deployable`
  - `core/modules/deploy.ExecuteEvent`
  - `core/modules/deploy.NewDeploy`
  - `core/modules/deploy.NewExecuteEvent`
  - `core/modules/deploy.NewPreviewEvent`
  - `core/modules/deploy.NewSelectEvent`
  - `core/modules/deploy.PreviewEvent`
  - `core/modules/deploy.SelectEvent`
  - `core/modules/deploy.Service`

`core/modules/deploy/pkg`:
  - `core/modules/deploy/pkg.Pkg`

`core/modules/fpslogger`:
  - `core/modules/fpslogger.Service`

`core/modules/fpslogger/pkg`:
  - `core/modules/fpslogger/pkg.Pkg`

`core/modules/generation`:
  - `core/modules/generation.Config`
  - `core/modules/generation.Entity`
  - `core/modules/generation.Generate`
  - `core/modules/generation.NewConfig`
  - `core/modules/generation.Seed`
  - `core/modules/generation.Service`
  - `core/modules/generation.Size`

`core/modules/generation/pkg`:
  - `core/modules/generation/pkg.Pkg`

`core/modules/loading`:
  - `core/modules/loading.Service`

`core/modules/loading/pkg`:
  - `core/modules/loading/pkg.Pkg`

`core/modules/obstruction`:
  - `core/modules/obstruction.AABB`
  - `core/modules/obstruction.Collisions`
  - `core/modules/obstruction.Component`
  - `core/modules/obstruction.Deployed`
  - `core/modules/obstruction.DeployedComponent`
  - `core/modules/obstruction.ErrPositionIsOccupied`
  - `core/modules/obstruction.Grid`
  - `core/modules/obstruction.NewAABB`
  - `core/modules/obstruction.NewDeployed`
  - `core/modules/obstruction.NewGrid`
  - `core/modules/obstruction.NewObstruction`
  - `core/modules/obstruction.Obstruction`
  - `core/modules/obstruction.Service`
  - `core/modules/obstruction.Tiles`

`core/modules/obstruction/pkg`:
  - `core/modules/obstruction/pkg.Pkg`

`core/modules/pathfind`:
  - `core/modules/pathfind.ApplyCoords`
  - `core/modules/pathfind.CanStep`
  - `core/modules/pathfind.Coords`
  - `core/modules/pathfind.Entity`
  - `core/modules/pathfind.ErrInvalidPath`
  - `core/modules/pathfind.FindPathEvent`
  - `core/modules/pathfind.InvSpeed`
  - `core/modules/pathfind.NewFindPathEvent`
  - `core/modules/pathfind.NewPreviewPathEvent`
  - `core/modules/pathfind.NewSelectEvent`
  - `core/modules/pathfind.NewSpeed`
  - `core/modules/pathfind.NewStep`
  - `core/modules/pathfind.NewTarget`
  - `core/modules/pathfind.PreviewPathEvent`
  - `core/modules/pathfind.SelectEvent`
  - `core/modules/pathfind.Service`
  - `core/modules/pathfind.Speed`
  - `core/modules/pathfind.SpeedComponent`
  - `core/modules/pathfind.Step`
  - `core/modules/pathfind.StepComponent`
  - `core/modules/pathfind.TargetComponent`

`core/modules/pathfind/pkg`:
  - `core/modules/pathfind/pkg.Pkg`

`core/modules/player`:
  - `core/modules/player.NewOwner`
  - `core/modules/player.Owner`
  - `core/modules/player.OwnerComponent`
  - `core/modules/player.Service`

`core/modules/player/pkg`:
  - `core/modules/player/pkg.Pkg`

`core/modules/settings`:
  - `core/modules/settings.EnterSettingsEvent`
  - `core/modules/settings.EnterSettingsForParentEvent`
  - `core/modules/settings.Parent`
  - `core/modules/settings.Service`

`core/modules/settings/pkg`:
  - `core/modules/settings/pkg.Pkg`

`core/modules/tile`:
  - `core/modules/tile.Aligned`
  - `core/modules/tile.ApplyCoords`
  - `core/modules/tile.ApplyCoordsEvent`
  - `core/modules/tile.BiomeAsset`
  - `core/modules/tile.ClickEntityEvent`
  - `core/modules/tile.Component`
  - `core/modules/tile.Coord`
  - `core/modules/tile.Entity`
  - `core/modules/tile.ErrInvalidPosition`
  - `core/modules/tile.ErrInvalidStep`
  - `core/modules/tile.ErrPositionAndSpeedIsRequiredToStep`
  - `core/modules/tile.GetTile`
  - `core/modules/tile.GetTileSize`
  - `core/modules/tile.Grid`
  - `core/modules/tile.HoverEvent`
  - `core/modules/tile.ID`
  - `core/modules/tile.Images`
  - `core/modules/tile.Layer`
  - `core/modules/tile.LayerComponent`
  - `core/modules/tile.NewBiomeAsset`
  - `core/modules/tile.NewClickEntityEvent`
  - `core/modules/tile.NewGrid`
  - `core/modules/tile.NewHoverEvent`
  - `core/modules/tile.NewLayer`
  - `core/modules/tile.NewPos`
  - `core/modules/tile.NewRot`
  - `core/modules/tile.NewSelectEvent`
  - `core/modules/tile.NewSize`
  - `core/modules/tile.NewTile`
  - `core/modules/tile.Pos`
  - `core/modules/tile.PosComponent`
  - `core/modules/tile.Quat`
  - `core/modules/tile.Renderer`
  - `core/modules/tile.Rot`
  - `core/modules/tile.RotComponent`
  - `core/modules/tile.SelectEvent`
  - `core/modules/tile.Service`
  - `core/modules/tile.Size`
  - `core/modules/tile.SizeComponent`
  - `core/modules/tile.Tile`
  - `core/modules/tile.X`
  - `core/modules/tile.Y`
  - `core/modules/tile.Z`

`core/modules/tile/pkg`:
  - `core/modules/tile/pkg.Pkg`

`core/modules/ui`:
  - `core/modules/ui.ActionComponent`
  - `core/modules/ui.Actions`
  - `core/modules/ui.AnimatedBackground`
  - `core/modules/ui.AnimatedBackgroundComponent`
  - `core/modules/ui.CursorCamera`
  - `core/modules/ui.CursorCameraComponent`
  - `core/modules/ui.Entities`
  - `core/modules/ui.NewSelect`
  - `core/modules/ui.NewSelectTick`
  - `core/modules/ui.NewUnselect`
  - `core/modules/ui.ObjectComponent`
  - `core/modules/ui.Objects`
  - `core/modules/ui.SelectEvent`
  - `core/modules/ui.SelectionGroup`
  - `core/modules/ui.Service`
  - `core/modules/ui.ShowMenu`
  - `core/modules/ui.UiCamera`
  - `core/modules/ui.UiCameraComponent`
  - `core/modules/ui.UnselectEvent`

`core/modules/ui/pkg`:
  - `core/modules/ui/pkg.Pkg`

`core/pkg`:
  - `core/pkg.Pkg`

`engine`:
  - `engine.Assets`
  - `engine.Audio`
  - `engine.Batcher`
  - `engine.Camera`
  - `engine.Clock`
  - `engine.Collider`
  - `engine.Connection`
  - `engine.Console`
  - `engine.Drag`
  - `engine.EngineWorld`
  - `engine.EntityRegistry`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Graphics`
  - `engine.Groups`
  - `engine.Hierarchy`
  - `engine.Inputs`
  - `engine.Layout`
  - `engine.Logger`
  - `engine.Loop`
  - `engine.Metadata`
  - `engine.NetSync`
  - `engine.Noise`
  - `engine.Prototype`
  - `engine.Record`
  - `engine.Render`
  - `engine.Scene`
  - `engine.Smooth`
  - `engine.Text`
  - `engine.Transform`
  - `engine.Transition`
  - `engine.UUID`
  - `engine.Window`
  - `engine.World`

`engine/modules/assets`:
  - `engine/modules/assets.Asset`
  - `engine/modules/assets.Cache`
  - `engine/modules/assets.GetAsset`
  - `engine/modules/assets.NewCache`
  - `engine/modules/assets.Path`
  - `engine/modules/assets.PathComponent`
  - `engine/modules/assets.Register`
  - `engine/modules/assets.Service`

`engine/modules/audio`:
  - `engine/modules/audio.Channel`
  - `engine/modules/audio.NewPlayEvent`

`engine/modules/batcher`:
  - `engine/modules/batcher.AddConcurrentBatch`
  - `engine/modules/batcher.AddOrderedBatch`
  - `engine/modules/batcher.Build`
  - `engine/modules/batcher.NewBatch`
  - `engine/modules/batcher.NewTask`
  - `engine/modules/batcher.Progress`
  - `engine/modules/batcher.Queue`
  - `engine/modules/batcher.Task`

`engine/modules/camera`:
  - `engine/modules/camera.Limits`
  - `engine/modules/camera.Mat4`
  - `engine/modules/camera.Mobile`
  - `engine/modules/camera.NewCameraLimits`
  - `engine/modules/camera.NewMobileCamera`
  - `engine/modules/camera.NewOrtho`
  - `engine/modules/camera.NewPriority`
  - `engine/modules/camera.Ortho`
  - `engine/modules/camera.OrthoComponent`
  - `engine/modules/camera.Priority`
  - `engine/modules/camera.ShootRay`
  - `engine/modules/camera.Zoom`

`engine/modules/collider`:
  - `engine/modules/collider.AABB`
  - `engine/modules/collider.Component`
  - `engine/modules/collider.Direction`
  - `engine/modules/collider.Leaf`
  - `engine/modules/collider.NewAABB`
  - `engine/modules/collider.NewCollider`
  - `engine/modules/collider.NewColliderAsset`
  - `engine/modules/collider.NewPolygon`
  - `engine/modules/collider.NewRange`
  - `engine/modules/collider.Polygon`
  - `engine/modules/collider.Pos`
  - `engine/modules/collider.Range`

`engine/modules/drag`:
  - `engine/modules/drag.DraggableEvent`
  - `engine/modules/drag.NewDraggable`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.GetRegistry`
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/graphics`:
  - `engine/modules/graphics.Bind`
  - `engine/modules/graphics.Buffer`
  - `engine/modules/graphics.FlipH`
  - `engine/modules/graphics.FlipHV`
  - `engine/modules/graphics.FlipV`
  - `engine/modules/graphics.Flush`
  - `engine/modules/graphics.FragmentShader`
  - `engine/modules/graphics.GeomShader`
  - `engine/modules/graphics.GetProgramLocations`
  - `engine/modules/graphics.ID`
  - `engine/modules/graphics.Image`
  - `engine/modules/graphics.Index`
  - `engine/modules/graphics.New`
  - `engine/modules/graphics.NewBuffer`
  - `engine/modules/graphics.NewImage`
  - `engine/modules/graphics.NewProgram`
  - `engine/modules/graphics.NewShader`
  - `engine/modules/graphics.NewVAO`
  - `engine/modules/graphics.NewVBO`
  - `engine/modules/graphics.Opaque`
  - `engine/modules/graphics.Program`
  - `engine/modules/graphics.Release`
  - `engine/modules/graphics.Scale`
  - `engine/modules/graphics.Service`
  - `engine/modules/graphics.Set`
  - `engine/modules/graphics.Texture`
  - `engine/modules/graphics.TextureArray`
  - `engine/modules/graphics.VAO`
  - `engine/modules/graphics.VBOFactory`
  - `engine/modules/graphics.VBOSetter`
  - `engine/modules/graphics.VertexShader`
  - `engine/modules/graphics.Wrap`

`engine/modules/grid`:
  - `engine/modules/grid.Component`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.GetCoords`
  - `engine/modules/grid.GetIndex`
  - `engine/modules/grid.GetLastIndex`
  - `engine/modules/grid.GetTile`
  - `engine/modules/grid.GetTiles`
  - `engine/modules/grid.Height`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.NewSquareGrid`
  - `engine/modules/grid.Service`
  - `engine/modules/grid.SetTile`
  - `engine/modules/grid.SquareGridComponent`
  - `engine/modules/grid.Width`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/modules/grid/pkg`:
  - `engine/modules/grid/pkg.NewConfig`
  - `engine/modules/grid/pkg.PkgT`

`engine/modules/groups`:
  - `engine/modules/groups.Component`
  - `engine/modules/groups.DefaultGroups`
  - `engine/modules/groups.EmptyGroups`
  - `engine/modules/groups.Enable`
  - `engine/modules/groups.Group`
  - `engine/modules/groups.Inherit`
  - `engine/modules/groups.InheritGroupsComponent`
  - `engine/modules/groups.Ptr`
  - `engine/modules/groups.SharesAnyGroup`
  - `engine/modules/groups.Val`

`engine/modules/inputs`:
  - `engine/modules/inputs.CaptureKeyboard`
  - `engine/modules/inputs.Drag`
  - `engine/modules/inputs.DragEvent`
  - `engine/modules/inputs.FocusEvent`
  - `engine/modules/inputs.KeepSelected`
  - `engine/modules/inputs.KeepSelectedComponent`
  - `engine/modules/inputs.KeyboardEvent`
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewCaptureKeyboard`
  - `engine/modules/inputs.NewDragComponent`
  - `engine/modules/inputs.NewLeftClick`
  - `engine/modules/inputs.NewTextInputEvent`
  - `engine/modules/inputs.Stack`
  - `engine/modules/inputs.StackComponent`

`engine/modules/layout`:
  - `engine/modules/layout.Align`
  - `engine/modules/layout.Gap`
  - `engine/modules/layout.NewAlign`
  - `engine/modules/layout.NewGap`
  - `engine/modules/layout.NewOrder`
  - `engine/modules/layout.Order`
  - `engine/modules/layout.OrderVectical`

`engine/modules/logger`:
  - `engine/modules/logger.ErrFatal`
  - `engine/modules/logger.ErrInfo`
  - `engine/modules/logger.Info`
  - `engine/modules/logger.IsWarning`
  - `engine/modules/logger.Log`

`engine/modules/logger/pkg`:
  - `engine/modules/logger/pkg.Config`

`engine/modules/loop`:
  - `engine/modules/loop.FrameEvent`
  - `engine/modules/loop.NewConfigureEvent`
  - `engine/modules/loop.NewStopEvent`
  - `engine/modules/loop.Run`
  - `engine/modules/loop.TickEvent`

`engine/modules/metadata`:
  - `engine/modules/metadata.Entity`
  - `engine/modules/metadata.Link`
  - `engine/modules/metadata.Name`
  - `engine/modules/metadata.NewName`

`engine/modules/netsync/pkg`:
  - `engine/modules/netsync/pkg.AddEvent`
  - `engine/modules/netsync/pkg.AddTransparentEvent`
  - `engine/modules/netsync/pkg.Config`
  - `engine/modules/netsync/pkg.RecordConfig`
  - `engine/modules/netsync/pkg.SetMaxPredictions`

`engine/modules/noise`:
  - `engine/modules/noise.AddPerlin`
  - `engine/modules/noise.AddValue`
  - `engine/modules/noise.Build`
  - `engine/modules/noise.NewLayer`
  - `engine/modules/noise.NewNoise`
  - `engine/modules/noise.Read`

`engine/modules/record`:
  - `engine/modules/record.AddToConfig`
  - `engine/modules/record.ComponentGetter`
  - `engine/modules/record.Config`
  - `engine/modules/record.Entities`
  - `engine/modules/record.Entity`
  - `engine/modules/record.NewConfig`
  - `engine/modules/record.RecordingID`
  - `engine/modules/record.StartBackwardsRecording`
  - `engine/modules/record.Stop`

`engine/modules/relation`:
  - `engine/modules/relation.Get`
  - `engine/modules/relation.Service`

`engine/modules/relation/pkg`:
  - `engine/modules/relation/pkg.SpatialRelationPkg`

`engine/modules/render`:
  - `engine/modules/render.AspectRatio`
  - `engine/modules/render.Camera`
  - `engine/modules/render.Color`
  - `engine/modules/render.ErrTextureAssetImagesHasToMatchResolution`
  - `engine/modules/render.Images`
  - `engine/modules/render.Mesh`
  - `engine/modules/render.NewColor`
  - `engine/modules/render.NewMesh`
  - `engine/modules/render.NewMeshAsset`
  - `engine/modules/render.NewTexture`
  - `engine/modules/render.NewTextureAsset`
  - `engine/modules/render.NewTextureFrame`
  - `engine/modules/render.RenderEvent`
  - `engine/modules/render.Renderer`
  - `engine/modules/render.Texture`
  - `engine/modules/render.TextureAsset`
  - `engine/modules/render.TextureFrameComponent`
  - `engine/modules/render.Vertex`

`engine/modules/scene`:
  - `engine/modules/scene.NewChangeSceneEvent`
  - `engine/modules/scene.NewSceneId`
  - `engine/modules/scene.Scene`
  - `engine/modules/scene.Service`
  - `engine/modules/scene.SetScene`

`engine/modules/seed`:
  - `engine/modules/seed.New`
  - `engine/modules/seed.Seed`

`engine/modules/text`:
  - `engine/modules/text.Align`
  - `engine/modules/text.Break`
  - `engine/modules/text.BreakNone`
  - `engine/modules/text.Color`
  - `engine/modules/text.Content`
  - `engine/modules/text.FontFamily`
  - `engine/modules/text.FontSize`
  - `engine/modules/text.NewAlign`
  - `engine/modules/text.NewBreak`
  - `engine/modules/text.NewColor`
  - `engine/modules/text.NewFontFamily`
  - `engine/modules/text.NewFontSize`
  - `engine/modules/text.NewText`
  - `engine/modules/text.Renderer`
  - `engine/modules/text.Service`

`engine/modules/text/pkg`:
  - `engine/modules/text/pkg.Config`
  - `engine/modules/text/pkg.UsedGlyphs`

`engine/modules/transform`:
  - `engine/modules/transform.Absolute`
  - `engine/modules/transform.AspectRatio`
  - `engine/modules/transform.Mat4`
  - `engine/modules/transform.MaxSize`
  - `engine/modules/transform.NewAspectRatio`
  - `engine/modules/transform.NewMaxSize`
  - `engine/modules/transform.NewParent`
  - `engine/modules/transform.NewParentPivotPoint`
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.NewRotation`
  - `engine/modules/transform.NewSize`
  - `engine/modules/transform.Parent`
  - `engine/modules/transform.ParentPivotPoint`
  - `engine/modules/transform.PivotPoint`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.PosComponent`
  - `engine/modules/transform.PrimaryAxisX`
  - `engine/modules/transform.RelativePos`
  - `engine/modules/transform.RelativeSizeX`
  - `engine/modules/transform.RelativeSizeXY`
  - `engine/modules/transform.RelativeSizeXYZ`
  - `engine/modules/transform.Rotation`
  - `engine/modules/transform.Size`
  - `engine/modules/transform.SizeComponent`

`engine/modules/transition`:
  - `engine/modules/transition.EasingFunction`
  - `engine/modules/transition.Lerp`
  - `engine/modules/transition.NewDelayedEvent`
  - `engine/modules/transition.NewEasingFunction`
  - `engine/modules/transition.NewTransition`
  - `engine/modules/transition.NewTransitionEvent`
  - `engine/modules/transition.Progress`
  - `engine/modules/transition.TransitionComponent`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.New`

`engine/modules/window`:
  - `engine/modules/window.GetMousePos`
  - `engine/modules/window.Service`
  - `engine/modules/window.Window`

`engine/pkg`:
  - `engine/pkg.Pkg`

`engine/services/datastructures`:
  - `engine/services/datastructures.Add`
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.GetIndices`
  - `engine/services/datastructures.NewSparseArray`
  - `engine/services/datastructures.Remove`
  - `engine/services/datastructures.Set`
  - `engine/services/datastructures.Size`
  - `engine/services/datastructures.SparseArray`

`engine/services/ecs`:
  - `engine/services/ecs.AddDirtySet`
  - `engine/services/ecs.BeforeGet`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.DirtySet`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.NewDirtySet`
  - `engine/services/ecs.NewEntity`
  - `engine/services/ecs.NewSystemRegister`
  - `engine/services/ecs.OnRemove`
  - `engine/services/ecs.OnUpsert`
  - `engine/services/ecs.RegisterSystems`
  - `engine/services/ecs.Remove`
  - `engine/services/ecs.RemoveEntity`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SetEmpty`
  - `engine/services/ecs.SystemRegister`
  - `engine/services/ecs.World`

### Third Party
- `github.com/go-gl/gl/v4.5-core/gl`
- `github.com/go-gl/mathgl/mgl32`
- `github.com/go-gl/mathgl/mgl64`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/veandco/go-sdl2/sdl`
- `golang.org/x/exp/constraints`