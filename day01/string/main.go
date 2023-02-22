package main

import (
	"fmt"
)

func testString() {
	var str = "hello"
	fmt.Printf("str[0]=%c len(str)=%d\n", str[0], len(str))

	for index, val := range str {
		fmt.Printf("str[%d]=%c\n", index, val)
	}

	var byteSlice []byte
	byteSlice = []byte(str)
	byteSlice[0] = '0'
	str = string(byteSlice)
	fmt.Println("after modify:", str)
}
func testReverseStrV1() {
	var str = "hello "
	var bytes []byte = []byte(str)

	for i := 0; i < len(str)/2; i++ {
		tmp := bytes[len(str)-i-1]
		bytes[len(str)-i-1] = bytes[i]
		bytes[i] = tmp
	}
	str = string(bytes)
	fmt.Println(str)
}
func testReverseStrV2() {
	var str = "hello 中文"
	var r []rune = []rune(str)

	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}
	str = string(r)
	fmt.Println(str)
}

func testHUiWen() {
	var str = "上海自来水来自海上"
	var r []rune = []rune(str)

	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}
	str2 := string(r)
	
	if str2==str{
		fmt.Println(str,"is huiwen")
	}else{
		fmt.Println(str,"is not huiwen")
	}
}

func main() {
	// testString()
	// testReverseStrV1()
	testReverseStrV2()
	testHUiWen()
}
