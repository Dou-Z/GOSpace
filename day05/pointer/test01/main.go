package main

import (
	"fmt"
)

func testPointer01() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr :%p\n", a, &a)
	fmt.Printf("b:%p type :%T\n", b, &b)
	fmt.Println(&b)

}

/*
1.对变量进行取地址（&）操作，可以获得这个变量的指针变量。
2.指针变量的值是指针地址。
3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
*/
func testPoint02() {
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Println("value of c: %v\n", c)
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func testPointer03() {
	a := 10
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}

func testPoint04() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	fmt.Printf("p的指针是%p\n", &p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空")
	}
}

func testPoint05() {
	var a *int
	*a = 100
	fmt.Println(*a)
	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)
}

/*
执行上面的代码会引发panic，为什么呢？
在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，
还要为它分配内存空间，否则我们的值就没办法存储。
而对于值类型的声明不需要分配内存空间，
是因为它们在声明的时候已经默认分配好了内存空间。
要分配内存，就引出来今天的new和make。
Go语言中new和make是内建的两个函数，主要用来分配内存
*/

/*
func new(Type) *Type
1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
*/

func testNew() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*a)
	fmt.Println(*b)
}

func testMake() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)
}

/*new与make的区别*/
/*
	1.二者都是用来做内存分配的。
    2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
    3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/

func main() {
	// testPoint05()
	testNew()
	testMake()
}
