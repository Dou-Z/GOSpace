package main

import (
	"fmt"
)

func main() {
	arrayA := []int{100, 200}
	// a1 := 1
	testArrayPoint(&arrayA) // 1.传数组指针
	arrayB := arrayA[:]
	testArrayPoint(&arrayB) // 2.传切片
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
}

func testArrayPoint(x *[]int) {
	fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
}
