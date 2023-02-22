package main

import (
	"encoding/json"
	"fmt"
)

func testGoMarshal() {
	v1 := []interface{}{
		"WPQ",
		3.1415926,
		true,
		map[string]interface{}{
			"name": "Douz",
			"age":  22,
		},
		[3]interface{}{1, "成都", false},
	}
	res, ok := json.Marshal(v1)
	data := string(res)
	fmt.Println(data, ok)
}

func testGoUnMarshal() {
	content := `["WPQ",3.1415926,true,{"age":22,"name":"Douz"},[1,"成都",false]]`
	var value []interface{}
	json.Unmarshal([]byte(content), &value)
	fmt.Print(value)

}

func main() {
	testGoMarshal()
	testGoUnMarshal()
}
