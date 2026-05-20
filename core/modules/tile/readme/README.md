# tile
## Architecture
This module contains:
- dual-grid system **renderer**
- `.biome` extension
- lazy mapping from tile `Pos`, `Size`, `Rot` to transform `Pos`, `Size`, `Rot`

Current version stores the entire grid in a contiguous slice and in the future it'll store the world in chunks
which will allow basic optimizations and the ability to store chunks on disk while loading the necessary into memory.

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
Flag benchmark:
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

Standard benchmark:
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

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               9            180             69            976
GLSL                             3             29              2            114
Markdown                         2              9              0             79
-------------------------------------------------------------------------------
SUM:                            14            218             71           1169
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `core/modules/tile.Service`

#### method Service Component
Type: `func() engine/services/ecs.ComponentsArray[core/modules/tile.Component]`

#### method Service GetPos
Type: `func(coords engine/modules/grid.Coords) engine/modules/transform.PosComponent`

#### method Service GetTile
Type: `func(core/modules/tile.ID) (engine/services/ecs.EntityID, bool)`

#### method Service GetTileSize
Type: `func() engine/modules/transform.SizeComponent`
transform 1x1 tile size.
can be used for graphics or collisions.

#### method Service Grid
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/grid.SquareGridComponent[core/modules/tile.ID]]`

#### method Service Layer
Type: `func() engine/services/ecs.ComponentsArray[core/modules/tile.LayerComponent]`

#### method Service NewBiomeAsset
Type: `func(srcImages [6][]image.Image) (core/modules/tile.BiomeAsset, error)`
src images should be:
- 1111
- 1110
- 1010
- 1001
- 0001

#### method Service Pos
Type: `func() engine/services/ecs.ComponentsArray[core/modules/tile.PosComponent]`

#### method Service Register
Type: `func() error`

#### method Service Renderer
Type: `func() engine/services/ecs.SystemRegister`

#### method Service Rot
Type: `func() engine/services/ecs.ComponentsArray[core/modules/tile.RotComponent]`

#### method Service Size
Type: `func() engine/services/ecs.ComponentsArray[core/modules/tile.SizeComponent]`

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

### type SelectEvent
Type: `core/modules/tile.SelectEvent`
changes event emitted on tile hover

#### property SelectEvent HoverEvent
Type: `any`

### type HoverEvent
Type: `core/modules/tile.HoverEvent`

#### property HoverEvent Grid
Type: `engine/services/ecs.EntityID`

#### property HoverEvent Tile
Type: `engine/modules/grid.Index`

### type ClickEntityEvent
Type: `core/modules/tile.ClickEntityEvent`

#### property ClickEntityEvent Entity
Type: `engine/services/ecs.EntityID`

#### method ClickEntityEvent ApplyEntity
Type: `func(entity engine/services/ecs.EntityID) any`

## Variables
### var ErrInvalidPosition
Type: `error`
error logged when grid.GetIndex returns !ok

### var ErrInvalidStep
Type: `error`

### var ErrPositionAndSpeedIsRequiredToStep
Type: `error`

### var Tau
Type: `untyped float`

## Functions
### func NewGrid
Type: `func(w engine/modules/grid.Coord, h engine/modules/grid.Coord) engine/modules/grid.SquareGridComponent[core/modules/tile.ID]`

### func NewTile
Type: `func(id core/modules/tile.ID) core/modules/tile.Component`

### func NewPos
Type: `func[Number golang.org/x/exp/constraints.Integer | golang.org/x/exp/constraints.Float](x Number, y Number) core/modules/tile.PosComponent`

### func NewLayer
Type: `func[Number golang.org/x/exp/constraints.Integer | golang.org/x/exp/constraints.Float](z Number) core/modules/tile.LayerComponent`

### func NewSize
Type: `func[Number golang.org/x/exp/constraints.Integer](x Number, y Number) core/modules/tile.SizeComponent`

### func NewRot
Type: `func(radians float32) core/modules/tile.RotComponent`

### func NewSelectEvent
Type: `func(hoverEvent any) core/modules/tile.SelectEvent`

### func NewHoverEvent
Type: `func(grid engine/services/ecs.EntityID, tile engine/modules/grid.Index) any`

### func NewClickEntityEvent
Type: `func() core/modules/tile.ClickEntityEvent`


## Dependencies
`core/game`:
  - `core/game.Definitions`
  - `core/game.Deploy`
  - `core/game.GameWorld`
  - `core/game.Pathfind`
  - `core/game.Player`
  - `core/game.Tile`
  - `core/game.Ui`

`core/modules/definitions`:
  - `core/modules/definitions.Assets`
  - `core/modules/definitions.Btn`
  - `core/modules/definitions.ConstructLayer`
  - `core/modules/definitions.GameGroup`
  - `core/modules/definitions.Hud`
  - `core/modules/definitions.Selected`
  - `core/modules/definitions.SquareCollider`
  - `core/modules/definitions.SquareMesh`
  - `core/modules/definitions.Text`
  - `core/modules/definitions.TileLayer`
  - `core/modules/definitions.UnitLayer`

`core/modules/deploy`:
  - `core/modules/deploy.Component`
  - `core/modules/deploy.Deployable`
  - `core/modules/deploy.NewSelectEvent`

`core/modules/pathfind`:
  - `core/modules/pathfind.NewSelectEvent`
  - `core/modules/pathfind.Speed`

`core/modules/tile`:
  - `core/modules/tile.ApplyCoords`
  - `core/modules/tile.ApplyCoordsEvent`
  - `core/modules/tile.BiomeAsset`
  - `core/modules/tile.ClickEntityEvent`
  - `core/modules/tile.Component`
  - `core/modules/tile.Coord`
  - `core/modules/tile.Entity`
  - `core/modules/tile.GetTileSize`
  - `core/modules/tile.Grid`
  - `core/modules/tile.HoverEvent`
  - `core/modules/tile.ID`
  - `core/modules/tile.Images`
  - `core/modules/tile.Layer`
  - `core/modules/tile.LayerComponent`
  - `core/modules/tile.NewBiomeAsset`
  - `core/modules/tile.NewHoverEvent`
  - `core/modules/tile.NewLayer`
  - `core/modules/tile.NewRot`
  - `core/modules/tile.NewSize`
  - `core/modules/tile.NewTile`
  - `core/modules/tile.Pos`
  - `core/modules/tile.PosComponent`
  - `core/modules/tile.Quat`
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

`core/modules/ui`:
  - `core/modules/ui.Entities`
  - `core/modules/ui.NewSelect`
  - `core/modules/ui.ObjectComponent`
  - `core/modules/ui.Objects`
  - `core/modules/ui.SelectEvent`
  - `core/modules/ui.ShowMenu`
  - `core/modules/ui.UnselectEvent`

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
  - `engine/modules/grid.Component`
  - `engine/modules/grid.Coord`
  - `engine/modules/grid.Coords`
  - `engine/modules/grid.GetCoords`
  - `engine/modules/grid.GetTiles`
  - `engine/modules/grid.Height`
  - `engine/modules/grid.Index`
  - `engine/modules/grid.NewCoords`
  - `engine/modules/grid.NewSquareGrid`
  - `engine/modules/grid.Service`
  - `engine/modules/grid.SquareGridComponent`
  - `engine/modules/grid.Width`
  - `engine/modules/grid.X`
  - `engine/modules/grid.Y`

`engine/modules/grid/pkg`:
  - `engine/modules/grid/pkg.NewConfig`
  - `engine/modules/grid/pkg.PkgT`

`engine/modules/groups`:
  - `engine/modules/groups.Component`
  - `engine/modules/groups.EmptyGroups`
  - `engine/modules/groups.Enable`
  - `engine/modules/groups.Ptr`
  - `engine/modules/groups.SharesAnyGroup`
  - `engine/modules/groups.Val`

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

`engine/modules/text`:
  - `engine/modules/text.Content`
  - `engine/modules/text.NewText`

`engine/modules/transform`:
  - `engine/modules/transform.Mat4`
  - `engine/modules/transform.NewParent`
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.NewRotation`
  - `engine/modules/transform.NewSize`
  - `engine/modules/transform.Parent`
  - `engine/modules/transform.PivotPoint`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.PosComponent`
  - `engine/modules/transform.RelativePos`
  - `engine/modules/transform.RelativeSizeXYZ`
  - `engine/modules/transform.Rotation`
  - `engine/modules/transform.Size`
  - `engine/modules/transform.SizeComponent`

`engine/modules/transition`:
  - `engine/modules/transition.Lerp`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/services/datastructures`:
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.GetIndices`
  - `engine/services/datastructures.NewSparseArray`
  - `engine/services/datastructures.Remove`
  - `engine/services/datastructures.Set`
  - `engine/services/datastructures.Size`
  - `engine/services/datastructures.SparseArray`

`engine/services/ecs`:
  - `engine/services/ecs.AddDirtySet`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.DirtySet`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.NewDirtySet`
  - `engine/services/ecs.NewEntity`
  - `engine/services/ecs.NewSystemRegister`
  - `engine/services/ecs.OnRemove`
  - `engine/services/ecs.OnUpsert`
  - `engine/services/ecs.RegisterSystems`
  - `engine/services/ecs.Remove`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SetEmpty`
  - `engine/services/ecs.SystemRegister`
  - `engine/services/ecs.World`

### Third Party
- `github.com/go-gl/gl/v4.5-core/gl`
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `golang.org/x/exp/constraints`