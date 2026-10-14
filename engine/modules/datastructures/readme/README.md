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
BenchmarkMapGetByIndex-8                      	143546320	         8.358 ns/op
BenchmarkIndexTrackerGetByIndex-8             	15578006	        71.48 ns/op
BenchmarkIndexTrackerGetByValue-8             	12646851	        87.21 ns/op
BenchmarkMapIterate-8                         	  109822	     10921 ns/op
BenchmarkIndexTrackerIterate-8                	 4143535	       289.8 ns/op
BenchmarkMapAdd-8                             	 3964046	       354.1 ns/op
BenchmarkIndexTrackerAdd-8                    	 1351284	       832.7 ns/op
BenchmarkMapDelete-8                          	 2966322	       465.4 ns/op
BenchmarkIndexTrackerDelete-8                 	 1000000	      1221 ns/op
BenchmarkSparseSetGetWithoutPaging-8          	472984779	         2.516 ns/op
BenchmarkSparseSetGetWithPaging-8             	252377475	         4.753 ns/op
BenchmarkSparseSetGetIndicesWithoutPaging-8   	714335947	         1.684 ns/op
BenchmarkSparseSetGetIndicesWithPaging-8      	696928789	         1.689 ns/op
BenchmarkSparseSetAddWithoutPaging-8          	297178858	         4.195 ns/op
BenchmarkSparseSetAddWithPaging-8             	132219685	         9.060 ns/op
BenchmarkSparseSetRemoveWithoutPaging-8       	291432214	         3.794 ns/op
BenchmarkSparseSetRemoveWithPaging-8          	134787138	         8.799 ns/op
PASS
ok  	engine/modules/datastructures/test	38.083s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              13            163             59            817
Markdown                         1              0              0              1
-------------------------------------------------------------------------------
SUM:                            14            163             59            818
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

### func NewStableSparseArray
Type: `func[Index golang.org/x/exp/constraints.Integer, Value any]() engine/modules/datastructures.SparseArray[Index, Value]`
doesn't reorder elements on removal

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