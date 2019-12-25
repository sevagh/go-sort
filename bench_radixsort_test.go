package gosort_test

import (
	"sort"
	"testing"

	"github.com/sevagh/go-sort"
)

func BenchmarkRadixSortRandom8(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled8(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled16Values8(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAllEqual8(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAscending8(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortDescending8(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPipeOrgan8(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushFront8(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushMiddle8(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortRandom32(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled32(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled16Values32(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAllEqual32(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAscending32(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortDescending32(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPipeOrgan32(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushFront32(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushMiddle32(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(32)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortRandom128(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled128(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled16Values128(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAllEqual128(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAscending128(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortDescending128(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPipeOrgan128(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushFront128(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushMiddle128(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(128)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortRandom1024(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled1024(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled16Values1024(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAllEqual1024(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAscending1024(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortDescending1024(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPipeOrgan1024(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushFront1024(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushMiddle1024(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortRandom8192(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled8192(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled16Values8192(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAllEqual8192(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAscending8192(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortDescending8192(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPipeOrgan8192(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushFront8192(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushMiddle8192(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(8192)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortRandom65536(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled65536(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled16Values65536(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAllEqual65536(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAscending65536(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortDescending65536(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPipeOrgan65536(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushFront65536(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushMiddle65536(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(65536)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortRandom1048576(b *testing.B) {
	b.StopTimer()
	nums := RandomInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled1048576(b *testing.B) {
	b.StopTimer()
	nums := ShuffledInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortShuffled16Values1048576(b *testing.B) {
	b.StopTimer()
	nums := Shuffled16ValuesInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAllEqual1048576(b *testing.B) {
	b.StopTimer()
	nums := AllEqualInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortAscending1048576(b *testing.B) {
	b.StopTimer()
	nums := AscendingInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortDescending1048576(b *testing.B) {
	b.StopTimer()
	nums := DescendingInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPipeOrgan1048576(b *testing.B) {
	b.StopTimer()
	nums := PipeOrganInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushFront1048576(b *testing.B) {
	b.StopTimer()
	nums := PushFrontInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}

func BenchmarkRadixSortPushMiddle1048576(b *testing.B) {
	b.StopTimer()
	nums := PushMiddleInt(1048576)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gosort.RadixSort(nums)
	}
	b.StopTimer()
	if !sort.IntsAreSorted(nums) {
		b.Fatalf("fail!")
	}
}
