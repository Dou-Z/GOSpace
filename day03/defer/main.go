package main

import (
	"fmt"
)

func testDefer() {
	defer fmt.Println("hello")
	defer fmt.Println("hello v1")
	defer fmt.Println("hello v2")

	fmt.Println("Good")
	fmt.Println("Good2")
	res := 2 / 10
	fmt.Printf("%d\n", res)
}

func testDefer2() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("i=%d\n", i)
	}

	fmt.Printf("Running!!!\n")
	fmt.Printf("return\n")

}

func main() {
	testDefer()
	// testDefer2()
}
