package main

import (
	"fmt"
)

// 值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。
func test(x [2]int) {
	fmt.Printf("x:%p\n", &x)
	x[1] = 1000
}
func testLenArr() {
	a := [2]int{}
	println(len(a), cap(a))
}

// 数组拷贝和传参

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func testArray12(){
	a := [3]int{10,20,30}
	b := a
	b[0] = 1000
	fmt.Printf("a=%v\n",a)
	fmt.Printf("b=%v\n",b)
}

func testArr() {
	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}

func main() {
	// a := [2]int{}
	// fmt.Printf("a:%p\n", &a)

	// test(a)
	// fmt.Println(a)
	// testLenArr()
	// testArr()
	testArray12()
}
