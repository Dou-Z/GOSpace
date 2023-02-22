package main

import (
	"fmt"
)

func test(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Printf("a.()type.value:%d\n", a.(int))
	default:
		fmt.Println("其他类型")
	}

	switch v := a.(type) {
	case string:
		fmt.Printf("a的值为：%s\n", v)
	case int:
		fmt.Printf("a的值为：%d\n", v)
	default:
		fmt.Println("其他类型")
	}

	// s, o := a.(int)
	// if o {
	// 	fmt.Println(s)
	// 	return
	// }
	// f1, ok := a.(string)
	// if ok {
	// 	fmt.Println(f1)
	// 	return
	// }
	// f2, ok := a.(float32)
	// if ok {
	// 	fmt.Println(f2)
	// 	return
	// }

}
func main() {
	var a = 100
	test(a)

	var b = "hello"
	test(b)
	var c float32 = 31.5
	test(c)

}
