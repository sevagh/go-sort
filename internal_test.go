package gosort

import (
	"sort"
	"testing"
)

func TestIsBitonic(t *testing.T) {
	type testcase struct {
		dat []int
		res bool
	}
	var testcases = []testcase{
		testcase{
			[]int{0, 1, 1},
			true,
		},
		testcase{
			[]int{1, 0, 1, 0},
			false,
		},
		testcase{
			[]int{1, 4, 6, 8, 3, 2},
			true,
		},
		testcase{
			[]int{6, 9, 4, 2, 3, 5},
			true,
		},
		testcase{
			[]int{9, 8, 3, 2, 4, 6},
			true,
		},
		testcase{
			[]int{9, 8, 3, 2, 4, 3, 6},
			false,
		},
	}

	for _, testcase := range testcases {
		realResult, _ := isBitonic(testcase.dat)
		if testcase.res != realResult {
			t.Errorf("%+v is bitonic: %t, expected: %t\n", testcase.dat, realResult, testcase.res)
		}
	}
}

func TestBitonicSeqBasic(t *testing.T) {
	myarr := []int{5, 4, 2, 8, 9, -1, -1, -3}

	t.Logf("myarr before: %+v\n", myarr)
	err := bitonicSeqSort(&myarr)
	if err == nil {
		t.Errorf("bitSeqSort expected err not bitonic\n")
	}
	return
}

func TestBitonicSeqBasic2(t *testing.T) {
	myarr := []int{-32, 5, 13, 4, 85, 2, 8, 337}

	t.Logf("myarr before: %+v\n", myarr)
	err := bitonicSeqSort(&myarr)
	if err == nil {
		t.Errorf("bitSeqSort err: %+v\n", err)
	}
	return
}

func TestBitonicSeqBasic3(t *testing.T) {
	myarr := []int{5, 8, 3}

	err := bitonicSeqSort(&myarr)
	if err == nil {
		t.Fatalf("bitSeqSort should have had err\n")
	}
}

func TestBitonicSeqZeroOneEvenLen(t *testing.T) {
	myarr := []int{0, 1, 1, 0}

	err := bitonicSeqSort(&myarr)
	if err != nil {
		t.Errorf("bitSeqSort err: %+v\n", err)
	}

	prev := myarr[0]
	for _, num := range myarr[1:] {
		if num < prev {
			t.Errorf("bitonicSeqSort failed to sort, got incorrectly ordered pair '%d,%d'", prev, num)
		}
		prev = num
	}
}

// the zero-one principle according to CLRS
// sorting networks are correct for all integers
// if they can be shown to work with 0s and 1s
func TestBitonicSeqZeroOneUnevenLen(t *testing.T) {
	myarr := []int{0, 1, 0}

	err := bitonicSeqSort(&myarr)
	if err == nil {
		t.Fatalf("bitSeqSort should have had error")
	}
}

func TestBitonicBuildPowerOf2Len8(t *testing.T) {
	// 8, power of two for simplicity
	// TODO: add code later for non-power-of-two sizes
	testarr := []int{3, -19, -12, 1, -4, -4, -8, -6}

	t.Logf("len test data: %d", len(testarr))

	if amIBitonic, _ := isBitonic(testarr); amIBitonic {
		t.Errorf("test array is already bitonic! we're testing nothing")
	}

	bitonicBuild(&testarr, increasing)

	if amIBitonic, _ := isBitonic(testarr); !amIBitonic {
		t.Errorf("bitonic build didn't build a bitonic seq!")
	}
}

func TestBitonicBuildPowerOf2Len128(t *testing.T) {
	// 128, power of two for simplicity
	// TODO: add code later for non-power-of-two sizes
	testarr := []int{7, -4, 9, -10, 19, 11, -3, -3, -10, 6, 8, -11, 13, -11, -18, -6, 7, 19, 12, -11, -9, 4, -3, -17, -2, 18, 13, -12, 0, 4, 7, 3, -20, -14, -18, 12, -8, 12, 0, -11, -14, -5, 10, 2, 11, 1, -6, -11, 2, 11, 5, -16, -1, -14, 3, 13, 9, 0, 10, 14, 17, -16, 11, -8, -9, -8, 6, -20, -19, -7, -14, -16, -17, -16, 15, -2, -16, 1, 11, 8, -18, 18, -6, -9, 19, 2, 13, 16, 7, 9, -12, -16, 13, -3, -8, 16, -2, -3, 15, -4, 13, -10, -9, -6, -9, 2, -5, -6, 8, 14, 5, 2, 11, -9, 3, 3, -1, 0, -13, 16, -3, -8, 12, -7, -5, -13, -19, -3}

	t.Logf("len test data: %d", len(testarr))

	if amIBitonic, _ := isBitonic(testarr); amIBitonic {
		t.Errorf("test array is already bitonic! we're testing nothing")
	}

	bitonicBuild(&testarr, increasing)

	if amIBitonic, _ := isBitonic(testarr); !amIBitonic {
		t.Errorf("bitonic build didn't build a bitonic seq!")
	}
}

func TestBitonicBentleyMcIlroyFailedCase(t *testing.T) {
	testarr := []int{
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
		0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
	}

	BitonicSortNaive(&testarr)

	if !sort.IntsAreSorted(testarr) {
		t.Fatalf("ints not sorted\n")
	}
}
