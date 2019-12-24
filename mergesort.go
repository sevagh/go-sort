package mergesort

var Sentinel int = int(^uint(0) >> 1)

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

	left[n1] = Sentinel
	right[n2] = Sentinel

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

// CLRS implementation
func MergeSort1(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func merge4(in, out []int, start, inc int) {
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

// G&T implementation, iterative bottom up
func MergeSort4(nums []int) {
	in := make([]int, len(nums))
	copy(in, nums)
	out := make([]int, len(nums))

	var tmp []int
	n := len(in)

	for i := 1; i < n; i *= 2 {
		for j := 0; j < n; j += 2 * i {
			merge4(in, out, j, i)
		}
		tmp = in
		in = out
		out = tmp
	}
	copy(nums, in)
}
