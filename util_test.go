package gomapbench

import (
	"strconv"
	"testing"
)

var (
	case1 = []int{1 << 4, 1 << 8, 1 << 16}
	case2 = []int{6, 12, 18, 24, 30, 64, 128, 256, 1 << 16}
)

func runWith(f func(*testing.B, int), v ...int) func(*testing.B) {
	return func(b *testing.B) {
		for _, n := range v {
			b.Run(strconv.Itoa(n), func(b *testing.B) { f(b, n) })
		}
	}
}
