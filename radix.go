package gosort

// RadixSort is based on a radix sort + counting sort for decimal digits from CLRS
func RadixSort(nums []int) {
	count := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	result := make([]int, len(nums))

	// max int is 2,147,483,647 > we care about 10 digits
	for i := 0; i < 10; i++ {
		for _, num := range nums {
			count[(num/powersOf10[i])%10]++
		}

		resultIdx := 0
		for j, c := range count {
			if c != 0 {
				for k := 0; k < c; k++ {
					result[resultIdx] += j * powersOf10[i]
					resultIdx++
				}
			}
		}

		count = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		resultIdx = 0
	}

	copy(nums, result)
}

var powersOf10 = []int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 100000000000}
