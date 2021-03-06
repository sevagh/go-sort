package gosort_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/sevagh/go-sort"
)

// copied from https://github.com/golang/go/blob/master/src/sort/sort_test.go

const (
	_Sawtooth = iota
	_Rand
	_Stagger
	_Plateau
	_Shuffle
	_NDist
)

const (
	_Copy = iota
	_Reverse
	_ReverseFirstHalf
	_ReverseSecondHalf
	_Sorted
	_Dither
	_NMode
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func testBentleyMcIlroy(t *testing.T, sortFn func([]int)) {
	sizes := []int{100, 1023, 1024, 1025}
	dists := []string{"sawtooth", "rand", "stagger", "plateau", "shuffle"}
	modes := []string{"copy", "reverse", "reverse1", "reverse2", "sort", "dither"}
	var tmp1, tmp2 [1025]int
	for _, n := range sizes {
		for m := 1; m < 2*n; m *= 2 {
			for dist := 0; dist < _NDist; dist++ {
				j := 0
				k := 1
				data := tmp1[0:n]
				for i := 0; i < n; i++ {
					switch dist {
					case _Sawtooth:
						data[i] = i % m
					case _Rand:
						data[i] = rand.Intn(m)
					case _Stagger:
						data[i] = (i*m + i) % n
					case _Plateau:
						data[i] = min(i, m)
					case _Shuffle:
						if rand.Intn(m) != 0 {
							j += 2
							data[i] = j
						} else {
							k += 2
							data[i] = k
						}
					}
				}

				mdata := tmp2[0:n]
				for mode := 0; mode < _NMode; mode++ {
					switch mode {
					case _Copy:
						for i := 0; i < n; i++ {
							mdata[i] = data[i]
						}
					case _Reverse:
						for i := 0; i < n; i++ {
							mdata[i] = data[n-i-1]
						}
					case _ReverseFirstHalf:
						for i := 0; i < n/2; i++ {
							mdata[i] = data[n/2-i-1]
						}
						for i := n / 2; i < n; i++ {
							mdata[i] = data[i]
						}
					case _ReverseSecondHalf:
						for i := 0; i < n/2; i++ {
							mdata[i] = data[i]
						}
						for i := n / 2; i < n; i++ {
							mdata[i] = data[n-(i-n/2)-1]
						}
					case _Sorted:
						for i := 0; i < n; i++ {
							mdata[i] = data[i]
						}
						// Ints is known to be correct
						// because mode Sort runs after mode _Copy.
						sort.Ints(mdata)
					case _Dither:
						for i := 0; i < n; i++ {
							mdata[i] = data[i] + i%5
						}
					}

					desc := fmt.Sprintf("n=%d m=%d dist=%s mode=%s", n, m, dists[dist], modes[mode])
					//t.Logf("testing %s", desc)
					d := mdata[0:n]
					sortFn(d)
					if !sort.IntsAreSorted(d) {
						t.Fatalf("%s: ints not sorted\n\t%v", desc, mdata)
					} else {
						t.Logf("%s: ints are sorted\n\t", desc)
					}
				}
			}
		}
	}
}

func TestMergeCLRSBentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.MergeSort1)
}

func TestMergeStdlibCLRSBentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.MergeSort2)
}

func TestMergeGTIterativeBentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.MergeSort3)
}

func TestTimSort1BentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.TimSort1)
}

func TestTimSort2BentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.TimSort2)
}

func TestRadixSortBentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.RadixSort)
}

func TestQuickSort1BentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.QuickSort1)
}

func TestQuickSort2BentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.QuickSort2)
}

func TestQuickSort3BentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.QuickSort3)
}

func TestPdqSort2BentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.PdqSort2)
}

func TestPdqSort1BentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.PdqSort1)
}

func TestPdqSort3BentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.PdqSort3)
}

func TestBlockQuickSort1BentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.BlockQuickSort1)
}

func TestBlockQuickSort2BentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.BlockQuickSort2)
}

func TestBlockQuickSortControl1BentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.BlockQuickSortControl1)
}

func TestBlockQuickSortControl2BentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.BlockQuickSortControl2)
}

func TestBitonicSortBentleyMcIlroy(t *testing.T) {
	testBentleyMcIlroy(t, gosort.BitonicSort)
}
