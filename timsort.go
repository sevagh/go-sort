package gosort

import "fmt"

// TimSort is a simplistic implementation of timsort with no galloping
//
// sources:
// https://github.com/python/cpython/blob/master/Objects/listobject.c
// https://medium.com/@rylanbauermeister/understanding-timsort-191c758a42f3
// https://wiki.c2.com/?TimSort
func TimSort(nums []int) {
	n := len(nums)
	lo := 0
	prev := 0
	iter := 0

	bigMerges := []int{}

	minrun := minRun(n - lo)
	currRun := countRun(nums, lo)
	hi := lo + currRun

	if currRun < minrun {
		// run is not big enough, insertion sort it
		hi = lo + minrun
		insertionSort(nums, lo, hi)
	}

	prev = lo
	lo = hi

	for lo < n {
		minrun = minRun(n - lo)
		currRun = countRun(nums, lo)
		hi = lo + currRun

		if currRun < minrun {
			// run is not big enough, insertion sort it
			hi = lo + minrun
			insertionSort(nums, lo, hi)
		}

		if iter%2 == 0 {
			fmt.Printf("merge: %d %d %d\n", prev, lo, hi)
			symMerge(nums, prev, lo, hi)
			bigMerges = append(bigMerges, hi)
		}

		prev = lo
		lo = hi
		iter++
	}
	symMerge(nums, prev, lo, hi)

	fmt.Printf("big merges: %+v\n", bigMerges)
	//mergeOnBoundaries(nums, &mergeBoundaries)
	//mergeSort(nums, 0, len(nums)-1)
}

func minRun(n int) int {
	r := 0 // Becomes 1 if any 1 bits are shifted off
	for n >= 64 {
		r |= (n & 1)
		n >>= 1
	}
	return n + r
}

// linear time run counter
// starting from the offset a
func countRun(nums []int, a int) int {
	descRun := 1
	prevDescRun := 1

	// strictly descending only, for stability
	for i := a + 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			descRun++
		}

		if descRun == prevDescRun {
			break
		}
		prevDescRun = descRun
	}

	ascRun := 1

	if descRun <= (len(nums)-a)/2 {
		// only here is there a chance a bigger ascending run could exist
		prevAscRun := 1

		for i := a + 1; i < len(nums); i++ {
			// ascending can be looser
			if nums[i] >= nums[i-1] {
				ascRun++
			}

			if ascRun == prevAscRun {
				break
			}
			prevAscRun = ascRun
		}
	}

	// return the max run
	if descRun > ascRun {
		ascRun = descRun
		// reverse the strictly descending run in-place
		// to make it ascending

		for i, j := a, a+ascRun-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return ascRun
}
