package main

import "fmt"

func testMap() {
	var a map[string]int
	// fmt.Println(a)
	fmt.Printf("a:%v\n", a)
	// 只声明不能直接试用，需要初始化
	if a == nil {
		a = make(map[string]int, 16)
		fmt.Printf("a=%v\n", a)
		a["student"] = 18
		a["student1"] = 19
		a["student2"] = 68
		fmt.Printf("a=%v\n", a)
	}
}

func testMap1() {
	var b map[string]int = map[string]int{
		"stu01": 18,
		"stu02": 19,
		"stu03": 20,
	}
	fmt.Printf("b:%v\n",b)
	fmt.Println("b[stu01]:",b["stu01"])
}



func main() {
	// testMap()
	testMap1()
}
