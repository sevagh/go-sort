package gosort_test

import (
	"sort"
	"testing"

	"github.com/sevagh/go-sort"
)

type adversaryTestingData struct {
	t         *testing.T
	data      []int // item values, initialized to special gas value and changed by Less
	maxcmp    int   // number of comparisons allowed
	ncmp      int   // number of comparisons (calls to Less)
	nsolid    int   // number of elements that have been set to non-gas values
	candidate int   // guess at current pivot
	gas       int   // special value for unset elements, higher than everything else
}

func (d *adversaryTestingData) Len() int { return len(d.data) }

func (d *adversaryTestingData) Less(i, j int) bool {
	if d.ncmp >= d.maxcmp {
		d.t.Fatalf("used %d comparisons sorting adversary data with size %d", d.ncmp, len(d.data))
	}
	d.ncmp++

	if d.data[i] == d.gas && d.data[j] == d.gas {
		if i == d.candidate {
			// freeze i
			d.data[i] = d.nsolid
			d.nsolid++
		} else {
			// freeze j
			d.data[j] = d.nsolid
			d.nsolid++
		}
	}

	if d.data[i] == d.gas {
		d.candidate = i
	} else if d.data[j] == d.gas {
		d.candidate = j
	}

	return d.data[i] < d.data[j]
}

func (d *adversaryTestingData) Swap(i, j int) {
	d.data[i], d.data[j] = d.data[j], d.data[i]
}

func newAdversaryTestingData(t *testing.T, size int, maxcmp int) *adversaryTestingData {
	gas := size - 1
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = gas
	}
	return &adversaryTestingData{t: t, data: data, maxcmp: maxcmp, gas: gas}
}

func lg(n int) int {
	i := 0
	for 1<<uint(i) < n {
		i++
	}
	return i
}

func TestQuickSortIface1Adversary(t *testing.T) {
	const size = 10000            // large enough to distinguish between O(n^2) and O(n*log(n))
	maxcmp := size * lg(size) * 4 // the factor 4 was found by trial and error
	d := newAdversaryTestingData(t, size, maxcmp)
	gosort.QuickSortIface1(d) // This should degenerate to heapsort.
	// Check data is fully populated and sorted.
	for i, v := range d.data {
		if v != i {
			t.Fatalf("adversary data not fully sorted")
		}
	}
}

func TestPdqSortIface1Adversary(t *testing.T) {
	const size = 10000            // large enough to distinguish between O(n^2) and O(n*log(n))
	maxcmp := size * lg(size) * 4 // the factor 4 was found by trial and error
	d := newAdversaryTestingData(t, size, maxcmp)
	gosort.PdqSortIface1(d) // This should degenerate to heapsort.
	// Check data is fully populated and sorted.
	for i, v := range d.data {
		if v != i {
			t.Fatalf("adversary data not fully sorted")
		}
	}
}

func TestAdversaryStdSort(t *testing.T) {
	const size = 10000            // large enough to distinguish between O(n^2) and O(n*log(n))
	maxcmp := size * lg(size) * 4 // the factor 4 was found by trial and error
	d := newAdversaryTestingData(t, size, maxcmp)
	sort.Sort(d) // This should degenerate to heapsort.
	// Check data is fully populated and sorted.
	for i, v := range d.data {
		if v != i {
			t.Fatalf("adversary data not fully sorted")
		}
	}
}
