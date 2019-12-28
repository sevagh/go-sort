package gosort_test

import (
	"testing"
	"time"

	"github.com/sevagh/go-sort"
	"github.com/sevagh/go-sort/bigO"
)

func TestBigOTimSort(t *testing.T) {
	sizes := []int{10, 100, 1000, 10000, 100000, 10000000}
	times := []float64{}

	for _, size := range sizes {
		timeTaken := 0.0
		nums := RandomInt(size)

		start := time.Now()
		gosort.TimSort(nums)
		timeTaken += time.Since(start).Seconds()

		times = append(times, timeTaken)
	}

	complexity, _ := bigO.Estimate(sizes, times)
	t.Logf(complexity.String())
}

func TestBigOQuickSort1WorstCase(t *testing.T) {
	sizes := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}
	times := []float64{}

	for _, size := range sizes {
		timeTaken := 0.0
		nums := DescendingInt(size)

		start := time.Now()
		gosort.QuickSort1(nums)
		timeTaken += time.Since(start).Seconds()

		times = append(times, timeTaken)
	}

	complexity, _ := bigO.Estimate(sizes, times)
	t.Logf(complexity.String())
}

func TestBigOInsertionSortAscending(t *testing.T) {
	sizes := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}
	times := []float64{}

	for _, size := range sizes {
		timeTaken := 0.0
		nums := AscendingInt(size)

		start := time.Now()
		gosort.QuickSort1(nums)
		timeTaken += time.Since(start).Seconds()

		times = append(times, timeTaken)
	}

	complexity, _ := bigO.Estimate(sizes, times)
	t.Logf(complexity.String())
}

func TestBigOInsertionSortRandom(t *testing.T) {
	sizes := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}
	times := []float64{}

	for _, size := range sizes {
		timeTaken := 0.0
		nums := RandomInt(size)

		start := time.Now()
		gosort.QuickSort1(nums)
		timeTaken += time.Since(start).Seconds()

		times = append(times, timeTaken)
	}

	complexity, _ := bigO.Estimate(sizes, times)
	t.Logf(complexity.String())
}

func TestBigOPdqSort1WorstCaseQuickSort(t *testing.T) {
	sizes := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}
	times := []float64{}

	for _, size := range sizes {
		timeTaken := 0.0
		nums := DescendingInt(size)

		start := time.Now()
		gosort.PdqSort1(nums)
		timeTaken += time.Since(start).Seconds()

		times = append(times, timeTaken)
	}

	complexity, _ := bigO.Estimate(sizes, times)
	t.Logf(complexity.String())
}
