package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var a = 10
	var b = 20
	t.Logf("a = %d,b=%d\n", a, b)
	c := Add(a, b)
	if c == 30 {
		t.Fatalf("invalid a + b,c=%d\n", c)
	}
	t.Logf("a = %d,b=%d,sum=%d\n", a, b, a+b)
}
