# Tile module
## Architecture
This module demonstrates the project's focus on performance. See [Benchmarks](#Benchmarks).
This module contains:
- dual-grid system **renderer**
- `.biom` extension
- lazy mapping from tile `Pos`, `Size`, `Rot` to transform `Pos`, `Size`, `Rot`

Currently, this is a demo version.\
It stores the entire grid in a contiguous slice and in the future it'll store the world in chunks
which will allow basic optimizations and the ability to store chunks on disk while loading the necessary into memory.

### Biom extension
Integration with `entityregistry` allows us to define biome assets in a **single** `struct tag`.
```go
type Tiles struct {
	BiomName    ecs.EntityID `path:"tiles/biom_directory.biom"`
}
```
The snippet trims suffix (`.biom`) and expects a path without suffix (`tiles/biom_directory`)
to be a directory and contain 6 directories with names from `1` to `6` each with different shapes
where each shape can have any amount of tile variants (variants can only have `.png` extension).\
To minimize the number of assets needed, images are flipped to make 16 images from 6 images.\
During flipping, axes are never swapped only `Y-axis` can become `-Y` and `X-axis` can become `-X`.

```
tiles/biom_directory/
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
```
type Tiles struct {
	BiomName    ecs.EntityID `path:"tiles/biom_directory.biom"`
}
```
and biome file structure like [example file structure presented](#biome-extension).

### Under the hood
`.biom` extension is translated to:
```go
type BiomAsset interface {
	Images() [15][]image.Image
	Res() image.Rectangle
	AspectRatio() image.Rectangle
	Release()
}
```

## Benchmarks
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

## Usage examples
```go
type Service interface {
	ecs.SystemRegister
	Renderer() ecs.SystemRegister

	Component() ecs.ComponentsArray[Component]
	Grid() ecs.ComponentsArray[grid.SquareGridComponent[ID]]
	GetTile(ID) (ecs.EntityID, bool)

	Pos() ecs.ComponentsArray[PosComponent]
	Size() ecs.ComponentsArray[SizeComponent]
	Rot() ecs.ComponentsArray[RotComponent]
	Layer() ecs.ComponentsArray[LayerComponent]

	// src images should be:
	// - 1111
	// - 1110
	// - 1010
	// - 1001
	// - 0001
	NewBiomAsset(srcImages [6][]image.Image) (BiomAsset, error)

	GetPos(coords grid.Coords) transform.PosComponent
	// transform 1x1 tile size.
	// can be used for graphics or collisions.
	GetTileSize() transform.SizeComponent
}
```

### `Component`
Refers specifically to biome

### `Grid`
Stores whole grid in continuous slice

### `GetTile`
Gets entity with tile metadata by its id

### `Pos`
Entities with this are positioned on specific tile position

### `Size`
Entities with this are positioned to cover `X`x`Y` tiles

### `Rot`
Entities with this are rotated in 2D on grid axis

### `Layer`
Entities with this have specific `z`.

### `NewBiomAsset`
Parses [biome file structure](#biome-extension) into underlying [data structure](#under-the-hood) used for biomes

### `GetPos`
Function used by [pos component](#pos) to determine `transform.Pos` from `tile.Pos`

### `GetTileSize`
Function used by [size component](#size) to determine `transform.Size` from `tile.Size`
