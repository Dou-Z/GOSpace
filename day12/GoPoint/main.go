package main

import "fmt"

func testPoint(){
	v1 := "武沛琪"
	v2 := &v1

	fmt.Println(v1)
	fmt.Println(v2, *v2)
	v1 = "Axel"
	fmt.Println(v1)
	fmt.Println(v2, *v2)
}
func test2(){
	name := "hello在吗"
	var p1 *string =&name
	var p2 **string = &p1
	var p3 ***string = &p2

	fmt.Println(name,&name)
	fmt.Println(p1,p2,p3)
}

func main() {
	test2()
}
