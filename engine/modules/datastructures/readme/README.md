# datastructures
## Architecture
Defines efficient data structures for more specific use cases than golang built-in ones.

## Benchmarks
```
$ go test ./... -bench=.
goos: linux
goarch: amd64
pkg: engine/modules/datastructures/test
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkMapGetByIndex-8                      	100000000	        10.26 ns/op
BenchmarkIndexTrackerGetByIndex-8             	12761145	        93.21 ns/op
BenchmarkIndexTrackerGetByValue-8             	10375312	       105.9 ns/op
BenchmarkMapIterate-8                         	   98338	     12453 ns/op
BenchmarkIndexTrackerIterate-8                	 3669985	       319.8 ns/op
BenchmarkMapAdd-8                             	 3741642	       403.4 ns/op
BenchmarkIndexTrackerAdd-8                    	 1000000	      1066 ns/op
BenchmarkMapDelete-8                          	 2653327	       500.7 ns/op
BenchmarkIndexTrackerDelete-8                 	 1000000	      1415 ns/op
BenchmarkSparseSetGetWithoutPaging-8          	475007954	         2.547 ns/op
BenchmarkSparseSetGetWithPaging-8             	252428830	         4.778 ns/op
BenchmarkSparseSetGetIndicesWithoutPaging-8   	714019465	         1.688 ns/op
BenchmarkSparseSetGetIndicesWithPaging-8      	698018500	         1.703 ns/op
BenchmarkSparseSetAddWithoutPaging-8          	283678922	         5.791 ns/op
BenchmarkSparseSetAddWithPaging-8             	131760746	         9.040 ns/op
BenchmarkSparseSetRemoveWithoutPaging-8       	257807088	         4.341 ns/op
BenchmarkSparseSetRemoveWithPaging-8          	131474908	        10.25 ns/op
PASS
ok  	engine/modules/datastructures/test	43.474s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              12            142             57            748
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                            13            142             57            749
-------------------------------------------------------------------------------
```
## Types
### type SetReader
Type: `engine/modules/datastructures.SetReader[Stored comparable]`
set

#### method SetReader Get
Type: `func() []Stored`

#### method SetReader GetIndex
Type: `func(element Stored) (index int, ok bool)`

#### method SetReader GetStored
Type: `func(index int) (element Stored, ok bool)`

### type Set
Type: `engine/modules/datastructures.Set[Stored comparable]`

#### method Set Add
Type: `func(elements ...Stored)`

#### method Set Get
Type: `func() []Stored`

#### method Set GetIndex
Type: `func(element Stored) (index int, ok bool)`

#### method Set GetStored
Type: `func(index int) (element Stored, ok bool)`

#### method Set Remove
Type: `func(indices ...int)`

#### method Set RemoveElements
Type: `func(elements ...Stored)`

#### method Set Set
Type: `func(index int, e Stored)`

### type SparseArray
Type: `engine/modules/datastructures.SparseArray[Index golang.org/x/exp/constraints.Integer, Value any]`
sparse array

#### method SparseArray Get
Type: `func(index Index) (value Value, ok bool)`

#### method SparseArray GetIndices
Type: `func() []Index`

#### method SparseArray GetValues
Type: `func() []Value`

#### method SparseArray Remove
Type: `func(index Index) (removed bool)`

#### method SparseArray Set
Type: `func(index Index, value Value) (added bool)`

#### method SparseArray Size
Type: `func() int`

### type SparseSetReader
Type: `engine/modules/datastructures.SparseSetReader[Index golang.org/x/exp/constraints.Integer]`
sparse set

#### method SparseSetReader Get
Type: `func(index Index) (ok bool)`

#### method SparseSetReader GetIndices
Type: `func() []Index`

### type SparseSet
Type: `engine/modules/datastructures.SparseSet[Index golang.org/x/exp/constraints.Integer]`

#### method SparseSet Add
Type: `func(index Index) (added bool)`

#### method SparseSet Get
Type: `func(index Index) (ok bool)`

#### method SparseSet GetIndices
Type: `func() []Index`

#### method SparseSet Remove
Type: `func(index Index) (removed bool)`

### type TrackingArray
Type: `engine/modules/datastructures.TrackingArray[Stored comparable]`

#### method TrackingArray Add
Type: `func(elements ...Stored)`

#### method TrackingArray Changes
Type: `func() []engine/modules/datastructures/internal/types.Change[Stored]`

#### method TrackingArray ClearChanges
Type: `func()`

#### method TrackingArray Get
Type: `func() []Stored`

#### method TrackingArray Remove
Type: `func(indices ...int)`

#### method TrackingArray Set
Type: `func(index int, e Stored)`

## Functions
### func NewSet
Type: `func[Stored comparable]() engine/modules/datastructures.Set[Stored]`

### func NewSparseArray
Type: `func[Index golang.org/x/exp/constraints.Integer, Value any]() engine/modules/datastructures.SparseArray[Index, Value]`

### func NewSparseSet
Type: `func[Index golang.org/x/exp/constraints.Integer]() engine/modules/datastructures.SparseSet[Index]`

### func NewSparseSetWithPaging
Type: `func[Index golang.org/x/exp/constraints.Integer]() engine/modules/datastructures.SparseSet[Index]`

### func NewTrackingArray
Type: `func[Stored comparable]() engine/modules/datastructures.TrackingArray[Stored]`

### func NewThreadSafeTrackingArray
Type: `func[Stored comparable](mutex sync.Locker) engine/modules/datastructures.TrackingArray[Stored]`


## Dependencies
### Third Party
- `golang.org/x/exp/constraints`