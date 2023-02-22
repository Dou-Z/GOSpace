package main

import (
	"fmt"
)

// GO接口
type Base interface {
}

func testInterface() {
	// 定义切片，内部可以存放任意类型
	dataList := make([]Base, 0) // 推荐简写：dataList := make([]interface{},0)
	// 切片中添加字符串类型
	dataList = append(dataList, "成都高新")
	dataList = append(dataList, 450)
	dataList = append(dataList, 99.99)
	fmt.Println(dataList)
}

type Person struct {
	name string
	age  int
}

func testInterfaceTran(arg interface{}) {
	//接口转换Person成功，ok=true
	// fmt.Println()
	tp, ok := arg.(Person)
	if ok {
		fmt.Println(tp.name, tp.age)
	} else {
		fmt.Println("转换失败")
	}

}

type Role struct {
	title string
	count int
}

func testInterfaceTran2(arg interface{}) {
	switch tp := arg.(type) {
	case Person:
		fmt.Println(tp.name)
	case Role:
		fmt.Println(tp.title)
	case string:
		fmt.Println(tp)
	default:
		fmt.Println(tp)
	}

}

func main() {
	// testInterface()
	// testInterfaceTran(Person{name: "douz", age: 18})
	testInterfaceTran2("douz")
	testInterfaceTran2(666)
	testInterfaceTran2(9.99)
	testInterfaceTran2(Person{name: "Axle", age: 11})
	testInterfaceTran2(Role{title: "高新", count: 2})
}
