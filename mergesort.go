package gosort

// MergeSort1 implements CLRS recursive mergesort
func MergeSort1(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

// MergeSort2 is pkg/sort's mergesort
//
// it does insertion sort in blocks and merges the blocks
//
// it looks a _little bit_ like timsort
func MergeSort2(nums []int) {
	stable(nums, len(nums))
}

// MergeSort3 implements G&T iterative bottom up mergesort
func MergeSort3(nums []int) {
	in := make([]int, len(nums))
	copy(in, nums)
	out := make([]int, len(nums))

	var tmp []int
	n := len(in)

	for i := 1; i < n; i *= 2 {
		for j := 0; j < n; j += 2 * i {
			merge3(in, out, j, i)
		}
		tmp = in
		in = out
		out = tmp
	}
	copy(nums, in)
}

func stable(nums []int, n int) {
	blockSize := 20 // must be > 0
	a, b := 0, blockSize
	for b <= n {
		insertionSort(nums, a, b)
		a = b
		b += blockSize
	}
	insertionSort(nums, a, n)

	for blockSize < n {
		a, b = 0, 2*blockSize
		for b <= n {
			symMerge(nums, a, a+blockSize, b)
			a = b
			b += 2 * blockSize
		}
		if m := a + blockSize; m < n {
			symMerge(nums, a, m, n)
		}
		blockSize *= 2
	}
}

var sentinel int = int(^uint(0) >> 1)

func merge(nums []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q

	left := make([]int, n1+1)
	right := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		left[i] = nums[p+i]
	}
	for j := 0; j < n2; j++ {
		right[j] = nums[q+j+1]
	}

	left[n1] = sentinel
	right[n2] = sentinel

	i := 0
	j := 0

	for k := p; k <= r; k++ {
		if left[i] <= right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
	}
}

func mergeSort(nums []int, a, b int) {
	if a < b {
		m := (a + b) / 2
		mergeSort(nums, a, m)
		mergeSort(nums, m+1, b)
		merge(nums, a, m, b)
	}
}

// SymMerge merges the two sorted subsequences data[a:m] and data[m:b] using
// the SymMerge algorithm from Pok-Son Kim and Arne Kutzner, "Stable Minimum
// Storage Merging by Symmetric Comparisons", in Susanne Albers and Tomasz
// Radzik, editors, Algorithms - ESA 2004, volume 3221 of Lecture Notes in
// Computer Science, pages 714-723. Springer, 2004.
//
// Let M = m-a and N = b-n. Wolog M < N.
// The recursion depth is bound by ceil(log(N+M)).
// The algorithm needs O(M*log(N/M + 1)) calls to data.Less.
// The algorithm needs O((M+N)*log(M)) calls to data.Swap.
//
// The paper gives O((M+N)*log(M)) as the number of assignments assuming a
// rotation algorithm which uses O(M+N+gcd(M+N)) assignments. The argumentation
// in the paper carries through for Swap operations, especially as the block
// swapping rotate uses only O(M+N) Swaps.
//
// symMerge assumes non-degenerate arguments: a < m && m < b.
// Having the caller check this condition eliminates many leaf recursion calls,
// which improves performance.
func symMerge(nums []int, a, m, b int) {
	// Avoid unnecessary recursions of symMerge
	// by direct insertion of data[a] into data[m:b]
	// if data[a:m] only contains one element.
	if m-a == 1 {
		// Use binary search to find the lowest index i
		// such that data[i] >= data[a] for m <= i < b.
		// Exit the search loop with i == b in case no such index exists.
		i := m
		j := b
		for i < j {
			h := int(uint(i+j) >> 1)
			if nums[h] < nums[a] {
				i = h + 1
			} else {
				j = h
			}
		}
		// Swap values until data[a] reaches the position before i.
		for k := a; k < i-1; k++ {
			nums[k], nums[k+1] = nums[k+1], nums[k]
		}
		return
	}

	// Avoid unnecessary recursions of symMerge
	// by direct insertion of data[m] into data[a:m]
	// if data[m:b] only contains one element.
	if b-m == 1 {
		// Use binary search to find the lowest index i
		// such that data[i] > data[m] for a <= i < m.
		// Exit the search loop with i == m in case no such index exists.
		i := a
		j := m
		for i < j {
			h := int(uint(i+j) >> 1)
			if !(nums[m] < nums[h]) {
				i = h + 1
			} else {
				j = h
			}
		}
		// Swap values until data[m] reaches the position i.
		for k := m; k > i; k-- {
			nums[k], nums[k-1] = nums[k-1], nums[k]
		}
		return
	}

	mid := int(uint(a+b) >> 1)
	n := mid + m
	var start, r int
	if m > mid {
		start = n - b
		r = mid
	} else {
		start = a
		r = m
	}
	p := n - 1

	for start < r {
		c := int(uint(start+r) >> 1)
		if !(nums[p-c] < nums[c]) {
			start = c + 1
		} else {
			r = c
		}
	}

	end := n - start
	if start < m && m < end {
		rotate(nums, start, m, end)
	}
	if a < start && start < mid {
		symMerge(nums, a, start, mid)
	}
	if mid < end && end < b {
		symMerge(nums, mid, end, b)
	}
}

func swapRange(nums []int, a, b, n int) {
	for i := 0; i < n; i++ {
		nums[a+i], nums[b+i] = nums[b+i], nums[a+i]
	}
}

func rotate(nums []int, a, m, b int) {
	i := m - a
	j := b - m

	for i != j {
		if i > j {
			swapRange(nums, m-i, m, j)
			i -= j
		} else {
			swapRange(nums, m-i, m+j-i, i)
			j -= i
		}
	}
	// i == j
	swapRange(nums, m-i, m, i)
}

func merge3(in, out []int, start, inc int) {
	x := start
	end1 := min(len(in), start+inc)
	end2 := min(len(in), start+2*inc)
	y := start + inc
	z := start
	for x < end1 && y < end2 {
		if in[x] < in[y] {
			out[z] = in[x]
			z++
			x++
		} else {
			out[z] = in[y]
			z++
			y++
		}
	}
	if x < end1 {
		for i := 0; i < end1-x; i++ {
			out[z+i] = in[x+i]
		}
	} else {
		for i := 0; i < end2-y; i++ {
			out[z+i] = in[y+i]
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
