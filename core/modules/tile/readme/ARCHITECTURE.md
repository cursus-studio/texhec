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
