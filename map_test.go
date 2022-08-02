package gomapbench

import (
	"strconv"
	"testing"
)

var capSlice = []int{1 << 8, 1 << 16}

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

func BenchmarkMapAssignGrow(b *testing.B) {
	b.Run("Int64", runWith(benchmarkMapAssignGrowInt64, capSlice...))
	b.Run("Int32", runWith(benchmarkMapAssignGrowInt32, capSlice...))
	b.Run("Str", runWith(benchmarkMapAssignGrowStr, capSlice...))
}

func benchmarkMapAssignWithoutGrowInt32(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := make(map[int32]int, n)
		for j := 0; j < n; j++ {
			m[int32(j)] = j
		}
	}
}

func benchmarkMapAssignWithoutGrowInt64(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := make(map[int64]int, n)
		for j := 0; j < n; j++ {
			m[int64(j)] = j
		}
	}
}

func benchmarkMapAssignWithoutGrowStr(b *testing.B, n int) {
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

func BenchmarkMapAssignWithoutGrow(b *testing.B) {
	b.Run("Int64", runWith(benchmarkMapAssignWithoutGrowInt64, capSlice...))
	b.Run("Int32", runWith(benchmarkMapAssignWithoutGrowInt32, capSlice...))
	b.Run("Str", runWith(benchmarkMapAssignWithoutGrowStr, capSlice...))
}

func benchmarkMapAccessMissInt64(b *testing.B, n int) {
	m := make(map[int64]int)
	for j := 0; j < n; j++ {
		m[int64(j)] = j
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[int64(i)+int64(n)]
	}
}

func benchmarkMapAccessMissInt32(b *testing.B, n int) {
	m := make(map[int32]int)
	for j := 0; j < n; j++ {
		m[int32(j)] = j
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[int32(i)+int32(n)]
	}
}

func benchmarkMapAccessMissStr(b *testing.B, n int) {
	m := make(map[string]int)
	for j := 0; j < n; j++ {
		m[strconv.Itoa(j)] = j
	}
	miss := make([]string, n)
	for j := 0; j < n; j++ {
		miss[j] = strconv.Itoa(j * (-1))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[miss[i&(n-1)]]
	}
}

func BenchmarkMapAccessMiss(b *testing.B) {
	b.Run("Int64", runWith(benchmarkMapAccessMissInt64, capSlice...))
	b.Run("Int32", runWith(benchmarkMapAccessMissInt32, capSlice...))
	b.Run("Str", runWith(benchmarkMapAccessMissStr, capSlice...))
}

func runWith(f func(*testing.B, int), v ...int) func(*testing.B) {
	return func(b *testing.B) {
		for _, n := range v {
			b.Run(strconv.Itoa(n), func(b *testing.B) { f(b, n) })
		}
	}
}
