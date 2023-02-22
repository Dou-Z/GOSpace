package main

import (
	"fmt"
	"os"
	"time"
)

func testTimeStr() {
	// 获取当前时间（本地时间）Time类型-结构体
	currentDate := time.Now()
	// 格式化处理
	dateString := currentDate.Format("2006-01-02-15-04-05")

	fmt.Println(dateString)
	// 创建文件夹
	os.Mkdir(dateString, 0755)
}

func testTimeUse() {
	//  1，获取当前时间
	t1 := time.Now()
	fmt.Printf("当前时间（Time类型）：%v\n", t1)
	// fmt.Println(t1)
	// 2,获取当前UTC时间 t2.Local()
	t2 := time.Now().UTC()
	fmt.Println("当前UTC时间（Time类型）：", t2)

	// 3，创建一个时间字符串类型 -》Time类型
	t3, _ := time.Parse("2006-01-02", "2022-11-01")
	fmt.Println("根据字符串转换为时间（Time）：", t3)

	// 4，创建一个时间
	t4 := time.Date(2022, 11, 01, 11, 12, 13, 11, time.Local)
	fmt.Println("根据字符串转换为时间（Time）：", t4)
	t5 := time.Date(2022, 12, 02, 12, 12, 12, 12, time.UTC)
	fmt.Println("根据字符串转换为时间（Time）：", t5)

	//5，时间格式化，time类型-》字符串类型
	// Mon Jan 2 15:04:05 -0700 MST 2006
	fmt.Println("格式化之后的时间为（String）：", t1.Format("2006-01-02 15:04:05.00000 -0700 MST"))
	fmt.Println("格式化之后的时间为（String）：", t2.Format("2006-01-02 15:04:05.00000 -0700 MST"))
	fmt.Println("格式化之后的时间为（String）：", t2.Format("2006-01-02 15:04:05"))
	fmt.Println("格式化之后的时间为（String）：", t2.Format("2006-01-03 15:04:06"))

	// 6,时间增加
	t6 := t1.Add(time.Hour * 1)
	fmt.Println("当前时间+1小时（Time类型）：", t6)
	// 7,时间减少
	t7 := t1.Add(-time.Minute * 1)
	fmt.Println(t7)

	// 8，时间间隔
	t8 := t1.Sub(t4)
	fmt.Printf("t1:%s\n", t1)
	fmt.Printf("t4:%s\n", t4)
	fmt.Println("时间间隔为（Duration类型）：", t8)
	fmt.Println("时间间隔小时为（Duration类型）：", t8.Hours())
	fmt.Println("时间间隔分钟为（Duration类型）：", t8.Minutes())
	fmt.Println("时间间隔秒为（Duration类型）：", t8.Seconds())

}
func main() {
	// testTimeStr()
	testTimeUse()
}
