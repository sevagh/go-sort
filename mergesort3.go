package gosort

// G&T implementation, iterative bottom up
func MergeSort3(nums []int) {
	in := make([]int, len(nums))
	copy(in, nums)
	out := make([]int, len(nums))

	var tmp []int
	n := len(in)

	for i := 1; i < n; i *= 2 {
		for j := 0; j < n; j += 2 * i {
			merge3(in, out, j, i)
		}
		tmp = in
		in = out
		out = tmp
	}
	copy(nums, in)
}

func merge3(in, out []int, start, inc int) {
	x := start
	end1 := min(len(in), start+inc)
	end2 := min(len(in), start+2*inc)
	y := start + inc
	z := start
	for x < end1 && y < end2 {
		if in[x] < in[y] {
			out[z] = in[x]
			z++
			x++
		} else {
			out[z] = in[y]
			z++
			y++
		}
	}
	if x < end1 {
		for i := 0; i < end1-x; i++ {
			out[z+i] = in[x+i]
		}
	} else {
		for i := 0; i < end2-y; i++ {
			out[z+i] = in[y+i]
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
