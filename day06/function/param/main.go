package main

import "fmt"

/* 定义相互交换值的函数 */
func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func testn(s string,n ...int)string{
	var x int
	fmt.Println(n)
	for _,i := range n{
		x+= i
	}
	return fmt.Sprintf(s,x)
}

func main() {
	// var a, b int = 1, 2

	// swap(&a, &b)
	// fmt.Println(a, b)
	println(testn("sum: %d", 1, 2, 3))

}
