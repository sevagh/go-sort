package gosort

import "math/rand"

// BlockQuickSort1 implements CLRS quicksort with bad pivot and block partitioning [1]
//
// [1]: S. Edelkamp, A. Weiß, "BlockQuicksort: Avoiding Branch Mispredictions in Quicksort" https://pdfs.semanticscholar.org/b24e/f8021811cd4ef0fcc96a770657b664ee5b52.pdf
func BlockQuickSort1(nums []int) {
	blockQuickSort1(nums, 0, len(nums))
}

// BlockQuickSort2 implements CLRS quicksort with randomized pivot and block partitioning [1]
//
// [1]: S. Edelkamp, A. Weiß, "BlockQuicksort: Avoiding Branch Mispredictions in Quicksort" https://pdfs.semanticscholar.org/b24e/f8021811cd4ef0fcc96a770657b664ee5b52.pdf
func BlockQuickSort2(nums []int) {
	blockQuickSort2(nums, 0, len(nums))
}

// BlockQuickSortControl1 is almost identical to BlockQuickSort1 with
// blocks removed, for fair bench comparisons
func BlockQuickSortControl1(nums []int) {
	blockQuickSortControl1(nums, 0, len(nums))
}

// BlockQuickSortControl2 is almost identical to BlockQuickSort2 with
// blocks removed, for fair bench comparisons
func BlockQuickSortControl2(nums []int) {
	blockQuickSortControl2(nums, 0, len(nums))
}

const blockSize = 128

var offsetsL [blockSize]int
var offsetsR [blockSize]int

func blockQuickSort1(nums []int, p, r int) {
	if p < r {
		pivot := r - 1 //nums[r]
		q := partitionWithBlock(nums, p, r, pivot)
		blockQuickSort1(nums, p, q)
		blockQuickSort1(nums, q+1, r)
	}
}

func blockQuickSortControl1(nums []int, p, r int) {
	if p < r {
		pivot := r - 1 //nums[r]
		q := partitionWithoutBlock(nums, p, r, pivot)
		blockQuickSortControl1(nums, p, q)
		blockQuickSortControl1(nums, q+1, r)
	}
}

func blockQuickSort2(nums []int, p, r int) {
	if p < r {
		pivot := r - 1 //nums[r]
		i := rand.Intn(r-p) + p
		nums[i], nums[pivot] = nums[pivot], nums[i]
		q := partitionWithBlock(nums, p, r, pivot)
		blockQuickSort2(nums, p, q)
		blockQuickSort2(nums, q+1, r)
	}
}

func blockQuickSortControl2(nums []int, p, r int) {
	if p < r {
		pivot := r - 1 //nums[r]
		i := rand.Intn(r-p) + p
		nums[i], nums[pivot] = nums[pivot], nums[i]
		q := partitionWithoutBlock(nums, p, r, pivot)
		blockQuickSortControl2(nums, p, q)
		blockQuickSortControl2(nums, q+1, r)
	}
}

func partitionWithoutBlock(nums []int, a, b, pivot int) int {
	nums[a], nums[pivot] = nums[pivot], nums[a]
	pivot = a

	l := a + 1
	r := b
	for l < r && nums[l] <= nums[pivot] {
		l++
	}

	for l < r && nums[r-1] > nums[pivot] {
		r--
	}

	i := l
	for j := l; j < r; j++ {
		if nums[j] <= nums[pivot] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	mid := i - 1

	nums[a], nums[mid] = nums[mid], nums[a]
	return mid
}

// attribution: https://github.com/MnO2/go-pdqsort
func partitionWithBlock(nums []int, a, b, pivot int) int {
	nums[a], nums[pivot] = nums[pivot], nums[a]
	pivot = a

	l := a + 1
	r := b
	for l < r && nums[l] <= nums[pivot] {
		l++
	}

	for l < r && nums[r-1] > nums[pivot] {
		r--
	}

	numOfSmallerThanPivotElems := partitionInBlock(nums, l, r, pivot)
	mid := (l - 1 + numOfSmallerThanPivotElems)

	nums[a], nums[mid] = nums[mid], nums[a]
	return mid
}

// attribution: https://github.com/MnO2/go-pdqsort
func partitionInBlock(nums []int, a, b, pivot int) int {
	l := a
	blockL := blockSize
	startL := 0
	endL := 0

	r := b
	blockR := blockSize
	startR := 0
	endR := 0

	for {
		isDone := (r - l) <= 2*blockSize

		if isDone {
			rem := r - l
			if startL < endL || startR < endR {
				rem -= blockSize
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

			for i := 0; i < blockL; i++ {
				offsetsL[endL] = l + i

				if nums[elem] > nums[pivot] {
					endL++
				}

				elem++
			}
		}

		if startR == endR {
			startR = 0
			endR = 0
			elem := r

			for i := 0; i < blockR; i++ {
				elem--
				offsetsR[endR] = r - i - 1

				if nums[elem] <= nums[pivot] {
					endR++
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
			endL--
			nums[offsetsL[endL]], nums[r-1] = nums[r-1], nums[offsetsL[endL]]
			r--
		}
		return (r - a)
	} else if startR < endR {
		for startR < endR {
			endR--
			nums[l], nums[offsetsR[endR]] = nums[offsetsR[endR]], nums[l]
			l++
		}
		return (l - a)
	} else {
		return (l - a)
	}
}
