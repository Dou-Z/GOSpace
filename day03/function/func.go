package main

import (
	"fmt"
)

func add(a, b int) int {
	// fmt.Printf("hello world\n")
	sum := a + b
	return sum
}

// 返回多个值
func moreReturnVal(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub

}

func add2(a, b int) (sum int, sub int) {
	// fmt.Printf("hello world\n")
	sum = a + b
	sub = a - b
	return
}

func calc(a, b int) (sum int, sub int) {
	// fmt.Printf("hello world\n")
	sum = a + b
	sub = a - b
	return
}

// ... 表示收一个或多个可变参数
func calc_v1(b ...int) int {
	sum := 0
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]

	}
	return sum
}

func main() {
	// s := add(100, 200)
	// fmt.Println(s)
	// val1,val2 := moreReturnVal(100,90)
	// fmt.Println(val1,val2)
	// val1, val2 := add2(200, 100)
	// fmt.Println(val1, val2)

	// 接受返回值的方式
	// _, _ = calc(100, 200)      // 表示不接受返回值
	// vall1, _ := calc(100, 300) // 接收第一个返回值
	// fmt.Printf("%d _", vall1)
	sum := calc_v1(10, 20, 30, 40, 50)
	fmt.Printf("sum=%d\n", sum)
}
