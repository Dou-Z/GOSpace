package main

import (
	"fmt"

	"github.com/study/day16/calc"
)

func main() {
	var sum = 0
	sum = calc.Add(1, 2)
	res := calc.Say("hello")
	fmt.Println(sum,res)
}
