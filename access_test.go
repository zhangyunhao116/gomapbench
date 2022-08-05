package gomapbench

import (
	"strconv"
	"testing"
)

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

func benchmarkMapAccessHitInt64(b *testing.B, n int) {
	type ttype = int64
	m := make(map[ttype]int, n)
	for i := 0; i < n; i++ {
		m[ttype(i)] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[ttype(i&(n-1))]
	}
}

func benchmarkMapAccessHitInt32(b *testing.B, n int) {
	type ttype = int32
	m := make(map[ttype]int, n)
	for i := 0; i < n; i++ {
		m[ttype(i)] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[ttype(i&(n-1))]
	}
}

func benchmarkMapAccessHitStr(b *testing.B, n int) {
	type ttype = string
	m := make(map[ttype]int, n)
	ss := make([]string, n)
	for i := 0; i < n; i++ {
		m[strconv.Itoa(i)] = i
		ss[i] = strconv.Itoa(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[ss[i&(n-1)]]
	}
}
