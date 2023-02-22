package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"name"`
	Sex      string `json:"性别"`
	Score    float32
}

func main() {
	user := &User{
		UserName: "user01",
		Sex:      "man",
		Score:    99.9,
	}
	data, _ := json.Marshal(user)
	fmt.Printf("json data:%s\n", string(data))
}
