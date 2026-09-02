# tile
## Architecture
This module contains:
- dual-grid system **renderer**
- `.biome` extension
- lazy mapping from tile `Pos`, `Size`, `Rot` to transform `Pos`, `Size`, `Rot`

### Biome extension
Integration with `entityregistry` allows us to define biome assets in a **single** `struct tag`.
```go
type Tiles struct {
	BiomeName    ecs.EntityID `path:"tiles/biome_directory.biome"`
}
```

The snippet trims suffix (`.biome`) and expects a path without suffix (`tiles/biome_directory`)
to be a directory and contain 6 directories with names from `1` to `6` each with different shapes
where each shape can have any amount of tile variants (variants can only have `.png` extension).\
To minimize the number of assets needed, images are flipped to make 16 images from 6 images.\
During flipping, axes are never swapped only `Y-axis` can become `-Y` and `X-axis` can become `-X`.

```
tiles/biome_directory/
├─ 1/         # first shape directory
│   ├── a.png # variant `a`
│   ├── b.png # variant `b`
│   └── ...   # we can go on with as many variants as we want with any names.
│             # Alphabetical variant naming is recommended
│
├─ 2/         # second shape directory with its own variants
│   └── ...   # variants...
├─ 3/         # third shape directory with its own variants
│   └── ...   # variants...
├─ 4/         # fourth shape directory with its own variants
│   └── ...   # variants...
├─ 5/         # fifth shape directory with its own variants
│   └── ...   # variants...
└─ 6/         # sixth shape directory with its own variants
    └── ...   # variants...
```

Expected shapes in directories from `1` to `6`:
- ![first image](grass/1/a.png)
- ![second image](grass/2/a.png)
- ![third image](grass/3/a.png)
- ![fourth image](grass/4/a.png)
- ![fifth image](grass/5/a.png)
- ![sixth image](grass/6/a.png)

Example [biome file structure](grass/).

### How to import a biome
One line using `entityregistry` is enough in codebase:
```go
type Tiles struct {
	BiomeName    ecs.EntityID `path:"tiles/biome_directory.biome"`
}
```

and biome file structure like [example file structure presented](#biome-extension).

## Benchmarks
### Flag benchmark
```sh
goos: windows
goarch: amd64
pkg: core/modules/tile/test
cpu: Intel(R) Core(TM) i7-14700KF
BenchmarkRendering36MTilesMap
gpu: Meta Virtual Monitor
gpu 2: NVIDIA GeForce RTX 4080 SUPER
BenchmarkRendering36MTilesMap-28    	     135	   8510424 ns/op
PASS
ok  	core/modules/tile/test	57.055s
```
Rendering 36 million tiles on `NVIDIA GeForce RTX 4080 SUPER` in less than **8.6ms**.

### Standard benchmark
```sh
$ go test . -bench=. -benchtime=10s
Failed to load plugin 'libdecor-gtk.so': failed to init
gpu: Kaby Lake-R GT2 [UHD Graphics 620]
goos: linux
goarch: amd64
pkg: core/modules/tile/test
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkRendering1MTilesMap-8   	    2221	   5062408 ns/op
```
Rendering 1 million tiles on `UHD Graphics 620` in less than **5.1ms**.

### Custom benchmark
To run benchmark on your machine with map with custom size then find `BenchmarkRendering1MTilesMap`.
```go
func BenchmarkRendering1MTilesMap(b *testing.B) { benchmarkRenderingXTilesMap(b, 1000) }
func BenchmarkRendering4MTilesMap(b *testing.B) { benchmarkRenderingXTilesMap(b, 2000) }
```

These benchmarks create map with custom size `n`*`n` (in first example 1,000*1,000 = 1,000,000).

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               8            183             95           1000
GLSL                             3             31              2            112
Markdown                         3             10              0             85
-------------------------------------------------------------------------------
SUM:                            14            224             97           1197
-------------------------------------------------------------------------------
```
## TODO
Currently animated tiles aren't supported and there is a big chance that they won't be supported

## Types
### type Service
Type: `core/modules/tile.Service`

#### method Service Blueprint
Type: `func() engine/modules/uuid.LinkService[core/modules/tile.BlueprintLink]`

#### method Service Component
Type: `func() engine/modules/ecs.ComponentArray[core/modules/tile.Component]`

#### method Service GetPos
Type: `func(coords engine/modules/grid.Coords) engine/modules/transform.PosComponent`

#### method Service GetTile
Type: `func(core/modules/tile.ID) (engine/modules/ecs.EntityID, bool)`

#### method Service Grid
Type: `func() engine/modules/grid.ServiceT[core/modules/tile.ID]`

#### method Service Layer
Type: `func() engine/modules/ecs.ComponentArray[core/modules/tile.LayerComponent]`

#### method Service Name
Type: `func() engine/modules/ecs.ComponentArray[core/modules/tile.NameComponent]`

#### method Service NewBiomeAsset
Type: `func(srcImages [6][]image.Image) (core/modules/tile.BiomeAsset, error)`
src images should be:
- 1111
- 1110
- 1010
- 1001
- 0001

#### method Service Pos
Type: `func() engine/modules/ecs.ComponentArray[core/modules/tile.PosComponent]`
is often used to check is element a prototype or a placed object

#### method Service Register
Type: `func() error`

#### method Service Renderer
Type: `func() engine/modules/ecs.SystemRegister`

#### method Service Rot
Type: `func() engine/modules/ecs.ComponentArray[core/modules/tile.RotComponent]`

#### method Service Size
Type: `func() engine/modules/ecs.ComponentArray[core/modules/tile.SizeComponent]`

### type BiomeAsset
Type: `core/modules/tile.BiomeAsset`

#### method BiomeAsset AspectRatio
Type: `func() image.Rectangle`

#### method BiomeAsset Images
Type: `func() [15][]image.Image`

#### method BiomeAsset Release
Type: `func()`

#### method BiomeAsset Res
Type: `func() image.Rectangle`

### type ApplyCoordsEvent
Type: `core/modules/tile.ApplyCoordsEvent`

#### method ApplyCoordsEvent ApplyCoords
Type: `func(engine/modules/grid.Coords) any`

### type ID
Type: `core/modules/tile.ID`

### type Component
Type: `core/modules/tile.Component`

#### property Component ID
Type: `core/modules/tile.ID`

### type Coord
Type: `core/modules/tile.Coord`

### type PosComponent
Type: `core/modules/tile.PosComponent`

#### property PosComponent X
Type: `core/modules/tile.Coord`

#### property PosComponent Y
Type: `core/modules/tile.Coord`

#### method PosComponent Smooth
Type: `func()`

#### method PosComponent Lerp
Type: `func(c2 core/modules/tile.PosComponent, mix32 float32) core/modules/tile.PosComponent`

#### method PosComponent Aligned
Type: `func() (coords engine/modules/grid.Coords, isAligned bool)`

### type LayerComponent
Type: `core/modules/tile.LayerComponent`

#### property LayerComponent Z
Type: `core/modules/tile.Coord`

### type NameComponent
Type: `core/modules/tile.NameComponent`

#### property NameComponent Name
Type: `string`

### type BlueprintLink
Type: `core/modules/tile.BlueprintLink`

### type SizeComponent
Type: `core/modules/tile.SizeComponent`

#### property SizeComponent X
Type: `engine/modules/grid.Coord`

#### property SizeComponent Y
Type: `engine/modules/grid.Coord`

#### method SizeComponent Size
Type: `func() (engine/modules/grid.Coord, engine/modules/grid.Coord)`

### type RotComponent
Type: `core/modules/tile.RotComponent`

#### property RotComponent Radians
Type: `float32`

#### method RotComponent Smooth
Type: `func()`

#### method RotComponent Lerp
Type: `func(c2 core/modules/tile.RotComponent, mix32 float32) core/modules/tile.RotComponent`

#### method RotComponent Quat
Type: `func() github.com/go-gl/mathgl/mgl32.Quat`

### type MissingChunkEvent
Type: `core/modules/tile.MissingChunkEvent`

#### property MissingChunkEvent Coords
Type: `engine/modules/grid.ChunkCoordsComponent`

### type UnloadChunkEvent
Type: `core/modules/tile.UnloadChunkEvent`

#### property UnloadChunkEvent Coords
Type: `engine/modules/grid.ChunkCoordsComponent`

### type ClickEntityEvent
Type: `core/modules/tile.ClickEntityEvent`

#### property ClickEntityEvent Entity
Type: `engine/modules/ecs.EntityID`

#### method ClickEntityEvent ApplyEntity
Type: `func(entity engine/modules/ecs.EntityID) any`

### type ClickBlueprintEvent
Type: `core/modules/tile.ClickBlueprintEvent`

#### property ClickBlueprintEvent Entity
Type: `engine/modules/ecs.EntityID`

## Variables
### var ErrInvalidPosition
Type: `error`
error logged when grid.GetIndex returns !ok

### var ErrInvalidStep
Type: `error`

### var ErrPositionAndSpeedIsRequiredToStep
Type: `error`

### var ErrBlueprintIsMissingUUID
Type: `error`

### var Tau
Type: `untyped float`

## Functions
### func NewTile
Type: `func(id core/modules/tile.ID) core/modules/tile.Component`

### func NewPos
Type: `func[Number golang.org/x/exp/constraints.Integer | golang.org/x/exp/constraints.Float](x Number, y Number) core/modules/tile.PosComponent`

### func NewLayer
Type: `func[Number golang.org/x/exp/constraints.Integer | golang.org/x/exp/constraints.Float](z Number) core/modules/tile.LayerComponent`

### func NewName
Type: `func(name string) core/modules/tile.NameComponent`

### func NewSize
Type: `func[Number golang.org/x/exp/constraints.Integer](x Number, y Number) core/modules/tile.SizeComponent`

### func NewRot
Type: `func(radians float32) core/modules/tile.RotComponent`

### func NewMissingChunkEvent
Type: `func(coords engine/modules/grid.ChunkCoordsComponent) core/modules/tile.MissingChunkEvent`

### func NewUnloadChunkEvent
Type: `func(coords engine/modules/grid.ChunkCoordsComponent) core/modules/tile.UnloadChunkEvent`

### func NewClickEntityEvent
Type: `func() core/modules/tile.ClickEntityEvent`

### func NewClickBlueprintEvent
Type: `func(deployed engine/modules/ecs.EntityID) core/modules/tile.ClickBlueprintEvent`


## Dependencies
`core/game`:
  - `core/game.Definitions`
  - `core/game.GameWorld`
  - `core/game.Obstruction`
  - `core/game.Tile`

`core/modules/definitions`:
  - `core/modules/definitions.Assets`
  - `core/modules/definitions.ConstructLayer`
  - `core/modules/definitions.SquareCollider`
  - `core/modules/definitions.SquareMesh`
  - `core/modules/definitions.TileLayer`
  - `core/modules/definitions.UnitLayer`

`core/modules/tile`:
  - `core/modules/tile.BiomeAsset`
  - `core/modules/tile.Blueprint`
  - `core/modules/tile.BlueprintLink`
  - `core/modules/tile.Component`
  - `core/modules/tile.Coord`
  - `core/modules/tile.Grid`
  - `core/modules/tile.ID`
  - `core/modules/tile.Images`
  - `core/modules/tile.Layer`
  - `core/modules/tile.LayerComponent`
  - `core/modules/tile.Name`
  - `core/modules/tile.NameComponent`
  - `core/modules/tile.NewBiomeAsset`
  - `core/modules/tile.NewClickEntityEvent`
  - `core/modules/tile.NewLayer`
  - `core/modules/tile.NewName`
  - `core/modules/tile.NewRot`
  - `core/modules/tile.NewSize`
  - `core/modules/tile.NewTile`
  - `core/modules/tile.Pos`
  - `core/modules/tile.PosComponent`
  - `core/modules/tile.Quat`
  - `core/modules/tile.Rot`
  - `core/modules/tile.RotComponent`
  - `core/modules/tile.Service`
  - `core/modules/tile.Size`
  - `core/modules/tile.SizeComponent`
  - `core/modules/tile.X`
  - `core/modules/tile.Y`
  - `core/modules/tile.Z`

`engine/modules/assets`:
  - `engine/modules/assets.Asset`
  - `engine/modules/assets.GetAsset`
  - `engine/modules/assets.Path`
  - `engine/modules/assets.PathComponent`
  - `engine/modules/assets.Register`
  - `engine/modules/assets.Service`

`engine/modules/collider`:
  - `engine/modules/collider.Component`
  - `engine/modules/collider.NewCollider`

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
  - `engine/modules/ecs.World`

`engine/modules/entityregistry`:
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
  - `engine/modules/graphics.Get`
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
  - `engine/modules/graphics.Opaque`
  - `engine/modules/graphics.Program`
  - `engine/modules/graphics.Release`
  - `engine/modules/graphics.Scale`
  - `engine/modules/graphics.Set`
  - `engine/modules/graphics.TextureArray`
  - `engine/modules/graphics.VAO`
  - `engine/modules/graphics.VBOFactory`
  - `engine/modules/graphics.VBOSetter`
  - `engine/modules/graphics.VertexShader`

`engine/modules/grid`:
  - `engine/modules/grid.Chunk`
  - `engine/modules/grid.ChunkCoordsComponent`
  - `engine/modules/grid.ChunkSize`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.GetChunk`
  - `engine/modules/grid.GetTileSize`
  - `engine/modules/grid.GetTiles`
  - `engine/modules/grid.NewChunkCoords`
  - `engine/modules/grid.NewCoord`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.ServiceT`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/modules/grid/pkg`:
  - `engine/modules/grid/pkg.PkgT`

`engine/modules/inputs`:
  - `engine/modules/inputs.LeftClick`
  - `engine/modules/inputs.NewLeftClick`
  - `engine/modules/inputs.Stack`

`engine/modules/relation`:
  - `engine/modules/relation.Get`
  - `engine/modules/relation.Service`

`engine/modules/relation/pkg`:
  - `engine/modules/relation/pkg.SpatialRelationPkg`

`engine/modules/render`:
  - `engine/modules/render.Camera`
  - `engine/modules/render.ErrTextureAssetImagesHasToMatchResolution`
  - `engine/modules/render.Mesh`
  - `engine/modules/render.NewMesh`
  - `engine/modules/render.NewTexture`
  - `engine/modules/render.RenderEvent`
  - `engine/modules/render.Texture`

`engine/modules/transform`:
  - `engine/modules/transform.Mat4`
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.NewRotation`
  - `engine/modules/transform.NewSize`
  - `engine/modules/transform.PivotPoint`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.PosComponent`
  - `engine/modules/transform.Rotation`
  - `engine/modules/transform.Size`
  - `engine/modules/transform.SizeComponent`

`engine/modules/transition`:
  - `engine/modules/transition.Lerp`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.Get`
  - `engine/modules/uuid.LinkService`
  - `engine/modules/uuid.New`
  - `engine/modules/uuid.NewUUID`
  - `engine/modules/uuid.UUID`

`engine/modules/uuid/pkg`:
  - `engine/modules/uuid/pkg.LinkPkgT`

### Third Party
- `github.com/go-gl/gl/v4.5-core/gl`
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`