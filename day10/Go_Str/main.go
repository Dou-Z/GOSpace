package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func TestStr1() {
	// 本质是utf-8 编码的序列
	var name string = "武佩奇"
	// 武 -》 11100110 10101101 10100110
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))

	// 佩 -》 11100100 10111101 10101001
	fmt.Println(name[3], strconv.FormatInt(int64(name[3]), 2))
	fmt.Println(name[4], strconv.FormatInt(int64(name[4]), 2))
	fmt.Println(name[5], strconv.FormatInt(int64(name[5]), 2))

	// 奇 -》 11100101 10100101 10000111
	fmt.Println(name[6], strconv.FormatInt(int64(name[6]), 2))
	fmt.Println(name[7], strconv.FormatInt(int64(name[7]), 2))
	fmt.Println(name[8], strconv.FormatInt(int64(name[8]), 2))

	fmt.Println(len(name))

	// 字符串转换为一个“字符集合”
	byteSet := []byte(name)
	fmt.Println(byteSet) //[230,173,166,228,189,169,229,165,135]
	//字节集合转换为字符串
	byteList := []byte{230, 173, 166, 228, 189, 169, 229, 165, 135}
	targetStr := string(byteList)
	fmt.Println(targetStr)
	//rune 方法
	tempSet := []rune(name)
	fmt.Println(tempSet)
	fmt.Println(tempSet[0], strconv.FormatInt(int64(tempSet[0]), 16))
	fmt.Println(tempSet[1], strconv.FormatInt(int64(tempSet[1]), 16))
	fmt.Println(tempSet[2], strconv.FormatInt(int64(tempSet[2]), 16))

	// 获取字符串的长度
	str_len := utf8.RuneCountInString(name)
	fmt.Println(str_len)

}

func Str_test_use() {

	name := "go语言"
	// 以xx开头
	result := strings.HasPrefix(name, "go")
	fmt.Println(result)
	//以XX结尾
	res2 := strings.HasSuffix(name, "言")
	fmt.Println(res2)
	// 是否包含
	res3 := strings.Contains(name, "语言")
	fmt.Println(res3)
	// 大小写转换
	str1 := "qweasd"
	res4 := strings.ToUpper(str1) //转为大写
	fmt.Println(res4)
	res5 := strings.ToLower(res4)
	fmt.Println(res5)

	//去除两边 Trim
	eg1 := "go学习go语言，gogogo"
	result1 := strings.TrimRight(eg1, "go") //去除右边的go
	result2 := strings.TrimLeft(eg1, "go")  //去除左边的go
	result3 := strings.Trim(eg1, "go")      //去除两边的
	fmt.Println(result1, result2, result3)
}

func Replace_Use() {

	name := "wupeiqipeipeiqiqiq"

	res1 := strings.Replace(name, "pei", "PE", 1)  // 找到pei替换成PE，从左到右替换第一个
	res2 := strings.Replace(name, "pei", "PE", 2)  // 找到pei替换成PE，从左到右替换第2个
	res3 := strings.Replace(name, "pei", "PE", -1) // 找到pei替换成PE，全部 替换

	fmt.Println(res1, res2, res3)
	// 分割
	str1 := "go学习go语言gocome on"
	result1 := strings.Split(str1, "go")
	fmt.Println(result1)

}

func PinJie_Str() {
	// 不建议使用
	message := "我爱" + "北京"
	fmt.Println(message)
	// 建议：效率高一些
	dataList := []string{"我爱", "北京"}
	result := strings.Join(dataList, "-join-")
	fmt.Println(result)

	// 效率更高(go 1.10 之前的)
	var buffer bytes.Buffer
	buffer.WriteString("你想")
	buffer.WriteString("wogan")
	buffer.WriteString("ta")

	data := buffer.String()
	fmt.Println(data)

	// 效率更高2(go 1.10 之后的)
	var builder strings.Builder
	builder.WriteString("哈哈哈哈")
	builder.WriteString("去你的吧")
	value := builder.String()
	fmt.Println(value)
}
func string_int() {
	v1 := string(65)
	fmt.Println(v1)

	v2 := string(27494)
	fmt.Println(v2)

	// 字符串转数字
	v3, size := utf8.DecodeRuneInString("A")
	fmt.Println(v3, size)
}

func main() {
	// TestStr1()
	// Str_test_use()
	// Replace_Use()
	// PinJie_Str()
	string_int()
}
