package main

import (
	"fmt"
	"unsafe"
)

func main() {
	dataList := [3]int8{11, 22, 33}
	// 获取数组第一个元素地址
	var firstDataPtr *int8 = &dataList[0]
	// 转换成pointer类型
	ptr := unsafe.Pointer(firstDataPtr)
	// 转换为uintptr类型，然后进行内存地址的计算（即：增加一个字节，取第二个索引位置的值）
	targetAddress := uintptr(ptr) + 1
	// 根据地址重新转换为pointer类型
	newPtr := unsafe.Pointer(targetAddress)
	// pointer对象转换为int8指针类型
	value := (*int8)(newPtr)
	// 根据指针获取值
	fmt.Printf("最后结果为：%v\n", *value)
}
