package main

import (
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 10
		sum := Add(a, i)
		b.Logf("sum=%d\n", sum)
	}
}
