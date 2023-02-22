package main

import (
	"fmt"
)

func add(a, b int) int {

	return a + b
}
func test1() {
	f1 := add
	fmt.Printf("type of f1 = %T\n", f1)
	sum := f1(2, 4)
	fmt.Printf("sum=%d\n", sum)
}

func testFunc() {
	f1 := func(a, b int) int {
		return a - b
	}
	fmt.Printf("type of f1 = %T\n", f1)
	sub := f1(2, 5)
	fmt.Printf("sub=%d\n", sub)
}

// defer 与匿名函数的使用
func testFunc3() {
	var i int = 1
	defer fmt.Printf("defer i=%d\n", i)
	i = 100
	fmt.Printf("i=%d\n", i)
	return
}

func testFunc4() {
	var i int = 1
	defer func() {
		fmt.Printf("defer i=%d\n", i)
	}()

	i = 100
	fmt.Printf("i=%d\n", i)
	return
}

// 闭包函数的使用
func sub(a, b int) int {
	return a - b
}

func add1(a, b int) int {
	return a - b
}

func testC(a, b int32) int32 {
	return a + b
}

func calc(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func testFunc5() {
	sum := calc(100, 200, add1)
	sub := calc(200, 500, sub)
	fmt.Printf("sum=%d,sub=%d\n", sum, sub)
}

func main() {
	// test1()
	// testFunc()
	// testFunc3()
	// testFunc4()
	testFunc5()

}
