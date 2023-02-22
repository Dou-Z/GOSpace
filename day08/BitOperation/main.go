package main

import "fmt"

func main() {
	// 与运算
	res := 5 & 99
	fmt.Println(res)

	// 或运算
	res2 := 5 | 99
	fmt.Println(res2)

	// 异或运算
	res3 := 5 ^ 99
	fmt.Println(res3)

	// 左移运算
	res4 := 5 << 2
	// 右移运算
	res5 := 5 >> 1
	fmt.Println(res4, res5)

	// 比较清楚 &^
	res6 := 5 &^ 99
	fmt.Println(res6)
}
