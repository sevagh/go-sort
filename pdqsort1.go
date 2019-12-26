package gosort

// CLRS with basic pivot and pdqsort bad pattern breaking

func PdqSort1(nums []int) {
	pdqSort1(nums, 0, len(nums)-1, allowedBadPartitions(len(nums)-1, 0))
}

func pdqSort1(nums []int, p, r, badAllowed int) {
	if p < r {
		q, pivot := partition(nums, p, r)

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
