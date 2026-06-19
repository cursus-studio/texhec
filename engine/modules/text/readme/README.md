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

## Types
### type Service
Type: `engine/modules/text.Service`

#### method Service AddDirtySet
Type: `func(engine/modules/ecs.DirtySet)`

#### method Service Align
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/text.AlignComponent]`

#### method Service Break
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/text.BreakComponent]`

#### method Service Color
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/text.ColorComponent]`

#### method Service Content
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/text.TextComponent]`

#### method Service FontFamily
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/text.FontFamilyComponent]`

#### method Service FontSize
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/text.FontSizeComponent]`

#### method Service Renderer
Type: `func() engine/modules/ecs.SystemRegister`

### type FontFaceAsset
Type: `engine/modules/text.FontFaceAsset`

#### method FontFaceAsset Font
Type: `func() *golang.org/x/image/font/opentype.Font`

#### method FontFaceAsset Glyphs
Type: `func() engine/modules/text.Glyphs`

#### method FontFaceAsset Release
Type: `func()`

### type Glyphs
Type: `engine/modules/text.Glyphs`

#### property Glyphs GlyphsWidth
Type: `engine/services/datastructures.SparseArray[uint32, float32]`

#### property Glyphs Images
Type: `engine/services/datastructures.SparseArray[uint32, image.Image]`

### type TextComponent
Type: `engine/modules/text.TextComponent`
this is required to render text
every other component is optional and has default value

#### property TextComponent Text
Type: `string`

### type AlignComponent
Type: `engine/modules/text.AlignComponent`

#### property AlignComponent Vertical
Type: `float32`
value between 0 and 1 where 0 means aligned to left and 1 aligned to right

#### property AlignComponent Horizontal
Type: `float32`
value between 0 and 1 where 0 means aligned to left and 1 aligned to right

### type ColorComponent
Type: `engine/modules/text.ColorComponent`

#### property ColorComponent Color
Type: `github.com/go-gl/mathgl/mgl32.Vec4`

### type FontFamilyComponent
Type: `engine/modules/text.FontFamilyComponent`

#### property FontFamilyComponent FontFamily
Type: `engine/modules/ecs.EntityID`

### type FontSizeComponent
Type: `engine/modules/text.FontSizeComponent`

#### property FontSizeComponent FontSize
Type: `uint`

### type BreakComponent
Type: `engine/modules/text.BreakComponent`

#### property BreakComponent Break
Type: `uint8`

## Variables
### var BreakNone
Type: `uint8`

### var BreakWord
Type: `uint8`

### var BreakAny
Type: `uint8`

## Functions
### func NewFontAsset
Type: `func(raw golang.org/x/image/font/opentype.Font, glyphs engine/modules/text.Glyphs) engine/modules/text.FontFaceAsset`

### func NewText
Type: `func(text string) engine/modules/text.TextComponent`

### func NewAlign
Type: `func(vertical float32, horizontal float32) engine/modules/text.AlignComponent`

### func NewColor
Type: `func(color github.com/go-gl/mathgl/mgl32.Vec4) engine/modules/text.ColorComponent`

### func NewFontFamily
Type: `func(fontFamily engine/modules/ecs.EntityID) engine/modules/text.FontFamilyComponent`

### func NewFontSize
Type: `func(fontSize uint) engine/modules/text.FontSizeComponent`

### func NewBreak
Type: `func(b uint8) engine/modules/text.BreakComponent`


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

`engine/services/datastructures`:
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.GetIndices`
  - `engine/services/datastructures.GetValues`
  - `engine/services/datastructures.NewSparseArray`
  - `engine/services/datastructures.NewSparseSet`
  - `engine/services/datastructures.Remove`
  - `engine/services/datastructures.Set`
  - `engine/services/datastructures.SparseArray`
  - `engine/services/datastructures.SparseSet`

### Third Party
- `github.com/go-gl/gl/v4.5-core/gl`
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/image/font`
- `golang.org/x/image/font/opentype`
- `golang.org/x/image/math/fixed`