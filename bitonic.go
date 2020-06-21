package gosort

import (
	"fmt"
	"sync"
)

func compare(nums *[]int, i, j, dir int) {
	if !compareWithDirection((*nums)[i], (*nums)[j], dir) {
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
	}
}

func bitonicSort(nums *[]int, lo, n, dir, maxDepth int) {
	if maxDepth == 0 {
		heapSort(*nums, lo, n)
		return
	}
	maxDepth--

	if n > 1 {
		m := n / 2

		bitonicSort(nums, lo, m, 1-dir, maxDepth)
		bitonicSort(nums, lo+m, n-m, dir, maxDepth)

		bitonicMerge(nums, lo, n, dir, maxDepth)
	}
}

func bitonicMerge(nums *[]int, lo, n, dir, maxDepth int) {
	if maxDepth == 0 {
		heapSort(*nums, lo, n)
		return
	}
	maxDepth--

	if n > 1 {
		m := greatestPowerOfTwoLessThan(n)
		for i := lo; i < lo+n-m; i++ {
			compare(nums, i, i+m, dir)
		}
		bitonicMerge(nums, lo, m, dir, maxDepth)
		bitonicMerge(nums, lo+m, n-m, dir, maxDepth)
	}
}

// BitonicSort implements https://www.inf.fh-flensburg.de/lang/algorithmen/sortieren/bitonic/oddn.htm
func BitonicSort(nums []int) {
	bitonicSort(&nums, 0, len(nums), increasing, maxDepth(len(nums)))
}

// n>=2  and  n<=Integer.MAX_VALUE
func greatestPowerOfTwoLessThan(n int) int {
	k := 1
	for k > 0 && k < n {
		k <<= 1
	}
	return k >> 1
}

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

const (
	increasing = iota
	decreasing
)

func comparisonToDirection(num1, num2 int) int {
	if num1 <= num2 {
		return increasing
	}
	return decreasing
}

func compareWithDirection(num1, num2, direction int) bool {
	if direction == increasing {
		return num1 <= num2
	}
	return num1 > num2
}

func isBitonic(btcseq []int) (bool, int) {
	n := len(btcseq)
	if n <= 2 {
		if n <= 1 {
			return true, increasing
		}
		return true, comparisonToDirection(btcseq[0], btcseq[1])
	}

	dir := increasing
	if btcseq[1] < btcseq[0] {
		dir = decreasing
	}

	transitions := 0
	for i := 1; i < n; i++ {
		if !compareWithDirection(btcseq[i-1], btcseq[i], dir) {
			dir = 1 - dir // changed direction
			transitions++
		}
	}

	if transitions > 1 {
		// check if there is a circular monotonic sequence
		if compareWithDirection(btcseq[len(btcseq)-1], btcseq[0], dir) {
			transitions--
		}
	}

	return transitions <= 1, dir
}

// https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-851-advanced-data-structures-spring-2012/calendar-and-notes/MIT6_851S12_L14.pdf
func bitonicSeqSortRecurse(btcseq []int) {
	n := len(btcseq)

	if n <= 1 {
		return
	}

	for i := 0; i < n/2; i++ {
		if btcseq[i] > btcseq[i+n/2] {
			btcseq[i], btcseq[i+n/2] = btcseq[i+n/2], btcseq[i]
		}
	}

	bitonicSeqSortRecurse(btcseq[:n/2])
	bitonicSeqSortRecurse(btcseq[n/2:])
}

// sorting bitonic sequences only
func bitonicSeqSort(nums *[]int) error {
	if len(*nums)&(len(*nums)-1) != 0 {
		return fmt.Errorf("not power of two")
	}

	isB, _ := isBitonic(*nums)
	if !isB {
		return fmt.Errorf("input is not bitonic")
	}

	bitonicSeqSortRecurse(*nums)
	return nil
}

// convert the original sequence into a bitonic one
/*
•
Build a single bitonic sequence from the given sequence
—any sequence of length 2 is a bitonic sequence.
—build bitonic sequence of length 4
– sort first two elements using ⊕BM[2]
– sort next two using ӨBM[2]
- repeatedly merge

https://wiki.rice.edu/confluence/download/attachments/4435861/comp322-s12-lec28-slides-JMC.pdf?version=1&modificationDate=1333163955158

lengths power of two for convenience
*/
func bitonicBuildRecurse(nums *[]int, k int) error {
	// in steps of k
	for i := 0; i < len(*nums); i += 2 * k {
		// first half swaps with positive comparator
		for j := 0; j < k-1; j++ {
			if !compareWithDirection((*nums)[i+j+0], (*nums)[i+j+1], increasing) {
				(*nums)[i+j+0], (*nums)[i+j+1] = (*nums)[i+j+1], (*nums)[i+j+0]
			}
		}

		// second half swaps with negative comparator
		for j := k; j < 2*k-1; j++ {
			if !compareWithDirection((*nums)[i+j+0], (*nums)[i+j+1], decreasing) {
				(*nums)[i+j+0], (*nums)[i+j+1] = (*nums)[i+j+1], (*nums)[i+j+0]
			}
		}
	}
	return nil
}

func mergeBitonicGoroutineDecreasing(nums *[]int, offset, l int, wg *sync.WaitGroup) {
	storage := make([]int, l)
	i := l/2 - 1
	j := 0
	k := 0

	seq1 := (*nums)[offset : offset+l/2]
	seq2 := (*nums)[offset+l/2 : offset+l]

	for {
		if k == l {
			break
		}

		// assign from j if we've exhausted i
		if j >= l/2 {
			storage[k] = seq1[i]
			i--
			k++
			continue
		}

		// assign from i if we've exhausted j
		if i < 0 {
			storage[k] = seq2[j]
			j++
			k++
			continue
		}

		// assign from i if it's smaller
		if compareWithDirection(seq1[i], seq2[j], decreasing) {
			storage[k] = seq1[i]
			i--
			k++
		} else {
			storage[k] = seq2[j]
			j++
			k++
		}
	}

	for i, storedElem := range storage {
		(*nums)[offset+i] = storedElem
	}
	wg.Done()
}

func mergeBitonicGoroutineIncreasing(nums *[]int, offset, l int, wg *sync.WaitGroup) {
	storage := make([]int, l)
	i := 0
	j := l/2 - 1
	k := 0

	seq1 := (*nums)[offset : offset+l/2]
	seq2 := (*nums)[offset+l/2 : offset+l]

	for {
		if k == l {
			break
		}

		// assign from j if we've exhausted i
		if i >= l/2 {
			storage[k] = seq2[j]
			j--
			k++
			continue
		}

		// assign from i if we've exhausted j
		if j < 0 {
			storage[k] = seq1[i]
			i++
			k++
			continue
		}

		// assign from i if it's smaller
		if compareWithDirection(seq1[i], seq2[j], increasing) {
			storage[k] = seq1[i]
			i++
			k++
		} else {
			storage[k] = seq2[j]
			j--
			k++
		}
	}

	for i, storedElem := range storage {
		(*nums)[offset+i] = storedElem
	}
	wg.Done()
}

func mergeBitonic(nums *[]int, k, startingDirection int) {
	nChunks := len(*nums) / k

	var wg sync.WaitGroup
	wg.Add(nChunks)

	dir := startingDirection

	for chunkOffset := 0; chunkOffset < len(*nums); chunkOffset += k {
		if dir == 0 {
			go mergeBitonicGoroutineIncreasing(nums, chunkOffset, k, &wg)
		} else {
			go mergeBitonicGoroutineDecreasing(nums, chunkOffset, k, &wg)
		}
		dir = 1 - dir //invert comparator every consecutive block
	}

	wg.Wait()
}

func bitonicBuild(nums *[]int, direction int) error {
	err := bitonicBuildRecurse(nums, 2)

	for k := 4; k <= len(*nums); k *= 2 {
		// parallel merging
		mergeBitonic(nums, k, direction)
	}

	return err
}

// BitonicSortNaive is my handwritten version of the CLRS bitonic sorting network with goroutines and out-of-place mergesort
// default ascending
// the real full sort of regular integers
// need to take a slice cause of sentinel padding
func BitonicSortNaive(nums *[]int) {
	origLen := len(*nums)
	desiredLen := nextPowerOfTwo(origLen)

	// pad with sentinels
	for i := 0; i < desiredLen-origLen; i++ {
		*nums = append(*nums, maxInt)
	}

	bitonicBuild(nums, increasing)
	bitonicSeqSort(nums)

	// drop the sentinels
	*nums = (*nums)[:origLen]
}

func nextPowerOfTwo(n int) int {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++
	return n
}
