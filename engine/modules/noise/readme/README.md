# noise
## Architecture
generates normalized noises.
extremely useful for generating maps

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/noise/test	1.050s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              11             93             70            392
Markdown                         2              3              0             45
-------------------------------------------------------------------------------
SUM:                            13             96             70            437
-------------------------------------------------------------------------------
```
## TODO
Improve noise flatenning bell curve.
Avoid equasions for flattening and instead create 10k sample slice to create a bell curve.

## Dependencies
`engine`:
  - `engine.EngineWorld`

`engine/modules/noise`:
  - `engine/modules/noise.CellSize`
  - `engine/modules/noise.Factory`
  - `engine/modules/noise.LayerConfig`
  - `engine/modules/noise.NewLayer`
  - `engine/modules/noise.NewNoise`
  - `engine/modules/noise.Noise`
  - `engine/modules/noise.Read`
  - `engine/modules/noise.Service`
  - `engine/modules/noise.Weight`

`engine/modules/seed`:
  - `engine/modules/seed.New`
  - `engine/modules/seed.Seed`
  - `engine/modules/seed.Value`

`engine/modules/transition`:
  - `engine/modules/transition.Lerp`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/go-gl/mathgl/mgl64`
- `github.com/ogiusek/ioc/v2`