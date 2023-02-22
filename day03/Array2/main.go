package main

import (
	"fmt"
)

var arr0 [5][3]int
var arr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {7, 8, 9}}

func main() {
	a := [2][3]int{{1, 2, 3}, {23, 4, 5}}
	b := [...][2]int{{1, 2}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
	fmt.Println(arr0, arr1)
	fmt.Println(a, b)
	for index, val := range a {
		for i, j := range val {
			fmt.Println(i, "===", j)
		}
		fmt.Println(index, "====", val)
	}

}
