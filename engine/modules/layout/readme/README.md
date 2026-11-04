# layout
## Architecture
positions and re-size children

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/layout/test	0.015s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             79             30            286
-------------------------------------------------------------------------------
SUM:                             6             79             30            286
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Hierarchy`
  - `engine.Transform`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`

`engine/modules/layout`:
  - `engine/modules/layout.AlignComponent`
  - `engine/modules/layout.Gap`
  - `engine/modules/layout.GapComponent`
  - `engine/modules/layout.NewAlign`
  - `engine/modules/layout.NewGap`
  - `engine/modules/layout.Order`
  - `engine/modules/layout.OrderComponent`
  - `engine/modules/layout.OrderHorizontal`
  - `engine/modules/layout.Primary`
  - `engine/modules/layout.Secondary`
  - `engine/modules/layout.Service`

`engine/modules/transform`:
  - `engine/modules/transform.AbsolutePos`
  - `engine/modules/transform.AbsoluteSize`
  - `engine/modules/transform.AddDirtySet`
  - `engine/modules/transform.NewParentPivotPoint`
  - `engine/modules/transform.NewPivotPoint`
  - `engine/modules/transform.NewPos`
  - `engine/modules/transform.ParentPivotPoint`
  - `engine/modules/transform.ParentPivotPointComponent`
  - `engine/modules/transform.PivotPoint`
  - `engine/modules/transform.PivotPointComponent`
  - `engine/modules/transform.Point`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.PosComponent`
  - `engine/modules/transform.Size`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/ioc/v2`