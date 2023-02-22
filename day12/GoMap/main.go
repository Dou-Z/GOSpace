package main

import "fmt"

func testMap1() {
	data := make(map[string]int)
	data["100"] = 998
	data["200"] = 999

	// 声明 ，nil
	value := new(map[string]int)
	// value["k1"] = 123 // 报错
	value = &data
	fmt.Println(value)
}

func testMapLen() {
	data := map[string]string{"n1": "武沛琪", "n2": "Alex"}
	val := len(data)
	fmt.Println(val)
}

func testMapCap() {
	dict := make(map[string]int, 10)

	dict["n1"] = 1
	dict["n2"] = 2
	dict["n3"] = 3

	val := len(dict)
	// val2 := cap(dict) //不支持计算map的容量

	fmt.Println(val)

}

func testMapFor(){
	data := map[string]string{"n1":"q","n2":"w"}
	for key,value := range data{
		fmt.Println(key)
		fmt.Println(value)
	}
}
/*
v1 := make(map[int]int)
v2 := make(map[string]int)
v3 := make(map[float32]int)
v4 := make(map[[2]int]int)
v5 := make(map[[]int]int)  //键错误 切片不可哈希
v6 := make(map[map[string]int]int)  //键错误 map不可哈希

v7 := make(map[ [2][]int ]int)  //报错 存在切片
v8 := make(map[[2]map[string]string ]int)	//报错 存在map

*/


func main() {
	// testMap1()
	// testMapLen()
	// testMapCap()
	testMapFor()
}
