package gomapbench

//go:generate go run cmd/main.go cmd/config.go

import (
	"strconv"
	"testing"

	_ "github.com/cockroachdb/swiss"
)

var (
	cases = []int{6, 12, 18, 24, 30,
		64,
		128,
		256,
		512,
		1024,
		2048,
		4096,
		8192,
		1 << 16}
)

func runWith(f func(*testing.B, int), v ...int) func(*testing.B) {
	return func(b *testing.B) {
		for _, n := range v {
			b.Run(strconv.Itoa(n), func(b *testing.B) { f(b, n) })
		}
	}
}
