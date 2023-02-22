package main

import "fmt"

func TestSlice() {
	var name string = "A武沛琪"

	v1 := name[0]
	fmt.Println(v1)

	v2 := name[0:4]
	fmt.Println(v2)

	for i := 0; i < len(name); i++ {
		fmt.Println(i, name[i])

	}

	for index, item := range name {
		fmt.Println(index, item, string(item))
	}

	dataList := []rune(name)
	fmt.Println(dataList[0], string(dataList[0]))
}

func main() {
	TestSlice()
}
