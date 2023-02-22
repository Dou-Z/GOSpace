package main

import (
	"github.com/study/hello/hello"
	_ "github.com/study/hello/hello"
)

func main() {
	hello.Print()
	//编译报错：./main.go:6:5: undefined: hello
}
