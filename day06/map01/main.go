package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func testMap01() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a :%T\n", scoreMap)

	// 如果key存在ok为true，v为对应的值；不存在ok为false，v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("None")
	}
}

// map也支持在声明的时候填充元素
func testMap02() {
	useInfo := map[string]string{
		"usesrname": "pprot.cn",
		"password":  "123456",
	}
	fmt.Println(useInfo)
}

func testMap03() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王二"] = 110
	for k, _ := range scoreMap {
		fmt.Println(k)
	}

}

// 使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：
// delete(map, key)
/*
map:表示要删除键值对的map
key:表示要删除的键值对的键
*/
func testMapDelete01() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王二"] = 110
	delete(scoreMap, "小明") //将小明：100删除
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

}

func testMap05() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value

	}
	// 取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

// 元素为map类型的切片
func testMapSlice() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value :%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小红"
	mapSlice[0]["passwd"] = "123546"
	mapSlice[0]["adress"] = "武侯大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
func testMapSlice02() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

func testMap06() {
	//直接创建初始化一个mao
	var mapInit = map[string]string{"xiaoli": "湖南", "xiaoliu": "天津"}
	//声明一个map类型变量,
	fmt.Println(mapInit)
	//map的key的类型是string，value的类型是string
	var mapTemp map[string]string
	//使用make函数初始化这个变量,并指定大小(也可以不指定)
	mapTemp = make(map[string]string, 10)
	//存储key ，value
	mapTemp["xiaoming"] = "北京"
	mapTemp["xiaowang"] = "河北"
	//根据key获取value,
	//如果key存在，则ok是true，否则是flase
	//v1用来接收key对应的value,当ok是false时，v1是nil
	v1, ok := mapTemp["xiaoming"]
	fmt.Println(ok, v1)
	//当key=xiaowang存在时打印value
	if v2, ok := mapTemp["xiaowang"]; ok {
		fmt.Println(v2)
	}
	//遍历map,打印key和value
	for k, v := range mapTemp {
		fmt.Println(k, v)
	}
	//删除map中的key
	delete(mapTemp, "xiaoming")
	//获取map的大小
	l := len(mapTemp)
	fmt.Println(l)
}

func main() {
	// testMap01()
	// testMap03()
	// testMap02()
	// testMapDelete01()
	// testMap05()
	// testMapSlice()
	// testMapSlice02()
	testMap06()
}
