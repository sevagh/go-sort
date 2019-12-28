package gosort

import "math"

// PdqSort1 implements CLRS bad pivot quicksort with pdq bad pivot pattern breaking
//
// Bad-pattern busting is copied from the Rust pdqsort implementation, https://docs.rs/pdqsort/0.1.0/src/pdqsort/lib.rs.html#427
func PdqSort1(nums []int) {
	pdqSort1(nums, 0, len(nums)-1, allowedBadPartitions(len(nums)-1, 0))
}

// PdqSort2 modifies pkg/sort's introsort quicksort with pdq bad pivot pattern breaking
func PdqSort2(nums []int) {
	n := len(nums)
	pdqSort2(nums, 0, n, maxDepth(n), allowedBadPartitions(0, n))
}

// PdqSort3 implements CLRS randomized pivot quicksort with pdq bad pivot pattern breaking
func PdqSort3(nums []int) {
	pdqSort3(nums, 0, len(nums)-1, allowedBadPartitions(len(nums)-1, 0))
}

func allowedBadPartitions(a, b int) int {
	return int(math.Log2(float64(b - a)))
}

func breakPatterns(nums []int) {
	n := len(nums)

	if n >= 4 {
		nums[0], nums[n/2] = nums[n/2], nums[0]
		nums[n-1], nums[n-n/2] = nums[n-n/2], nums[n-1]

		if n >= 8 {
			nums[1], nums[n/2+1] = nums[n/2+1], nums[1]
			nums[2], nums[n/2+2] = nums[n/2+2], nums[2]

			nums[n-2], nums[n-n/2-1] = nums[n-n/2-1], nums[n-2]
			nums[n-3], nums[n-n/2-2] = nums[n-n/2-2], nums[n-3]
		}
	}
}

func pdqSort2(nums []int, a, b, maxDepth, badAllowed int) {
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if badAllowed == 0 {
			heapSort(nums, a, b)
			return
		}

		// inherently always trying to drop back down to heapsort
		// so it's a bit different from the C++ pdqsort
		maxDepth--

		mlo, mhi := doPdqPivot(nums, a, b)
		pivotPos := mlo // this is where the pivot landed after the partitioning

		lSize := pivotPos - a
		rSize := b - (pivotPos + 1)
		highlyUnbalanced := lSize < len(nums)/8 || rSize < len(nums)/8

		if highlyUnbalanced {
			// do some specific swaps here per pdqsort
			// https://github.com/orlp/pdqsort/blob/master/pdqsort.h#L467
			badAllowed--

			// too many bad partitions encountered, time to bail
			if badAllowed == 0 {
				heapSort(nums, a, b)
				return
			}

			breakPatterns(nums[a:mlo])
			breakPatterns(nums[mhi:b])
		}

		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			pdqSort2(nums, a, mlo, maxDepth, badAllowed)
			a = mhi // i.e., quickSort(data, mhi, b)
		} else {
			pdqSort2(nums, mhi, b, maxDepth, badAllowed)
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

func pdqSort1(nums []int, p, r, badAllowed int) {
	if p < r {
		q, pivot := partition1(nums, p, r)

		lSize := pivot - p
		rSize := r - (pivot + 1)
		highlyUnbalanced := lSize < len(nums)/8 || rSize < len(nums)/8

		if highlyUnbalanced {
			// do some specific swaps here per pdqsort
			// https://github.com/orlp/pdqsort/blob/master/pdqsort.h#L467
			badAllowed--

			// too many bad partitions encountered, time to bail
			if badAllowed == 0 {
				heapSort(nums, p, r)
				return
			}

			breakPatterns(nums[p:q])
			breakPatterns(nums[q+1 : r+1])
		}

		pdqSort1(nums, p, q-1, badAllowed)
		pdqSort1(nums, q+1, r, badAllowed)
	}
}

func doPdqPivot(nums []int, lo, hi int) (mlo, mhi int) {
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

func pdqSort3(nums []int, p, r, badAllowed int) {
	if p < r {
		q, pivot := partition3(nums, p, r)

		lSize := pivot - p
		rSize := r - (pivot + 1)
		highlyUnbalanced := lSize < len(nums)/8 || rSize < len(nums)/8

		if highlyUnbalanced {
			// do some specific swaps here per pdqsort
			// https://github.com/orlp/pdqsort/blob/master/pdqsort.h#L467
			badAllowed--

			// too many bad partitions encountered, time to bail
			if badAllowed == 0 {
				heapSort(nums, p, r)
				return
			}

			breakPatterns(nums[p:q])
			breakPatterns(nums[q+1 : r+1])
		}

		pdqSort3(nums, p, q-1, badAllowed)
		pdqSort3(nums, q+1, r, badAllowed)
	}
}
