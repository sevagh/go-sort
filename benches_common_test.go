package mergesort_test

import "math/rand"

func ShuffledInt(size int) []int {
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = i
	}
	rand.Shuffle(size, func(i, j int) {
		ints[i], ints[j] = ints[j], ints[i]
	})
	return ints
}

func Shuffled16ValuesInt(size int) []int {
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = i % 16
	}
	rand.Shuffle(size, func(i, j int) {
		ints[i], ints[j] = ints[j], ints[i]
	})
	return ints
}

func AllEqualInt(size int) []int {
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = 0
	}
	return ints
}

func AscendingInt(size int) []int {
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = i
	}
	return ints
}

func DescendingInt(size int) []int {
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = (size - 1) - i
	}
	return ints
}

func PipeOrganInt(size int) []int {
	ints := make([]int, size)
	for i := 0; i < size/2; i++ {
		ints[i] = i
	}
	for i := size / 2; i < size; i++ {
		ints[i] = size - 1
	}
	return ints
}

func PushFrontInt(size int) []int {
	ints := make([]int, size)
	for i := 1; i < size; i++ {
		ints[i-1] = i
	}
	ints[size-1] = 0
	return ints
}

func PushMiddleInt(size int) []int {
	ints := make([]int, size)
	idx := 0
	for i := 0; i < size; i++ {
		if i != size/2 {
			ints[idx] = i
			idx++
		}
	}
	ints[idx] = 0
	return ints
}
