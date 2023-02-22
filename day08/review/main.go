package main

import "fmt"

func testscanf() {
	var name string
	fmt.Scanf("%s", &name)

	fmt.Println(name)
}


func val_variable(){
	var (
		name = "武侯区"
		age = 11
		hobby = "大保健"
		salary = 1000000
		gender string
		length int
		sb bool
	)
	fmt.Println(name,age,hobby,salary,gender,length,sb)
}

var sh = "daxue"

func val_test1(){
	fmt.Println(sh)

}

func addr_test(){
	name := "哈哈"
	nickname := name

	fmt.Println(name,&name)
	fmt.Println(nickname,&nickname)
}

/*常量*/
func testConst(){
	// 变量
	name := "go lang"
	name = "alex"
	fmt.Println(name)

	// 常量
	const age = 10
	fmt.Println(age)

}

func test_ioat(){
	const (
		v1 = iota
		v2
		v3
		v4 = "hello 1"
		v5


	)
	fmt.Println(v1,v2,v3,v4,v5)
}

func main() {
	// testscanf()
	// val_variable()
	// addr_test()
	// testConst()
	test_ioat()
}
