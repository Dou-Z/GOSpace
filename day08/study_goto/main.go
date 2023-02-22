package main

import "fmt"

func main() {
	var name string
	fmt.Print("请输入：")
	fmt.Scan(&name)

	if name == "SVIP"{
		goto SVIP
	}else if name == "VIP"{
		goto VIP
	}

	fmt.Println("预约。。。")
VIP:
	fmt.Println("排队。。。")
SVIP:
	fmt.Println("进入。。。")
}