package main

import (
	"fmt"
)

func test_input() {
	var name string
	fmt.Println("请输入用户名：")
	fmt.Scan(&name)
	fmt.Printf("%s\n", name)
}
func test_input2() {
	var name string
	var age int

	fmt.Println("请输入用户名、密码：")
	// 当使用scan时，会提示用户输入
	// 用户输入完后，会得到两个值。count：用户输入了几个值，err：输入错误时的异常信息
	count, err := fmt.Scan(&name, &age)
	// 没有输入完，会一直等待
	fmt.Println(count, err)
	fmt.Println(name, age)
}

func test_scanln() {
	var name string
	var age int

	fmt.Print("请输入用户名、密码：")
	// 当使用scanln时，会提示用户输入
	// 用户输入完后，会得到两个值。count：用户输入了几个值，err：输入错误时的异常信息
	count, err := fmt.Scanln(&name, &age)
	fmt.Println(count, err)
	fmt.Println(name, age)
}


func test_scanf(){
	var name string
	var age int

	fmt.Print("请输入用户名：")
	// 占位符后要用空格隔开，否则会当作一个值
	fmt.Scanf("I am %s,i am %d ",&name,&age) 
	fmt.Println(name,age)
}
func test_scanfln(){
	var name string
	
	fmt.Print("请输入用户名：")
	// 使用scanln，输入时不能输入空格，否则只会取空格前的值
	fmt.Scanln(&name) 
	fmt.Println(name)
}


func main() {
	// test_input2()
	// test_scanln()
	// test_scanf()
	test_scanfln()

}
