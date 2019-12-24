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
