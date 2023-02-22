package main

import (
	"fmt"
	"math/big"
)

func BigIntTest() {
	// 创建超大整型对象
	var v1 big.Int

	// 在超大整型对象中写入值
	v1.SetInt64(9223372036854775807)
	fmt.Println(v1)
	v1.SetString("922337203685477588202021551", 10)
	fmt.Println(v1)
	fmt.Println(v1.String())
	v1.SetString("101110101011101101", 2)
	fmt.Println(v1)

	/*使用超大整型，推荐使用指针创建*/
	//创建一个超大整型指针对象
	v3 := new(big.Int)
	v4 := new(big.Int)
	//写入值
	v3.SetInt64(545897512852331863)
	fmt.Println(v3)
	v4.SetString("99999999999999999999999999999",10)
	fmt.Println(v4)
	


}

func BigIntMath(){
	n1 := new(big.Int)
	n1.SetString("99999999999999999999999999999998",10)


	n2 := new(big.Int)
	n2.SetString("10000000000000000000000000000002",10)

	result := new(big.Int)
	result.Add(n1,n2)

	fmt.Println(result.String())
}

func main() {
	// BigIntTest()
	BigIntMath()
}
