# go-mergesort

A collection of merge sort implementations, including a naive timsort I wrote from scratch. The goal of this project was to write a version of timsort that's still similar to a basic mergesort (no galloping, minimal magic constants, etc.) and validate the performance improvement.

| Function   | Source | Notes |
|------------|--------|-------|
| MergeSort1 | CLRS | Basic recursive implementation |
| MergeSort2 | [golang pkg/sort](https://golang.org/pkg/sort/#Stable) | Implements in-place symmetric merge ([[1]](https://www.semanticscholar.org/paper/Stable-Minimum-Storage-Merging-by-Symmetric-Kim-Kutzner/d664cee462cb8e6a8ae2a1a7c6bab1b5f81e0618)) and does insertion sort by blocks |
| MergeSort3 | Various timsort sources | No galloping, and uses the same in-place symmetric merge ([1]) as above |
| MergeSort4 | Goodrich & Tamassia | Bottom-up iterative merge sort |

Parts of TimSort implemented in MergeSort3:

1. Calculate minimum ascending runs of at least 32 in length (insertion sort to 32 if less)
    1. If the descending run is bigger, reverse it in-place and use that as the minrun
2. Record the boundaries between runs on which to merge the sorted minruns
3. Merge on the boundaries as follows (pseudocode):

```
boundaries = [[a, b], [b, c], [c, d], [d, e]]

for len(boundaries) > 1 {
    for i := 0; i < len(boundaries); i += 2 {
        // first iteration: i = 0
        // consider [a, b] and [b, c]
        symMerge(a, b, c)
        boundaries[0/2] = [a, c]

        // second iteration: i = 1
        // consider [c, d] and [d, e]
        symMerge(c, d, e)
        boundaries[1/2] = [c, e]
    }
    boundaries = boundaries[:len(boundaries)/2]
    // boundaries = [[a, c], [c, e]]
}
```

Benches for 1 million random ints:

```
sevagh:go-mergesort $ go test -benchmem -run=^a -bench='.*Random1048576$' -v
goos: linux
goarch: amd64
pkg: github.com/sevagh/go-mergesort
BenchmarkMergeSort1Random1048576-8            10         112270897 ns/op        204194092 B/op   2097151 allocs/op
BenchmarkMergeSort2Random1048576-8           264           4090389 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort3Random1048576-8           758           1390610 ns/op            2359 B/op          1 allocs/op
BenchmarkMergeSort4Random1048576-8            30          34546053 ns/op        16777232 B/op          2 allocs/op
PASS
ok      github.com/sevagh/go-mergesort  12.053s
```

I'm surprised by how well my naive implementation of timsort is performing - I'm sure it can be made better.
