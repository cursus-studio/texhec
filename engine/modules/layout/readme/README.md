# layout
## Architecture
positions and re-size children

## Types
### type Service
Type: `engine/modules/layout.Service`

#### method Service Align
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/layout.AlignComponent]`

#### method Service Gap
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/layout.GapComponent]`
Wrap() ecs.ComponentsArray[WrapComponent]

#### method Service Order
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/layout.OrderComponent]`

### type AlignComponent
Type: `engine/modules/layout.AlignComponent`
centering
Y axis is reversed for primary axis

#### property AlignComponent Primary
Type: `float32`
value between 0 and 1 where 0 means aligned to left and 1 aligned to right

#### property AlignComponent Secondary
Type: `float32`
value between 0 and 1 where 0 means aligned to left and 1 aligned to right

### type Order
Type: `engine/modules/layout.Order`
order

### type OrderComponent
Type: `engine/modules/layout.OrderComponent`

#### property OrderComponent Order
Type: `engine/modules/layout.Order`
default horizontal

#### method OrderComponent Primary
Type: `func() engine/modules/layout.Order`

#### method OrderComponent Secondary
Type: `func() engine/modules/layout.Order`

### type GapComponent
Type: `engine/modules/layout.GapComponent`
gaps

#### property GapComponent Gap
Type: `float32`

## Variables
### var OrderHorizontal
Type: `engine/modules/layout.Order`

### var OrderVectical
Type: `engine/modules/layout.Order`

## Functions
### func NewAlign
Type: `func(primary float32, secondary float32) engine/modules/layout.AlignComponent`

### func NewOrder
Type: `func(order engine/modules/layout.Order) engine/modules/layout.OrderComponent`

### func NewGap
Type: `func(x float32) engine/modules/layout.GapComponent`


## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/layout/test	0.010s
```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Hierarchy`
  - `engine.Transform`
  - `engine.World`

`engine/modules/layout`:
  - `engine/modules/layout.AlignComponent`
  - `engine/modules/layout.Gap`
  - `engine/modules/layout.GapComponent`
  - `engine/modules/layout.NewAlign`
  - `engine/modules/layout.NewGap`
  - `engine/modules/layout.Order`
  - `engine/modules/layout.OrderComponent`
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

`engine/services/ecs`:
  - `engine/services/ecs.AddDependency`
  - `engine/services/ecs.AddDirtySet`
  - `engine/services/ecs.BeforeGet`
  - `engine/services/ecs.Clear`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.Dirty`
  - `engine/services/ecs.DirtySet`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.NewDirtySet`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SetEmpty`

### Third Party
`github.com/go-gl/mathgl/mgl32`
`github.com/ogiusek/ioc/v2`