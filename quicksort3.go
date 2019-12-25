package gosort

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// CLRS with randomized pivot

func QuickSort3(nums []int) {
	quickSort3(nums, 0, len(nums)-1)
}

func quickSort3(nums []int, p, r int) {
	if p < r {
		q := randomizedPartition(nums, p, r)
		quickSort3(nums, p, q-1)
		quickSort3(nums, q+1, r)
	}
}

func randomizedPartition(nums []int, p, r int) int {
	i := rand.Intn(r-p) + p
	nums[i], nums[r] = nums[r], nums[i]
	return partition(nums, p, r)
}
