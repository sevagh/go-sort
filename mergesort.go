package mergesort

var Sentinel int = int(^uint(0) >> 1)

// CLRS implementation
func merge(nums []int, a, m, b int) {
	n1 := m - a + 1
	n2 := b - m

	l := make([]int, n1+1)
	r := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		l[i] = nums[a+i]
	}
	for j := 0; j < n2; j++ {
		r[j] = nums[m+j]
	}
	l[n1] = Sentinel
	r[n2] = Sentinel

	i := 0
	j := 0

	for k := a; k < b; k++ {
		if l[i] <= r[j] {
			nums[k] = l[i]
			i++
		} else {
			nums[k] = r[j]
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

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums))
}
