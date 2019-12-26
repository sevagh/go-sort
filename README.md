# go-sort

A collection of Go sort experiments and implementations.

The goal of this project was to implement minimal versions of **timsort** and **pdqsort** and validate their performance improvement over standard mergesort and quicksort. Also included are:

* RadixSort (using counting sort on decimal digits)
* heapSort (from go/pkg/sort)
* insertionSort (from go/pkg/sort)

### Quicksort

| Function   | Source | Notes |
|------------|--------|-------|
| QuickSort1 | CLRS | Naive pivot selection (A[r]) |
| QuickSort2 | [golang pkg/sort](https://golang.org/pkg/sort/#Sort) | Uses median-of-three partitioning, drops to insertionSort and heapSort in some cases. It looks like [introsort](https://en.wikipedia.org/wiki/Introsort). Partitioning looks like Hoare's partitioning scheme |
| QuickSort3 | CLRS | Randomized pivot selection |
| PdqSort1 | [pdqsort](https://github.com/orlp/pdqsort) | QuickSort1 with added bad partition detection and elimination |
| PdqSort2 | [pdqsort](https://github.com/orlp/pdqsort) | QuickSort2 with added bad partition detection and elimination |
| PdqSort3 | [pdqsort](https://github.com/orlp/pdqsort) | QuickSort3 with added bad partition detection and elimination |

QuickSort1 vs QuickSort3 shows the importance of picking good pivots in quicksort.

Parts of PdqSort implemented:

1. Check partitions that are bad and do swaps to fix them (swaps taken from the [Rust pdqsort implementation](https://docs.rs/pdqsort/0.1.0/src/pdqsort/lib.rs.html#427))

Results for PdqSort2 were expectedly disappointing. Some gains, some losses, given my amateurish modifications on top of Go's standard, optimized quicksort. Same goes for PdqSort3 vs. QuickSort3, where the random pivot selection seems to be good enough to not benefit from the addition of partition busting.

However, PdqSort1 vs QuickSort1 on an array of random integers is significant (as per my benchmarks, which may be flawed), which may be a demonstration of the isolated benefits of purely bad partition busting bolted onto a basic quicksort. This might be a tautological conclusion since QuickSort1 picks almost the worst partition:

```
sevagh:go-sort $ go test -benchmem -run=^a -bench='.*(QuickSort1|PdqSort1).*Random(8192|65536)$' -v
goos: linux
goarch: amd64
pkg: github.com/sevagh/go-sort
BenchmarkPdqSort1Random8192-8               3015            383876 ns/op               0 B/op          0 allocs/op
BenchmarkPdqSort1Random65536-8               121           9831749 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSort1Random8192-8              100          25272420 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSort1Random65536-8             100        1673251259 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/sevagh/go-sort       173.248s
```

Additional reading for quicksort:

1. S. Edelkamp, A. Weiß, "BlockQuicksort: Avoiding Branch Mispredictions in Quicksort", [link](https://pdfs.semanticscholar.org/b24e/f8021811cd4ef0fcc96a770657b664ee5b52.pdf)
2. M. D. McIlroy, "A Killer Adversary for Quicksort", [link](https://www.cs.dartmouth.edu/~doug/mdmspe.pdf)
3. The above adversary test in Go's sort pkg [tests](https://github.com/golang/go/blob/master/src/sort/sort_test.go#L455)

### Mergesort

| Function   | Source | Notes |
|------------|--------|-------|
| MergeSort1 | CLRS | Basic recursive implementation |
| MergeSort2 | [golang pkg/sort](https://golang.org/pkg/sort/#Stable) | Implements in-place symmetric merge ([[1]](https://www.semanticscholar.org/paper/Stable-Minimum-Storage-Merging-by-Symmetric-Kim-Kutzner/d664cee462cb8e6a8ae2a1a7c6bab1b5f81e0618)) and does insertion sort by blocks |
| MergeSort3 | Goodrich & Tamassia | Bottom-up iterative merge sort |
| TimSort | Various timsort sources | No galloping, and uses the same in-place symmetric merge as MergeSort2 |

Parts of TimSort implemented:

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
sevagh:go-mergesort $ go test -benchmem -run=^a -bench='.*Merge.*Random1048576$' -v
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

Additional reading for timsort:

1. Python [implementation](https://github.com/python/cpython/blob/master/Objects/listobject.c) and [description](https://github.com/python/cpython/blob/master/Objects/listsort.txt)
2. High level descriptions of timsort I used to write my implementation: [here](https://medium.com/@rylanbauermeister/understanding-timsort-191c758a42f3?) and [here](https://wiki.c2.com/?TimSort)
