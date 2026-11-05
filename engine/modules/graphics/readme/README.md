# graphics
## Architecture
integrates opengl into the engine

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              17            188             60            833
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                            18            188             60            834
-------------------------------------------------------------------------------
```
## TODO
Check `vulcan` and `wasm` as opengl alternatives

## Dependencies
`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.NewThreadSafeTrackingArray`
  - `engine/modules/datastructures.SparseArray`
  - `engine/modules/datastructures.TrackingArray`

`engine/modules/graphics`:
  - `engine/modules/graphics.Configure`
  - `engine/modules/graphics.EBO`
  - `engine/modules/graphics.ErrInvalidLocation`
  - `engine/modules/graphics.ErrNotALocation`
  - `engine/modules/graphics.ErrProgramHasOtherLocations`
  - `engine/modules/graphics.ErrTexturesHaveToShareSize`
  - `engine/modules/graphics.FlipV`
  - `engine/modules/graphics.Image`
  - `engine/modules/graphics.Index`
  - `engine/modules/graphics.Name`
  - `engine/modules/graphics.Parameter`
  - `engine/modules/graphics.Program`
  - `engine/modules/graphics.Release`
  - `engine/modules/graphics.Service`
  - `engine/modules/graphics.Shader`
  - `engine/modules/graphics.Texture`
  - `engine/modules/graphics.TextureArray`
  - `engine/modules/graphics.TextureArrayFactory`
  - `engine/modules/graphics.TextureFactory`
  - `engine/modules/graphics.VAO`
  - `engine/modules/graphics.VBO`
  - `engine/modules/graphics.Value`

### Third Party
- `github.com/go-gl/gl/v4.5-core/gl`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/image/draw`