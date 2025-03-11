package gomapbench

//go:generate go run cmd/main.go cmd/config.go

import (
	"strconv"
	"testing"
)

var (
	cases = []int{1, 3, 7, 8, 12, 18,
		30, 200,
		1024,
		65536}
)

func runWith(f func(*testing.B, int), v ...int) func(*testing.B) {
	return func(b *testing.B) {
		for _, n := range v {
			b.Run(strconv.Itoa(n), func(b *testing.B) { f(b, n) })
		}
	}
}
