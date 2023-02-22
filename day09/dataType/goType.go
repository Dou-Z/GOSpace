package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func testType() {
	var v1 int8 = 10
	var v2 int16 = 20
	v3 := v1 + int8(v2)

	fmt.Println(v3, reflect.TypeOf(v3))
}

// 高位向低位转换，高位超过低位的最大值，会出现轮回转换
func testType2() {
	var v0 int16 = 130
	v2 := int8(v0)
	fmt.Println(v2)
}

func testTasnpore() {
	// 整型转为字符串类型
	v1 := 20
	result := strconv.Itoa(v1)
	fmt.Println(result, reflect.TypeOf(result))
	var v2 int8 = 18
	data := strconv.Itoa(int(v2))
	fmt.Println(data, reflect.TypeOf(data))
}

func testTran() {
	// 字符串转整型
	v1 := "6"
	res, err := strconv.Atoi(v1)
	if err == nil {
		fmt.Println("转换成功", res, reflect.TypeOf(res))
	} else {
		fmt.Println("转换失败", res)
	}
}


func bitTran(){
	v1 := 99
	// 整型转换2，8，10，16 进制
	r1 := strconv.FormatInt(int64(v1),2)
	fmt.Printf("2进制：%s\n",r1)

	r2 := strconv.FormatInt(int64(v1),8)
	fmt.Printf("8进制：%s\n",r2)

	r3 := strconv.FormatInt(int64(v1),16)
	fmt.Printf("16进制：%s\n",r3)

	r4 := strconv.FormatInt(int64(v1),10)
	fmt.Printf("10进制：%s\n",r4)
}

func bitTran2(){

	// 其他进制转化为10进制
	/*strconv.ParseInt(data,2,16)
		2,把data当作二进制转换成十进制（整型）
		16，转换过程中对结果进行约束
		结果：如果转换成功，err=nul，result则以int64的类型返回
		
	*/
	data:="10110101"
	result,err := strconv.ParseInt(data,2,16)
	fmt.Println(result,err,reflect.TypeOf(result))
}


func main1() {
	// testType2()
	// testTasnpore()
	// testTran()
	// bitTran()
	bitTran2()
}
