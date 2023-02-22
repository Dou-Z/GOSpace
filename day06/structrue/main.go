package main

import (
	"fmt"
)

// 类型定义
type NewInt int

// 类型别名
type MyInt = int

type person struct {
	name string
	city string
	age  int8
}

func testStruct() {

	var p1 person
	p1.name = "ABC"
	p1.city = "成都"
	p1.age = 10
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)

}

func testNiMing() {
	var user struct {
		Name string
		age  int
	}
	user.Name = "douz"
	user.age = 15
	fmt.Printf("%#v\n", user)
}

// 创建指针类型结构体

func testPointStruct() {
	var p2 = new(person)
	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	fmt.Printf("p2=%#v\n", p2)
}

// 取结构体的地址实例化
func teststructExp() {
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)
	p3.name = "bb"
	p3.age = 20
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3)
}

func testStructInit() {
	var p4 person
	fmt.Printf("p4=%#v\n", p4)
}

// 结构体内存布局
func testStructRom() {
	type test struct {
		a int
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
}

func testexce(){
	type student struct {
		name string
		age  int
	}
	m := make(map[string]*student)
    stus := []student{
        {name: "pprof.cn", age: 18},
        {name: "测试", age: 23},
        {name: "博客", age: 28},
    }

    for _, stu := range stus {
		fmt.Println(stu.name)
		fmt.Printf("stu=%v\n",&stu)
        m[stu.name] = &stu
    }
	fmt.Println(m)
    for k, v := range m {
        fmt.Println(k, "=>", v.name)
    }
}

type Person struct {
    name string
    age  int8
}

func NewPerson(name string,age int8) *Person{
	return &Person{
		name:name,
		age:age,
	}
}

func (p Person) Dream(){
	fmt.Printf("%s====s%\n",p.name)
}

func main() {
	// var a NewInt
	// var b MyInt

	// fmt.Printf("type of a :%T\n", a)
	// fmt.Printf("type of b :%T\n", b)
	// testStruct()
	// testNiMing()
	// testPointStruct()
	// teststructExp()
	// testStructInit()
	// testStructRom()
	// testexce()
	p1 := NewPerson("测试",22)
	p1.Dream()
}
