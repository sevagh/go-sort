package bigO_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sevagh/go-sort"
	"github.com/sevagh/go-sort/bigO"
)

func genRandomInt(size int) []int {
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = rand.Int()
	}
	return ints
}

func ExampleBigOMergeSort() {
	sizes := []int{10, 100, 1000, 10000, 100000, 10000000}
	iters := 5
	times := []float64{}

	for _, size := range sizes {
		timeTaken := 0.0
		nums := genRandomInt(size)

		currIter := 0
		for {
			start := time.Now()
			gosort.MergeSort3(nums)
			timeTaken += time.Since(start).Seconds()
			currIter++
			if currIter > iters {
				break
			}
		}

		timeTaken /= float64(iters)

		times = append(times, timeTaken)
	}

	complexity, _ := bigO.Estimate(sizes, times)
	fmt.Printf(complexity.String())
	// Output: O(Nlg(N))
}
