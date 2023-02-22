package main

import (
	"fmt"
	"strings"
	"time"
)

func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}
func testclosure() {
	f := Adder()
	ret := f(1)
	fmt.Printf("ret:%T\n", f)
	fmt.Printf("f(1):ret=%d\n", ret)
	ret = f(20)
	fmt.Printf("f(20):ret=%d\n", ret)
	ret = f(100)
	fmt.Printf("f(100):ret=%d\n", ret)
}

func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}
func test1() {
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))
}

// example3
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

}
func test2() {
	func1 := makeSuffixFunc(".bmp")
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func1("test"))
	fmt.Printf(func2("test"))
}

// 返回两个闭包函数
func calcFunc(base int) (func(int) int, func(int) int) {

	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func testCalc() {
	f1, f2 := calcFunc(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))

}

func testGo() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)

}
func testGo2() {
	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Println(index)
		}(i)
	}
	time.Sleep(time.Second)

}

func main() {
	// testclosure()
	// test2()
	// testCalc()
	testGo()
	testGo2()
}
