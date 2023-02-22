package main

import (
	"fmt"
)

func testSlince1() {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int
	b = a[1:4]
	fmt.Printf("Slice b:%v\n", b)
}

func testSlice2() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("array a:%v type of b:%T\n", a, a)
	b := a[2:5]
	fmt.Printf("slice b:%v type of b:%T\n", b, b)
	b[0] = b[0] + 10
	b[1] = b[1] + 20
	b[2] = b[2] + 30
	fmt.Printf("after modify silce b,array a:%v type of a:%T\n", a, a)

}

func testArraySlice2() {
	numa := [3]int{70, 71, 72}
	// 创建切片，包含整个数组的所有元素
	num1 := numa[:]
	num2 := numa[:]
	fmt.Println("array before change 1", numa)
	num1[0] = 100
	fmt.Println("array after modify to slice num1", numa)
	num2[1] = 1020
	fmt.Println("array after modify to slice num2", numa)
}

func main() {
	// testSlince1()
	// testSlice2()
	testArraySlice2()
}
