package main

import "fmt"

func main() {
	var name, addr, action string
	fmt.Print("请输入name：")
	fmt.Scan(&name)

	fmt.Print("请输入addr：")
	fmt.Scan(&addr)

	fmt.Print("请输入action：")
	fmt.Scan(&action)

	result := fmt.Sprintf("我叫%s,现在在%s正在%s\n",name,addr,action)
	fmt.Println(result)

}