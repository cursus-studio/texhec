# noise
## Architecture
generates normalized noises.
extremely useful for generating maps

## Types
### type Service
Type: `engine/modules/noise.Service`

#### method Service NewNoise
Type: `func(engine/modules/seed.Seed) engine/modules/noise.Factory`

### type Noise
Type: `engine/modules/noise.Noise`

#### method Noise Read
Type: `func(github.com/go-gl/mathgl/mgl64.Vec2) float64`
returns normalized value <0, 1> with uniform distribution

### type Factory
Type: `engine/modules/noise.Factory`
each layer offset is `mgl64.Vec2{math.Pi, math.Pi}.Mul(layerIndex)`

#### method Factory AddPerlin
Type: `func(...engine/modules/noise.LayerConfig) engine/modules/noise.Factory`

#### method Factory AddValue
Type: `func(...engine/modules/noise.LayerConfig) engine/modules/noise.Factory`

#### method Factory Build
Type: `func() engine/modules/noise.Noise`

### type LayerConfig
Type: `engine/modules/noise.LayerConfig`

#### property LayerConfig CellSize
Type: `float64`
default size is 1

#### property LayerConfig Weight
Type: `float64`

## Functions
### func NewNoise
Type: `func(fn func(github.com/go-gl/mathgl/mgl64.Vec2) float64) engine/modules/noise.Noise`

### func NewLayer
Type: `func(cellSize float64, weight float64) engine/modules/noise.LayerConfig`


## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/noise/test	1.023s
```
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
`github.com/go-gl/mathgl/mgl64`
`github.com/ogiusek/ioc/v2`