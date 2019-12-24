# go-mergesort

```
sevagh:go-mergesort $ go test -benchmem -run=^a -bench='.*1048576$' -v
goos: linux
goarch: amd64
pkg: github.com/sevagh/go-mergesort
BenchmarkMergeSort1Shuffled1048576-8                  10         122966919 ns/op        204194137 B/op   2097153 allocs/op
BenchmarkMergeSort1Shuffled16Values1048576-8          10         110442522 ns/op        204194128 B/op   2097153 allocs/op
BenchmarkMergeSort1AllEqual1048576-8                   9         112532196 ns/op        204194016 B/op   2097152 allocs/op
BenchmarkMergeSort1Ascending1048576-8                  9         112216945 ns/op        204194208 B/op   2097154 allocs/op
BenchmarkMergeSort1Descending1048576-8                 9         137265331 ns/op        204194037 B/op   2097152 allocs/op
BenchmarkMergeSort1PipeOrgan1048576-8                  9         128164760 ns/op        204194090 B/op   2097153 allocs/op
BenchmarkMergeSort1PushFront1048576-8                  9         137694488 ns/op        204194101 B/op   2097153 allocs/op
BenchmarkMergeSort1PushMiddle1048576-8                 8         147227042 ns/op        204194116 B/op   2097153 allocs/op
BenchmarkMergeSort2Shuffled1048576-8                 285           3856389 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort2Shuffled16Values1048576-8         393           3032438 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort2AllEqual1048576-8                 428           2735934 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort2Ascending1048576-8                393           2744665 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort2Descending1048576-8               457           2835564 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort2PipeOrgan1048576-8                465           2561045 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort2PushFront1048576-8                433           2555345 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort2PushMiddle1048576-8               429           2439072 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort3Shuffled1048576-8                 570           1944582 ns/op            3132 B/op          1 allocs/op
BenchmarkMergeSort3Shuffled16Values1048576-8         481           2262359 ns/op            3708 B/op          1 allocs/op
BenchmarkMergeSort3AllEqual1048576-8                 694           1544059 ns/op              16 B/op          1 allocs/op
BenchmarkMergeSort3Ascending1048576-8                777           1512064 ns/op              16 B/op          1 allocs/op
BenchmarkMergeSort3Descending1048576-8               780           1480964 ns/op              16 B/op          1 allocs/op
BenchmarkMergeSort3PipeOrgan1048576-8                771           1501609 ns/op              16 B/op          1 allocs/op
BenchmarkMergeSort3PushFront1048576-8                745           1619041 ns/op              16 B/op          1 allocs/op
BenchmarkMergeSort3PushMiddle1048576-8               771           1509540 ns/op              16 B/op          1 allocs/op
PASS
ok      github.com/sevagh/go-mergesort  49.364s
```
