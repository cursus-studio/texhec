# text
## Architecture
this module defines text and renderers it

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              11            157             49            776
GLSL                             3             20              4             57
Markdown                         1              0              0              2
-------------------------------------------------------------------------------
SUM:                            15            177             53            835
-------------------------------------------------------------------------------
```
## TODO
1. Implement more letters and make them easier to expand (emojis, hiragana, etc.).
2. Write own glyph rendering (solves first issue).

## Dependencies
`engine`:
  - `engine.Assets`
  - `engine.Camera`
  - `engine.EngineWorld`
  - `engine.EventsBuilder`
  - `engine.Graphics`
  - `engine.Groups`
  - `engine.Logger`
  - `engine.Text`
  - `engine.Transform`
  - `engine.World`

`engine/modules/assets`:
  - `engine/modules/assets.Asset`
  - `engine/modules/assets.GetAsset`
  - `engine/modules/assets.Path`
  - `engine/modules/assets.PathComponent`
  - `engine/modules/assets.Register`
  - `engine/modules/assets.Service`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.NewSparseSet`
  - `engine/modules/datastructures.SparseArray`
  - `engine/modules/datastructures.SparseSet`

`engine/modules/ecs`:
  - `engine/modules/ecs.AnyComponentArray`
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/graphics`:
  - `engine/modules/graphics.Bind`
  - `engine/modules/graphics.Buffer`
  - `engine/modules/graphics.FlipV`
  - `engine/modules/graphics.Flush`
  - `engine/modules/graphics.FragmentShader`
  - `engine/modules/graphics.GeomShader`
  - `engine/modules/graphics.GetProgramLocations`
  - `engine/modules/graphics.ID`
  - `engine/modules/graphics.Image`
  - `engine/modules/graphics.New`
  - `engine/modules/graphics.NewBuffer`
  - `engine/modules/graphics.NewImage`
  - `engine/modules/graphics.NewProgram`
  - `engine/modules/graphics.NewShader`
  - `engine/modules/graphics.NewVAO`
  - `engine/modules/graphics.NewVBO`
  - `engine/modules/graphics.Program`
  - `engine/modules/graphics.Release`
  - `engine/modules/graphics.Service`
  - `engine/modules/graphics.Set`
  - `engine/modules/graphics.SetVertices`
  - `engine/modules/graphics.TextureArray`
  - `engine/modules/graphics.VAO`
  - `engine/modules/graphics.VBOFactory`
  - `engine/modules/graphics.VBOSetter`
  - `engine/modules/graphics.VertexShader`

`engine/modules/render`:
  - `engine/modules/render.Camera`
  - `engine/modules/render.RenderEvent`

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
  - `engine/modules/text.Service`
  - `engine/modules/text.Text`
  - `engine/modules/text.TextComponent`
  - `engine/modules/text.Vertical`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/go-gl/gl/v4.5-core/gl`
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/image/font`
- `golang.org/x/image/font/opentype`
- `golang.org/x/image/math/fixed`