package main

import (
	"fmt"
	"reflect"
)

func testTruct() {
	// 定义一个结构体（类型），每个结构体包含name，age，hobby三个元素
	type Person struct {
		name  string
		age   int
		hobby []string
	}
	// 方式1：先后顺序
	var p1 = Person{"武沛琪", 19, []string{"篮球", "鸡哥"}}
	fmt.Println(p1.name, p1.age, p1.hobby)
	// 方式二：关键字
	var p2 = Person{name: "QWE", age: 19, hobby: []string{"饺子", "嫂子"}}
	fmt.Println(p2.name, p2.age, p2.hobby)
	// 方式二：先声明在赋值
	var p3 Person
	p3.name = "ABC"
	p3.age = 20
	p3.hobby = []string{"游戏", "篮球"}
	fmt.Println(p3.name, p3.age, p3.hobby)

}
func testStructPointer() {
	type Person struct {
		name string
		age  int
	}
	// 初始化结构体
	p1 := Person{"大浪", 10}
	fmt.Println(p1.name, p1.age)

	// 初始化结构体指针
	// var p2 *Person = &Person{"大浪",10}
	p2 := &Person{"大浪", 10}
	fmt.Println(p2, p2.name, p2.age)

	var p3 *Person = new(Person)
	p3.name = "武沛琪"
	p3.age = 18
	fmt.Println(p3.name, p3.age)
}

// 结构体嵌套

func testTruct1() {
	type Address struct {
		city, state string
	}
	type Person struct {
		name    string
		age     int
		Address Address
	}

	p1 := Person{name: "二狗子", age: 19, Address: Address{"北京", "BJ"}}
	p2 := p1

	fmt.Println(p1.Address)
	fmt.Println(p2.Address)
	p1.Address.city = "上海"
	fmt.Println(p1.Address)
	fmt.Println(p2.Address)
}

func testTruct2() {
	type Person struct {
		name   string
		age    int
		hobby  [2]string
		num    []int
		parent map[string]string
	}
	p1 := Person{
		name:   "二狗子",
		age:    19,
		hobby:  [2]string{"qaz", "qwe"},
		num:    []int{11, 22, 33, 44},
		parent: map[string]string{"father": "Douz", "mother": "dd"},
	}
	p2 := p1
	fmt.Println(p1)
	fmt.Println(p2)
	p1.age = 20
	p1.hobby[0] = "打游戏"
	p1.num[0] = 6666
	p1.parent["mother"] = "moth"
	fmt.Println(p1)
	fmt.Println(p2)

}

func testTructTag() {
	type Person struct {
		name string "姓名"
		age  int    "年龄"
		blog string "博客"
	}
	p1 := Person{
		name: "二狗子",
		age:  19,
		blog: "https://www.baidu.com",
	}

	p1Type := reflect.TypeOf(p1)

	// 1，
	filed1 := p1Type.Field(0)
	fmt.Println(filed1.Tag)
	// 2
	filed2, _ := p1Type.FieldByName("blog")
	fmt.Println(filed2.Tag)
	//  循环
	filedNum := p1Type.NumField() // 总共有多少个字段
	for index := 0; index < filedNum; index++ {
		filed := p1Type.Field(index)
		fmt.Println(filed.Name, filed.Tag)
	}
}

func main() {
	// testTruct()
	// testStructPointer()
	// testTruct1()
	// testTruct2()
	testTructTag()
}
