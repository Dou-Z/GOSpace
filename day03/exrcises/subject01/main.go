package main

import (
	"fmt"
)

// 个人练习
func task1() {
	var res int
	for i := 1; i < 101; i++ {
		if i == 3 || i == 2 {
			fmt.Println(i)
		} else {
			for j := 2; j < i; j++ {
				res = i % j
				if res == 0 {
					break
				}
			}
			if res == 0 {
				continue
			} else {
				fmt.Println(i)
			}
		}
	}
}

// 课堂代码
func justify(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func example1() {
	for i := 2; i < 100; i++ {
		if justify(i) == true {
			fmt.Printf("[%d] is prime\n", i)
		}
	}
}

func main() {
	// task1()
	example1()
}
