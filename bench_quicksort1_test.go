package gosort_test

import (
	"sort"
	"testing"

	"github.com/sevagh/go-sort"
)

func BenchmarkQuickSort1Random8(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled8(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled16Values8(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1AllEqual8(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Ascending8(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Descending8(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PipeOrgan8(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushFront8(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushMiddle8(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Random32(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled32(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled16Values32(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1AllEqual32(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Ascending32(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Descending32(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PipeOrgan32(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushFront32(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushMiddle32(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Random128(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled128(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled16Values128(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1AllEqual128(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Ascending128(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Descending128(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PipeOrgan128(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushFront128(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushMiddle128(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Random1024(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled1024(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled16Values1024(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1AllEqual1024(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Ascending1024(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Descending1024(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PipeOrgan1024(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushFront1024(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushMiddle1024(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Random8192(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled8192(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled16Values8192(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1AllEqual8192(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Ascending8192(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Descending8192(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PipeOrgan8192(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushFront8192(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushMiddle8192(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Random65536(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled65536(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled16Values65536(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1AllEqual65536(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Ascending65536(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Descending65536(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PipeOrgan65536(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushFront65536(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushMiddle65536(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Random1048576(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled1048576(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Shuffled16Values1048576(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1AllEqual1048576(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Ascending1048576(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1Descending1048576(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PipeOrgan1048576(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushFront1048576(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort1PushMiddle1048576(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort1(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}