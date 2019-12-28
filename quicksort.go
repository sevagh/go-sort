package gosort

import (
	"math/rand"
	"sort"
)

// QuickSort1 implements CLRS quicksort with basic pivot selection
func QuickSort1(nums []int) {
	quickSort1(nums, 0, len(nums)-1)
}

// QuickSortIface1 implements CLRS quicksort with basic pivot selection
//
// using the sorting.Interface
func QuickSortIface1(data sort.Interface) {
	quickSortIface1(data, 0, data.Len()-1)
}

// QuickSort2 is copied from pkg/sort - looks like introsort
func QuickSort2(nums []int) {
	n := len(nums)
	quickSort2(nums, 0, n, maxDepth(n))
}

// QuickSort3 implements CLRS quicksort with randomized pivot selection
func QuickSort3(nums []int) {
	quickSort3(nums, 0, len(nums)-1)
}

func quickSort1(nums []int, p, r int) {
	if p < r {
		q, _ := partition1(nums, p, r)
		quickSort1(nums, p, q-1)
		quickSort1(nums, q+1, r)
	}
}

func quickSortIface1(data sort.Interface, p, r int) {
	if p < r {
		q, _ := partitionIface1(data, p, r)
		quickSortIface1(data, p, q-1)
		quickSortIface1(data, q+1, r)
	}
}

func partition1(nums []int, p, r int) (int, int) {
	x := nums[r]
	i := p - 1
	for j := p; j < r; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1, r
}

func partitionIface1(data sort.Interface, p, r int) (int, int) {
	i := p - 1
	for j := p; j < r; j++ {
		if data.Less(j, r) {
			i++
			data.Swap(i, j)
		}
	}
	data.Swap(i+1, r)
	return i + 1, r
}

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

// medianOfThree moves the median of the three values data[m0], data[m1], data[m2] into data[m1].
func medianOfThree(nums []int, m1, m0, m2 int) {
	// sort 3 elements
	if nums[m1] <= nums[m0] {
		nums[m1], nums[m0] = nums[m0], nums[m1]
	}
	if nums[m2] <= nums[m1] {
		nums[m1], nums[m2] = nums[m2], nums[m1]
		if nums[m1] <= nums[m0] {
			nums[m1], nums[m0] = nums[m0], nums[m1]
		}
	}
}

func doPivot(nums []int, lo, hi int) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree(nums, lo, lo+s, lo+2*s)
		medianOfThree(nums, m, m-s, m+s)
		medianOfThree(nums, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThree(nums, lo, m, hi-1)

	// Invariants are:
	//	nums[lo] = pivot (set up by ChoosePivot)
	//	nums[lo < i < a] < pivot
	//	nums[a <= i < b] <= pivot
	//	nums[b <= i < c] unexamined
	//	nums[c <= i < hi-1] > pivot
	//	nums[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && nums[a] <= nums[pivot]; a++ {
	}
	b := a
	for {
		for ; b < c && !(nums[pivot] <= nums[b]); b++ {
		}
		for ; b < c && nums[pivot] <= nums[c-1]; c-- {
		}
		if b >= c {
			break
		}
		nums[b], nums[c-1] = nums[c-1], nums[b]
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let's be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		//if !(nums[pivot] <= nums[hi-1]) {
		if nums[pivot] == nums[hi-1] {
			nums[c], nums[hi-1] = nums[hi-1], nums[c]
			c++
			dups++
		}
		//if !(nums[b-1] <= nums[pivot]) {
		if nums[b-1] == nums[pivot] {
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> data[m] <= pivot
		//if !(nums[m] <= nums[pivot]) {
		if nums[m] == nums[pivot] {
			nums[m], nums[b-1] = nums[b-1], nums[m]
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		//	data[a <= i < b] unexamined
		//	data[b <= i < c] = pivot
		for {
			for ; a < b && nums[b-1] == nums[pivot]; b-- {
			}
			for ; a < b && nums[a] < nums[pivot]; a++ {
			}
			if a >= b {
				break
			}
			// data[a] == pivot; data[b-1] < pivot
			nums[a], nums[b-1] = nums[b-1], nums[a]
			a++
			b--
		}
	}
	// Swap pivot into middle
	nums[pivot], nums[b-1] = nums[b-1], nums[pivot]
	return b - 1, c
}

func quickSort2(nums []int, a, b, maxDepth int) {
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			heapSort(nums, a, b)
			return
		}
		maxDepth--

		mlo, mhi := doPivot(nums, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSort2(nums, a, mlo, maxDepth)
			a = mhi // i.e., quickSort(data, mhi, b)
		} else {
			quickSort2(nums, mhi, b, maxDepth)
			b = mlo // i.e., quickSort(data, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if nums[i] <= nums[i-6] {
				nums[i], nums[i-6] = nums[i-6], nums[i]
			}
		}
		insertionSort(nums, a, b)
	}
}

func quickSort3(nums []int, p, r int) {
	if p < r {
		q, _ := partition3(nums, p, r)
		quickSort3(nums, p, q-1)
		quickSort3(nums, q+1, r)
	}
}

func partition3(nums []int, p, r int) (int, int) {
	i := rand.Intn(r-p) + p
	nums[i], nums[r] = nums[r], nums[i]
	return partition1(nums, p, r)
}
