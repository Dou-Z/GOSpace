package main

import (
	"fmt"
)

func main() {

	// var a int
	// var b bool
	// var c string
	// var d float32
	var (
		a int
		b bool
		c string
		d float32
	)
	fmt.Printf("a=%d b=%t c=%s d=%f\n", a, b, c, d)

	a = 1
	b = true
	c = "go"
	d = 0.1
	fmt.Printf("a=%d b=%t c=%s d=%f\n", a, b, c, d)
}
