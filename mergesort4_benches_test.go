package mergesort_test

import (
	"sort"
	"testing"

	"github.com/sevagh/go-sort"
)

func BenchmarkMergeSort4Random8(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled8(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled16Values8(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4AllEqual8(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Ascending8(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Descending8(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PipeOrgan8(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushFront8(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushMiddle8(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Random32(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled32(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled16Values32(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4AllEqual32(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Ascending32(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Descending32(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PipeOrgan32(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushFront32(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushMiddle32(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Random128(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled128(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled16Values128(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4AllEqual128(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Ascending128(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Descending128(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PipeOrgan128(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushFront128(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushMiddle128(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Random1024(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled1024(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled16Values1024(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4AllEqual1024(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Ascending1024(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Descending1024(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PipeOrgan1024(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushFront1024(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushMiddle1024(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Random8092(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(8092)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled8092(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(8092)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled16Values8092(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(8092)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4AllEqual8092(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(8092)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Ascending8092(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(8092)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Descending8092(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(8092)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PipeOrgan8092(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(8092)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushFront8092(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(8092)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushMiddle8092(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(8092)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Random65536(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled65536(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled16Values65536(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4AllEqual65536(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Ascending65536(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Descending65536(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PipeOrgan65536(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushFront65536(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushMiddle65536(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Random1048576(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled1048576(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Shuffled16Values1048576(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4AllEqual1048576(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Ascending1048576(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4Descending1048576(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PipeOrgan1048576(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushFront1048576(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkMergeSort4PushMiddle1048576(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergesort.MergeSort4(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}
