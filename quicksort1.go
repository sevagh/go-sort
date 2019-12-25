package gosort

// CLRS with basic pivot

func QuickSort1(nums []int) {
	quickSort1(nums, 0, len(nums)-1)
}

func quickSort1(nums []int, p, r int) {
	if p < r {
		q := partition(nums, p, r)
		quickSort1(nums, p, q-1)
		quickSort1(nums, q+1, r)
	}
}

func partition(nums []int, p, r int) int {
	x := nums[r]
	i := p - 1
	for j := p; j < r; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
