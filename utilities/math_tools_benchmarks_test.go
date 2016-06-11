package math_tools

import (
	"testing"
)

func BenchmarkIsPrime53(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(53)
	}
}
