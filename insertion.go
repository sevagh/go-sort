package mergesort

// copied from https://raw.githubusercontent.com/golang/go/master/src/sort/sort.go

func InsertionSort(nums []int) {
	insertionSort(nums, 0, len(nums))
}

func insertionSort(nums []int, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
