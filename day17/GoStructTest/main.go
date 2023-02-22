package main

import "fmt"

type Address struct {
	Name string
	area string
}
type User struct {
	Name string
	Age  int
	add  Address
}

func main() {
	var user01 = User{}
	user01.Name = "douz"
	user01.Age = 11
	user01.add.Name = "成都"
	user01.add.area = "GX"
	fmt.Println(user01)
}
