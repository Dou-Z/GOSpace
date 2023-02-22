package main

import "fmt"
//定义接口
type IBase interface {
	f1() string
}
// 定义结构体Person
type Person struct {
	name string
}
// 为结构体Person定义方法
func (p Person) f1() string {
	return p.name
}
// 定义结构体user
type User struct {
	name string
}
// 为user结构体定义方法
func (p User) f1() string {
	return p.name
}
// 基于接口的参数，可以实现传入多种类型（多态）参数，也同时具有约束对象必须实现接口方法的功能
func DoSomething(base IBase) {
	result := base.f1()  //直接调用  接口.f1（） ->找到对应的类型并执行其方法
	fmt.Println(result)

}

func main() {
	per := Person{name: "武沛琪"}
	user := User{name: "wpq"}

	DoSomething(per)
	DoSomething(user)
}
