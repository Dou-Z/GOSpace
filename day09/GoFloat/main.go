package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func testFloat() {
	var v1 float32
	v1 = 3.14
	v2 := 99.3

	v3 := v1 + float32(v2)
	fmt.Println(v3)

	v4 := 0.1
	v5 := 0.2
	result := v4 + v5
	fmt.Println(result)

	v6 := 0.3
	v7 := 0.2
	data := v6 + v7
	fmt.Println(data)
}

func testDecimal() {
	var price = decimal.NewFromFloat(3.4626)
	var data1 = price.Round(1)    // 保留小数点后一位（四舍五入）
	var data2 = price.Truncate(1) // 保留小数点后一位
	fmt.Println(data1, data2)
}

func main() {
	// testFloat()
	testDecimal()
}
