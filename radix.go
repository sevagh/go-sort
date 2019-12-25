package gosort

// a radix sort with counting sort from CLRS
//
// assume the numbers are decimal i.e. digits = 0-9
// enabling us to use counting sort easily

var PowersOf10 = []int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 100000000000}

func RadixSort(nums []int) {
	count := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	result := make([]int, len(nums))

	// max int is 2,147,483,647 > we care about 10 digits
	for i := 0; i < 10; i++ {
		for _, num := range nums {
			count[(num/PowersOf10[i])%10] += 1
		}

		resultIdx := 0
		for j, c := range count {
			if c != 0 {
				for k := 0; k < c; k++ {
					result[resultIdx] += j * PowersOf10[i]
					resultIdx++
				}
			}
		}

		count = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		resultIdx = 0
	}

	copy(nums, result)
}
