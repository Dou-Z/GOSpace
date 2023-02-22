package main

import (
	"fmt"

	"github.com/study/day17/user"
)

type City struct {
	name     string
	areaList []string
	int
	string
}
func testanonyoums(){
	var rong City
	rong.name = "成都"
	rong.areaList = []string{"武侯区","金牛区","双流区","高新区"}
	rong.int = 100
	rong.string = "hhh"
	fmt.Printf("rong=%#v\n",rong)
}

func testUser(){
	u := user.NewUser("user01", "男", 18, "https://www.baidu.com/xxx.jpg")
	fmt.Println(u)
	fmt.Printf("u=%#v\n", u)
}

func main() {
	testanonyoums()
}