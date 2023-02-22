package main

import (
	"fmt"
	"time"
)

func switchTest() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n)
	switch n {
	case 1:
		fmt.Println("I am 1")
	case 2:
		fmt.Println("I am 2")
	case 3:
		fmt.Println("I am 333")
	case 4:
		fmt.Println("I am ")
	default:
		fmt.Println("I am default")
	}

}

// for循环
// 示例1：
func test_for1() {
	for {
		fmt.Println("死循环")
		time.Sleep(time.Second * 1)
	}
}

// 示例2：
func test_for2() {
	// num := 1
	// fmt.Println("start")
	// for num < 5{
	// 	fmt.Println("test")
	// 	num = 10
	// }
	// fmt.Println("end")
	number := 1
	for number < 5 {
		fmt.Println("学习")
		number++
	}
	fmt.Println("结束")
}

func test_break() {
	for {
		fmt.Println("hhhhhh")
		break
		fmt.Println("aaaaaaaa")
	}
}

/* 对for循环打标签 */
func test_For_Lable() {
f1:
	for i := 1; i < 3; i++ {

		for j := 1; j < 3; j++ {
			if j == 3 {
				continue f1
			}
			fmt.Println(i,j)
		}
	}
}

func test_For_Lable2() {
	f2:
		for i := 1; i < 3; i++ {
	
			for j := 1; j < 5; j++ {
				if j == 3 {
					break f2
				}
				fmt.Println(i,j)
			}
		}
	}
	

func main() {
	// switchTest()
	// test_break()
	// test_For_Lable()
	test_For_Lable2()
}
