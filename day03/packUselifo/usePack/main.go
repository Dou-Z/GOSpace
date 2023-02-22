package main

import (
	"fmt"

	"github.com/study/day03/calc"
)

func main() {
	var s1 int = 10
	var s2 int = 20
	sum := calc.Add(s1, s2)
	fmt.Printf("s1+s2=%d\n", sum)
}
