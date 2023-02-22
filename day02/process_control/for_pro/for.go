package main

import (
	"fmt"
)

func testFor1() {
	var i int
	for i = 1; i <= 10; i++ {
		fmt.Printf("i=%d\n", i)
	}
	fmt.Printf("Final i=%d", i)
}

// for 语句的另一种写法
func testFor2() {
	i := 1
	for i <= 10 { // 分号也可以不写
		fmt.Printf("i=%d\n", i)
		i += 2
	}
}

func testFor3() {
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d*%d=%d\n", no, i, no*i)
	}
}

// 无限循环

func testLoop() {
	for {
		fmt.Printf("hello\n")
	}
}

func main() {
	// testFor1()
	// testFor2()
	testFor3()
	testLoop()
}
