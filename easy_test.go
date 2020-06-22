package gosort_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/sevagh/go-sort"
)

func TestTimSort1Easy(t *testing.T) {
	easyArr := make([]int, 400)
	for j := 50; j < 100; j++ {
		easyArr[j] = rand.Int()
	}
	for j := 130; j < 200; j++ {
		easyArr[j] = 1000 - rand.Int()
	}
	for j := 320; j < 400; j++ {
		easyArr[j] = rand.Int() / rand.Int()
	}

	gosort.TimSort1(easyArr)

	if !sort.IntsAreSorted(easyArr) {
		t.Fatalf("timsort didn't work")
	} else {
		t.Logf("test concluded successfuly")
	}
}

func TestTimSort2Easy(t *testing.T) {
	easyArr := make([]int, 400)
	for j := 50; j < 100; j++ {
		easyArr[j] = rand.Int()
	}
	for j := 130; j < 200; j++ {
		easyArr[j] = 1000 - rand.Int()
	}
	for j := 320; j < 400; j++ {
		easyArr[j] = rand.Int() / rand.Int()
	}

	gosort.TimSort2(easyArr)

	if !sort.IntsAreSorted(easyArr) {
		t.Fatalf("timsort didn't work")
	} else {
		t.Logf("test concluded successfuly")
	}
}
