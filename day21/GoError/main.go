package main

import "fmt"

// 系统抛
func test01() {
	a := [5]int{0, 1, 2, 3, 4}
	a[1] = 123
	fmt.Println(a)
	index := 10
	a[index] = 10
	fmt.Println(a)

}
func getCircleArea(radius float32) (area float32, err error) {
	if radius < 0 {
		// 自己抛
		panic("半径不能小于0")
	}
	area = 3.14 * radius * radius
	return
}
func test02() {
	getCircleArea(-5)
}
func test03() {
	// 延时执行匿名函数
	// 延时到何时？（1）程序正常结束（2）发生异常
	defer func() {
		// recover() 复活 恢复
		// 会返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getCircleArea(-3)
	fmt.Println("这里有没有执行")
}
func test04() {
	test03()
	fmt.Println("test04")
}

func main() {
	// test04()
	area, err := getCircleArea(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area)
	}
}
