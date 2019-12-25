package gosort_test

import (
	"sort"
	"testing"

	"github.com/sevagh/go-sort"
)

func BenchmarkQuickSort2Random8(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled8(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled16Values8(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2AllEqual8(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Ascending8(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Descending8(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PipeOrgan8(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushFront8(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushMiddle8(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Random32(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled32(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled16Values32(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2AllEqual32(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Ascending32(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Descending32(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PipeOrgan32(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushFront32(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushMiddle32(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Random128(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled128(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled16Values128(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2AllEqual128(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Ascending128(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Descending128(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PipeOrgan128(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushFront128(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushMiddle128(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Random1024(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled1024(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled16Values1024(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2AllEqual1024(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Ascending1024(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Descending1024(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PipeOrgan1024(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushFront1024(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushMiddle1024(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Random8192(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled8192(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled16Values8192(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2AllEqual8192(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Ascending8192(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Descending8192(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PipeOrgan8192(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushFront8192(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushMiddle8192(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Random65536(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled65536(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled16Values65536(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2AllEqual65536(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Ascending65536(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Descending65536(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PipeOrgan65536(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushFront65536(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushMiddle65536(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Random1048576(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled1048576(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Shuffled16Values1048576(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2AllEqual1048576(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Ascending1048576(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2Descending1048576(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PipeOrgan1048576(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushFront1048576(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkQuickSort2PushMiddle1048576(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.QuickSort2(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}
