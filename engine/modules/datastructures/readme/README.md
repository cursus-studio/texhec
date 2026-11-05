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
BenchmarkMapGetByIndex-8                      	142136802	         8.420 ns/op
BenchmarkIndexTrackerGetByIndex-8             	17745988	        69.94 ns/op
BenchmarkIndexTrackerGetByValue-8             	13944583	        85.43 ns/op
BenchmarkMapIterate-8                         	  111024	     10850 ns/op
BenchmarkIndexTrackerIterate-8                	 4135228	       290.0 ns/op
BenchmarkMapAdd-8                             	 4084909	       351.3 ns/op
BenchmarkIndexTrackerAdd-8                    	 1374637	       816.1 ns/op
BenchmarkMapDelete-8                          	 2864173	       481.5 ns/op
BenchmarkIndexTrackerDelete-8                 	  990952	      1205 ns/op
BenchmarkSparseSetGetWithoutPaging-8          	477387853	         2.517 ns/op
BenchmarkSparseSetGetWithPaging-8             	252547494	         4.775 ns/op
BenchmarkSparseSetGetIndicesWithoutPaging-8   	712959960	         1.681 ns/op
BenchmarkSparseSetGetIndicesWithPaging-8      	714232222	         1.681 ns/op
BenchmarkSparseSetAddWithoutPaging-8          	299762292	         4.722 ns/op
BenchmarkSparseSetAddWithPaging-8             	133732274	         9.531 ns/op
BenchmarkSparseSetRemoveWithoutPaging-8       	268311296	         3.797 ns/op
BenchmarkSparseSetRemoveWithPaging-8          	132810042	         8.925 ns/op
PASS
ok  	engine/modules/datastructures/test	40.739s
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
## Dependencies
### Third Party
- `golang.org/x/exp/constraints`