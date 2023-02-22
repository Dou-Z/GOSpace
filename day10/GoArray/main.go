package main

import "fmt"

func test1() {
	nums := [3]int8{11, 22, 33}
	// nums := [3]int32{11, 22, 33}

	fmt.Printf("数组的内存地址：%p\n", &nums)
	fmt.Printf("数组第一个元素的内存地址：%p\n", &nums[0])
	fmt.Printf("数组第2个元素的内存地址：%p\n", &nums[1])
	fmt.Printf("数组第3个元素的内存地址：%p\n", &nums[2])

	name := [2]string{"武沛琪", "Alex"}
	fmt.Printf("数组的内存地址：%p\n", &name)
	fmt.Printf("数组第一个元素的内存地址：%p\n", &name[0])
	fmt.Printf("数组第2个元素的内存地址：%p\n", &name[1])
}

//拷贝，重新拷贝
func testCopy() {
	name := [2]string{"武沛琪", "Alex"}
	name2 := name
	name2[1] = "Douz"
	fmt.Println(name, name2)
}

func TestArr() {
	var nesrData [2][3]int
	nesrData[0] = [3]int{11, 22, 33}
	nesrData[1][1] = 666
	fmt.Println(nesrData)
}

func main() {
	testCopy()
	TestArr()
}
