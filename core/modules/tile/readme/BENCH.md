### Flag benchmark
```sh
goos: windows
goarch: amd64
pkg: core/modules/tile/test
cpu: Intel(R) Core(TM) i7-14700KF
BenchmarkRendering36MTilesMap
gpu: Meta Virtual Monitor
gpu 2: NVIDIA GeForce RTX 4080 SUPER
BenchmarkRendering36MTilesMap-28    	     135	   8510424 ns/op
PASS
ok  	core/modules/tile/test	57.055s
```
Rendering 36 million tiles on `NVIDIA GeForce RTX 4080 SUPER` in less than **8.6ms**.

### Standard benchmark
```sh
$ go test . -bench=. -benchtime=10s
Failed to load plugin 'libdecor-gtk.so': failed to init
gpu: Kaby Lake-R GT2 [UHD Graphics 620]
goos: linux
goarch: amd64
pkg: core/modules/tile/test
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkRendering1MTilesMap-8   	    2221	   5062408 ns/op
```
Rendering 1 million tiles on `UHD Graphics 620` in less than **5.1ms**.

### Custom benchmark
To run benchmark on your machine with map with custom size then find `BenchmarkRendering1MTilesMap`.
```go
func BenchmarkRendering1MTilesMap(b *testing.B) { benchmarkRenderingXTilesMap(b, 1000) }
func BenchmarkRendering4MTilesMap(b *testing.B) { benchmarkRenderingXTilesMap(b, 2000) }
```

These benchmarks create map with custom size `n`*`n` (in first example 1,000*1,000 = 1,000,000).
