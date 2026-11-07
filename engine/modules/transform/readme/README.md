# transform
## Architecture
This module contains many components which allow us to define relative position
and this module transforms these relative components to absolute position which
is used in collisions and rendering

## Benchmarks
```
$ go test ./... -bench=.
goos: linux
goarch: amd64
pkg: engine/modules/transform/test
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkGetPos-8                 	42703231	        28.15 ns/op
BenchmarkRawGetPos-8              	100000000	        10.90 ns/op
BenchmarkSetAbsolutePos-8         	 1946289	       617.0 ns/op
BenchmarkSetAndGetAbsolutePos-8   	 2051758	       590.7 ns/op
PASS
ok  	engine/modules/transform/test	5.952s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              14            163             23            792
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                            15            163             23            793
-------------------------------------------------------------------------------
```
## TODO
Start using Fixed-Point math

## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Hierarchy`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.AnyComponentArray`
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`
  - `engine/modules/ecs.World`

`engine/modules/hierarchy`:
  - `engine/modules/hierarchy.Children`
  - `engine/modules/hierarchy.Parent`
  - `engine/modules/hierarchy.Service`

`engine/modules/transform`:
  - `engine/modules/transform.AbsolutePos`
  - `engine/modules/transform.AbsolutePosComponent`
  - `engine/modules/transform.AbsoluteRotationComponent`
  - `engine/modules/transform.AbsoluteSize`
  - `engine/modules/transform.AbsoluteSizeComponent`
  - `engine/modules/transform.AspectRatio`
  - `engine/modules/transform.AspectRatioComponent`
  - `engine/modules/transform.InheritComponent`
  - `engine/modules/transform.MaxSizeComponent`
  - `engine/modules/transform.MinSizeComponent`
  - `engine/modules/transform.NewAspectRatio`
  - `engine/modules/transform.NewInherit`
  - `engine/modules/transform.NewMaxSize`
  - `engine/modules/transform.NewMinSize`
  - `engine/modules/transform.NewParentPivotPoint`
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.NewRotation`
  - `engine/modules/transform.NewSize`
  - `engine/modules/transform.ParentPivotPointComponent`
  - `engine/modules/transform.PivotPointComponent`
  - `engine/modules/transform.Point`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.PosComponent`
  - `engine/modules/transform.PrimaryAxis`
  - `engine/modules/transform.PrimaryAxisX`
  - `engine/modules/transform.PrimaryAxisY`
  - `engine/modules/transform.PrimaryAxisZ`
  - `engine/modules/transform.RelativeMask`
  - `engine/modules/transform.RelativePos`
  - `engine/modules/transform.RelativeRotation`
  - `engine/modules/transform.RelativeSizeX`
  - `engine/modules/transform.RelativeSizeXYZ`
  - `engine/modules/transform.RelativeSizeY`
  - `engine/modules/transform.RelativeSizeZ`
  - `engine/modules/transform.Rotation`
  - `engine/modules/transform.RotationComponent`
  - `engine/modules/transform.Service`
  - `engine/modules/transform.Size`
  - `engine/modules/transform.SizeComponent`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/ioc/v2`