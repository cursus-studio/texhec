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
BenchmarkGetPos-8                 	43297045	        26.70 ns/op
BenchmarkRawGetPos-8              	100000000	        10.97 ns/op
BenchmarkSetAbsolutePos-8         	 1842477	       633.9 ns/op
BenchmarkSetAndGetAbsolutePos-8   	 2019024	       610.9 ns/op
PASS
ok  	engine/modules/transform/test	5.940s
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

## Types
### type Service
Type: `engine/modules/transform.Service`

#### method Service AbsolutePos
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/transform.AbsolutePosComponent]`

#### method Service AbsoluteRotation
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/transform.AbsoluteRotationComponent]`

#### method Service AbsoluteSize
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/transform.AbsoluteSizeComponent]`

#### method Service AddDirtySet
Type: `func(engine/modules/ecs.DirtySet)`

#### method Service AspectRatio
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/transform.AspectRatioComponent]`

#### method Service Inherit
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/transform.InheritComponent]`

#### method Service Mat4
Type: `func(engine/modules/ecs.EntityID) github.com/go-gl/mathgl/mgl32.Mat4`

#### method Service MaxSize
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/transform.MaxSizeComponent]`

#### method Service MinSize
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/transform.MinSizeComponent]`

#### method Service ParentPivotPoint
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/transform.ParentPivotPointComponent]`

#### method Service PivotPoint
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/transform.PivotPointComponent]`

#### method Service Pos
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/transform.PosComponent]`

#### method Service Rotation
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/transform.RotationComponent]`

#### method Service Size
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/transform.SizeComponent]`

### type InheritFlag
Type: `engine/modules/transform.InheritFlag`
parent

### type PrimaryAxis
Type: `engine/modules/transform.PrimaryAxis`
aspect ratio

### type PosComponent
Type: `engine/modules/transform.PosComponent`

#### property PosComponent Pos
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### method PosComponent Smooth
Type: `func()`

#### method PosComponent Lerp
Type: `func(c2 engine/modules/transform.PosComponent, mix32 float32) engine/modules/transform.PosComponent`

### type RotationComponent
Type: `engine/modules/transform.RotationComponent`

#### property RotationComponent Rotation
Type: `github.com/go-gl/mathgl/mgl32.Quat`

#### method RotationComponent Smooth
Type: `func()`

#### method RotationComponent Lerp
Type: `func(c2 engine/modules/transform.RotationComponent, mix32 float32) engine/modules/transform.RotationComponent`

### type SizeComponent
Type: `engine/modules/transform.SizeComponent`

#### property SizeComponent Size
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### method SizeComponent Smooth
Type: `func()`

#### method SizeComponent Lerp
Type: `func(c2 engine/modules/transform.SizeComponent, mix32 float32) engine/modules/transform.SizeComponent`

### type AbsolutePosComponent
Type: `engine/modules/transform.AbsolutePosComponent`

#### property AbsolutePosComponent Pos
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

### type AbsoluteRotationComponent
Type: `engine/modules/transform.AbsoluteRotationComponent`

#### property AbsoluteRotationComponent Rotation
Type: `github.com/go-gl/mathgl/mgl32.Quat`

### type AbsoluteSizeComponent
Type: `engine/modules/transform.AbsoluteSizeComponent`

#### property AbsoluteSizeComponent Size
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

### type MinSizeComponent
Type: `engine/modules/transform.MinSizeComponent`

#### property MinSizeComponent Size
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### method MinSizeComponent Smooth
Type: `func()`

#### method MinSizeComponent Lerp
Type: `func(c2 engine/modules/transform.MinSizeComponent, mix32 float32) engine/modules/transform.MinSizeComponent`

### type MaxSizeComponent
Type: `engine/modules/transform.MaxSizeComponent`

#### property MaxSizeComponent Size
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### method MaxSizeComponent Smooth
Type: `func()`

#### method MaxSizeComponent Lerp
Type: `func(c2 engine/modules/transform.MaxSizeComponent, mix32 float32) engine/modules/transform.MaxSizeComponent`

### type AspectRatioComponent
Type: `engine/modules/transform.AspectRatioComponent`

#### property AspectRatioComponent AspectRatio
Type: `github.com/go-gl/mathgl/mgl32.Vec3`
0 means ignore axis

#### property AspectRatioComponent PrimaryAxis
Type: `engine/modules/transform.PrimaryAxis`

#### method AspectRatioComponent Smooth
Type: `func()`

#### method AspectRatioComponent Lerp
Type: `func(c2 engine/modules/transform.AspectRatioComponent, mix32 float32) engine/modules/transform.AspectRatioComponent`

### type PivotPointComponent
Type: `engine/modules/transform.PivotPointComponent`
pivot refers to object center.
default center is (.5, .5, .5).
each axis value should be between 0 and 1.

example: to align to left use (0, .5, .5)

#### property PivotPointComponent Point
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### method PivotPointComponent Smooth
Type: `func()`

#### method PivotPointComponent Lerp
Type: `func(c2 engine/modules/transform.PivotPointComponent, mix32 float32) engine/modules/transform.PivotPointComponent`

### type ParentPivotPointComponent
Type: `engine/modules/transform.ParentPivotPointComponent`

#### property ParentPivotPointComponent Point
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### method ParentPivotPointComponent Smooth
Type: `func()`

#### method ParentPivotPointComponent Lerp
Type: `func(c2 engine/modules/transform.ParentPivotPointComponent, mix32 float32) engine/modules/transform.ParentPivotPointComponent`

### type InheritComponent
Type: `engine/modules/transform.InheritComponent`

#### property InheritComponent RelativeMask
Type: `engine/modules/transform.InheritFlag`

## Variables
### var RelativePos
Type: `engine/modules/transform.InheritFlag`

### var RelativeRotation
Type: `engine/modules/transform.InheritFlag`

### var RelativeSizeX
Type: `engine/modules/transform.InheritFlag`

### var RelativeSizeY
Type: `engine/modules/transform.InheritFlag`

### var RelativeSizeZ
Type: `engine/modules/transform.InheritFlag`

### var RelativeSizeXY
Type: `engine/modules/transform.InheritFlag`

### var RelativeSizeXZ
Type: `engine/modules/transform.InheritFlag`

### var RelativeSizeXYZ
Type: `engine/modules/transform.InheritFlag`

### var RelativeSizeYZ
Type: `engine/modules/transform.InheritFlag`

### var Absolute
Type: `engine/modules/transform.InheritFlag`

### var Relative
Type: `engine/modules/transform.InheritFlag`

### var PrimaryAxisX
Type: `engine/modules/transform.PrimaryAxis`

### var PrimaryAxisY
Type: `engine/modules/transform.PrimaryAxis`

### var PrimaryAxisZ
Type: `engine/modules/transform.PrimaryAxis`

## Functions
### func NewPos
Type: `func(x float32, y float32, z float32) engine/modules/transform.PosComponent`

### func NewRotation
Type: `func(rotation github.com/go-gl/mathgl/mgl32.Quat) engine/modules/transform.RotationComponent`

### func NewSize
Type: `func(x float32, y float32, z float32) engine/modules/transform.SizeComponent`

### func NewMinSize
Type: `func(x float32, y float32, z float32) engine/modules/transform.MinSizeComponent`

### func NewMaxSize
Type: `func(x float32, y float32, z float32) engine/modules/transform.MaxSizeComponent`

### func NewAspectRatio
Type: `func(x float32, y float32, z float32, primaryAxis engine/modules/transform.PrimaryAxis) engine/modules/transform.AspectRatioComponent`

### func NewPivotPoint
Type: `func(x float32, y float32, z float32) engine/modules/transform.PivotPointComponent`

### func NewParentPivotPoint
Type: `func(x float32, y float32, z float32) engine/modules/transform.ParentPivotPointComponent`

### func NewInherit
Type: `func(mask engine/modules/transform.InheritFlag) engine/modules/transform.InheritComponent`


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
  - `engine/modules/hierarchy.Component`
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