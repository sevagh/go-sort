package gosort

import "fmt"

// TimSort is a simplistic implementation of timsort with no galloping
func TimSort(nums []int) {
	n := len(nums)
	lo := 0
	prev := 0

	for lo < n {
		minrun := minRun(n - lo)
		currRun := countRun(nums, lo)
		hi := lo + currRun

		if currRun < minrun {
			// run is not big enough, insertion sort it
			hi = lo + minrun
			insertionSort(nums, lo, hi)
		}

		prev = lo
		symMerge(nums, 0, prev, hi)
		lo = hi
	}
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

var currentMerge [2]int
var nextMerge [2]int

func reduceOddMergeBoundaries(nums []int, mergeBoundaries *[][2]int) {
	lenMB := len(*mergeBoundaries)

	if lenMB <= 1 {
		return
	}

	// odd merge boundaries, eliminate the middle
	if lenMB%2 == 1 {
		mid := lenMB / 2
		symMerge(nums, (*mergeBoundaries)[mid][0], (*mergeBoundaries)[mid][1], (*mergeBoundaries)[mid+1][1])
		toAdd := [2]int{(*mergeBoundaries)[mid][0], (*mergeBoundaries)[mid+1][1]}
		*mergeBoundaries = append((*mergeBoundaries)[:mid], (*mergeBoundaries)[mid+1:]...)
		(*mergeBoundaries)[mid] = toAdd
	}
}

func mergeOnBoundaries(nums []int, mergeBoundaries *[][2]int) {
	// here, we have to repeatedly merge on the boundaries
	// merge duples from the beginning and end

	reduceOddMergeBoundaries(nums, mergeBoundaries)

	for len(*mergeBoundaries) > 1 {
		for i := 0; i < len(*mergeBoundaries); i += 2 {
			//fmt.Printf("len(mb): %d, i: %d, i+1: %d\n", len(*mergeBoundaries), i, i+1)
			currentMerge, nextMerge = (*mergeBoundaries)[i], (*mergeBoundaries)[i+1]
			fmt.Printf("mergeOnBoundaries doing %d %d %d\n", currentMerge[0], nextMerge[0], nextMerge[1])
			symMerge(nums, currentMerge[0], nextMerge[0], nextMerge[1])
			(*mergeBoundaries)[i/2] = [2]int{currentMerge[0], nextMerge[1]}
		}
		(*mergeBoundaries) = (*mergeBoundaries)[:len(*mergeBoundaries)/2]
		reduceOddMergeBoundaries(nums, mergeBoundaries)
	}
}
