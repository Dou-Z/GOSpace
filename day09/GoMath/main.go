package main

import (
	"fmt"
	"math"
)

func TestMath() {
	//math的用法
	fmt.Println(math.Abs(-19))
	fmt.Println(math.Floor(3.14))
	fmt.Println(math.Ceil(3.14))
	fmt.Println(math.Round(3.14))
	fmt.Println(math.Round(3.14159*100)/100)
	fmt.Println(math.Mod(11,3))
	fmt.Println(math.Pow(2,5))
	fmt.Println(math.Pow10(2))
	fmt.Println(math.Max(1,2))
	fmt.Println(math.Min(1,2))
	fmt.Println(math.Jn(3,3.15))


}

func test1(){

	var v2 *int

	// *v2 = 9
	fmt.Println(&v2)
}

func main() {
	// TestMath()
	test1()
}