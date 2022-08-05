package gomapbench

import (
	"strconv"
	"testing"
)

func BenchmarkMapAssignGrow(b *testing.B) {
	b.Run("Int64", runWith(benchmarkMapAssignGrowInt64, cases...))
	b.Run("Int32", runWith(benchmarkMapAssignGrowInt32, cases...))
	b.Run("Str", runWith(benchmarkMapAssignGrowStr, cases...))
}

func BenchmarkMapAssignPreAllocate(b *testing.B) {
	b.Run("Int64", runWith(benchmarkMapAssignPreAllocateInt64, cases...))
	b.Run("Int32", runWith(benchmarkMapAssignPreAllocateInt32, cases...))
	b.Run("Str", runWith(benchmarkMapAssignPreAllocateStr, cases...))
}

func benchmarkMapAssignGrowInt32(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := make(map[int32]int)
		for j := 0; j < n; j++ {
			m[int32(j)] = j
		}
	}
}

func benchmarkMapAssignGrowInt64(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := make(map[int64]int)
		for j := 0; j < n; j++ {
			m[int64(j)] = j
		}
	}
}

func benchmarkMapAssignGrowStr(b *testing.B, n int) {
	k := make([]string, n)
	for i := 0; i < len(k); i++ {
		k[i] = strconv.Itoa(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := make(map[string]int)
		for j := 0; j < n; j++ {
			a[k[j]] = i
		}
	}
}

func benchmarkMapAssignPreAllocateInt32(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := make(map[int32]int, n)
		for j := 0; j < n; j++ {
			m[int32(j)] = j
		}
	}
}

func benchmarkMapAssignPreAllocateInt64(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := make(map[int64]int, n)
		for j := 0; j < n; j++ {
			m[int64(j)] = j
		}
	}
}

func benchmarkMapAssignPreAllocateStr(b *testing.B, n int) {
	k := make([]string, n)
	for i := 0; i < len(k); i++ {
		k[i] = strconv.Itoa(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := make(map[string]int, n)
		for j := 0; j < n; j++ {
			a[k[j]] = i
		}
	}
}
