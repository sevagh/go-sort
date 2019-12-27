package gosort

// CLRS with basic pivot and block partitioning
// copied from https://github.com/MnO2/go-pdqsort

const BLOCK = 128

var offsetsL [BLOCK]int
var offsetsR [BLOCK]int

func BlockQuickSort1(nums []int) {
	blockQuickSort1(nums, 0, len(nums))
}

func blockQuickSort1(nums []int, p, r int) {
	if p < r {
		pivot := r - 1 //nums[r]
		q := partitionWithBlock(nums, p, r, pivot)
		blockQuickSort1(nums, p, q)
		blockQuickSort1(nums, q+1, r)
	}
}

func partitionWithBlock(nums []int, a, b, pivot int) int {
	nums[a], nums[pivot] = nums[pivot], nums[a]
	pivot = a

	l := a + 1
	r := b
	for l < r && nums[l] <= nums[pivot] {
		l += 1
	}

	for l < r && nums[r-1] > nums[pivot] {
		r -= 1
	}

	numOfSmallerThanPivotElems := partitionInBlock(nums, l, r, pivot)
	mid := (l - 1 + numOfSmallerThanPivotElems)

	nums[a], nums[mid] = nums[mid], nums[a]
	return mid
}

func partitionInBlock(nums []int, a, b, pivot int) int {
	l := a
	blockL := BLOCK
	startL := 0
	endL := 0

	r := b
	blockR := BLOCK
	startR := 0
	endR := 0

	for {
		isDone := (r - l) <= 2*BLOCK

		if isDone {
			rem := r - l
			if startL < endL || startR < endR {
				rem -= BLOCK
			}

			if startL < endL {
				blockR = rem
			} else if startR < endR {
				blockL = rem
			} else {
				blockL = rem / 2
				blockR = rem - blockL
			}
		}

		if startL == endL {
			startL = 0
			endL = 0
			elem := l

			for i := 0; i < blockL; i += 1 {
				offsetsL[endL] = l + i

				if nums[elem] > nums[pivot] {
					endL += 1
				}

				elem += 1
			}
		}

		if startR == endR {
			startR = 0
			endR = 0
			elem := r

			for i := 0; i < blockR; i += 1 {
				elem -= 1
				offsetsR[endR] = r - i - 1

				if nums[elem] <= nums[pivot] {
					endR += 1
				}
			}
		}

		count := min(endL-startL, endR-startR)
		if count > 0 {
			tmp := nums[offsetsL[startL]]
			nums[offsetsL[startL]] = nums[offsetsR[startR]]
			for i := 1; i < count; i++ {
				swap1 := offsetsL[startL+i]
				nums[offsetsR[startR+i-1]] = nums[swap1]
				nums[swap1] = nums[offsetsR[startR+i]]
			}
			nums[offsetsR[startR+count-1]] = tmp
			startL += count
			startR += count
		}

		if startL == endL {
			l += blockL
		}

		if startR == endR {
			r -= blockR
		}

		if isDone {
			break
		}
	}

	if startL < endL {
		for startL < endL {
			endL -= 1
			nums[offsetsL[endL]], nums[r-1] = nums[r-1], nums[offsetsL[endL]]
			r -= 1
		}
		return (r - a)
	} else if startR < endR {
		for startR < endR {
			endR -= 1
			nums[l], nums[offsetsR[endR]] = nums[offsetsR[endR]], nums[l]
			l += 1
		}
		return (l - a)
	} else {
		return (l - a)
	}
}
