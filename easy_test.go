package gosort_test

import (
	"testing"
	"math/rand"
	"sort"

	"github.com/sevagh/go-sort"
)

func TestTimSortEasy(t *testing.T) {
	easyArr := make([]int, 400)
	for j := 50; j < 100; j++ {
		easyArr[j] = rand.Int()
	}
	for j := 130; j < 200; j++ {
		easyArr[j] = 1000-rand.Int()
	}
	for j := 320; j < 400; j++ {
		easyArr[j] = rand.Int()/rand.Int()
	}

	gosort.TimSort(easyArr)

	if !sort.IntsAreSorted(easyArr) {
		t.Fatalf("timsort didn't work")
	} else {
		t.Logf("test concluded successfuly")
	}
}
