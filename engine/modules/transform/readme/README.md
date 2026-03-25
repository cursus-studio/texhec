# Transform
## Architecture
We have many components which allow us to define relative position\
and this module transforms these relative components to absolute position.

## Benchmarks
```sh
$ go test ./... -bench=. -benchmem
goos: linux
goarch: amd64
pkg: engine/modules/transform/tests
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkGetPos-8                 	62290279	        17.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkRawGetPos-8              	100000000	        10.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkSetAbsolutePos-8         	 2190777	       538.1 ns/op	     128 B/op	       4 allocs/op
BenchmarkSetAndGetAbsolutePos-8   	 2272837	       516.6 ns/op	     128 B/op	       4 allocs/op
```

## Usage examples
```go
type Service interface {
	AbsolutePos() ecs.ComponentsArray[AbsolutePosComponent]
	AbsoluteRotation() ecs.ComponentsArray[AbsoluteRotationComponent]
	AbsoluteSize() ecs.ComponentsArray[AbsoluteSizeComponent]

	Pos() ecs.ComponentsArray[PosComponent]
	Rotation() ecs.ComponentsArray[RotationComponent]
	Size() ecs.ComponentsArray[SizeComponent]

	MaxSize() ecs.ComponentsArray[MaxSizeComponent]
	MinSize() ecs.ComponentsArray[MinSizeComponent]

	AspectRatio() ecs.ComponentsArray[AspectRatioComponent]
	PivotPoint() ecs.ComponentsArray[PivotPointComponent]

	Parent() ecs.ComponentsArray[ParentComponent]
	ParentPivotPoint() ecs.ComponentsArray[ParentPivotPointComponent]

	Mat4(ecs.EntityID) mgl32.Mat4
	AddDirtySet(ecs.DirtySet)
}
```

### Components
```go
type PosComponent struct{ Pos mgl32.Vec3 }
type RotationComponent struct{ Rotation mgl32.Quat }
type SizeComponent struct{ Size mgl32.Vec3 }

type AbsolutePosComponent struct{ Pos mgl32.Vec3 }
type AbsoluteRotationComponent struct{ Rotation mgl32.Quat }
type AbsoluteSizeComponent struct{ Size mgl32.Vec3 }

type MinSizeComponent SizeComponent // refers to absolute size. 0 means ignore axis
type MaxSizeComponent SizeComponent // refers to absolute size. 0 means ignore axis

type AspectRatioComponent struct {
	// 0 means ignore axis
	AspectRatio mgl32.Vec3
	PrimaryAxis PrimaryAxis
}

// pivot refers to object center.
// default center is (.5, .5, .5).
// each axis value should be between 0 and 1.
//
// example: to align to left use (0, .5, .5)
type PivotPointComponent struct{ Point mgl32.Vec3 }

type ParentComponent struct{ RelativeMask ParentFlag }
type ParentPivotPointComponent PivotPointComponent
```

## Dependencies
- [ecs](/engine/services/ecs/readme/README.md)
- [codec](/engine/services/codec/readme/README.md)
- [hierarchy](/engine/modules/hierarchy/readme/README.md)
- [logger](/engine/services/logger/readme/README.md)
- [transition](/engine/modules/transition/readme/README.md)
