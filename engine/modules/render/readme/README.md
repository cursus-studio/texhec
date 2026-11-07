# render
## Architecture
this module is respinsible for defining unified events and components for rendering and
for providing basic instancing renderer

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              11            136             73            639
GLSL                             2             15              0             42
Markdown                         1              1              0              2
-------------------------------------------------------------------------------
SUM:                            14            152             73            683
-------------------------------------------------------------------------------
```
## TODO
Add frustrum culling. For each camera store all colliding objects.

Abstract shaders.

## Dependencies
`engine`:
  - `engine.Assets`
  - `engine.Camera`
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Graphics`
  - `engine.Groups`
  - `engine.Logger`
  - `engine.Render`
  - `engine.Transform`
  - `engine.Window`
  - `engine.World`

`engine/modules/assets`:
  - `engine/modules/assets.Asset`
  - `engine/modules/assets.GetAsset`
  - `engine/modules/assets.Path`
  - `engine/modules/assets.PathComponent`
  - `engine/modules/assets.Register`
  - `engine/modules/assets.Service`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSet`
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.Set`
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

`engine/modules/graphics`:
  - `engine/modules/graphics.Bind`
  - `engine/modules/graphics.Buffer`
  - `engine/modules/graphics.EBO`
  - `engine/modules/graphics.FlipV`
  - `engine/modules/graphics.Flush`
  - `engine/modules/graphics.FragmentShader`
  - `engine/modules/graphics.GetProgramLocations`
  - `engine/modules/graphics.ID`
  - `engine/modules/graphics.Image`
  - `engine/modules/graphics.ImagesCount`
  - `engine/modules/graphics.Index`
  - `engine/modules/graphics.Len`
  - `engine/modules/graphics.NewBuffer`
  - `engine/modules/graphics.NewEBO`
  - `engine/modules/graphics.NewFromSlice`
  - `engine/modules/graphics.NewImage`
  - `engine/modules/graphics.NewProgram`
  - `engine/modules/graphics.NewShader`
  - `engine/modules/graphics.NewVAO`
  - `engine/modules/graphics.NewVBO`
  - `engine/modules/graphics.Program`
  - `engine/modules/graphics.Release`
  - `engine/modules/graphics.Remove`
  - `engine/modules/graphics.Set`
  - `engine/modules/graphics.SetIndices`
  - `engine/modules/graphics.SetVertices`
  - `engine/modules/graphics.TextureArray`
  - `engine/modules/graphics.TrimTransparentBackground`
  - `engine/modules/graphics.VAO`
  - `engine/modules/graphics.VBOFactory`
  - `engine/modules/graphics.VBOSetter`
  - `engine/modules/graphics.VertexShader`

`engine/modules/loop`:
  - `engine/modules/loop.FrameEvent`

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

`engine/modules/transition`:
  - `engine/modules/transition.Lerp`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/go-gl/gl/v4.5-core/gl`
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`