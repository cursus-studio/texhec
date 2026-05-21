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
-------------------------------------------------------------------------------
SUM:                            17            188             60            833
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/graphics.Service`

#### method Service NewEBO
Type: `func() engine/modules/graphics.EBO`
NewVBO[Vertex any](configure func()) VBOSetter[Vertex]

#### method Service NewImage
Type: `func(image image.Image) engine/modules/graphics.Image`

#### method Service NewProgram
Type: `func(p uint32, parameters []engine/modules/graphics.Parameter) (engine/modules/graphics.Program, error)`

#### method Service NewShader
Type: `func(shaderSource string, shaderType uint32) (engine/modules/graphics.Shader, error)`

#### method Service NewVAO
Type: `func(VBO engine/modules/graphics.VBO, EBO engine/modules/graphics.EBO) engine/modules/graphics.VAO`

#### method Service Texture
Type: `func() engine/modules/graphics.TextureFactory`

#### method Service TextureArray
Type: `func() engine/modules/graphics.TextureArrayFactory`

### type TextureArrayFactory
Type: `engine/modules/graphics.TextureArrayFactory`

#### method TextureArrayFactory New
Type: `func(engine/services/datastructures.SparseArray[uint32, image.Image]) (engine/modules/graphics.TextureArray, error)`

#### method TextureArrayFactory NewFromSlice
Type: `func([]image.Image) (engine/modules/graphics.TextureArray, error)`

#### method TextureArrayFactory Wrap
Type: `func(func(engine/modules/graphics.TextureArray))`

### type TextureFactory
Type: `engine/modules/graphics.TextureFactory`

#### method TextureFactory New
Type: `func(image.Image) (engine/modules/graphics.Texture, error)`

#### method TextureFactory Wrap
Type: `func(func(engine/modules/graphics.Texture))`

### type Program
Type: `engine/modules/graphics.Program`

#### method Program Bind
Type: `func()`

#### method Program ID
Type: `func() uint32`

#### method Program Locations
Type: `func(reflect.Type) (any, error)`

#### method Program Release
Type: `func()`

### type Buffer
Type: `engine/modules/graphics.Buffer[Stored comparable]`

#### method Buffer Add
Type: `func(elements ...Stored)`

#### method Buffer Bind
Type: `func()`

#### method Buffer Flush
Type: `func()`

#### method Buffer Get
Type: `func() []Stored`

#### method Buffer ID
Type: `func() uint32`

#### method Buffer Release
Type: `func()`

#### method Buffer Remove
Type: `func(indices ...int)`

#### method Buffer Set
Type: `func(index int, e Stored)`

### type Shader
Type: `engine/modules/graphics.Shader`

#### method Shader ID
Type: `func() uint32`

#### method Shader Release
Type: `func()`

### type Image
Type: `engine/modules/graphics.Image`

#### method Image FlipH
Type: `func() engine/modules/graphics.Image`

#### method Image FlipHV
Type: `func() engine/modules/graphics.Image`
horizontally and vertically

#### method Image FlipV
Type: `func() engine/modules/graphics.Image`

#### method Image Image
Type: `func() image.Image`

#### method Image Opaque
Type: `func() engine/modules/graphics.Image`

#### method Image RotateClockwise
Type: `func(times int) engine/modules/graphics.Image`
rotates 90 deg clockwise

#### method Image Scale
Type: `func(w int, h int) engine/modules/graphics.Image`

#### method Image TrimTransparentBackground
Type: `func() engine/modules/graphics.Image`

### type Texture
Type: `engine/modules/graphics.Texture`

#### method Texture Bind
Type: `func()`

#### method Texture ID
Type: `func() uint32`

#### method Texture Release
Type: `func()`

### type TextureArray
Type: `engine/modules/graphics.TextureArray`

#### method TextureArray Bind
Type: `func()`

#### method TextureArray ImagesCount
Type: `func() int16`

#### method TextureArray Release
Type: `func()`

#### method TextureArray Texture
Type: `func() uint32`

### type VAO
Type: `engine/modules/graphics.VAO`

#### method VAO Bind
Type: `func()`

#### method VAO EBO
Type: `func() engine/modules/graphics.EBO`

#### method VAO ID
Type: `func() uint32`

#### method VAO Release
Type: `func()`

#### method VAO VBO
Type: `func() engine/modules/graphics.VBO`

### type VBO
Type: `engine/modules/graphics.VBO`

#### method VBO Configure
Type: `func()`

#### method VBO ID
Type: `func() uint32`

#### method VBO Len
Type: `func() int32`

#### method VBO Release
Type: `func()`

### type VBOSetter
Type: `engine/modules/graphics.VBOSetter[Vertex any]`

#### method VBOSetter Configure
Type: `func()`

#### method VBOSetter ID
Type: `func() uint32`

#### method VBOSetter Len
Type: `func() int32`

#### method VBOSetter Release
Type: `func()`

#### method VBOSetter SetVertices
Type: `func(vertices []Vertex) error`

### type EBO
Type: `engine/modules/graphics.EBO`

#### method EBO Configure
Type: `func()`

#### method EBO ID
Type: `func() uint32`

#### method EBO Len
Type: `func() int32`

#### method EBO Release
Type: `func()`

#### method EBO SetIndices
Type: `func(indices []engine/modules/graphics.Index) error`

### type Parameter
Type: `engine/modules/graphics.Parameter`

#### property Parameter Name
Type: `uint32`

#### property Parameter Value
Type: `int32`

### type VBOFactory
Type: `engine/modules/graphics.VBOFactory[Vertex any]`

### type Index
Type: `engine/modules/graphics.Index`

## Variables
### var ErrProgramHasOtherLocations
Type: `error`

### var ErrNotALocation
Type: `error`

### var ErrInvalidLocation
Type: `error`

### var VertexShader
Type: `uint32`

### var GeomShader
Type: `uint32`

### var FragmentShader
Type: `uint32`

### var ComputeShader
Type: `uint32`

### var ErrTexturesHaveToShareSize
Type: `error`

## Functions
### func NewVBO
Type: `func[Vertex any](configure func()) engine/modules/graphics.VBOSetter[Vertex]`
methods outside lean Service abstraction because of golang underdeveloped generics {

### func NewBuffer
Type: `func[Stored comparable](target uint32, usage uint32, index uint32) engine/modules/graphics.Buffer[Stored]`

### func GetProgramLocations
Type: `func[Locations any](p engine/modules/graphics.Program) (Locations, error)`


## Dependencies
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

`engine/services/datastructures`:
  - `engine/services/datastructures.Changes`
  - `engine/services/datastructures.ClearChanges`
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.GetIndices`
  - `engine/services/datastructures.GetValues`
  - `engine/services/datastructures.NewSparseArray`
  - `engine/services/datastructures.NewThreadSafeTrackingArray`
  - `engine/services/datastructures.Set`
  - `engine/services/datastructures.Size`
  - `engine/services/datastructures.SparseArray`
  - `engine/services/datastructures.TrackingArray`

### Third Party
- `github.com/go-gl/gl/v4.5-core/gl`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/image/draw`