package gomapbench

import (
	{{.Imports}}
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
	m := {{New "int" "int" "n"}}
	for i := 0; i < n; i++ {
		{{Store "m" "i" "i"}}
	}
	b.ResetTimer()
	var tmp int
	for i := 0; i < b.N; i++ {
		{{RangeAll "m" "k" "v" "tmp += k + v"}}
	}
}

func benchmarkMapAccessMissInt64(b *testing.B, n int) {
	m := {{New "int64" "int" "0"}}
	for j := 0; j < n; j++ {
		{{Store "m" "int64(j)" "j"}}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = {{Load "m" "int64(i)+int64(n)"}}
	}
}

func benchmarkMapAccessMissInt32(b *testing.B, n int) {
	m := {{New "int32" "int" "0"}}
	for j := 0; j < n; j++ {
		{{Store "m" "int32(j)" "j"}}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = {{Load "m" "int32(i)+int32(n)"}}
	}
}

func benchmarkMapAccessMissStr(b *testing.B, n int) {
	m := {{New "string" "int" "0"}}
	for j := 0; j < n; j++ {
		{{Store "m" "strconv.Itoa(j)" "j"}}
	}
	miss := make([]string, n)
	for j := 0; j < n; j++ {
		miss[j] = strconv.Itoa(j * (-1))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = {{Load "m" "miss[i&(n-1)]"}}
	}
}

func benchmarkMapAccessHitInt64(b *testing.B, n int) {
	type ttype = int64
	m := {{New "ttype" "int" "n"}}
	for i := 0; i < n; i++ {
		{{Store "m" "ttype(i)" "i"}}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = {{Load "m" "ttype(i&(n-1))"}}
	}
}

func benchmarkMapAccessHitInt32(b *testing.B, n int) {
	type ttype = int32
	m := {{New "ttype" "int" "n"}}
	for i := 0; i < n; i++ {
		{{Store "m" "ttype(i)" "i"}}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = {{Load "m" "ttype(i&(n-1))"}}
	}
}

func benchmarkMapAccessHitStr(b *testing.B, n int) {
	type ttype = string
	m := {{New "ttype" "int" "n"}}
	ss := make([]string, n)
	for i := 0; i < n; i++ {
		{{Store "m" "strconv.Itoa(i)" "i"}}
		ss[i] = strconv.Itoa(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = {{Load "m" "ss[i&(n-1)]"}}
	}
}


func benchmarkMapAssignGrowInt32(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := {{New "int32" "int" "0"}}
		for j := 0; j < n; j++ {
			{{Store "m" "int32(j)" "j"}}
		}
	}
}

func benchmarkMapAssignGrowInt64(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := {{New "int64" "int" "0"}}
		for j := 0; j < n; j++ {
			{{Store "m" "int64(j)" "j"}}
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
		m := {{New "string" "int" "0"}}
		for j := 0; j < n; j++ {
			{{Store "m" "k[j]" "j"}}
		}
	}
}

func benchmarkMapAssignPreAllocateInt32(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := {{New "int32" "int" "n"}}
		for j := 0; j < n; j++ {
			{{Store "m" "int32(j)" "j"}}
		}
	}
}

func benchmarkMapAssignPreAllocateInt64(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := {{New "int64" "int" "n"}}
		for j := 0; j < n; j++ {
			{{Store "m" "int64(j)" "j"}}
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
		m := {{New "string" "int" "n"}}
		for j := 0; j < n; j++ {
			{{Store "m" "k[j]" "j"}}
		}
	}
}

func benchmarkMapAssignReuseInt32(b *testing.B, n int) {
	m := {{New "int32" "int" "n"}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			{{Store "m" "int32(j)" "j"}}
		}
		{{DeleteAll "m"}}
	}
}

func benchmarkMapAssignReuseInt64(b *testing.B, n int) {
	m := {{New "int64" "int" "n"}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			{{Store "m" "int64(j)" "j"}}
		}
		{{DeleteAll "m"}}
	}
}

func benchmarkMapAssignReuseStr(b *testing.B, n int) {
	k := make([]string, n)
	for i := 0; i < len(k); i++ {
		k[i] = strconv.Itoa(i)
	}
	m := {{New "string" "int" "n"}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			{{Store "m" "k[j]" "j"}}
		}
		{{DeleteAll "m"}}
	}
}
