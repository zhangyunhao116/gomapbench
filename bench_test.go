package gomapbench

import (
	"github.com/cockroachdb/swiss"
	"strconv"
	"testing"
)

func BenchmarkMapIter(b *testing.B) {
	b.Run("Int", runWith(benchmarkMapIter, cases...))
}

func BenchmarkMapAccessHit(b *testing.B) {
	b.Run("Int64", runWith(benchmarkMapAccessHitInt64, cases...))
	b.Run("Int32", runWith(benchmarkMapAccessHitInt32, cases...))
	b.Run("Str", runWith(benchmarkMapAccessHitStr, cases...))
}

func BenchmarkMapAccessMiss(b *testing.B) {
	b.Run("Int64", runWith(benchmarkMapAccessMissInt64, cases...))
	b.Run("Int32", runWith(benchmarkMapAccessMissInt32, cases...))
	b.Run("Str", runWith(benchmarkMapAccessMissStr, cases...))
}

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

func BenchmarkMapAssignReuse(b *testing.B) {
	b.Run("Int64", runWith(benchmarkMapAssignReuseInt64, cases...))
	b.Run("Int32", runWith(benchmarkMapAssignReuseInt32, cases...))
	b.Run("Str", runWith(benchmarkMapAssignReuseStr, cases...))
}

func benchmarkMapIter(b *testing.B, n int) {
	m := swiss.New[int, int](n)
	for i := 0; i < n; i++ {
		m.Put(i, i)
	}
	b.ResetTimer()
	var tmp int
	for i := 0; i < b.N; i++ {
		m.All(func(k, v int) bool {
			tmp += k + v
			return true
		})
	}
}

func benchmarkMapAccessMissInt64(b *testing.B, n int) {
	m := swiss.New[int64, int](0)
	for j := 0; j < n; j++ {
		m.Put(int64(j), j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Get(int64(i) + int64(n))
	}
}

func benchmarkMapAccessMissInt32(b *testing.B, n int) {
	m := swiss.New[int32, int](0)
	for j := 0; j < n; j++ {
		m.Put(int32(j), j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Get(int32(i) + int32(n))
	}
}

func benchmarkMapAccessMissStr(b *testing.B, n int) {
	m := swiss.New[string, int](0)
	for j := 0; j < n; j++ {
		m.Put(strconv.Itoa(j), j)
	}
	miss := make([]string, n)
	for j := 0; j < n; j++ {
		miss[j] = strconv.Itoa(j * (-1))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Get(miss[i&(n-1)])
	}
}

func benchmarkMapAccessHitInt64(b *testing.B, n int) {
	type ttype = int64
	m := swiss.New[ttype, int](n)
	for i := 0; i < n; i++ {
		m.Put(ttype(i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Get(ttype(i & (n - 1)))
	}
}

func benchmarkMapAccessHitInt32(b *testing.B, n int) {
	type ttype = int32
	m := swiss.New[ttype, int](n)
	for i := 0; i < n; i++ {
		m.Put(ttype(i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Get(ttype(i & (n - 1)))
	}
}

func benchmarkMapAccessHitStr(b *testing.B, n int) {
	type ttype = string
	m := swiss.New[ttype, int](n)
	ss := make([]string, n)
	for i := 0; i < n; i++ {
		m.Put(strconv.Itoa(i), i)
		ss[i] = strconv.Itoa(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Get(ss[i&(n-1)])
	}
}

func benchmarkMapAssignGrowInt32(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := swiss.New[int32, int](0)
		for j := 0; j < n; j++ {
			m.Put(int32(j), j)
		}
	}
}

func benchmarkMapAssignGrowInt64(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := swiss.New[int64, int](0)
		for j := 0; j < n; j++ {
			m.Put(int64(j), j)
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
		m := swiss.New[string, int](0)
		for j := 0; j < n; j++ {
			m.Put(k[j], j)
		}
	}
}

func benchmarkMapAssignPreAllocateInt32(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := swiss.New[int32, int](n)
		for j := 0; j < n; j++ {
			m.Put(int32(j), j)
		}
	}
}

func benchmarkMapAssignPreAllocateInt64(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := swiss.New[int64, int](n)
		for j := 0; j < n; j++ {
			m.Put(int64(j), j)
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
		m := swiss.New[string, int](n)
		for j := 0; j < n; j++ {
			m.Put(k[j], j)
		}
	}
}

func benchmarkMapAssignReuseInt32(b *testing.B, n int) {
	m := swiss.New[int32, int](n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			m.Put(int32(j), j)
		}
		m.Clear()
	}
}

func benchmarkMapAssignReuseInt64(b *testing.B, n int) {
	m := swiss.New[int64, int](n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			m.Put(int64(j), j)
		}
		m.Clear()
	}
}

func benchmarkMapAssignReuseStr(b *testing.B, n int) {
	k := make([]string, n)
	for i := 0; i < len(k); i++ {
		k[i] = strconv.Itoa(i)
	}
	m := swiss.New[string, int](n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			m.Put(k[j], j)
		}
		m.Clear()
	}
}
