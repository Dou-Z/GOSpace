package main

import (
	"fmt"
)

func testIf1() {
	num := 10
	if num%2 == 0 {
		fmt.Printf("num:%d is even\n", num)
	} else {
		fmt.Printf("num :%d is odd\n")
	}
}

func testIf2() {
	// num := 10
	if num := 11; num%2 == 0 {
		fmt.Printf("num:%d is even\n", num)
	} else {
		fmt.Printf("num :%d is odd\n",num)
	}
}

func main() {
	testIf1()
	testIf2()
}
