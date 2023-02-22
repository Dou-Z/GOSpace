package main

import "fmt"

func testArray() {
	v1 := [2]int{1, 2}
	v2 := v1

	fmt.Printf("v1的内存地址：%p \n", &v1)
	fmt.Printf("v2的内存地址：%p \n", &v2)

	v1[0] = 6666
	fmt.Println(v1, v2)
}

func testSlice() {
	v1 := []int{1, 2}
	v2 := v1

	fmt.Printf("v1的内存地址：%p \n", &v1)
	fmt.Printf("v2的内存地址：%p \n", &v2)

	v1[0] = 6666
	fmt.Println(v1, v2)
}

func test1(){
	v1 := [][]int{make([]int, 2,5),make([]int, 2,3)}
	v2 := append(v1[0],99)

	fmt.Println(v1)
	fmt.Println(v2)

	v2[0] =69
	fmt.Println(v1)
	fmt.Println(v2)
}
func main() {
	// testArray()
	// testSlice()
	test1()
}
