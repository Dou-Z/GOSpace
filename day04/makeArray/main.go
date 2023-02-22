package main

import (
	"fmt"
)

func testMake1() {
	var a []int
	a = make([]int, 5, 10)
	a[0] = 10
	a[1] = 20
	fmt.Printf("a=%v\n", a)

}

func testMake2() {
	var a []int
	a = make([]int, 5, 10)
	a[0] = 10
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap((a)))
	a = append(a, 11)
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	for i := 0; i < 8; i++ {
		a = append(a, i)
		fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	}
	// 观察切片的扩容操作，扩容的策略是翻倍扩容
	a = append(a, 1000)
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
}

func testCap() {
	a := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	b := a[2:4]
	fmt.Printf("b:%v len:%d cap:%d\n", b, len(b), cap(b))
}

func testSlice3() {
	var a []int
	fmt.Printf("%p,len:%d,cap:%d\n", a, len(a), cap(a))
	if a == nil {
		fmt.Printf("a is nil\n")
	}

	a = append(a, 100)
	fmt.Printf("%p,len:%d,cap:%d\n", a, len(a), cap(a))
	a = append(a, 200)
	fmt.Printf("%p,len:%d,cap:%d\n", a, len(a), cap(a))
	a = append(a, 300)
	fmt.Printf("%p,len:%d,cap:%d\n", a, len(a), cap(a))
}

// append 插入操作
func testAppend() {
	var a []int = []int{1, 2, 3}
	var b []int = []int{4, 5, 6}

	a = append(a, 23, 45, 55)
	fmt.Printf("a=%v\n", a)
	a = append(a, b...)
	fmt.Printf("a=%v\n", a)
}

// copy的使用
func testCopy() {
	// var a []int = []int{1, 2, 3}
	var a []int = []int{1, 2}
	var b []int = []int{4, 5, 6}
	copy(a, b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func testLiXI() {
	var sa = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		sa = append(sa, string(fmt.Sprintf("%v", i)))
	}
	fmt.Println(sa)
}

func main() {
	// testMake1()
	// testMake2()
	// testCap()
	// testSlice3()
	// testAppend()
	// testCopy()
	testLiXI()
}
