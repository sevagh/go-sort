/*
Package gosort implements several sorting algorithms and experiments.
*/
package gosort

import "sort"

// copy-pasted from
// https://github.com/golang/go/blob/master/src/sort/sort.go

// HeapSort is copied from pkg/sort
func HeapSort(nums []int) {
	heapSort(nums, 0, len(nums))
}

// siftDown implements the heap property on nums[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown(nums []int, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && nums[first+child] < nums[first+child+1] {
			child++
		}
		if !(nums[first+root] < nums[first+child]) {
			return
		}
		nums[first+root], nums[first+child] = nums[first+child], nums[first+root]
		root = child
	}
}

func heapSort(nums []int, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(nums, i, hi, first)
	}

	// Pop elements, largest first, into end of nums.
	for i := hi - 1; i >= 0; i-- {
		nums[first], nums[first+i] = nums[first+i], nums[first]
		siftDown(nums, lo, i, first)
	}
}

// siftDown implements the heap property on nums[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownIface(data sort.Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

func heapSortIface(data sort.Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownIface(data, i, hi, first)
	}

	// Pop elements, largest first, into end of nums.
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		siftDownIface(data, lo, i, first)
	}
}
